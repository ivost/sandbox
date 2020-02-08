package main

import (
	"log"

	"github.com/ivost/sandbox/myservice/pkg/version"

	"github.com/ivost/sandbox/myservice/config"
	"github.com/ivost/sandbox/myservice/server"
)

func main() {
	log.Printf("%v client %v %v", version.Name, version.Version, version.Build)
	cf := config.DefaultConfigFile
	conf := config.New(cf)
	server := server.New(conf)
	if server != nil {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}
}
