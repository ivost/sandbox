package config

import (
	"regexp"
	"time"
)

type Run struct {
	CPUProfilePath      string
	MemProfilePath      string
	TracePath           string
	Concurrency         int
	PrintResourcesUsage bool `mapstructure:"print-resources-usage"`

	Args    []string
	Timeout time.Duration

	PrintVersion bool
}

type Config struct {
	Config    string `mapstructure:"config"`
	NoConfig  bool   `mapstructure:"no-config"`
	IsVerbose bool   `mapstructure:"verbose"`
	Silent    bool

	Run Run

	Input struct {
		RootDir       string           `mapstructure:"root"`
		HeaderDir     string           `mapstructure:"header-dir"`
		HeaderName    string           `mapstructure:"header"`
		Files         []string         `mapstructure:"files"` // files to includes
		FilesRex      []*regexp.Regexp `mapstructure:"-"`
		SkipFiles     []string         `mapstructure:"skip-files"` // files to exclude
		SkipDirs      []string         `mapstructure:"skip-dirs"`
		SkipFilesRex  []*regexp.Regexp `mapstructure:"-"`
		SkipDirsRex   []*regexp.Regexp `mapstructure:"-"`
		SkipText      []string         `mapstructure:"skip"`
		SkipTextRex   []*regexp.Regexp `mapstructure:"-"`
		EstimatedSize int              `mapstructure:"estimated-size"`
	}

	Output struct {
		Format string
		Color  string
	}
}

func NewDefault() *Config {
	return &Config{}
}
