package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/cloudflare/cfssl/api/client"
	"github.com/cloudflare/cfssl/helpers"
	"github.com/cloudflare/cfssl/info"
	"io/ioutil"
	"log"
)

/*
// A Remote points to at least one (but possibly multiple) remote
// CFSSL instances. It must be able to perform a authenticated and
// unauthenticated certificate signing requests, return information
// about the CA on the other end, and return a list of the hosts that
// are used by the remote.

type Remote interface {
	AuthSign(req, id []byte, provider auth.Provider) ([]byte, error)
	Sign(jsonData []byte) ([]byte, error)
	Info(jsonData []byte) (*info.Resp, error)
	Hosts() []string
	SetReqModifier(func(*http.Request, []byte))
	SetRequestTimeout(d time.Duration)
	SetProxy(func(*http.Request) (*url.URL, error))
}
*/

var bundleFile = "01DJBY9CAXSYN2SR9VM2MZF2P2-bundle.txt"

func main() {

	resp, err := getRootCert()
	log.Printf("err: %+v", err)
	log.Printf("resp: %+v", resp)

	csr := generateCsr("123")
	obj := map[string]interface{}{}

	/*
		Required parameters:

		    * certificate_request: the CSR bytes to be signed in PEM

		Optional parameters:

		    * hosts: an array of SAN (subject alternative names)
		    which overrides the ones in the CSR
		    * subject: the certificate subject which overrides
		    the ones in the CSR
		    * serial_sequence: a string specify the prefix which the generated
		    certificate serial should have
		    * label: a string specifying which signer to be appointed to sign
		    the CSR, useful when interacting with a remote multi-root CA signer
		    * profile: a string specifying the signing profile for the signer,
		    useful when interacting with a remote multi-root CA signer
		    * bundle: a boolean specifying whether to include an "optimal"
		    certificate bundle along with the certificate
	*/

	//if hosts != nil {
	//	obj["hosts"] = hosts
	//}

	obj["certificate_request"] = string(csr.Bytes())

	//if subject != nil {
	//	obj["subject"] = subject
	//}

	blob, err := json.Marshal(obj)
	if err != nil {
		log.Panic(err)
	}
	//log.Printf("CSR: %v", string(blob))
	cert := sign(blob)
	log.Printf("Cert subject: %+v", cert.Subject.String())
	log.Printf("      Issuer: %+v", cert.Issuer.String())
	log.Printf("  Key Ussage: %+v", cert.KeyUsage)
	log.Printf("    NotAfter: %+v", cert.NotAfter)
	log.Printf("SerialNumber: %+v", cert.SerialNumber.String())
}

//csr := readBundle(bundleFile)
//log.Printf("%v certificates in the bundle", len(csr))
//for i, c := range csr {
//	log.Printf("cert %v: %+v ", i, c.Subject)
//}

//remote1()

func generateCsr(id string) (result bytes.Buffer) {

	// Generate an ECDSA256 keypair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	// CreateCertificateRequest creates a new certificate request based on a
	// template. The following members of template are used:
	//
	//  - SignatureAlgorithm
	//  - Subject
	//  - DNSNames
	//  - EmailAddresses
	//  - IPAddresses
	//  - URIs
	//  - ExtraExtensions
	//  - Attributes (deprecated)
	//
	// priv is the private key to sign the CSR with, and the corresponding public
	// key will be included in the CSR. It must implement crypto.Signer and its
	// Public() method must return a *rsa.PublicKey or a *ecdsa.PublicKey. (A
	// *rsa.PrivateKey or *ecdsa.PrivateKey satisfies this.)
	//
	// The returned slice is the certificate request in DER encoding.

	//  certificate request using the deviceID as the CN

	template := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:         id,
			Country:            []string{"US"},
			Organization:       []string{"Brain Corporation"},
			OrganizationalUnit: []string{"ROC"},
		},
		SignatureAlgorithm: x509.ECDSAWithSHA256,
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, privateKey)
	if err != nil {
		log.Panic(err)
	}

	err = pem.Encode(&result, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
	if err != nil {
		log.Panic(err)
	}

	return result
}

func readBundle(fileName string) (bundle []*x509.Certificate) {
	bundle = make([]*x509.Certificate, 0)

	buffer := read(fileName)
	if buffer == nil {
		return bundle
	}

	block, rest := pem.Decode(buffer)
	for block != nil {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			log.Panic(err)
		}
		bundle = append(bundle, cert)
		if rest == nil {
			return
		}
		block, rest = pem.Decode(rest)
	}
	return bundle
}

func getRootCert() (resp *info.Resp, err error) {

	s := client.NewServer("http://127.0.0.1:8888")
	r, err := s.Info([]byte(`{"label": "primary"}`))
	return r, err
}

//log.Printf("Usage: %+#v", resp.Usage)
//log.Printf("Expiry: %v", resp.ExpiryString)

func remote2() {
	s := client.NewServer("http://127.0.0.1:8888")
	_ = s
	subj := pkix.Name{
		CommonName:         "example.com",
		Country:            []string{"US"},
		Province:           []string{"CA"},
		Locality:           []string{"MyCity"},
		Organization:       []string{"Company Ltd"},
		OrganizationalUnit: []string{"IT"},
	}
	var oidEmailAddress = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}
	email := "foo@bar.com"
	rawSubj := subj.ToRDNSequence()
	rawSubj = append(rawSubj, []pkix.AttributeTypeAndValue{
		{Type: oidEmailAddress, Value: email},
	})

	//asn1Subj, _ := asn1.Marshal(rawSubj)
	//template := x509.CertificateRequest{
	//	RawSubject:         asn1Subj,
	//	EmailAddresses:     []string{email},
	//	SignatureAlgorithm: x509.SHA256WithRSA,
	//}
	//
	//{"certificate_request":
	//resp, err := s.Sign()
	//
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//log.Printf("Usage: %+#v", resp.Usage)
	//log.Printf("Expiry: %v", resp.ExpiryString)
}

func sign(csr []byte) *x509.Certificate {
	s := client.NewServer("http://127.0.0.1:8888")
	//log.Printf("Certificate: %+v", resp.Certificate)
	r, err := s.Sign(csr)
	if err != nil {
		log.Panic(err)
	}

	crt, err := helpers.ParseCertificatePEM(r)
	if err != nil {
		log.Panic(err)
	}
	//log.Printf("Certificate: Issuer: %+#v, SAN: %+#v", crt.Issuer, crt.DNSNames)
	return crt
}

// copied from internal/ca/csr.go
/////////////////////////////////
// Return a PEM-encoded csr, and private key
/////////////////////////////////
func GenerateSigningRequest() (CSR string, ecKey string, err error) {
	// Generate an ECDSA256 keypair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate keypair: %v", err)
	}

	// And a certificate request using the deviceID as the CN
	template := x509.CertificateRequest{
		SignatureAlgorithm: x509.ECDSAWithSHA256,
	}

	// Create a signing request
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, privateKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to create CSR: %v", err)
	}

	// And to make things nice to work with, just pass back base64-encoded strings
	csrBuf := bytes.NewBuffer(nil)
	err = pem.Encode(csrBuf, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
	if err != nil {
		return "", "", fmt.Errorf("failed to PEM-encode CSR: %v", err)
	}

	key, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal EC private key: %v", err)
	}

	ecBuf := bytes.NewBuffer(nil)
	err = pem.Encode(ecBuf, &pem.Block{Type: "EC PRIVATE KEY", Bytes: key})
	if err != nil {
		return "", "", fmt.Errorf("failed to PEM-encode EC private key: %v", err)
	}

	return csrBuf.String(), ecBuf.String(), nil
}

func read(path string) []byte {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	return d
}

/*

	from ca.go

	// Load the provisioning cert
	caCert, err := x509.ParseCertificate(provCert.Certificate[0])
	if err != nil {
		return "", err
	}

	notBefore := time.Now()
	notAfter := notBefore.AddDate(5, 0, 0) // 5 years

	template := &x509.Certificate{
		SerialNumber: big.NewInt(0),
		Subject: pkix.Name{
			CommonName:         id,
			Country:            []string{"US"},
			Organization:       []string{"Brain Corporation"},
			OrganizationalUnit: []string{"ROC"},
		},
		PublicKey:          csrv.PublicKey,
		SignatureAlgorithm: csrv.SignatureAlgorithm,
		NotBefore:          notBefore,
		NotAfter:           notAfter,
		KeyUsage:           x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		IsCA:               isCa,
	}

	if isCa {
		template.KeyUsage = template.KeyUsage | x509.KeyUsageCertSign
		template.DNSNames = []string{caDnsName}
	}

	// Create a cert
	derBytes, err := x509.CreateCertificate(rand.Reader, template, caCert, csrv.PublicKey, provCert.PrivateKey)
	if err != nil {
		return "", err
	}



api/provisioning.go

package v0

type Provisioning struct {
	CSR  string `json:"csr,omitempty"`
	Cert string `json:"cert,omitempty"`
}


cmd/rooc/devices.go

func (d *Devices) Provision(id string) (*v0.Device, string, error) {
	// Generate a CSR for provisioning
	csr, key, err := ca.GenerateSigningRequest()
	if err != nil {
		return nil, "", err
	}

	resource, err := d.collection.Update(id, &v0.Device{
		CSR: csr,
	})
	if err != nil {
		return nil, "", err
	}

	return asDevice(resource), key, nil
}


int/services/provisions/handler.go

// Create handles requests:
// - POST /provisions[/]
func (s *Service) Create(ctx context.Context, w resource.ResponseWriter, rq *Request) {

	// Generate a new provision Id
	rq.Id = resource.NewID(nil)


	// Check to see if a principal exists in the context. Typically the call
	// to the provisions api is made via a local unauthenticated client. In this case
	// the device hosting the api should set the device as the principal
	if auth.GetPrincipal(ctx) == nil {
		ctx = auth.WithPrincipal(ctx, s.principal)
	}

	// TODO : Move client0 to client1 and move /v0/provisioning in boxd.go to /v1/provisioning (ROC-669)
	// Before starting first connect to the rocd url and see if we can get a CSR before proceeding
	rocdCtx := client0.ClientContext{
		Url:     &rq.RocdUrl,
		TLSCert: nil,
		Token:   nil, // This might need to be an actual token issued by rocd
	}

	rocdClient, err := client0.New(rocdCtx)
	if err != nil {
		w.WriteError(fmt.Errorf("Failed to create rocd client: %s", err.Error()))
		return
	}

	// Get the provisioning data (CSR) from the device
	pData, err := rocdClient.Provisioning().Get()
	if err != nil {
		w.WriteError(fmt.Errorf("Failed to get provisioning data: %s", err.Error()))
		return
	}

*******
internal/services/devices/handler.go:handleCsr()
internal/services/provisions/handler.go:Create()
*******


	// Check to see if a principal exists in the context. Typically the call
	// to the provisions api is made via a local unauthenticated client. In this case
	// the device hosting the api should set the device as the principal
	if auth.GetPrincipal(ctx) == nil {
		ctx = auth.WithPrincipal(ctx, s.principal)
	}

	// TODO : Move client0 to client1 and move /v0/provisioning in boxd.go to /v1/provisioning (ROC-669)
	// Before starting first connect to the rocd url and see if we can get a CSR before proceeding
	rocdCtx := client0.ClientContext{
		Url:     &rq.RocdUrl,
		TLSCert: nil,
		Token:   nil, // This might need to be an actual token issued by rocd
	}

	rocdClient, err := client0.New(rocdCtx)
	if err != nil {
		w.WriteError(fmt.Errorf("Failed to create rocd client: %s", err.Error()))
		return
	}

	// Get the provisioning data (CSR) from the device
	pData, err := rocdClient.Provisioning().Get()
	if err != nil {
		w.WriteError(fmt.Errorf("Failed to get provisioning data: %s", err.Error()))
		return
	}

...

// Create Local Device table entry
	devRq := devices.Request{
		Id:           rq.DeviceId, // If device Id is not set, the device svc will create one
		Name:         rq.DeviceName,
		SerialNumber: rq.SerialNumber,
		DeviceType:   rq.DeviceType,
		Csr:          pData.CSR,
		IsCa:         false,
	}

https://ssltools.digicert.com/checker/views/csrCheck.jsp

openssl req -in mycsr.csr -noout -text

-----BEGIN CERTIFICATE REQUEST-----
MIG5MGICAQAwADBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABL298X5Q0WaQt4oT
l+MbFbSlkt2DLf91mmlxm4/je3ifB3koc1bvLoZJiPPMRvWjLE7rKJzmUifRHVvm
bwOGw2WgADAKBggqhkjOPQQDAgNHADBEAiB9jxnHV/fsiMzCpYoS5bBcbkUkRRK1
Q/2VIlWCc60ewAIgNhquG7lYbQeisqpK3W+9ml017h0wVTT5C50Ww3VbwpM=
-----END CERTIFICATE REQUEST-----

-----BEGIN CERTIFICATE-----
MIIBsDCCAVWgAwIBAgIBADAKBggqhkjOPQQDAjBIMQswCQYDVQQGEwJVUzEaMBgG
A1UEChMRQnJhaW4gQ29ycG9yYXRpb24xDDAKBgNVBAsTA1JPQzEPMA0GA1UEAxMG
Uk9DIENBMB4XDTE5MDgxOTIzNTkwNVoXDTI0MDgxOTIzNTkwNVowZjELMAkGA1UE
BhMCVVMxGjAYBgNVBAoTEUJyYWluIENvcnBvcmF0aW9uMQwwCgYDVQQLEwNST0Mx
LTArBgNVBAMTJGU2NGNkY2QxLTgxMWQtNDEwMS1hZmEyLTc0NWQyMGFkNTc3MDBZ
MBMGByqGSM49AgEGCCqGSM49AwEHA0IABL298X5Q0WaQt4oTl+MbFbSlkt2DLf91
mmlxm4/je3ifB3koc1bvLoZJiPPMRvWjLE7rKJzmUifRHVvmbwOGw2WjEjAQMA4G
A1UdDwEB/wQEAwIFoDAKBggqhkjOPQQDAgNJADBGAiEAmqhb9wzRKNqJ2o583VDi
3Y4IR/nMdbAylR6VVfpmkiYCIQDcI6hpPbW+4+PzC5ex9YWuNlSlsiarxCxEDXhO
QISfvA==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIBxDCCAWqgAwIBAgIBADAKBggqhkjOPQQDAjBIMQswCQYDVQQGEwJVUzEaMBgG
A1UEChMRQnJhaW4gQ29ycG9yYXRpb24xDDAKBgNVBAsTA1JPQzEPMA0GA1UEAxMG
Uk9DIENBMB4XDTE3MDIwMjIxMjczOFoXDTIyMDIwMjIxMjczOFowSDELMAkGA1UE
BhMCVVMxGjAYBgNVBAoTEUJyYWluIENvcnBvcmF0aW9uMQwwCgYDVQQLEwNST0Mx
DzANBgNVBAMTBlJPQyBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC5dIk3H
6bD7lfsQ0MOG2W72/4hhlm+VWheQrYsscV2Hy4hAXOIM7INe1r+fNg32UNnh/QsT
ly9WwcgCB847hEujRTBDMA4GA1UdDwEB/wQEAwICpDAPBgNVHRMBAf8EBTADAQH/
MCAGA1UdEQQZMBeCFWh0dHA6Ly9sb2NhbGhvc3Q6ODA4MDAKBggqhkjOPQQDAgNI
ADBFAiEAw1WyiS8hssnced9y9ePwedqONEL7gcwgaIENSchMo5ICIBf9c60yk6WL
ZiVFb3sbgnu49vcPw+c6+Iycu+zcJwRE
-----END CERTIFICATE-----



-----BEGIN CERTIFICATE-----
MIIBsDCCAVWgAwIBAgIBADAKBggqhkjOPQQDAjBIMQswCQYDVQQGEwJVUzEaMBgG
A1UEChMRQnJhaW4gQ29ycG9yYXRpb24xDDAKBgNVBAsTA1JPQzEPMA0GA1UEAxMG
Uk9DIENBMB4XDTE5MDgxOTIzNTkwNVoXDTI0MDgxOTIzNTkwNVowZjELMAkGA1UE
BhMCVVMxGjAYBgNVBAoTEUJyYWluIENvcnBvcmF0aW9uMQwwCgYDVQQLEwNST0Mx
LTArBgNVBAMTJGU2NGNkY2QxLTgxMWQtNDEwMS1hZmEyLTc0NWQyMGFkNTc3MDBZ
MBMGByqGSM49AgEGCCqGSM49AwEHA0IABL298X5Q0WaQt4oTl+MbFbSlkt2DLf91
mmlxm4/je3ifB3koc1bvLoZJiPPMRvWjLE7rKJzmUifRHVvmbwOGw2WjEjAQMA4G
A1UdDwEB/wQEAwIFoDAKBggqhkjOPQQDAgNJADBGAiEAmqhb9wzRKNqJ2o583VDi
3Y4IR/nMdbAylR6VVfpmkiYCIQDcI6hpPbW+4+PzC5ex9YWuNlSlsiarxCxEDXhO
QISfvA==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIBxDCCAWqgAwIBAgIBADAKBggqhkjOPQQDAjBIMQswCQYDVQQGEwJVUzEaMBgG
A1UEChMRQnJhaW4gQ29ycG9yYXRpb24xDDAKBgNVBAsTA1JPQzEPMA0GA1UEAxMG
Uk9DIENBMB4XDTE3MDIwMjIxMjczOFoXDTIyMDIwMjIxMjczOFowSDELMAkGA1UE
BhMCVVMxGjAYBgNVBAoTEUJyYWluIENvcnBvcmF0aW9uMQwwCgYDVQQLEwNST0Mx
DzANBgNVBAMTBlJPQyBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC5dIk3H
6bD7lfsQ0MOG2W72/4hhlm+VWheQrYsscV2Hy4hAXOIM7INe1r+fNg32UNnh/QsT
ly9WwcgCB847hEujRTBDMA4GA1UdDwEB/wQEAwICpDAPBgNVHRMBAf8EBTADAQH/
MCAGA1UdEQQZMBeCFWh0dHA6Ly9sb2NhbGhvc3Q6ODA4MDAKBggqhkjOPQQDAgNI
ADBFAiEAw1WyiS8hssnced9y9ePwedqONEL7gcwgaIENSchMo5ICIBf9c60yk6WL
ZiVFb3sbgnu49vcPw+c6+Iycu+zcJwRE
-----END CERTIFICATE-----


provisioning cert from rocd provisioning.go

cert$ certdump cert497975968
--cert497975968 ---
CERTIFICATE
Subject: /123/C=US/O=Brain Corporation/OU=ROC
Issuer: /ROC CA/C=US/O=Brain Corporation/OU=ROC
	Signature algorithm: ECDSA / SHA256
Details:
	Public key: ECDSA-prime256v1
	Serial number: 0
	Valid from: 2019-08-20T00:56:32+0000
	     until: 2024-08-20T00:56:32+0000
	Key usages: digital signature, key encipherment
	Basic constraints: invalid
	SANs (0):

cert$ cat /var/folders/rh/qc1lvl3j5w30mk1tqs6p7f9r0000gp/T/cert497975968
-----BEGIN CERTIFICATE-----
MIIBjjCCATSgAwIBAgIBADAKBggqhkjOPQQDAjBIMQswCQYDVQQGEwJVUzEaMBgG
A1UEChMRQnJhaW4gQ29ycG9yYXRpb24xDDAKBgNVBAsTA1JPQzEPMA0GA1UEAxMG
Uk9DIENBMB4XDTE5MDgyMDAwNTYzMloXDTI0MDgyMDAwNTYzMlowRTELMAkGA1UE
BhMCVVMxGjAYBgNVBAoTEUJyYWluIENvcnBvcmF0aW9uMQwwCgYDVQQLEwNST0Mx
DDAKBgNVBAMTAzEyMzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPdHruDaIEvY
ioHu/fyIZ+iPdog8gGPqOyovJ/g71FIe7BNta6GrD6DY4D3HF6nLOqw9UvFch2ej
YAVxE1P8+FejEjAQMA4GA1UdDwEB/wQEAwIFoDAKBggqhkjOPQQDAgNIADBFAiAv
DcALdu/U0MvGNO8mSja4IqdAEWhMaMguEm/s4k8zrQIhAK+4XNX9OLjaLnWjr3Ij
CI1uTuvx/qnPlfsZcManuGE+
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIBxDCCAWqgAwIBAgIBADAKBggqhkjOPQQDAjBIMQswCQYDVQQGEwJVUzEaMBgG
A1UEChMRQnJhaW4gQ29ycG9yYXRpb24xDDAKBgNVBAsTA1JPQzEPMA0GA1UEAxMG
Uk9DIENBMB4XDTE3MDIwMjIxMjczOFoXDTIyMDIwMjIxMjczOFowSDELMAkGA1UE
BhMCVVMxGjAYBgNVBAoTEUJyYWluIENvcnBvcmF0aW9uMQwwCgYDVQQLEwNST0Mx
DzANBgNVBAMTBlJPQyBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC5dIk3H
6bD7lfsQ0MOG2W72/4hhlm+VWheQrYsscV2Hy4hAXOIM7INe1r+fNg32UNnh/QsT
ly9WwcgCB847hEujRTBDMA4GA1UdDwEB/wQEAwICpDAPBgNVHRMBAf8EBTADAQH/
MCAGA1UdEQQZMBeCFWh0dHA6Ly9sb2NhbGhvc3Q6ODA4MDAKBggqhkjOPQQDAgNI
ADBFAiEAw1WyiS8hssnced9y9ePwedqONEL7gcwgaIENSchMo5ICIBf9c60yk6WL
ZiVFb3sbgnu49vcPw+c6+Iycu+zcJwRE
-----END CERTIFICATE-----
******
*/

/*
	c := read("./crt/ca.crt")
	//log.Printf("Device cert: %+v", cert.Leaf)
	crt, err := helpers.ParseCertificatePEM(c)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Certificate: Issuer: %+#v, SAN: %+#v", crt.Issuer, crt.DNSNames)
*/
//d, err := helpers.ReadBytes("file:cert/csr-ca.json")
//if err != nil {
//	log.Panic(err)
//}
