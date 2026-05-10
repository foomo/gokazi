// Package gokazi is the daemonless process manager that powers the
// gokazi binary.
//
// There is no long-running supervisor and no PID file. Each operation
// enumerates the live operating-system processes via gopsutil and
// matches every configured [github.com/foomo/gokazi/pkg/config.Task]
// against them by name, executable path, working directory, and
// arguments. See [config.Task.Match] for the matching contract.
//
// Typical use is to construct a [Gokazi] with [New], register tasks
// with [Gokazi.Add], and then call [Gokazi.List], [Gokazi.Find],
// [Gokazi.Start], or [Gokazi.Stop]. Errors are wrapped with the task
// identifier; the sentinel values [ErrNotFound], [ErrNotRunning],
// [ErrMultipleFound], and [ErrAlreadyRunning] can be tested with
// [errors.Is].
//
// Start uses a platform-specific fork that detaches the child process
// group: on Unix the child is placed in its own process group via
// setpgid; on Windows it is created with CREATE_NEW_PROCESS_GROUP.
// The parent does not retain the child; subsequent List or Stop calls
// rediscover it through the OS process table.
package gokazi
