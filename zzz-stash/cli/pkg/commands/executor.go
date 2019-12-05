package commands

import (
	"cli/pkg/config"
	"cli/pkg/report"

	"cli/pkg/logutils"
	goutil "cli/pkg/sysutils"
	"cli/pkg/timeutils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Executor struct {
	rootCmd *cobra.Command
	runCmd  *cobra.Command
	exitCode              int
	version, commit, date string
	cfg    *config.Config
	log    logutils.Log
	reportData        report.Data
	goenv  *goutil.Env
	debugf logutils.DebugFunc
	sw     *timeutils.Stopwatch
}

func NewExecutor(version, commit, date string) *Executor {
	e := &Executor{
		cfg:     config.NewDefault(),
		version: version,
		commit:  commit,
		date:    date,
		debugf:  logutils.Debug("exec"),
	}

	e.debugf("Starting execution...")
	e.log = report.NewLogWrapper(logutils.NewStderrLog(""), &e.reportData)

	// to setup log level early we need to parse config from command line extra time to
	// find `-v` option
	commandLineCfg, err := e.getConfigForCommandLine()
	if err != nil && err != pflag.ErrHelp {
		e.log.Fatalf("Can't get config for command line: %s", err)
	}
	if commandLineCfg != nil {
		logutils.SetupVerboseLog(e.log, commandLineCfg.Run.IsVerbose)

		switch commandLineCfg.Output.Color {
		case "always":
			color.NoColor = false
		case "never":
			color.NoColor = true
		case "auto":
			// nothing
		default:
			e.log.Fatalf("invalid value %q for --color; must be 'always', 'auto', or 'never'", commandLineCfg.Output.Color)
		}
	}

	// init of commands must be done before config file reading because
	// init sets config with the default values of flags
	e.initRoot()
	e.initRun()
	e.initHelp()
	e.initConfig()
	e.initVersion()

	// init e.cfg by values from config: flags parse will see these values
	// like the default ones. It will overwrite them only if the same option
	// is found in command-line: it's ok, command-line has higher priority.

	r := config.NewFileReader(e.cfg, commandLineCfg, e.log.Child("config_reader"))
	if err = r.Read(); err != nil {
		e.log.Fatalf("Can't read config: %s", err)
	}

	e.goenv = goutil.NewEnv(e.log.Child("goenv"))

	e.sw = timeutils.NewStopwatch("some-scope", e.log.Child("stopwatch"))
	e.debugf("Initialized executor")
	return e
}

func (e *Executor) Execute() error {
	return e.rootCmd.Execute()
}
