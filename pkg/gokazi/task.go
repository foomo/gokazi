package gokazi

import (
	"github.com/foomo/gokazi/pkg/config"
)

type Task struct {
	Config  config.Task `json:"config" yaml:"config"`
	Pid     int32       `json:"pid" yaml:"pid"`
	Running bool        `json:"running" yaml:"running"`
}
