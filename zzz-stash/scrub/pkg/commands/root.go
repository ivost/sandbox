package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"scrub/pkg/config"
	"scrub/pkg/fsutils"
	"scrub/pkg/logutils"
)

const myName = "scrub"

func (e *Executor) persistentPreRun(_ *cobra.Command, _ []string) {
	if e.cfg.Run.PrintVersion {
		fmt.Fprintf(logutils.StdOut, "version %s built from %s on %s\n", e.version, e.commit, e.date)
		os.Exit(0)
	}

	runtime.GOMAXPROCS(e.cfg.Run.Concurrency)

	if e.cfg.Run.CPUProfilePath != "" {
		f, err := os.Create(e.cfg.Run.CPUProfilePath)
		if err != nil {
			e.log.Fatalf("Can't create file %s: %s", e.cfg.Run.CPUProfilePath, err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			e.log.Fatalf("Can't start CPU profiling: %s", err)
		}
	}

	if e.cfg.Run.MemProfilePath != "" {
		if rate := os.Getenv("GL_MEMPROFILE_RATE"); rate != "" {
			runtime.MemProfileRate, _ = strconv.Atoi(rate)
		}
	}

	if e.cfg.Run.TracePath != "" {
		f, err := os.Create(e.cfg.Run.TracePath)
		if err != nil {
			e.log.Fatalf("Can't create file %s: %s", e.cfg.Run.TracePath, err)
		}
		if err = trace.Start(f); err != nil {
			e.log.Fatalf("Can't start tracing: %s", err)
		}
	}
}

func (e *Executor) persistentPostRun(_ *cobra.Command, _ []string) {
	if e.cfg.Run.CPUProfilePath != "" {
		pprof.StopCPUProfile()
	}
	if e.cfg.Run.MemProfilePath != "" {
		f, err := os.Create(e.cfg.Run.MemProfilePath)
		if err != nil {
			e.log.Fatalf("Can't create file %s: %s", e.cfg.Run.MemProfilePath, err)
		}

		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)
		printMemStats(&ms, e.log)

		if err := pprof.WriteHeapProfile(f); err != nil {
			e.log.Fatalf("Can't write heap profile: %s", err)
		}
		f.Close()
	}
	if e.cfg.Run.TracePath != "" {
		trace.Stop()
	}

	e.log.Infof("Exit code (number of files without header) %v", e.exitCode)
	os.Exit(e.exitCode)
}

func printMemStats(ms *runtime.MemStats, logger logutils.Log) {
	logger.Infof("Mem stats: alloc=%s total_alloc=%s sys=%s "+
		"heap_alloc=%s heap_sys=%s heap_idle=%s heap_released=%s heap_in_use=%s "+
		"stack_in_use=%s stack_sys=%s "+
		"mspan_sys=%s mcache_sys=%s buck_hash_sys=%s gc_sys=%s other_sys=%s "+
		"mallocs_n=%d frees_n=%d heap_objects_n=%d gc_cpu_fraction=%.2f",
		formatMemory(ms.Alloc), formatMemory(ms.TotalAlloc), formatMemory(ms.Sys),
		formatMemory(ms.HeapAlloc), formatMemory(ms.HeapSys),
		formatMemory(ms.HeapIdle), formatMemory(ms.HeapReleased), formatMemory(ms.HeapInuse),
		formatMemory(ms.StackInuse), formatMemory(ms.StackSys),
		formatMemory(ms.MSpanSys), formatMemory(ms.MCacheSys), formatMemory(ms.BuckHashSys),
		formatMemory(ms.GCSys), formatMemory(ms.OtherSys),
		ms.Mallocs, ms.Frees, ms.HeapObjects, ms.GCCPUFraction)
}

func formatMemory(memBytes uint64) string {
	if memBytes < 1024 {
		return fmt.Sprintf("%db", memBytes)
	}
	if memBytes < 1024*1024 {
		return fmt.Sprintf("%dkb", memBytes/1024)
	}
	return fmt.Sprintf("%dmb", memBytes/1024/1024)
}

func getDefaultConcurrency() int {
	if os.Getenv("HELP_RUN") == "1" {
		return 8
	}
	return runtime.NumCPU()
}

func (e *Executor) initRoot() {
	rootCmd := &cobra.Command{
		Use:   "scrub",
		Short: "check/add copyright header",
		Long:  `check/add copyright header`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				e.log.Fatalf("Usage: %v ...", myName)
			}
			if err := cmd.Help(); err != nil {
				e.log.Fatalf("Can't run help: %s", err)
			}
		},
		PersistentPreRun:  e.persistentPreRun,
		PersistentPostRun: e.persistentPostRun,
	}
	initRootFlagSet(rootCmd.PersistentFlags(), e.cfg, e.needVersionOption())
	e.rootCmd = rootCmd
}

func (e *Executor) needVersionOption() bool {
	return e.date != ""
}

func initRootFlagSet(fs *pflag.FlagSet, cfg *config.Config, needVersionOption bool) {

	p := os.Getenv("SCRUB_CONFIG")
	if p == "" {
		p = "config.yaml"
	}
	fs.StringVarP(&cfg.Config, "config", "c", p, wh("Read config from file path `PATH`"))
	fs.BoolVar(&cfg.NoConfig, "no-config", false, wh("Don't read config"))

	fs.BoolVarP(&cfg.IsVerbose, "verbose", "v", false, wh("verbose output"))
	fs.StringVar(&cfg.Run.CPUProfilePath, "cpu-profile-path", "", wh("Path to CPU profile output file"))
	fs.StringVar(&cfg.Run.MemProfilePath, "mem-profile-path", "", wh("Path to memory profile output file"))
	fs.StringVar(&cfg.Run.TracePath, "trace-path", "", wh("Path to trace output file"))
	fs.IntVarP(&cfg.Run.Concurrency, "concurrency", "j", getDefaultConcurrency(), wh("Concurrency (default NumCPU)"))
	if needVersionOption {
		fs.BoolVar(&cfg.Run.PrintVersion, "version", false, wh("Print version"))
	}
	fs.StringVar(&cfg.Output.Color, "color", "auto", wh("Use color when printing; can be 'always', 'auto', or 'never'"))
	//fs.StringSliceVar(&cfg.Input.SkipDirs, "skip-dirs", nil, wh("Regexps of directories to skip"))
	//fs.StringSliceVar(&cfg.Input.SkipFiles, "skip-files", nil, wh("Regexps of files to skip"))
}

func ensureValidConfig(cfg *config.Config, log logutils.Log) {
	var err error

	if cfg.Input.RootDir == "" {
		log.Fatalf("no root dir set")
	}

	cfg.Input.RootDir, err = filepath.Abs(cfg.Input.RootDir)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	if cfg.Input.HeaderDir == "" {
		cfg.Input.HeaderDir = cfg.Input.RootDir
	}
	//cfg.Input.SkipDirs = append(cfg.Input.SkipDirs, fsutils.StdExcludeDirRegexps...)
	// compile regexps
	for _, r := range cfg.Input.SkipDirs {
		re := fsutils.OptionallyEscape(r)
		re = fsutils.PathElemRe(re)
		// todo: deduplicate
		rex := regexp.MustCompile(re)
		cfg.Input.SkipDirsRex = append(cfg.Input.SkipDirsRex, rex)
	}
	for _, r := range cfg.Input.SkipFiles {
		re := fsutils.OptionallyEscape(r)
		re = fsutils.PathElemRe(re)
		rex := regexp.MustCompile(re)
		cfg.Input.SkipFilesRex = append(cfg.Input.SkipFilesRex, rex)
	}
	// include files - check after exclusion check
	for _, r := range cfg.Input.Files {
		re := fsutils.OptionallyEscape(r)
		re = fsutils.PathElemRe(re)
		rex := regexp.MustCompile(re)
		cfg.Input.FilesRex = append(cfg.Input.FilesRex, rex)
	}
	for _, r := range cfg.Input.SkipText {
		re := strings.ToLower(r)
		rex := regexp.MustCompile(re)
		cfg.Input.SkipTextRex = append(cfg.Input.SkipTextRex, rex)
	}

}
