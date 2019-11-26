package logutils

import (
	"github.com/fatih/color"
	colorable "github.com/mattn/go-colorable"
)

var StdOut = color.Output
var StdErr = colorable.NewColorableStderr()
