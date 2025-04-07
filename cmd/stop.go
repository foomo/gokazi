package cmd

import (
	"log/slog"

	cmdx "github.com/foomo/gokazi/pkg/cmd"
	"github.com/foomo/gokazi/pkg/gokazi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewStop(l *slog.Logger) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop task by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := cmdx.NewConfig(l, c, cmd)
			if err != nil {
				return err
			}

			gk := gokazi.New(l)
			for id, task := range cfg.Tasks {
				gk.Add(id, task)
			}

			return gk.Stop(cmd.Context(), args[0])
		},
	}

	return cmd
}
