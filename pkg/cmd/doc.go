// Package cmd provides shared building blocks used by the gokazi
// binary and by cmd/docgen.
//
// [NewConfig] loads the gokazi YAML configuration through koanf. It
// reads one or more sources passed via the --config / -c flag, layers
// them in order, and unmarshals the result into a
// [github.com/foomo/gokazi/pkg/config.Config] using the koanf struct
// tag and "/" as the key delimiter. The literal source "-" reads from
// the command's standard input.
//
// [NewLogger] returns a [log/slog.Logger] backed by
// [github.com/foomo/gokazi/pkg/pterm.SlogHandler] so library and
// command code share the same pterm-styled output. The package init
// function configures pterm prefixes and applies the GOKAZI_SCOPE
// environment variable as the scope label for every logger.
package cmd
