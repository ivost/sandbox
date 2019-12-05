

to install:

https://github.com/bitnami-labs/sealed-secrets/releases

brew install kubeseal

kaf kubeseal.yaml

note changes for --all-namespaces false

    spec:
      containers:
      - args: [ "--all-namespaces", "false" ]
        command:
        - controller
        env:
        - name: SEALED_SECRETS_ALL_NAMESPACES
          value: "false"
        image: quay.io/bitnami/sealed-secrets-controller:v0.9.6


# Create a json/yaml-encoded Secret somehow:
# (note use of `--dry-run` - this is just a local file!)

echo -n bar | kubectl create secret generic mysecret --dry-run --from-file=foo=/dev/stdin -o json >mysecret.json

# This is the important bit:
kubeseal <mysecret.json >mysealedsecret.json

# mysealedsecret.json is safe to upload to github, post to twitter,
# etc.  Eventually:
kubectl create -f mysealedsecret.json

# Profit!
kubectl describe secret mysecret


