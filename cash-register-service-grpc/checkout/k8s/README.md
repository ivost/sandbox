
https://github.com/kubernetes-sigs/kustomize

kustomize build ~/someApp | kubectl apply -f -

kustomize build ~/someApp/overlays/production

https://github.com/kubernetes-sigs/kustomize/tree/master/examples/helloWorld

how to change image tag in deployment
```
kustomize edit set image ivostoy/store=ivostoy/store:0.12.4.0
```
