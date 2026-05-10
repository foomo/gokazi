package config

// Config is the root document of a gokazi configuration file.
//
// A Config is loaded from one or more YAML sources and validated against
// the package-level [Version] constant. The Tasks map keys are the task
// IDs used by the gokazi CLI (for example, "gokazi stop <id>").
type Config struct {
	// Version is the configuration schema version. It must equal the
	// package-level [Version] constant or loading fails.
	Version string `json:"version" yaml:"version" koanf:"version"`
	// Tasks maps a task ID to its definition. The ID is the value passed
	// to subcommands such as "gokazi stop".
	Tasks map[string]Task `json:"tasks" yaml:"tasks" koanf:"tasks"`
}
