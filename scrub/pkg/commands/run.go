package commands

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/google/renameio"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"scrub/pkg/config"
	"scrub/pkg/exitcodes"
	"scrub/pkg/logutils"
)

const welcomeMessage = "Run scrub to check/add copyright text to source code"

func wh(text string) string {
	return color.GreenString(text)
}

//nolint:funlen
func initFlagSet(fs *pflag.FlagSet, cfg *config.Config, isFinalInit bool) {
	// Run config
	rc := &cfg.Run
	fs.DurationVar(&rc.Timeout, "timeout", 10*time.Minute, wh("Timeout for total work"))

	fs.BoolVar(&rc.PrintResourcesUsage, "print-resources-usage", false,
		wh("Print avg and max memory usage and total time"))
}

func (e *Executor) initRunConfiguration(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.SortFlags = false // sort them as they are defined here
	initFlagSet(fs, e.cfg, true)
}

func (e *Executor) getConfigForCommandLine() (*config.Config, error) {
	// We use another pflag.FlagSet here to not set `changed` flag
	// on cli.Flags() options. Otherwise string slice options will be duplicated.
	fs := pflag.NewFlagSet("config flag set", pflag.ContinueOnError)

	var cfg config.Config
	// Don't do `fs.AddFlagSet(cli.Flags())` because it shares flags representations:
	// `changed` variable inside string slice vars will be shared.
	// Use another config variable here, not e.cfg, to not
	// affect main parsing by this parsing of only config option.
	initFlagSet(fs, &cfg, false)

	// Parse max options, even force version option: don't want
	// to get access to Executor here: it's error-prone to use
	// cfg vs e.cfg.
	initRootFlagSet(fs, &cfg, true)

	fs.Usage = func() {} // otherwise help text will be printed twice
	if err := fs.Parse(os.Args); err != nil {
		if err == pflag.ErrHelp {
			return nil, err
		}

		return nil, fmt.Errorf("can't parse args: %s", err)
	}

	return &cfg, nil
}

func (e *Executor) initRun() {
	e.runCmd = &cobra.Command{
		Use:   "run",
		Short: welcomeMessage,
		Run:   e.executeRun,
	}
	e.rootCmd.AddCommand(e.runCmd)
	e.runCmd.SetOut(logutils.StdOut) // use custom output to properly color it in Windows terminals
	e.initRunConfiguration(e.runCmd)
}

//nolint:unused
func fixSlicesFlags(fs *pflag.FlagSet) {
	// It's a dirty hack to set flag.Changed to true for every string slice flag.
	// It's necessary to merge config and command-line slices: otherwise command-line
	// flags will always overwrite ones from the config.
	fs.VisitAll(func(f *pflag.Flag) {
		if f.Value.Type() != "stringSlice" {
			return
		}

		s, err := fs.GetStringSlice(f.Name)
		if err != nil {
			return
		}

		if s == nil { // assume that every string slice flag has nil as the default
			return
		}

		// calling Set sets Changed to true: next Set calls will append, not overwrite
		_ = f.Value.Set(strings.Join(s, ","))
	})
}

//nolint:unused
func (e *Executor) setOutputToDevNull() (savedStdout, savedStderr *os.File) {
	savedStdout, savedStderr = os.Stdout, os.Stderr
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		e.log.Warnf("Can't open null device %q: %s", os.DevNull, err)
		return
	}

	os.Stdout, os.Stderr = devNull, devNull
	return
}

func RunVisitor(ex *Executor, path string, info os.FileInfo) {
	ex.RunVisitor(path, info)
}

func (e *Executor) RunVisitor(path string, info os.FileInfo) {
	//e.log.Infof("++++ Found file without header: %v", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		e.log.Infof("read %v error %v", path, err)
		return
	}

	f, err := renameio.TempFile("/tmp", path)
	if err != nil {
		e.log.Infof("renameio.TempFile error %v", err)
		return
	}
	defer f.Cleanup()
	_, err = f.WriteString(e.Header)
	if err != nil {
		e.log.Infof("Write header error %v", err)
		return
	}
	_, err = f.Write(data)
	if err != nil {
		e.log.Infof("Write file data error %v", err)
		return
	}

	err = f.CloseAtomicallyReplace()
	if err != nil {
		e.log.Infof("CloseAtomicallyReplace %v", err)
		return
	}

}

//func (e *Executor) setExitCodeIfIssuesFound(issues []result.Issue) {
//	if len(issues) != 0 {
//		e.exitCode = e.cfg.Run.ExitCodeIfIssuesFound
//	}
//}

func (e *Executor) Run(ctx context.Context, args []string) error {
	c := e.cfg.Input
	e.log.Debugf("RUN")
	e.SanityCheck()
	e.Traverse(c.RootDir, RunVisitor)
	e.log.Infof("FOUND %v file(s) without header",
		e.NumFoundFiles)
	e.log.Infof("checked %v dirs total (%v skipped), %v total files, %v skipped",
		e.NumTotalDirs, e.NumSkippedDirs, e.NumTotalFiles, e.NumSkippedFiles)
	return nil
}

func (e *Executor) executeRun(_ *cobra.Command, args []string) {

	needTrackResources := e.cfg.IsVerbose || e.cfg.Run.PrintResourcesUsage
	trackResourcesEndCh := make(chan struct{})
	defer func() { // XXX: this defer must be before ctx.cancel defer
		if needTrackResources { // wait until resource tracking finished to print properly
			<-trackResourcesEndCh
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), e.cfg.Run.Timeout)
	defer cancel()

	if needTrackResources {
		go watchResources(ctx, trackResourcesEndCh, e.log, e.debugf)
	}

	if err := e.Run(ctx, args); err != nil {
		e.log.Errorf("Run error: %s", err)
		if e.exitCode == exitcodes.Success {
			if exitErr, ok := errors.Cause(err).(*exitcodes.ExitError); ok {
				e.exitCode = exitErr.Code
			} else {
				e.exitCode = exitcodes.Failure
			}
		}
	}

	// set exit code to be number of files found
	e.exitCode = int(e.NumFoundFiles)
	e.setupExitCode(ctx)
}

func (e *Executor) setupExitCode(ctx context.Context) {
	if ctx.Err() != nil {
		e.exitCode = exitcodes.Timeout
		e.log.Errorf("Timeout exceeded: try increase it by passing --timeout option")
		return
	}

	if e.exitCode != exitcodes.Success {
		return
	}
}

func watchResources(ctx context.Context, done chan struct{}, logger logutils.Log, debugf logutils.DebugFunc) {
	startedAt := time.Now()
	debugf("Started tracking time")

	var maxRSSMB, totalRSSMB float64
	var iterationsCount int
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	logEveryRecord := os.Getenv("GL_MEM_LOG_EVERY") == "1"
	const MB = 1024 * 1024

	track := func() {
		debugf("Starting memory tracing iteration ...")
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		if logEveryRecord {
			debugf("Stopping memory tracing iteration, printing ...")
			printMemStats(&m, logger)
		}

		rssMB := float64(m.Sys) / MB
		if rssMB > maxRSSMB {
			maxRSSMB = rssMB
		}
		totalRSSMB += rssMB
		iterationsCount++
	}

	for {
		track()

		stop := false
		select {
		case <-ctx.Done():
			stop = true
			debugf("Stopped resources tracking")
		case <-ticker.C:
		}

		if stop {
			break
		}
	}
	track()

	avgRSSMB := totalRSSMB / float64(iterationsCount)

	logger.Infof("Memory: %d samples, avg is %.1fMB, max is %.1fMB",
		iterationsCount, avgRSSMB, maxRSSMB)
	logger.Infof("Execution took %s", time.Since(startedAt))
	close(done)
}
