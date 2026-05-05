package gokazi

import (
	"github.com/foomo/gokazi/pkg/config"
)

// Task represents a task with process information, configuration, and running status.
type Task struct {
	// Pid represents the process ID of the task. It is used to identify and manage the corresponding process.
	Pid int32 `json:"pid" yaml:"pid"`
	// Config represents the configuration details of the task, including name, description, path, args, and working directory.
	Config config.Task `json:"config" yaml:"config"`
	// Running represents the status of the task, indicating whether it is currently running or not.
	Running bool `json:"running" yaml:"running"`
}
