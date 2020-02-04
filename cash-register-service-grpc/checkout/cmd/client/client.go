package main

import (
	"log"

	"github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/checkout/pkg/client"
	"github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/pkg/config"
	"github.com/lt-schmidt-jr/cash-register-service-IVOSTOYANOV/shared/pkg/version"
)

func main() {
	log.Printf("%v client %v %v", version.Name, version.Version, version.Build)
	cf := config.DefaultConfigFile
	conf := config.New(cf)
	c := client.New(conf)
	log.Printf("grpc client endpoint:  %v, secure: %v", conf.GrpcAddr, conf.Secure)
	rsp, err := c.Health()
	if err != nil {
		log.Printf("error %v", err)
		return
	}
	log.Printf("response %+v", rsp)
}
