package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

var version = "latest"

func NewVersion(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	return cmd
}
