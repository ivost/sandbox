package commands

import (
	"os"
	"scrub/pkg/config"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ex *Executor

func TestMain(m *testing.M) {

	// must go to the scrub root - expects config.yaml there and header.txt too
	os.Chdir("../..")

	// use config-test.yml which points to test/testdata
	os.Setenv(config.EnvVarName, "test/config-test.yaml")
	ex = NewExecutor("0.0.0", "abcd", time.Now().Format(time.RFC3339))
	os.Exit(m.Run())
}

func TestNewExecutor(t *testing.T) {
	assert.NotNil(t, ex)
}
