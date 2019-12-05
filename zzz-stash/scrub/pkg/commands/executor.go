package commands

import (
	"io/ioutil"
	"os"
	"path"
	"scrub/pkg/config"
	"scrub/pkg/report"
	"sync"

	"github.com/cheggaaa/pb/v3"

	"scrub/pkg/logutils"
	"scrub/pkg/sysutils"
	"scrub/pkg/timeutils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Executor struct {
	rootCmd         *cobra.Command
	runCmd          *cobra.Command
	checkCmd        *cobra.Command
	exitCode        int
	version         string
	commit          string
	date            string
	cfg             *config.Config
	log             logutils.Log
	reportData      report.Data
	goenv           *sysutils.Env
	debugf          logutils.DebugFunc
	sw              *timeutils.Stopwatch
	Header          string
	HeaderLen       int64
	NumTotalFiles   int64
	NumFoundFiles   int64
	NumSkippedFiles int64
	NumTotalDirs    int64
	NumSkippedDirs  int64
	progress        *pb.ProgressBar
	locker          sync.Mutex
}

type VisitorFunc func(ex *Executor, path string, info os.FileInfo)

func NewExecutor(version, commit, date string) *Executor {
	cfg := config.NewDefault()
	e := &Executor{
		cfg:     cfg,
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
		logutils.SetupVerboseLog(e.log, commandLineCfg.IsVerbose)

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
	e.initCheck()

	// init e.cfg by values from config: flags parse will see these values
	// like the default ones. It will overwrite them only if the same option
	// is found in command-line: it's ok, command-line has higher priority.

	r := config.NewFileReader(e.cfg, commandLineCfg, e.log.Child("config_reader"))
	if err = r.Read(); err != nil {
		e.log.Fatalf("Can't read config: %s", err)
	}

	e.goenv = sysutils.NewEnv(e.log.Child("goenv"))
	//wd, _ := os.Getwd()
	//log.Printf("pwd %v", wd)
	ensureValidConfig(e.cfg, e.log)

	p := path.Join(e.cfg.Input.HeaderDir, e.cfg.Input.HeaderName)
	d, err := ioutil.ReadFile(p)
	if err != nil {
		e.log.Fatalf("Cannot read header %v", p)
	}
	e.Header = string(d)
	e.HeaderLen = int64(len(e.Header))
	e.progress = pb.StartNew(e.cfg.Input.EstimatedSize)

	e.sw = timeutils.NewStopwatch("some-scope", e.log.Child("stopwatch"))
	e.debugf("Initialized executor")
	return e
}

func (e *Executor) Execute() error {
	return e.rootCmd.Execute()
}

func (e *Executor) IncrementProgress() {
	e.locker.Lock()
	defer e.locker.Unlock()
	e.progress.Increment()
}
