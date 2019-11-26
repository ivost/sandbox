package main

import (
	"github.com/ivostoyanov-bc/sandbox/mygreet/config"
	"github.com/ivostoyanov-bc/sandbox/mygreet/server"
)

func main() {
	cf := config.DefaultConfigFile
	conf := config.New(cf)
	server := server.New(conf)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
