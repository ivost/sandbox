package main

import (
	"github.com/ivost/sandbox/myservice/config"
	"github.com/ivost/sandbox/myservice/server"
)

func main() {
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
