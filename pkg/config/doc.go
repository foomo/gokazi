// Package config defines the on-disk configuration schema for gokazi.
//
// A configuration declares a set of named tasks that the gokazi binary can
// list, start, and stop. Configurations are loaded from one or more YAML
// sources by [github.com/foomo/gokazi/pkg/cmd.NewConfig], which uses the
// koanf struct tag on each field.
//
// The package also acts as the source of truth for gokazi.schema.json:
// the generator at cmd/gokazi reflects [Config] using invopop/jsonschema
// and writes the resulting JSON Schema to disk. Doc comments on every
// exported field therefore appear verbatim as JSON Schema descriptions.
//
// Loading fails unless [Config.Version] equals [Version].
//
// Path-like fields on [Task] ([Task.Path], [Task.Cwd], [Task.Args]) are
// expanded with [os.ExpandEnv] when matched against running processes, so
// configurations may reference environment variables such as $PWD or $HOME.
package config
