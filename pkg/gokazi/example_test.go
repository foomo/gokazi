package gokazi_test

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/foomo/gokazi/pkg/config"
	"github.com/foomo/gokazi/pkg/gokazi"
)

// Absent is the name of an executable that is not expected to be
// running while the examples execute, used to make List output
// deterministic.
const absent = "gokazi-example-nonexistent-binary"

func ExampleNew() {
	gk := gokazi.New(slog.New(slog.NewTextHandler(io.Discard, nil)))
	gk.Add("web", config.Task{Name: absent, Args: []string{"server.js"}})

	fmt.Println("registered")
	// Output: registered
}

func ExampleGokazi_List() {
	gk := gokazi.New(slog.New(slog.NewTextHandler(io.Discard, nil)))
	gk.Add("web", config.Task{Name: absent})

	tasks, err := gk.List(context.Background())
	if err != nil {
		return
	}

	t := tasks["web"]

	fmt.Printf("running=%t pid=%d\n", t.Running, t.Pid)
	// Output: running=false pid=0
}

func ExampleGokazi_Find() {
	gk := gokazi.New(slog.New(slog.NewTextHandler(io.Discard, nil)))
	gk.Add("web", config.Task{Name: absent})

	t, err := gk.Find(context.Background(), "web")
	if err != nil {
		return
	}

	fmt.Printf("running=%t\n", t.Running)
	// Output: running=false
}

func ExampleGokazi_Stop() {
	gk := gokazi.New(slog.New(slog.NewTextHandler(io.Discard, nil)))
	gk.Add("web", config.Task{Name: absent})

	err := gk.Stop(context.Background(), "web")

	fmt.Println(err)
	// Output: web: task not running
}
