package config

import (
	"time"
)


type Run struct {
	IsVerbose           bool `mapstructure:"verbose"`
	Silent              bool
	CPUProfilePath      string
	MemProfilePath      string
	TracePath           string
	Concurrency         int
	PrintResourcesUsage bool `mapstructure:"print-resources-usage"`

	Config   string
	NoConfig bool

	Args    []string
	Timeout time.Duration

	PrintVersion       bool
	SkipFiles          []string `mapstructure:"skip-files"`
	SkipDirs           []string `mapstructure:"skip-dirs"`
}

type Config struct {
	Run Run

	Output struct {
		Format              string
		Color               string
	}
}

func NewDefault() *Config {
	return &Config{}
}
