#!/bin/bash

url=localhost:8080

echo
echo health check
echo
curl $url/health
echo
echo
echo ADD SOME ITEMS
echo
curl -X POST $url/v1/checkout/items -d '{ "item": { "id": "111", "name": "foo", "price": 10.25 } }'
curl -X POST $url/v1/checkout/items -d '{ "item": { "id": "222", "name": "bar", "price": 21.50 } }'
echo
echo
echo CHECKOUT
echo
curl -X POST $url/v1/checkout/receipt -d '{ "register": "1", "items": [ {"upc": "111"}, {"upc": "222"},  {"upc": "111"}] }'
echo
echo
echo "receipt.total should be 42 (: which is answer to everything :)"
echo

