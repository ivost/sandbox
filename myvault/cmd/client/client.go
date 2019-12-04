package main

import (
	"log"

	"github.com/ivost/sandbox/myvault/pkg/version"

	"github.com/ivost/sandbox/myvault/client"
	"github.com/ivost/sandbox/myvault/config"
)

func main() {
	log.Printf("%v client %v %v", version.Name, version.Version, version.Build)
	cf := config.DefaultConfigFile
	conf := config.New(cf)
	c := client.New(conf)

	log.Printf("grpc client endpoint:  %v, secure: %v", conf.GrpcAddr, conf.Secure)
	c.DoUnary()

	// c.DoServerStream()

	// c.DoClientStream()

	// c.DoBiDirStream(client)
}
