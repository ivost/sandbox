# Cash Register Receipt Printer Service

## [Requirements](./REQUIREMENTS.md)

I've decided to use gRPC + [grpc-gateway](http://github.com/grpc-ecosystem/grpc-gateway)
which gives us free REST and Swagger.

It takes some time to setup the boilerplate but it is worth the effort.

## Main entities

### Item

```
message Item {
    string id   = 2;     // ID / UPC (unique string code)
    string name = 6;    // Item Name
    float price = 8;    // Item Unit Price
    bool  valid = 10;   // false on errors
}

```

### Receipt
```
message Receipt {
    string register = 2;
    repeated LineItem items = 6;
    int64  time = 10;
    float total = 12;
}

message LineItem {
    string upc = 2;
    int32 qty = 4;
    string name = 6;
    float price = 8;
    float extprice = 10;    // qty * price
}

```

## checkout service

### AddItem

Allows adding new or updating existing (upsert)

* input: Item
* output: Item

### Checkout 

Given array of UPC scan codes, builds and returns receipt according to business requirements
Errors out if an UPC is not found

* input: array of ScanItem
* output: Receipt, error 

## Implementation

See [protobuf schema](./shared/api/checkout/checkout.proto)

and [Swagger](./shared/api/checkout/checkout.swagger.json)

gRPC Reflection is enabled - you can use tool like [evans](https://evans.syfm.me/about) to introspect the schema  

[Main code](./checkout/server/receipt.go)

[Unit tests](./checkout/server/receipt_test.go)
 (using buffconn package for easy unit-testing of grpc stack)

run ```make help``` in checkout dir to see build targets (include docker and k8s)

local endpoints:
* gRPC Server on 0.0.0.0:52052
* REST Server on 0.0.0.0:8080

no TLS - securing needs more work (mainly configuration, see ssl folder).

Service is published as Docker container on dockerhub.

To run it:

```
docker run --rm -d -p 8080:8080 -p 52052:52052 ivostoy/checkout_service:0.12.7.0
```

Demo script with curl against REST API (in checkout/test/test.sh)

```
url=localhost:8080

curl $url/health
#
# add some items
#
curl -X POST $url/v1/checkout/items -d '{ "item": { "id": "111", "name": "foo", "price": 10.25 } }'
curl -X POST $url/v1/checkout/items -d '{ "item": { "id": "222", "name": "bar", "price": 21.50 } }'
#
# checkout
#
curl -X POST $url/v1/checkout/receipt -d '{ "register": "1", "items": [ {"upc": "111"}, {"upc": "222"},  {"upc": "111"}] }'
#
# "receipt.total should be 42 (: the ultimate answer to everything :)"
#

```


