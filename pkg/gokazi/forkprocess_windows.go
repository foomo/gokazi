//go:build windows

package gokazi

import (
	"os/exec"
	"syscall"
)

func forkProcess(cmd *exec.Cmd) error {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	return cmd.Start()
}
