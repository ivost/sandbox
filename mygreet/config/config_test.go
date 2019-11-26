package config_test

import (
	"testing"

	"github.com/ivost/sandbox/mygreet/config"
	"github.com/stretchr/testify/require"
)

const configFile = "./config.yaml"

func TestConfig(t *testing.T) {
	exp := &config.Config{
		Endpoint: "0.0.0.0:1234",
		Secure:   0,
		CertFile: "../ssl/server.crt",
		KeyFile:  "../ssl/server.pem",
	}
	cfg := config.New(configFile)
	require.NotNil(t, cfg)
	require.EqualValues(t, exp, cfg)
}
