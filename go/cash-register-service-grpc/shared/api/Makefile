
.PHONY: all
.DEFAULT_GOAL:=all	
all:
	#./proto-gen.sh catalog/catalog.proto
	./proto-gen.sh checkout/checkout.proto

.PHONY: clean
clean:
	fd -e pb.go -e pb.gw.go -e swagger.json | xargs rm
