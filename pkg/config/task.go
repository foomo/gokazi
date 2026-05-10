package config

import (
	"os"
	"path"
	"slices"
	"strings"
)

// Task describes a single managed process.
//
// A Task is matched against running operating-system processes by the
// gokazi runtime: a process matches when its executable name equals
// [Task.Name] and, where set, its executable directory equals
// [Task.ExpandPath], its working directory equals [Task.ExpandCwd], and
// every expanded argument in [Task.ExpandArgs] appears in the process
// command line. See [Task.Match] for the exact rules.
type Task struct {
	// Name is the executable name of the process, as reported by the
	// operating system (for example, "node" or "python3"). It is matched
	// exactly and case-sensitively.
	Name string `json:"name" yaml:"name" koanf:"name"`
	// Description is a human-readable summary of the task. It is shown by
	// "gokazi list" and "gokazi config" but is not used for matching.
	Description string `json:"description" yaml:"description" koanf:"description"`
	// Path is the directory containing the executable. When non-empty,
	// only processes whose executable lives in this directory match.
	// The value is expanded with [os.ExpandEnv] before matching.
	Path string `json:"path" yaml:"path" koanf:"path"`
	// Cwd is the working directory of the process. When non-empty, only
	// processes started from this directory match. The value is expanded
	// with [os.ExpandEnv] before matching.
	Cwd string `json:"cwd" yaml:"cwd" koanf:"cwd"`
	// Args is the set of command-line arguments that must be present on
	// the running process. Order and additional arguments are ignored:
	// every entry must appear in the process command line for the task
	// to match. Each value is expanded with [os.ExpandEnv].
	Args []string `json:"args" yaml:"args" koanf:"args"`
}

// ExpandCwd returns [Task.Cwd] with environment variables expanded by
// [os.ExpandEnv] and the result cleaned by [path.Clean]. It returns the
// empty string when Cwd is empty.
func (t Task) ExpandCwd() string {
	if t.Cwd == "" {
		return ""
	}

	return path.Clean(os.ExpandEnv(t.Cwd))
}

// ExpandPath returns [Task.Path] with environment variables expanded by
// [os.ExpandEnv] and the result cleaned by [path.Clean]. It returns the
// empty string when Path is empty.
func (t Task) ExpandPath() string {
	if t.Path == "" {
		return ""
	}

	return path.Clean(os.ExpandEnv(t.Path))
}

// ExpandArgs returns a new slice containing each entry of [Task.Args]
// with environment variables expanded by [os.ExpandEnv].
func (t Task) ExpandArgs() []string {
	ret := make([]string, len(t.Args))
	for i, arg := range t.Args {
		ret[i] = os.ExpandEnv(arg)
	}

	return ret
}

// Match reports whether the given operating-system process attributes
// match the task. The arguments are the process name, executable path,
// working directory, and command-line arguments.
//
// A match requires Name to be equal, and, when set on the task, the
// expanded Path and Cwd to be equal to the corresponding process value.
// When Args is non-empty, every expanded entry must appear in args; the
// order of arguments and any extra arguments on the process are ignored.
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

// String returns a human-readable representation of the task in the form
// "<path>/<name> <args> in <cwd>". Path, Args, and Cwd are omitted when
// not set. Path-like fields are expanded with [os.ExpandEnv].
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
