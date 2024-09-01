package lib

import (
	"os"
	"os/exec"
	"syscall"
)

func StartServer(apiEntryPath string) (*os.Process, error) {
  cmd := exec.Command("go", "run", apiEntryPath)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

  err := cmd.Start()
  if err != nil {
    return nil, err
  }

  return cmd.Process, nil
}

func StopServer(process *os.Process) error {
  pgid, err := syscall.Getpgid(process.Pid)
  if err != nil {
    return err
  }
  err = syscall.Kill(-pgid, syscall.SIGTERM)
  if err != nil {
    return err
  }
  _, err = process.Wait()
  return err
}
