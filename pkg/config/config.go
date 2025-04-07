package config

type Config struct {
	// Config version
	Version string `json:"version" yaml:"version" koanf:"version"`
	// Tasks definitions
	Tasks map[string]Task `json:"tasks" yaml:"tasks" koanf:"tasks"`
}
