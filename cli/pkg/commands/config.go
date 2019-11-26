package commands

import (
	"cli/pkg/exitcodes"
	"cli/pkg/fsutils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (e *Executor) initConfig() {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Config",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				e.log.Fatalf("Usage: cli config")
			}
			if err := cmd.Help(); err != nil {
				e.log.Fatalf("Can't run help: %s", err)
			}
		},
	}
	e.rootCmd.AddCommand(cmd)

	pathCmd := &cobra.Command{
		Use:   "path",
		Short: "Print used config path",
		Run:   e.executePathCmd,
	}
	e.initRunConfiguration(pathCmd) // allow --config
	cmd.AddCommand(pathCmd)
}

func (e *Executor) executePathCmd(_ *cobra.Command, args []string) {
	if len(args) != 0 {
		e.log.Fatalf("Usage: cli config path")
	}

	usedConfigFile := viper.ConfigFileUsed()
	if usedConfigFile == "" {
		e.log.Warnf("No config file detected")
		os.Exit(exitcodes.NoConfigFileDetected)
	}

	usedConfigFile, err := fsutils.ShortestRelPath(usedConfigFile, "")
	if err != nil {
		e.log.Warnf("Can't pretty print config file path: %s", err)
	}

	fmt.Println(usedConfigFile)
	os.Exit(0)
}
