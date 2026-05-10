package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

var (
	version        = "dev"
	commitHash     = "none"
	buildTimestamp = "unknown"
)

// NewVersion returns the "gokazi version" subcommand, which prints the
// build version, commit hash, and build timestamp injected at link
// time via -ldflags.
func NewVersion(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version, commitHash, buildTimestamp)
		},
	}

	return cmd
}
