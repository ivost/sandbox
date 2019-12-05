package commands

import (
	"context"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	// executor set in TestMain
	assert.NotNil(t, ex)
	// must go the work root
	os.Chdir("../..")
	ctx := context.Background()

	// copy testdata dir to temp location
	// and scrub files there
	temp := "/tmp"
	root := ex.cfg.Input.RootDir
	out, err := exec.Command("rm", "-rf", temp+"/testdata").Output()
	out, err = exec.Command("cp", "-r", root, temp).Output()
	assert.NoError(t, err)
	_ = out

	// use new data dir
	ex.cfg.Input.RootDir = path.Join(temp, "testdata")
	ex.cfg.Input.HeaderDir = path.Join(temp, "testdata")

	ex.Check()
	assert.Equal(t, int64(5), ex.NumFoundFiles)
	// now scrub
	ex.Run(ctx, nil)
	assert.Equal(t, int64(5), ex.NumFoundFiles)
	// double check - should find 0 files
	ex.Check()
	assert.Equal(t, int64(0), ex.NumFoundFiles)
}
