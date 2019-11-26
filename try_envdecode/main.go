package main

import (
	"log"
	"net/url"
	"os"
	"time"

	"github.com/joeshaw/envdecode"
)

// https://github.com/joeshaw/envdecode

type Config struct {
	Hostname string   `env:"SERVER_HOSTNAME,default=localhost"`
	Port     uint16   `env:"SERVER_PORT,default=8080"`
	URL      *url.URL `env:"URL,required"`
	AWS      struct {
		ID        string   `env:"AWS_ACCESS_KEY_ID"`
		Secret    string   `env:"AWS_SECRET_ACCESS_KEY,required"`
		SnsTopics []string `env:"AWS_SNS_TOPICS"`
	}

	Timeout time.Duration `env:"TIMEOUT,default=1m,strict"`
}

func main() {
	var cfg Config
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "foo")
	_ = os.Setenv("TIMEOUT", "42s")
	_ = os.Setenv("URL", "http://example.com")
	_ = os.Setenv("AWS_SNS_TOPICS", "foo,bar")
	//os.Setenv("TIMEOUT", "foo") // error
	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	log.Printf("Config:%+v", cfg)
}
