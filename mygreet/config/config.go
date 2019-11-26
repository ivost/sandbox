package config

import (
	"flag"

	mconfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/yaml"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	mflag "github.com/micro/go-micro/config/source/flag"
)

// configuration
type Config struct {
	Endpoint string
	Secure   int
	CertFile string
	KeyFile  string
}

const DefaultEndpoint = "0.0.0.0:51051"
const DefaultConfigFile = "./config.yaml"
const DefaultCertFile = "./ssl/server.crt"
const DefaultKeyFile = "./ssl/server.pem"

func New(yamlFile string) *Config {
	conf := &Config{
		Endpoint: DefaultEndpoint,
		CertFile: DefaultCertFile,
		KeyFile:  DefaultKeyFile,
		Secure:   0,
	}
	//flag.StringVar(&conf.Endpoint, "endpoint", DefaultEndpoint, "endpoint")
	flag.IntVar(&conf.Secure, "secure", 0, "secure: 0=no TLS, 1=server, 2=mTLS")
	flag.Parse()
	//log.Printf("secure %+v", secure)
	//https://micro.mu/docs/go-config.html
	flags := mflag.NewSource(
	//mflag.IncludeUnset(true),
	)
	flags.Read()
	err := mconfig.Load(
		// base config from env
		env.NewSource(),
		// yaml override
		file.NewSource(
			source.WithEncoder(yaml.NewEncoder()),
			file.WithPath(yamlFile),
		),
		// flag override
		flags,
	)
	if err != nil {
		//panic(err)
	}
	err = mconfig.Scan(conf)
	if err != nil {
		panic(err)
	}
	return conf
}
