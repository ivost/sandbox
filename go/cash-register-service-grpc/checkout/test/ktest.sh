#!/bin/bash
#set -euo pipefail
#### kube test - using API gateway
APORT=30088
URL=192.168.99.100:$APORT/store

http $URL/health

# kubectl get mappings -A
# kubectl describe mapping store-map1
# kubectl describe mapping store-map2

# this doesn't work
#grpc=store.StoreService
#grpcurl -plaintext 192.168.99.100:30506  ${grpc}/Health

# but this is OK
../build/client  -config client-config.yaml

# elvis is OK too (something with reflection?)

#echo === REST endpoint on 8080/store
#curl -X POST  http://localhost:8080/hello -d '{ "storeing": { "first_name": "John" } }'
#
#echo === grpcurl
#
#grpcurl -plaintext localhost:52052 describe myredis.KVService
#
#grpcurl -d '{ "storeing": { "first_name": "Ivo" } }' \
#     -plaintext localhost:52052  myredis.KVService/store
#
#echo === grpc go client
#[[ -f build/client ]] && build/client
#
#
#echo " "
#
#echo === evans CLI
#
#echo '{ "storeing": { "first_name": "Ivo" } }' | evans --port 52052 --package store --service KVService --call store store/store.proto
#
#echo === evans REPL
#
#echo evans store/store.proto --repl --host localhost --port 52052 -r
#
#echo evans store/store.proto --repl --host 192.168.99.100 --port 30506 -r
#
#echo package store
#echo service KVService
#echo call store

#  pod=$(kubectl get pod -l app=hello -o  jsonpath='{.items[0].metadata.name}')
#  kubectl port-forward $pod 52052:52052 8080:8080

# amb.
# http://192.168.99.100:30506/store

# url=http://192.168.99.100:30506/store
# http $url/health
# curl -X POST $url/hello -d '{ "storeing": { "first_name": "John", "last_name": "Doe" } }'

# grpcurl -plaintext 192.168.99.100:30506 myredis.KVService/Health
