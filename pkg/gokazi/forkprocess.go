//go:build !windows

package gokazi

import (
	"os/exec"
	"syscall"
)

func forkProcess(cmd *exec.Cmd) error {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
		Pgid:    0,
	}
	return cmd.Start()
}
