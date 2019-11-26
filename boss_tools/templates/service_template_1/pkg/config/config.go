package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/braincorp/boss_pod_template/pkg/api"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
)

type config struct {
	v     map[string]interface{}
	Flags *pflag.FlagSet
	Cfg   api.Config
}

// A DecoderConfigOption can be passed to viper.Unmarshal to configure
// mapstructure.DecoderConfig options
type DecoderConfigOption func(*mapstructure.DecoderConfig)

func Configure() (c *config, err error) {
	c = new(config)
	// Flags definition
	c.Flags = pflag.NewFlagSet("default", pflag.ContinueOnError)
	c.Flags.Bool("version", false, "get version number")
	c.Flags.Bool("clone", false, "clone a new service")
	c.Flags.String("host-port", ":8080", "host:port")
	c.Flags.String("level", "info", "log level debug, info, warn, error, flat or panic")
	c.Flags.String("backend-url", "", "backend service URL")
	c.Flags.Duration("http-client-timeout", 40*time.Second, "client timeout duration")
	c.Flags.Duration("http-server-timeout", 30*time.Second, "server read and write timeout duration")
	c.Flags.Duration("http-server-shutdown-timeout", 5*time.Second, "server graceful shutdown timeout duration")
	c.Flags.String("data-path", "/data", "data local path")
	c.Flags.String("config-path", "", "config dir path")
	c.Flags.String("config", "config.yaml", "optional config file name")

	// parse Flags
	err = c.Flags.Parse(os.Args[1:])

	// viper adds 4 MB to 8 MB min binary
	//if err := viper.Unmarshal(&srvCfg); err != nil {
	//	logger.Panic("config unmarshal failed", zap.Error(err))
	//}
	if err := c.Unmarshal(c.Flags, &c.Cfg); err != nil {
		log.Panic(err)
	}
	return c, err
}

// defaultDecoderConfig returns default mapstructure.DecoderConfig with support
// of time.Duration values & string slices
func defaultDecoderConfig(result *api.Config, opts ...DecoderConfigOption) *mapstructure.DecoderConfig {
	c := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           result,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *config) Unmarshal(set *pflag.FlagSet, result *api.Config, opts ...DecoderConfigOption) error {
	err := c.decode(set, defaultDecoderConfig(result, opts...))
	return err
}

// BindPFlags binds a full flag set to the configuration, using each flag's long
// name as the config key.
func (c *config) BindPFlags(flags *pflag.FlagSet) (err error) {
	c.v = make(map[string]interface{})
	flags.VisitAll(func(flag *pflag.Flag) {
		c.v[strings.ToLower(flag.Name)] = flag.Value.String()
	})
	return nil
}

// A wrapper around mapstructure.Decode that mimics the WeakDecode functionality
func (c *config) decode(input interface{}, config *mapstructure.DecoderConfig) error {
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	c.BindPFlags(input.(*pflag.FlagSet))
	return decoder.Decode(c.v)
}
