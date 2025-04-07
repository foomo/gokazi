package gokazi

import (
	"context"
	"log/slog"
	"os/exec"
	"path"
	"slices"
	"syscall"
	"time"

	"github.com/foomo/gokazi/pkg/config"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shirou/gopsutil/v4/process"
)

var (
	ErrNotFound       = errors.New("task not found")
	ErrNotRunning     = errors.New("task not running")
	ErrAlreadyRunning = errors.New("task already running")
)

type Gokazi struct {
	l     *slog.Logger
	tasks map[string]config.Task
}

func New(l *slog.Logger) *Gokazi {
	return &Gokazi{
		l:     l,
		tasks: map[string]config.Task{},
	}
}

func (g *Gokazi) Add(id string, task config.Task) {
	g.tasks[id] = task
}

// Start a detached child process
func (g *Gokazi) Start(ctx context.Context, id string, cmd *exec.Cmd) error {
	if t, err := g.Find(ctx, id); err != nil {
		return err
	} else if t.Running {
		return ErrAlreadyRunning
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
		Pgid:    0,
	}

	g.l.Info("Starting: " + cmd.String())
	if err := cmd.Start(); err != nil {
		return err
	}

	if t, err := g.Find(ctx, id); err != nil {
		return err
	} else if !t.Running {
		return ErrNotRunning
	}

	return nil
}

func (g *Gokazi) Stop(ctx context.Context, id string) error {
	task, ok := g.tasks[id]
	if !ok {
		return errors.Errorf("task '%s' not found", id)
	}

	ps, err := g.listProcesses(ctx)
	if err != nil {
		return err
	}

	p, err := g.findProcess(ctx, task, ps)
	if err != nil {
		return err
	}

	running, err := p.IsRunningWithContext(ctx)
	if err != nil {
		return err
	} else if !running {
		return ErrNotRunning
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return p.KillWithContext(ctx)
}

func (g *Gokazi) Find(ctx context.Context, id string) (Task, error) {
	task, ok := g.tasks[id]
	if !ok {
		return Task{}, errors.Errorf("task '%s' not found", id)
	}

	ps, err := g.listProcesses(ctx)
	if err != nil {
		return Task{}, err
	}

	t := Task{
		Config:  task,
		Running: false,
		Pid:     0,
	}

	p, err := g.findProcess(ctx, task, ps)
	if errors.Is(err, ErrNotFound) {
		t.Running = false
	} else if err != nil {
		return Task{}, err
	} else {
		running, err := p.IsRunningWithContext(ctx)
		if err != nil {
			return Task{}, err
		}
		t.Running = running
		t.Pid = p.Pid
	}

	return t, nil
}

func (g *Gokazi) List(ctx context.Context) (map[string]Task, error) {
	ps, err := g.listProcesses(ctx)
	if err != nil {
		return nil, err
	}

	tasks := map[string]Task{}
	for id, task := range g.tasks {
		t := Task{
			Config:  task,
			Pid:     0,
			Running: false,
		}
		p, err := g.findProcess(ctx, task, ps)
		if errors.Is(err, ErrNotFound) {
			t.Running = false
		} else if err != nil {
			return nil, err
		} else {
			running, err := p.IsRunningWithContext(ctx)
			if err != nil {
				return nil, err
			}
			t.Running = running
			t.Pid = p.Pid
		}
		tasks[id] = t
	}

	return tasks, nil
}

// ------------------------------------------------------------------------------------------------
// ~ Private methods
// ------------------------------------------------------------------------------------------------

func (g *Gokazi) listProcesses(ctx context.Context) ([]*process.Process, error) {
	ps, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return nil, err
	}

	names := map[string]struct{}{}
	for _, value := range g.tasks {
		names[value.Name] = struct{}{}
	}

	return slices.DeleteFunc(ps, func(p *process.Process) bool {
		name, err := p.NameWithContext(ctx)
		if err != nil {
			return true
		}
		if _, ok := names[name]; !ok {
			return true
		}
		return false
	}), nil
}

func (g *Gokazi) findProcess(ctx context.Context, task config.Task, ps []*process.Process) (*process.Process, error) {
	for _, p := range ps {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		name, err := p.NameWithContext(ctx)
		if err != nil || task.Name != name {
			continue
		}

		exe, err := p.ExeWithContext(ctx)
		if err != nil || (task.ExpandPath() != "" && task.ExpandPath() != path.Dir(exe)) {
			continue
		}

		cwd, err := p.CwdWithContext(ctx)
		if err != nil || (task.ExpandCwd() != "" && task.ExpandCwd() != cwd) {
			continue
		}

		cmdline, err := p.CmdlineSliceWithContext(ctx)
		if err != nil || (len(task.Args) > 0 && len(lo.Intersect(cmdline[1:], task.ExpandArgs())) != len(cmdline[1:])) {
			continue
		}

		return p, nil
	}

	return nil, ErrNotFound
}
