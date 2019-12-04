protoc -I/usr/local/include -I. \
  -I"$GOPATH/src" \
  -I"$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" \
  --go_out=plugins=grpc:. \
  myvault/myvault.proto

# echo generate gateway

protoc -I/usr/local/include -I. \
  -I"$GOPATH/src" \
  -I"$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" \
  --grpc-gateway_out=logtostderr=true:. \
  myvault/myvault.proto

protoc -I/usr/local/include -I. \
  -I"$GOPATH/src" \
  -I"$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" \
  --swagger_out=logtostderr=true:. \
  myvault/myvault.proto

