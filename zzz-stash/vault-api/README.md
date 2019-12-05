
export VAULT_TOKEN=root
kubectl port-forward vault-0 8200:8200

alias pfv='kubectl port-forward vault-0 8200:8200'


