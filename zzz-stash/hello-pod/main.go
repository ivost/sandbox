package main

import (
	"fmt"
	"log"
	"net/http"
)

var Version string
var Build string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s !!! My Version: %v, build: %v", r.URL.Path[1:], Version, Build)
}

func main() {
	endpoint := ":8080"
	log.Printf("http Listen And Serve %v, ver: %v, build: %v", endpoint, Version, Build)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(endpoint, nil))
}
