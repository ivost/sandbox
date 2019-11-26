package main

import (
	"log"

	"github.com/ivost/sandbox/mygreet/client"
	"github.com/ivost/sandbox/mygreet/config"
)

func main() {
	cf := config.DefaultConfigFile
	conf := config.New(cf)
	c := client.New(conf)

	log.Printf("grpc client endpoint:  %v, secure: %v", conf.Endpoint, conf.Secure)
	c.DoUnary()

	// c.DoServerStream()

	// c.DoClientStream()

	// c.DoBiDirStream(client)
}
