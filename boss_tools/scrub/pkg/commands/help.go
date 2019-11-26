package commands

import (
	"github.com/spf13/cobra"
)

func (e *Executor) initHelp() {
	helpCmd := &cobra.Command{
		Use:   "help",
		Short: "Help",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				e.log.Fatalf("Usage: %v help", myName)
			}
			if err := cmd.Help(); err != nil {
				e.log.Fatalf("Can't run help: %s", err)
			}
		},
	}
	e.rootCmd.SetHelpCommand(helpCmd)
}
