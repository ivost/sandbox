package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
	"os"
)

/*

	https://github.com/hashicorp/vault/tree/master/api

	https://github.com/hashicorp/vault-ruby

200 - Success with data.
204 - Success, no data returned.
400 - Invalid request, missing or invalid data.
403 - Forbidden, your authentication details are either incorrect, you don't have access to this feature,
	or - if CORS is enabled - you made a cross-origin request from an origin that is not allowed to make such requests.
404 - Invalid path. This can both mean that the path truly doesn't exist or that you don't have permission to view a specific path. We use 404 in some cases to avoid state leakage.
429 - Default return code for health status of standby nodes. This will likely change in the future.
473 - Default return code for health status of performance standby nodes.
500 - Internal server error. An internal error has occurred, try again later. If the error persists, report a bug.
502 - A request to Vault required Vault making a request to a third party; the third party responded with an error of some kind.
503 - Vault is down for maintenance or is currently sealed. Try again later.

https://github.com/tuenti/pouch

Pouch and friends are a set of tools to manage provisioning of secrets on hosts
based on the AppRole authentication method of Vault.

https://github.com/DaspawnW/vault-crd


Vault-CRD is a custom resource definition for holding secrets that are stored in HashiCorp Vault up to date with Kubernetes secrets.

The following Secret engines of Vault are supported:

KV (Version 1)
KV (Version 2)
PKI
The following types of secrets can be managed by Vault-CRD:

Docker Pull Secret (DockerCfg)
Ingress Certificates
JKS Key Stores

https://vault.koudingspawn.de/supported-secret-types/secret-type-pki
https://koudingspawn.de/advanced-ingress/
https://github.com/DaspawnW/vault-crd-helm
https://vault.koudingspawn.de/install-vault-crd#kubernetes-service-account-authentication
*/
func init() {
	// Ensure our special envvars are not present
	os.Setenv("VAULT_ADDR", "http://127.0.0.1:8200")
	os.Setenv("VAULT_TOKEN", "root")
}

var res *api.Secret
var client *api.Client
var err error

// https://www.vaultproject.io/api/secret/pki/index.html#generate-root

type CertReq struct {
	CSR               string `json:"csr,omitempty"`
	CommonName        string `json:"common_name"`
	AltNames          string `json:"alt_names,omitempty"`
	IpSans            string `json:"op_sans,omitempty"`
	UriSans           string `json:"uri_sans,omitempty"`
	OtherSans         string `json:"other_sans,omitempty"`
	TTL               string `json:"ttl,omitempty`
	Format            string `json:"format,omitempty"`
	ExcludeCNFromSans bool   `json:"exclude_cn_from_sans,omitempty"`
}

func main() {
	log.Printf("Hello, Vault API")

	config := api.DefaultConfig()
	client, err = api.NewClient(config)
	if err != nil {
		panic("NewClient error: " + err.Error())
	}
	k := "pki"
	c, _ := ReadCaPem(k)
	log.Printf("pki: %v, CA cert: %v", k, c)

	k = "pki_sub1"
	role := "brain_sub1"

	c, _ = ReadCaPem(k)
	log.Printf("pki: %v, CA cert: %v", k, c)
	rq := CertReq{CommonName: "test1.rocbox.braincorp.com",
		Format: "pem",
		TTL:    "1h",
	}

	crt, err := IssueCert(k, role, rq)
	if err != nil {
		panic("IssueCert error: " + err.Error())
	}
	log.Printf("issued cert: %v", crt)

	csr, err := generateCSR(rq.CommonName)
	if err != nil {
		panic("generateCSR error: " + err.Error())
	}
	rq.CSR = string(csr)
	crt, err = SignCSR(k, role, rq)
	if err != nil {
		panic("SignCSR error: " + err.Error())
	}
	log.Printf("issued cert: %v", crt)

	certs, err := readStrList("pki_sub1/certs", "keys")
	if err != nil {
		panic("readList error: " + err.Error())
	}

	log.Printf("%v certs", len(certs))
	for i, c := range certs {
		log.Printf("cert %v, serial: %v", i, c)
		// - and : work, removing them doesn't
		//c = strings.Replace(c, "-", ":", -1)
		k := "pki_sub1/cert/" + c
		c, err := ReadCert(k)
		if err != nil {
			panic("ReadCert error: " + err.Error())
		}
		_ = c
		//log.Printf("cert: %v", c)
	}
}

func IssueCert(pki string, role string, req CertReq) (string, error) {
	rq, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	dynamic := make(map[string]interface{})
	json.Unmarshal(rq, &dynamic)
	path := fmt.Sprintf("%v/issue/%v", pki, role)
	resp, err := post(path, dynamic)
	//todo: define resp structs and json field names
	if err != nil {
		return "", err
	}
	//log.Printf("path: %v, PEM:\n\n%v\n\n", path, pem)
	return resp, nil
}

func SignCSR(pki string, role string, csr CertReq) (string, error) {
	rq, err := json.Marshal(csr)
	if err != nil {
		return "", err
	}
	dynamic := make(map[string]interface{})
	json.Unmarshal(rq, &dynamic)
	path := fmt.Sprintf("%v/sign/%v", pki, role)
	resp, err := post(path, dynamic)
	//todo: define resp structs and json field names
	if err != nil {
		return "", err
	}
	//log.Printf("path: %v, PEM:\n\n%v\n\n", path, pem)
	return string(resp), nil
}

func post(path string, data map[string]interface{}) (string, error) {
	res, err := client.Logical().Write(path, data)
	if err != nil {
		return "", err
	}
	if res == nil {
		return "", nil
	}
	if len(res.Warnings) > 0 {
		log.Printf("warnings: %+v", res.Warnings)
	}
	s, err := json.Marshal(res.Data)
	return string(s), err
}

//
func ReadCaPem(pki string) (string, error) {
	path := fmt.Sprintf("/v1/%v/ca/pem", pki)
	pem, err := readPEM(path)
	if err != nil {
		return "", err
	}
	//log.Printf("path: %v, PEM:\n\n%v\n\n", path, pem)
	return pem, nil
}

func readPEM(path string) (string, error) {
	//log.Printf("readPEM - path: %v", path)
	r := client.NewRequest("GET", path)
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	resp, err := client.RawRequestWithContext(ctx, r)
	if err != nil {
		return "", err
	}
	if resp == nil || resp.Body == nil {
		return "", nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		//log.Printf("not found")
		return "", nil
	}
	var buf bytes.Buffer
	n, err := buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	if n == 0 {
		return "", nil
	}
	d := buf.Bytes()
	//log.Printf("got %v bytes", len(d))
	return string(d), nil
}

func ReadCert(path string) (string, error) {
	//log.Printf("ReadCert path: %v", path)
	res, err = client.Logical().Read(path)
	if err != nil {
		return "", err
	}
	if res == nil {
		return "", nil
	}
	if len(res.Warnings) > 0 {
		log.Printf("warnings: %+v", res.Warnings)
	}
	//log.Printf("res resp: %+v", res)
	//log.Printf("data: %+v", res.Data)

	cert, ok := res.Data["certificate"]
	if ok {
		//log.Printf("data: %+v", res.Data)
		return cert.(string), nil
	}
	return "", nil
}

func read(path string) {
	log.Printf("path: %v", path)
	res, err = client.Logical().Read(path)
	if err != nil {
		panic("Read error: " + err.Error())
	}
	if res == nil {
		panic("nil result")
	}
	if len(res.Warnings) > 0 {
		log.Printf("warnings: %+v", res.Warnings)
	}
	//log.Printf("res resp: %+v", res)
	log.Printf("data: %+v", res.Data)
}

func readStrList(path string, key string) ([]string, error) {
	lst, err := readList(path, key)
	if err != nil {
		return nil, err
	}
	sl := make([]string, len(lst))
	for i := range lst {
		sl[i] = lst[i].(string)
	}
	return sl, nil
}

func readList(path string, key string) ([]interface{}, error) {
	// a nil slice is functionally equivalent to a zero-length slice,
	// even though it points to nothing.
	// It has length zero and can be appended to, with allocation.
	var e []interface{}
	//e = make([]interface{}, 0)
	res, err = client.Logical().List(path)
	if err != nil {
		return e, err
	}
	if res == nil {
		return e, errors.New("nil result")
	}
	if len(res.Warnings) > 0 {
		log.Printf("res warnings: %+v", res.Warnings)
	}
	// result of List is a map
	//  map[keys:[FOO foo]]
	//log.Printf("res resp: %+v", res)
	//log.Printf("res data: %+v", res.Data)
	lst, ok := res.Data[key].([]interface{})
	if !ok {
		return e, errors.New("result is not array")
	}
	return lst, nil
}

/*
	//testList()
	//testKVv1()
	// not working
	// https://www.vaultproject.io/api/secret/kv/kv-v2.html
	//testKVv2()

vault kv  res kv
Keys
----
FOO
foo
foo/

vault kv  res kv/foo
Keys
----
bar
*/
func testKVv1() {
	k := "kv/foo/bar"
	log.Printf("key: %v", k)
	m := make(map[string]interface{})
	m["aaa"] = "bbb"
	m["num"] = 123
	res, err = client.Logical().Write(k, m)
	if err != nil {
		panic("Write error: " + err.Error())
	}

	res, err = client.Logical().Read(k)
	if err != nil {
		panic("Logical error: " + err.Error())
	}
	if res == nil {
		panic("nil result")
	}
	if len(res.Warnings) > 0 {
		log.Printf("res warnings: %+v", res.Warnings)
	}
	//log.Printf("res resp: %+v", res)
	log.Printf("res data: %+v", res.Data)
}

func testKVv2() {
	k := "res/hello"
	log.Printf("key: %v", k)
	m := make(map[string]interface{})
	m["aaa"] = "bbb"
	m["num"] = 123
	res, err = client.Logical().Write(k, m)
	if err != nil {
		panic("Write error: " + err.Error())
	}

	res, err = client.Logical().Read(k)
	if err != nil {
		panic("Logical error: " + err.Error())
	}
	if res == nil {
		panic("nil result")
	}
	if len(res.Warnings) > 0 {
		log.Printf("res warnings: %+v", res.Warnings)
	}
	//log.Printf("res resp: %+v", res)
	log.Printf("res data: %+v", res.Data)
}

func generateCSR(cn string) ([]byte, error) {
	keyBytes, _ := rsa.GenerateKey(rand.Reader, 2048)
	subj := pkix.Name{
		CommonName: cn,
	}

	template := x509.CertificateRequest{
		Subject:            subj,
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, keyBytes)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	pem.Encode(&buf, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
	return buf.Bytes(), nil
}
