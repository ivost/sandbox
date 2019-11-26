package main

import (
	"github.com/ivostoyanov-bc/sandbox/mygreet/client"
	"github.com/ivostoyanov-bc/sandbox/mygreet/config"
)

func main() {
	cf := config.DefaultConfigFile
	conf := config.New(cf)
	c := client.New(conf)
	c.DoUnary()

	// c.DoServerStream()

	// c.DoClientStream()

	// c.DoBiDirStream(client)
}
