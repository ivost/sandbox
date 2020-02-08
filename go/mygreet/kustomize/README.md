
https://github.com/kubernetes-sigs/kustomize

kustomize build ~/someApp | kubectl apply -f -

kustomize build ~/someApp/overlays/production

https://github.com/kubernetes-sigs/kustomize/tree/master/examples/helloWorld

pushd $BASE
kustomize edit set image monopole/hello=monopole/hello:1
popd