package client_test

import (
	"testing"

	"github.com/ivost/sandbox/myservice/client"
	"github.com/ivost/sandbox/myservice/config"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	conf := config.New("../config/config.yaml")
	require.NotNil(t, conf)
	c := client.New(conf)
	require.NotNil(t, c)
}
