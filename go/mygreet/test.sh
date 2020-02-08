#!/bin/bash
#set -euo pipefail

echo === health
http :8080/health
grpcurl -plaintext localhost:52052  greet.GreetService/Health

echo === REST endpoint on 8080/greet
curl -X POST  http://localhost:8080/hello -d '{ "greeting": { "first_name": "John" } }'

echo === grpcurl

grpcurl -plaintext localhost:52052 describe greet.GreetService

grpcurl -d '{ "greeting": { "first_name": "Ivo" } }' \
     -plaintext localhost:52052  greet.GreetService/Greet

echo === grpc go client
[[ -f build/client ]] && build/client


echo " "

echo === evans CLI

echo '{ "greeting": { "first_name": "Ivo" } }' | evans --port 52052 --package greet --service GreetService --call Greet greet/greet.proto

echo === evans REPL

echo evans greet/greet.proto --repl --host localhost --port 52052 -r

echo evans greet/greet.proto --repl --host 192.168.99.100 --port 30506 -r

echo package greet
echo service GreetService
echo call Greet

#  pod=$(kubectl get pod -l app=hello -o  jsonpath='{.items[0].metadata.name}')
#  kubectl port-forward $pod 52052:52052 8080:8080

# amb.
# http://192.168.99.100:30506/greet

# url=http://192.168.99.100:30506/greet
# http $url/health
# curl -X POST $url/hello -d '{ "greeting": { "first_name": "John", "last_name": "Doe" } }'

# grpcurl -plaintext 192.168.99.100:30506 greet.GreetService/Health
