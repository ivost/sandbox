package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// todo: use config-test.yaml

func TestCheck(t *testing.T) {
	// executor set in TestMain
	assert.NotNil(t, ex)

	ex.Check()
	//log.Printf("%v dirs total (%v skipped), %v files found (%v total, %v skipped)",
	//	ex.NumTotalDirs, ex.NumSkippedDirs, ex.NumFoundFiles, ex.NumTotalFiles, ex.NumSkippedFiles)

	assert.Equal(t, int64(3), ex.NumTotalDirs)
	assert.Equal(t, int64(1), ex.NumSkippedDirs)
	//assert.Equal(t, ex.NumTotalFiles, ex.NumFoundFiles+ex.NumSkippedFiles)
	assert.Equal(t, int64(5), ex.NumFoundFiles)
	assert.Equal(t, int64(5), ex.NumSkippedFiles)
	//assert.Equal(t, int64(7), ex.NumTotalFiles)
}
