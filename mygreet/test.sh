#!/bin/bash
set -euo pipefail

echo grpcurl

grpcurl -plaintext localhost:52052 describe greet.GreetService

grpcurl -d '{ "greeting": { "first_name": "Ivo" } }' \
     -plaintext localhost:52052  greet.GreetService/Greet

echo grpc go client
[[ -f build/client ]] && build/client

echo REST endpoint on 8080/greet
curl -X POST \
  http://localhost:8080/greet \
  -H 'Accept: */*' \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
  "greeting": {
    "first_name": "John",
    "last_name": "Doe"
  }
}
'

echo evans CLI

echo '{ "greeting": { "first_name": "Ivo" } }' | evans --port 52052 --package greet --service GreetService --call Greet greet/greet.proto

echo evans REPL

echo package greet
echo service GreetService
echo call Greet

evans greet/greet.proto --port 52052 -r

