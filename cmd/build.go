package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

// New returns the fully-wired gokazi root command with all subcommands attached.
// Reused by the binary entry point and by cmd/docgen.
func New(l *slog.Logger) *cobra.Command {
	root := NewRoot(l)
	root.AddCommand(
		NewList(l),
		NewStop(l),
		NewConfig(l),
		NewVersion(l),
	)

	return root
}
