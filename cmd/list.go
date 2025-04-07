package cmd

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/alecthomas/chroma/quick"
	cmdx "github.com/foomo/gokazi/pkg/cmd"
	"github.com/foomo/gokazi/pkg/gokazi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewList(l *slog.Logger) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := cmdx.NewConfig(l, c, cmd)
			if err != nil {
				return err
			}

			gk := gokazi.New(l)
			for id, task := range cfg.Tasks {
				gk.Add(id, task)
			}
			tasks, err := gk.List(cmd.Context())
			if err != nil {
				return err
			}

			out, err := json.MarshalIndent(tasks, "", "  ")
			if err != nil {
				return err
			}

			return quick.Highlight(os.Stdout, string(out), "json", "terminal", "color")
		},
	}

	return cmd
}
