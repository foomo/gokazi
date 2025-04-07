package cmd

import (
	"log/slog"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewRoot represents the base command when called without any subcommands
func NewRoot(l *slog.Logger) *cobra.Command {
	c := viper.New()
	cmd := &cobra.Command{
		Use:           "gogazi",
		Short:         "CLI process manager",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			pterm.PrintDebugMessages = c.GetBool("debug")
		},
	}

	flags := cmd.PersistentFlags()

	flags.Bool("debug", false, "enable debug logs")
	_ = c.BindPFlag("debug", flags.Lookup("debug"))

	return cmd
}
