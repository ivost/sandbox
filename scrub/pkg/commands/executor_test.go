package commands

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ex *Executor

func TestMain(m *testing.M) {

	// must go to the scrub root - expect config.yaml there and header.txt too
	os.Chdir("../..")

	ex = NewExecutor("0.0.0", "abcd", time.Now().Format(time.RFC3339))
	os.Exit(m.Run())
}

func TestNewExecutor(t *testing.T) {
	assert.NotNil(t, ex)
}
