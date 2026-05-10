package config_test

import (
	"fmt"

	"github.com/foomo/gokazi/pkg/config"
)

func ExampleTask() {
	t := config.Task{
		Name: "node",
		Path: "/usr/local/bin",
		Args: []string{"server.js"},
		Cwd:  "/srv/app",
	}
	fmt.Println(t.String())
	// Output: /usr/local/bin/node server.js in /srv/app
}

func ExampleTask_Match() {
	t := config.Task{
		Name: "python3",
		Args: []string{"-m", "http.server"},
	}
	ok := t.Match("python3", "", "", []string{"python3", "-m", "http.server", "8080"})
	fmt.Println(ok)
	// Output: true
}

func ExampleConfig() {
	cfg := config.Config{
		Version: config.Version,
		Tasks: map[string]config.Task{
			"web": {
				Name: "node",
				Args: []string{"server.js"},
			},
		},
	}
	fmt.Println(cfg.Version, cfg.Tasks["web"].Name)
	// Output: 1.0 node
}
