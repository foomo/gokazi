package config

import (
	"os"
	"path"
	"slices"
	"strings"
)

type Task struct {
	// Task name
	Name string `json:"name" yaml:"name" koanf:"name"`
	// Task description
	Description string `json:"description" yaml:"description" koanf:"description"`
	// Task path
	Path string `json:"path" yaml:"path" koanf:"path"`
	// Task working directory
	Cwd string `json:"cwd" yaml:"cwd" koanf:"cwd"`
	// Task Args
	Args []string `json:"args" yaml:"args" koanf:"args"`
}

func (t Task) ExpandCwd() string {
	if t.Cwd == "" {
		return ""
	}
	return path.Clean(os.ExpandEnv(t.Cwd))
}

func (t Task) ExpandPath() string {
	if t.Path == "" {
		return ""
	}
	return path.Clean(os.ExpandEnv(t.Path))
}

func (t Task) ExpandArgs() []string {
	ret := make([]string, len(t.Args))
	for i, arg := range t.Args {
		ret[i] = os.ExpandEnv(arg)
	}
	return ret
}

func (t Task) Match(name, cpath, cwd string, args []string) bool {
	if t.Name != name {
		return false
	}
	if t.Path != "" && t.ExpandPath() != cpath {
		return false
	}
	if t.Cwd != "" && t.ExpandCwd() != cwd {
		return false
	}
	if len(t.Args) > 0 {
		for _, arg := range t.ExpandArgs() {
			if !slices.Contains(args, arg) {
				return false
			}
		}
	}
	return true
}

func (t Task) String() string {
	var ret string
	if t.Path != "" {
		ret += path.Join(t.ExpandPath(), t.Name)
	} else {
		ret += t.Name
	}
	if len(t.Args) > 0 {
		ret += " " + strings.Join(t.ExpandArgs(), " ")
	}
	if t.Cwd != "" {
		ret += " in " + t.ExpandCwd()
	}
	return ret
}
