protoc -I/usr/local/include -I. \
  -I"$GOPATH/src" \
  -I"$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" \
  --go_out=plugins=grpc:. \
  myservice/myservice.proto

# echo generate gateway

protoc -I/usr/local/include -I. \
  -I"$GOPATH/src" \
  -I"$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" \
  --grpc-gateway_out=logtostderr=true:. \
  myservice/myservice.proto

protoc -I/usr/local/include -I. \
  -I"$GOPATH/src" \
  -I"$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis" \
  --swagger_out=logtostderr=true:. \
  myservice/myservice.proto

