package gokazi

import (
	"context"
	"log/slog"
	"os/exec"
	"path"
	"slices"
	"time"

	gotime "github.com/foomo/go/time"
	"github.com/foomo/gokazi/pkg/config"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shirou/gopsutil/v4/process"
)

// Sentinel errors returned by [Gokazi] methods. They are wrapped with
// the task identifier and may be tested with [errors.Is].
var (
	// ErrNotFound is returned when no task with the given ID has been
	// registered via [Gokazi.Add].
	ErrNotFound = errors.New("task not found")
	// ErrNotRunning is returned when the task exists but no matching
	// process is currently running, or when a started process did not
	// become visible within the wait deadline.
	ErrNotRunning = errors.New("task not running")
	// ErrMultipleFound is returned when more than one running process
	// matches the same task definition. The matcher cannot disambiguate
	// them, so the operation is refused.
	ErrMultipleFound = errors.New("multiple tasks found")
	// ErrAlreadyRunning is returned by [Gokazi.Start] when a process
	// matching the task is already running.
	ErrAlreadyRunning = errors.New("task already running")
)

// Gokazi is a daemonless process manager. It holds a set of registered
// task definitions and answers list, find, start, and stop queries by
// scanning the operating-system process table on each call.
//
// A Gokazi is not safe for concurrent use.
type Gokazi struct {
	l     *slog.Logger
	tasks map[string]config.Task
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

// New returns a Gokazi with no registered tasks. The logger is used for
// debug output during process operations.
func New(l *slog.Logger) *Gokazi {
	return &Gokazi{
		l:     l,
		tasks: map[string]config.Task{},
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

// Add registers task under id. Adding a task with an existing id
// replaces the previous definition.
func (g *Gokazi) Add(id string, task config.Task) {
	g.tasks[id] = task
}

// Start launches cmd as a detached child process for the task
// identified by id. The child is placed in its own process group so it
// survives the parent. Start waits up to one second for a matching
// process to appear in the OS process table; on timeout it returns
// [ErrNotRunning]. It returns [ErrAlreadyRunning] when a process
// matching the task is already running, [ErrNotFound] when id is not
// registered, and any error from forking the process.
func (g *Gokazi) Start(ctx context.Context, id string, cmd *exec.Cmd) error {
	t, err := g.Find(ctx, id)
	if err != nil {
		return err
	}

	if t.Running {
		return ErrAlreadyRunning
	}

	g.l.Debug("Starting: " + cmd.String())

	if err := forkProcess(cmd); err != nil {
		return err
	}

	if err := g.WaitForRunning(ctx, id, time.Second); err != nil {
		return ErrNotRunning
	}

	return nil
}

// Stop terminates the running process for the task identified by id by
// sending it a kill signal. It returns [ErrNotFound] when id is not
// registered, [ErrNotRunning] when no matching process is alive, and
// [ErrMultipleFound] when more than one running process matches the
// task definition. The kill itself has a five-second deadline.
func (g *Gokazi) Stop(ctx context.Context, id string) error {
	task, ok := g.tasks[id]
	if !ok {
		return errors.Wrap(ErrNotFound, id)
	}

	ps, err := g.listProcesses(ctx)
	if err != nil {
		return err
	}

	found, err := g.findProcess(ctx, task, ps)
	if errors.Is(err, ErrNotFound) {
		return errors.Wrap(ErrNotRunning, id)
	} else if err != nil {
		return err
	}

	if len(found) > 1 {
		return errors.Wrap(ErrMultipleFound, id)
	}

	p := found[0]

	running, err := p.IsRunningWithContext(ctx)
	if err != nil {
		return errors.Wrap(err, id)
	} else if !running {
		return errors.Wrap(ErrNotRunning, id)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return p.KillWithContext(ctx)
}

// Find returns the current state of the task identified by id. The
// returned [Task] embeds the configured definition and reports the
// process ID and running flag observed from the OS process table.
// Find returns [ErrNotFound] when id is not registered and
// [ErrMultipleFound] when more than one process matches the task
// definition. When no process matches, Find returns the task with
// Running=false and no error.
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

	found, err := g.findProcess(ctx, task, ps)
	if errors.Is(err, ErrNotFound) {
		return t, nil
	} else if err != nil {
		return Task{}, err
	}

	if len(found) > 1 {
		return Task{}, errors.Wrap(ErrMultipleFound, id)
	}

	p := found[0]

	running, err := p.IsRunningWithContext(ctx)
	if err != nil {
		return Task{}, err
	}

	t.Running = running
	t.Pid = p.Pid

	return t, nil
}

// List returns the current state of every registered task, keyed by
// task id. Tasks with no matching process are returned with
// Running=false and Pid=0. List returns [ErrMultipleFound] when more
// than one process matches the same task definition.
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

		found, err := g.findProcess(ctx, task, ps)
		if errors.Is(err, ErrNotFound) {
			tasks[id] = t
			continue
		} else if err != nil {
			return nil, err
		}

		if len(found) > 1 {
			return nil, errors.Wrap(ErrMultipleFound, id)
		}

		p := found[0]

		running, err := p.IsRunningWithContext(ctx)
		if err != nil {
			return nil, err
		}

		t.Running = running
		t.Pid = p.Pid

		tasks[id] = t
	}

	return tasks, nil
}

// WaitForRunning blocks until the task identified by id has a matching
// running process, the context is cancelled, or timeout elapses. It
// polls at 100ms intervals. The error from the underlying scan is
// propagated; on timeout the caller observes the context error.
func (g *Gokazi) WaitForRunning(ctx context.Context, id string, timeout time.Duration) error {
	return gotime.WaitFor(ctx, func(ctx context.Context) (bool, error) {
		t, err := g.Find(ctx, id)
		if err != nil {
			return false, err
		}

		return t.Running, nil
	}, timeout, 100*time.Millisecond)
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

func (g *Gokazi) findProcess(ctx context.Context, task config.Task, ps []*process.Process) ([]*process.Process, error) {
	var found []*process.Process

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

		if len(task.Args) > 0 {
			cmdline, err := p.CmdlineSliceWithContext(ctx)
			if err != nil || len(lo.Intersect(cmdline[1:], task.ExpandArgs())) != len(task.Args) {
				continue
			}
		}

		found = append(found, p)
	}

	if len(found) == 0 {
		return nil, ErrNotFound
	}

	return found, nil
}
