
VAULT / PKI Readme

https://learn.hashicorp.com/vault/getting-started/apis

https://medium.com/@sufiyanghori/guide-using-hashicorp-vault-to-manage-pki-and-issue-certificates-e1981e7574e



vault server -config=config.hcl

init

http :8200/v1/sys/init secret_shares:=1 secret_threshold:=1

{
    "keys": [
        "d9db49571840533e3ad3711f766fa3da8b733252ff7ef9269b14657ce0adcbb9"
    ],
    "keys_base64": [
        "2dtJVxhAUz4603Efdm+j2otzMlL/fvkmmxRlfOCty7k="
    ],
    "root_token": "s.cSrLCBljbrIV3R0YABfhSh0A"
}

unseal

http :8200/v1/sys/unseal key=2dtJVxhAUz4603Efdm+j2otzMlL/fvkmmxRlfOCty7k=

enable AppRole

http :8200/v1/sys/auth/approle X-Vault-Token:$VAULT_TOKEN type=approle

export VAULT_TOKEN=s.cSrLCBljbrIV3R0YABfhSh0A

vault login

vault secrets list

vault secrets enable -path=rootca_store -description="PKI backend for Root CA" -max-lease-ttl=87600h pki


vault write rootca_store/root/generate/internal \
common_name="suf.com" \
ttl=87600h \
key_bits=4096


Key              Value
---              -----
certificate      -----BEGIN CERTIFICATE-----
MIIFKTCCAxGgAwIBAgIUFn/1LmeWX6noBYqbs3qW18k3vlIwDQYJKoZIhvcNAQEL
BQAwEjEQMA4GA1UEAxMHc3VmLmNvbTAeFw0xOTA3MjYwMjEyMDNaFw0yOTA3MjMw
MjEyMzJaMBIxEDAOBgNVBAMTB3N1Zi5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4IC
DwAwggIKAoICAQCrxR6JWfsGvsNLxlDfwu/9xutcNcYxyD52FE8ndCCFHdgDJdVS
xoUlcua7ngY2xkJHcmreGZ/estKEEx2Thc666ilvntoFohgftFbAnG5CLK7JtYt1
IfPciTDJ9TcbCtRSv9fe9oJCAGCjhqufv/KLX9+T/JkUDAUN5CxhbMwaTG6it03Q
gd0qMgCk7bWXRpJD98q74F1UAClm9jt6H2ipp8NVk/p4SP8XWX8r0sDVVLY7Tarq
gJ0jwVPpmVtdjgKkCZ0SxTmS8mSXF5LG2plEE56kCiEasrEyqX55K2JLFMC+SCsp
t2X7/3uHqPFT68Pe2oSKxvG0E9SQNkn4zt5EhIH+SVyV6l11tI6DWMKyySrv99WL
n4iw0VriZdOYRLmSAWma7hVROll9wt5IPKayJZcxBO+ox6W+SKlWJTTGXVUL39sS
H1XdXiPp1rm4vTc9AQvpNqVD/+oKj7Ct2CO7lQJ7eHYg5qxr8Q4gPFgzeGjXPAX0
U/tnSwJfq2MrsXu1B/MHLtJXjvz13IyFdPKfH5TddqW9R4hZVZbtgSUUgqa7AWvV
qlIUzgQmg9zn7kNbINrko6IFC06jFY0QYkhtK6WiUFZXP+4IaEJaKwo2Mk25fCJA
iSUgDMD2CEXNQDOgN2yJQ41WFoEQifoiUqizzkxEVc2WbTA7bDboBCuyWwIDAQAB
o3cwdTAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQU
rUb2eIeiPfyyfVj3txCKEDc7NEwwHwYDVR0jBBgwFoAUrUb2eIeiPfyyfVj3txCK
EDc7NEwwEgYDVR0RBAswCYIHc3VmLmNvbTANBgkqhkiG9w0BAQsFAAOCAgEAAL3O
qZRDp5hTBzQfzc52lxlfNKZdP1/xvKsehs3cO3VrhesNjuIiSEk/BUZ+KdST5S6Y
oqLodP4dxYw2XzVLPcr1ciKqUlv5HjwdgBkL5dtM5JiI9vqcFZaxjTru9AkKkh0B
EndDeaIWZii4USiz2Asldq/3WcuW6a8eS/BIDQyJBELiRDqMDgORoZeFwsoqC525
fUFwpZwxE9+iIZxObNd3DlCya6SENoFkFlEwKuue1LDXhAEyzV46rQjuQ5lac7+r
ty6UNlJ1nDppJX8EmTHOv8INm8F9EPynPL4WlqjeEfl3zH7ttgGujqd4gMF/GA0v
HQQO8huttIXk4sQJ+qK5ZzW0JAFOiOPdLowHlO82cG8OvmW9NAGtJtpOHIPBhQ2N
K2NFmsfBFG1LU1qBct09bgSOpHXHG0AyXavDsCeDDzxDLy2TX6xRNz/gp0TjmPDo
CA3MnXBY42lXNaH+E0HHn+kTAuFJS1NSsdcLE3VcksD9ro/7O0CXZq113DzUcxFR
ZEYcUPj53G61h9B5+FM7h3XnQVHoEaRvlOj71mnPY25hNuJG8gIYKRya9OXG8plB
9NiL76S0K5O70yUykwBL2MofCcR89LY/kPK4QhCRYIHL6Q+dzyYkXs/fQKYyfmjE
FMdYzxlw4UEdU8H1+kGf/Xhls9og1EJAWN2uo9s=
-----END CERTIFICATE-----
expiration       1879467152
issuing_ca       -----BEGIN CERTIFICATE-----
MIIFKTCCAxGgAwIBAgIUFn/1LmeWX6noBYqbs3qW18k3vlIwDQYJKoZIhvcNAQEL
BQAwEjEQMA4GA1UEAxMHc3VmLmNvbTAeFw0xOTA3MjYwMjEyMDNaFw0yOTA3MjMw
MjEyMzJaMBIxEDAOBgNVBAMTB3N1Zi5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4IC
DwAwggIKAoICAQCrxR6JWfsGvsNLxlDfwu/9xutcNcYxyD52FE8ndCCFHdgDJdVS
xoUlcua7ngY2xkJHcmreGZ/estKEEx2Thc666ilvntoFohgftFbAnG5CLK7JtYt1
IfPciTDJ9TcbCtRSv9fe9oJCAGCjhqufv/KLX9+T/JkUDAUN5CxhbMwaTG6it03Q
gd0qMgCk7bWXRpJD98q74F1UAClm9jt6H2ipp8NVk/p4SP8XWX8r0sDVVLY7Tarq
gJ0jwVPpmVtdjgKkCZ0SxTmS8mSXF5LG2plEE56kCiEasrEyqX55K2JLFMC+SCsp
t2X7/3uHqPFT68Pe2oSKxvG0E9SQNkn4zt5EhIH+SVyV6l11tI6DWMKyySrv99WL
n4iw0VriZdOYRLmSAWma7hVROll9wt5IPKayJZcxBO+ox6W+SKlWJTTGXVUL39sS
H1XdXiPp1rm4vTc9AQvpNqVD/+oKj7Ct2CO7lQJ7eHYg5qxr8Q4gPFgzeGjXPAX0
U/tnSwJfq2MrsXu1B/MHLtJXjvz13IyFdPKfH5TddqW9R4hZVZbtgSUUgqa7AWvV
qlIUzgQmg9zn7kNbINrko6IFC06jFY0QYkhtK6WiUFZXP+4IaEJaKwo2Mk25fCJA
iSUgDMD2CEXNQDOgN2yJQ41WFoEQifoiUqizzkxEVc2WbTA7bDboBCuyWwIDAQAB
o3cwdTAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQU
rUb2eIeiPfyyfVj3txCKEDc7NEwwHwYDVR0jBBgwFoAUrUb2eIeiPfyyfVj3txCK
EDc7NEwwEgYDVR0RBAswCYIHc3VmLmNvbTANBgkqhkiG9w0BAQsFAAOCAgEAAL3O
qZRDp5hTBzQfzc52lxlfNKZdP1/xvKsehs3cO3VrhesNjuIiSEk/BUZ+KdST5S6Y
oqLodP4dxYw2XzVLPcr1ciKqUlv5HjwdgBkL5dtM5JiI9vqcFZaxjTru9AkKkh0B
EndDeaIWZii4USiz2Asldq/3WcuW6a8eS/BIDQyJBELiRDqMDgORoZeFwsoqC525
fUFwpZwxE9+iIZxObNd3DlCya6SENoFkFlEwKuue1LDXhAEyzV46rQjuQ5lac7+r
ty6UNlJ1nDppJX8EmTHOv8INm8F9EPynPL4WlqjeEfl3zH7ttgGujqd4gMF/GA0v
HQQO8huttIXk4sQJ+qK5ZzW0JAFOiOPdLowHlO82cG8OvmW9NAGtJtpOHIPBhQ2N
K2NFmsfBFG1LU1qBct09bgSOpHXHG0AyXavDsCeDDzxDLy2TX6xRNz/gp0TjmPDo
CA3MnXBY42lXNaH+E0HHn+kTAuFJS1NSsdcLE3VcksD9ro/7O0CXZq113DzUcxFR
ZEYcUPj53G61h9B5+FM7h3XnQVHoEaRvlOj71mnPY25hNuJG8gIYKRya9OXG8plB
9NiL76S0K5O70yUykwBL2MofCcR89LY/kPK4QhCRYIHL6Q+dzyYkXs/fQKYyfmjE
FMdYzxlw4UEdU8H1+kGf/Xhls9og1EJAWN2uo9s=
-----END CERTIFICATE-----
serial_number    16:7f:f5:2e:67:96:5f:a9:e8:05:8a:9b:b3:7a:96:d7:c9:37:be:52


curl -s http://127.0.0.1:8200/v1/rootca_store/ca/pem

vault write rootca_store/config/urls issuing_certificates="http://127.0.0.1:8200/v1/rootca_store/ca" \
crl_distribution_points="http://127.0.0.1:8200/v1/rootca_store/crl"


vault secrets enable -path=interca -description="PKI backend for Intermediate CA" \
-max-lease-ttl=87600h pki


GENERATE SUB-CA

vault write interca/intermediate/generate/internal ttl=26280h key_bits=4096 > signing_request.csr

*****

ROOT SIGNS AND ISSUES SUB-CA CERT

vault write rootca_store/root/sign-intermediate csr=@signing_request.csr ttl=8760h format=pem_bundle


Key              Value
---              -----
certificate      -----BEGIN CERTIFICATE-----
MIIFhzCCA2+gAwIBAgIUfeB4HkgBRuy3BLO6SA8OhlxOXw0wDQYJKoZIhvcNAQEL
BQAwEjEQMA4GA1UEAxMHc3VmLmNvbTAeFw0xOTA3MjYwMjI0NTFaFw0yMDA3MjUw
MjI1MjFaMAAwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDeD5Nv5l3c
umshJskmORpSn2P1VCpYpYlTSVfXA3iUvUYQn185WyCm5tgVjbKcmRzlIbxcvxaN
juTtSB4fTxvwl4cJrISGF1huENhXxNg++rHtFsJ5n53/ZvJElnHre9u6pm2pZENg
0e/0YbYELYAhnX0kvCmV2MHMwp8KYTs7OSwz/sGFLa7WSuxgAMW0CIgyZoxBc894
LAhaAzDzJ0mO/4pepSjvQ3cYLQycinEEWOr+I2ofkvehznvv7s6UwEtkR5dzTLtc
iYJyC+2R05vDLBkd6fU3oHmKTvlreRa6FpubFIP8otAAqPLVKc9YnMwrFRGIuqFm
3KuM648uW0wQ3I87tOHoLKpNIBLUpVHGPVZ9ugCIALTtKObILDOKZIrdCDg0xgIU
YxO60jHUM3ifHaG9n+FeSk6hApElRxadiWsopdwHQtRhaQ0ib8RHh6OL3j30B1gX
1jjKnvFd3Zioqw/OrbXKIzT8LRJ1oxvcsS9Unw47UQfQlGqGaU0LqwZ1+i1tx/mN
vqSBpzo8B+JUPWPW0cY7OZop5iS8ai2mvZJ7FlTkZ9m0VIK0N6Cf2ZLPh4nEapSS
koFj+iuz4gK++iTVEK5WkA4dg9QOAoklkX7OzLfzTUEtAd3adTgxH/hL3SX8iM8O
nFOB/v2oCbqur5UyQ422wDsK0DzgPB+rcQIDAQABo4HmMIHjMA4GA1UdDwEB/wQE
AwIBBjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTpyFccAt9+dDe6N+CzRxce
2RRr7zAfBgNVHSMEGDAWgBStRvZ4h6I9/LJ9WPe3EIoQNzs0TDBEBggrBgEFBQcB
AQQ4MDYwNAYIKwYBBQUHMAKGKGh0dHA6Ly8xMjcuMC4wLjE6ODIwMC92MS9yb290
Y2Ffc3RvcmUvY2EwOgYDVR0fBDMwMTAvoC2gK4YpaHR0cDovLzEyNy4wLjAuMTo4
MjAwL3YxL3Jvb3RjYV9zdG9yZS9jcmwwDQYJKoZIhvcNAQELBQADggIBAIeXDr7N
qAUPOhFhWx7Pr0apZtZ/SQlD2pQgPjcyxCQKh9reaZaBXyeo05g3XOxeNY2MOsF/
2JsDNxMKSyXbkkZXFjLPyMhWClMSxasw/Ni7u6y8dpQG7rhmCF0jN2w/4+HJzLzk
dqA3Ymf61gatt2/WNQKRdoAW2u9/KVDMVxzZruO+8+9FitEBlEQbMQtWuO7esMhR
MZMB5yjPZiW4Hil7VpVBFyRNksOI5+VXxfSqhxLjid9l8bflav79J8Qm8YuU+YiZ
yYtSU3BLiDPVH3PhGRZZpzp9BytHt6kanUmGDQGWisP80XPW54ABba3U33AmEZ5Y
OT2nz0wMQ1+ZZCLguaneroZEbiqCMldJGBa+p9JLQbnH5HC+mzNId+oRVRtWuTDA
3n3E36f5454E/qxwgTiGo7V2OWRqGJnXUmTzF9BqnN24Owncxf9lFdQ9XmAKofq/
wkwKBt4Zi+CLHHXuu7S/cq4KzMrtqrQSTvuOgr2qpq5BSaWZHPsdsaTrAsHKpNib
2tOLUahyR7Zd0P3WD7+yc7Z8eG4V8XOJPCDbOJ6ShrogVGXrD+TGq+5LluPMA2R8
T7Arn9PlFK9eAxPYJ9G8Big3NY2g1H62HkHdnfPdU882wjCXqLbjTEOOJsGK9qQG
2FpsguFdzN5t4DfsGVDlcqmazbj+zIlP7SRx
-----END CERTIFICATE-----
expiration       1595643921
issuing_ca       -----BEGIN CERTIFICATE-----
MIIFKTCCAxGgAwIBAgIUFn/1LmeWX6noBYqbs3qW18k3vlIwDQYJKoZIhvcNAQEL
BQAwEjEQMA4GA1UEAxMHc3VmLmNvbTAeFw0xOTA3MjYwMjEyMDNaFw0yOTA3MjMw
MjEyMzJaMBIxEDAOBgNVBAMTB3N1Zi5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4IC
DwAwggIKAoICAQCrxR6JWfsGvsNLxlDfwu/9xutcNcYxyD52FE8ndCCFHdgDJdVS
xoUlcua7ngY2xkJHcmreGZ/estKEEx2Thc666ilvntoFohgftFbAnG5CLK7JtYt1
IfPciTDJ9TcbCtRSv9fe9oJCAGCjhqufv/KLX9+T/JkUDAUN5CxhbMwaTG6it03Q
gd0qMgCk7bWXRpJD98q74F1UAClm9jt6H2ipp8NVk/p4SP8XWX8r0sDVVLY7Tarq
gJ0jwVPpmVtdjgKkCZ0SxTmS8mSXF5LG2plEE56kCiEasrEyqX55K2JLFMC+SCsp
t2X7/3uHqPFT68Pe2oSKxvG0E9SQNkn4zt5EhIH+SVyV6l11tI6DWMKyySrv99WL
n4iw0VriZdOYRLmSAWma7hVROll9wt5IPKayJZcxBO+ox6W+SKlWJTTGXVUL39sS
H1XdXiPp1rm4vTc9AQvpNqVD/+oKj7Ct2CO7lQJ7eHYg5qxr8Q4gPFgzeGjXPAX0
U/tnSwJfq2MrsXu1B/MHLtJXjvz13IyFdPKfH5TddqW9R4hZVZbtgSUUgqa7AWvV
qlIUzgQmg9zn7kNbINrko6IFC06jFY0QYkhtK6WiUFZXP+4IaEJaKwo2Mk25fCJA
iSUgDMD2CEXNQDOgN2yJQ41WFoEQifoiUqizzkxEVc2WbTA7bDboBCuyWwIDAQAB
o3cwdTAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQU
rUb2eIeiPfyyfVj3txCKEDc7NEwwHwYDVR0jBBgwFoAUrUb2eIeiPfyyfVj3txCK
EDc7NEwwEgYDVR0RBAswCYIHc3VmLmNvbTANBgkqhkiG9w0BAQsFAAOCAgEAAL3O
qZRDp5hTBzQfzc52lxlfNKZdP1/xvKsehs3cO3VrhesNjuIiSEk/BUZ+KdST5S6Y
oqLodP4dxYw2XzVLPcr1ciKqUlv5HjwdgBkL5dtM5JiI9vqcFZaxjTru9AkKkh0B
EndDeaIWZii4USiz2Asldq/3WcuW6a8eS/BIDQyJBELiRDqMDgORoZeFwsoqC525
fUFwpZwxE9+iIZxObNd3DlCya6SENoFkFlEwKuue1LDXhAEyzV46rQjuQ5lac7+r
ty6UNlJ1nDppJX8EmTHOv8INm8F9EPynPL4WlqjeEfl3zH7ttgGujqd4gMF/GA0v
HQQO8huttIXk4sQJ+qK5ZzW0JAFOiOPdLowHlO82cG8OvmW9NAGtJtpOHIPBhQ2N
K2NFmsfBFG1LU1qBct09bgSOpHXHG0AyXavDsCeDDzxDLy2TX6xRNz/gp0TjmPDo
CA3MnXBY42lXNaH+E0HHn+kTAuFJS1NSsdcLE3VcksD9ro/7O0CXZq113DzUcxFR
ZEYcUPj53G61h9B5+FM7h3XnQVHoEaRvlOj71mnPY25hNuJG8gIYKRya9OXG8plB
9NiL76S0K5O70yUykwBL2MofCcR89LY/kPK4QhCRYIHL6Q+dzyYkXs/fQKYyfmjE
FMdYzxlw4UEdU8H1+kGf/Xhls9og1EJAWN2uo9s=
-----END CERTIFICATE-----
serial_number    7d:e0:78:1e:48:01:46:ec:b7:04:b3:ba:48:0f:0e:86:5c:4e:5f:0d


Now that we have a Root CA signed certificate, we need to import it into our Intermediate CA backend.

vault write interca/intermediate/set-signed certificate=@SUB-CA-BUNDLE.PEM

VERIFY

curl -s http://127.0.0.1:8200/v1/interca/ca/pem | openssl x509 -text


vault write interca/config/urls issuing_certificates="http://127.0.0.1:8200/v1/interca/ca" \ crl_distribution_points="http://127.0.0.1:8200/v1/interca/crl"

====

http POST http://example.com/posts/3 \
    Origin:example.com \  # :   HTTP headers
    name="John Doe" \     # =   string
    q=="search" \         # ==  URL parameters (?q=search)
    age:=29 \             # :=  for non-strings
    list:='[1,3,4]' \     # :=  json
    file@file.bin \       # @   attach file
    token=@token.txt \    # =@  read from file (text)
    user:=@user.json      # :=@ read from file (json)
