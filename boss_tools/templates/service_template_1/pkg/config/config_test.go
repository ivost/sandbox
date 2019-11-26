package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	cfg, err := Configure()
	require.NoError(t, err)
	require.NotNil(t, cfg)
}
