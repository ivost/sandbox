backend "file" {
  path = "vault"
}

listener "tcp" {
  address = "127.0.0.1:8200"
  api_addr = "http://127.0.0.1:8200"
  tls_disable = 1
}
