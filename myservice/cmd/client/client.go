package main

import (
	"log"

	"github.com/ivost/sandbox/myservice/client"
	"github.com/ivost/sandbox/myservice/config"
)

func main() {
	cf := config.DefaultConfigFile
	conf := config.New(cf)
	c := client.New(conf)

	log.Printf("grpc client endpoint:  %v, secure: %v", conf.GrpcAddr, conf.Secure)
	c.DoUnary()

	// c.DoServerStream()

	// c.DoClientStream()

	// c.DoBiDirStream(client)
}
