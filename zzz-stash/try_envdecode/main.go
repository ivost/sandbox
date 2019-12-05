package main

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"net/url"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joeshaw/envdecode"
)

// https://github.com/joeshaw/envdecode

type Config struct {
	Hostname string   `env:"SERVER_HOSTNAME,default=localhost"`
	Port     uint16   `env:"SERVER_PORT,default=8080"`
	URL      *url.URL `env:"URL,required"`
	AWS      struct {
		ID        string   `env:"AWS_ACCESS_KEY_ID"`
		Secret    string   `env:"AWS_SECRET_ACCESS_KEY,required"`
		SnsTopics []string `env:"AWS_SNS_TOPICS"`
	}

	Timeout time.Duration `env:"TIMEOUT,default=1m,strict"`
}

type MyClaims struct {
	jwt.StandardClaims
	Scopes []string `json:"scopes,omitempty"`
}

var CaCert = `-----BEGIN CERTIFICATE-----
MIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTE5MDgwMTAyMzIzNVoXDTI5MDcyOTAyMzIzNVowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKZj
5rMRc+loKOT13yP3gH8367WXBdFeA1Szm+uXai22Am5TUCZeMy33xMMfI3nJCNE+
JhFF42P8iFxybyZl4IBhApfRkCkh6/dxQqR0WuO7J8luJjtkc1D9J4sCB0vj/uDe
ybh/62V+jmRCiWqwOo+0puVZZZN7ZvsSCC6pzPtPlfZnxx3dNvbsmrTvMTuueO2l
kf11+MGK1Hs82kfa1Jl4BjMKZydbiJBFbmCuCwvhnrYtuLEjidYk9RlUVklAiKxF
8WobxLEoJU/JN/+YQoRY9MOhk302IpKpj1a9T3F5VRfQ2yYjjICZb+R72oV6KNGD
7wADYa8LHl28u/mZ7hECAwEAAaMjMCEwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAHVaA3iMz10vkpkxRe6/6X9q0Wul
5w3YP0V0Hac+/0m+CQSee9uhBilX/y31kP+emXzDEOVuL40DTlGpzFqpp/GrcHNz
b9Kx+/roSSNNA7zrauvdW0R+j4TljkgNVZVxld3mfW4FBHEhj37ZhPTDPy8y88Py
24zg3dO5AC23ckZcdhNjhb02cpkmieu3iG++iDlHTQOXwTo9+IDtyu0o97HQ3RZz
Yo7dMdLlQPzSOa7qSjWUpH8W5a10Ie2RBKiUOWTaruwap+Y5ABmJIjQBHAy2+S7P
O9AIx1UVrQ2THbuN2bRcydVFWhCXhX8W7tnoV9AkJlvNrjqCQNtfQ5REN84=
-----END CERTIFICATE-----`

var token = `eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InZhdWx0LXRva2VuLWdmeDdxIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiMGNmMzhiYmMtMTNmZC0xMWVhLWIzOTctMDI1MDAwMDAwMDAxIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6dmF1bHQifQ.bq_OPXX9UWVGs-Yc4sLSOVncAheXavI1hmb9n5Dy4PnRln8zlsbcNpxe_0FEf7kqpV43S0yktBg5zpdYi1W4SQq0cMUuLxMJ-QJ4sl--yJnGsGSwRSH27_9jQ8N1muUL7_Yf62JEXWLYKlAacA0_HvV8DNKsoSN6hGqXwu2khubZjXShgmIzXC7lgwdRYIAcohJMzB6Nt6rF8yLo7yFdxaY5IxkI_eTT_qP8ChIcGzp8_206Ov9vWqJixB6lK0-_eztFWlaR0n_C3re-mxGuSC0CyIQZqnd74FJ7LXFrFLYKpA0JcIlUvwZ0aUB4QxJb7OQRMV2bbTpKH_5-rmfgcw`

func ParseCert(certPEMBytes []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(certPEMBytes)
	if block == nil {
		return nil, errors.New("invalid PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func VerifyToken(jwtString string) (*jwt.Token, error) {
	crt, err := ParseCert([]byte(CaCert))
	if err != nil {
		return nil, err
	}
	tok, err := jwt.Parse(jwtString, func(tok *jwt.Token) (interface{}, error) {
		//key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(CaCert))
		//return key, err
		return crt.PublicKey, nil
	})
	return tok, err
}

func main() {

	var cfg Config
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "foo")
	_ = os.Setenv("TIMEOUT", "42s")
	_ = os.Setenv("URL", "http://example.com")
	_ = os.Setenv("AWS_SNS_TOPICS", "foo,bar")
	//os.Setenv("TIMEOUT", "foo") // error
	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	//log.Printf("Config:%+v", cfg)

	tok, err := VerifyToken(token)
	log.Printf("token:%+v, err: %v", tok, err)
}
