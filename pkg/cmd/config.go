package cmd

import (
	"io"
	"log/slog"

	"github.com/foomo/gokazi/pkg/config"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewConfig(l *slog.Logger, c *viper.Viper, cmd *cobra.Command) (*config.Config, error) {
	cmd.Flags().StringSliceP("config", "c", []string{"gokazi.yaml"}, "config files (default is gokazi.yaml)")
	if err := c.BindPFlag("config", cmd.Flags().Lookup("config")); err != nil {
		return nil, err
	}

	k := koanf.NewWithConf(koanf.Conf{
		Delim: "/",
	})

	for _, source := range c.GetStringSlice("config") {
		var p koanf.Provider
		switch source {
		case "-":
			l.Debug("reading config from stdin")
			if b, err := io.ReadAll(cmd.InOrStdin()); err != nil {
				return nil, err
			} else {
				p = rawbytes.Provider(b)
			}
		default:
			l.Debug("reading config from file: " + source)
			p = file.Provider(source)
		}
		if err := k.Load(p, yaml.Parser()); err != nil {
			return nil, errors.Wrap(err, "failed to parse config")
		}
	}

	var cfg *config.Config
	if err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{
		Tag: "koanf",
	}); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}

	if cfg.Version != config.Version {
		return nil, errors.New("missing or invalid config version: " + cfg.Version + " != '" + config.Version + "'")
	}

	return cfg, nil
}

// func ReadConfig(l *slog.Logger, cmd *cobra.Command) (*config.Config, error) {
// 	filenames := viper.GetStringSlice("config")
//
// 	for _, filename := range filenames {
// 		var p koanf.Provider
// 		switch {
// 		case filename == "-":
// 			pterm.Debug.Println("reading config from stdin")
// 			if b, err := io.ReadAll(cmd.InOrStdin()); err != nil {
// 				return nil, err
// 			} else {
// 				p = rawbytes.Provider(b)
// 			}
// 		default:
// 			pterm.Debug.Println("reading config from filename: " + filename)
// 			p = file.Provider(filename)
// 		}
// 		if err := k.Load(p, yaml.Parser()); err != nil {
// 			return nil, errors.Wrap(err, "error loading config file: "+filename)
// 		}
// 	}
//
// 	var cfg *config.Config
// 	pterm.Debug.Println("unmarshalling config")
// 	if err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{
// 		Tag: "yaml",
// 	}); err != nil {
// 		return nil, errors.Wrap(err, "failed to unmarshal config")
// 	}
//
// 	if cfg.Version != config.Version {
// 		return nil, errors.New("missing or invalid config version: " + cfg.Version + " != '" + config.Version + "'")
// 	}
//
// 	return cfg, nil
// }
