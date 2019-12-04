package config

import (
	"flag"
	"log"
	"os"

	mconfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	mflag "github.com/micro/go-micro/config/source/flag"
)

// configuration
type Config struct {
	GrpcAddr string
	RestAddr string
	Secure   int
	CertFile string
	KeyFile  string
}

const DefaultGrpc = "0.0.0.0:52052"
const DefaultRest = "0.0.0.0:8080"
const DefaultConfigFile = "./config.yaml"
const DefaultCertFile = "./ssl/server.crt"
const DefaultKeyFile = "./ssl/server.pem"

func New(yamlFile string) *Config {
	configFile := yamlFile
	conf := &Config{
		GrpcAddr: DefaultGrpc,
		RestAddr: DefaultRest,
		CertFile: DefaultCertFile,
		KeyFile:  DefaultKeyFile,
		Secure:   0,
	}
	flag.StringVar(&configFile, "config", DefaultConfigFile, "config file")
	flag.StringVar(&conf.GrpcAddr, "grpc", DefaultGrpc, "grpc address")
	flag.StringVar(&conf.RestAddr, "rest", DefaultRest, "rest address")
	flag.IntVar(&conf.Secure, "secure", 0, "secure: 0=no TLS, 1=server, 2=mTLS")
	flag.Parse()
	//https://micro.mu/docs/go-config.html
	flags := mflag.NewSource(
	//mflag.IncludeUnset(true),
	)
	flags.Read()
	err := mconfig.Load(
		// base config from env
		env.NewSource(),
		// flag override
		flags,
	)

	if err = mconfig.Scan(conf); err != nil {
		panic(err)
	}

	_, err = os.Stat(configFile)
	if err == nil {
		log.Printf("Using config file %v", configFile)
		mconfig.LoadFile(configFile)
	}

	if err = mconfig.Scan(conf); err != nil {
		panic(err)
	}

	return conf
}
