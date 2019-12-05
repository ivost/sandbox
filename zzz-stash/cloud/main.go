package main

import (
	"github.com/ivostoyanov-bc/sandbox/foo"
)

/*
export GO111MODULE=off
go clean -modcache
go get github.com/ivostoyanov-bc/sandbox/foo

import "github.com/ivostoyanov-bc/sandbox/foo"
// works in GO111MODULE=off
go run main.go
*/

func main() {
	foo.Hello()
}
