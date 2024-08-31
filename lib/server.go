package lib

import (
	"os"
	"os/exec"
)

func StartServer(apiEntryPath string) (*os.Process, error) {
  cmd := exec.Command("go", "run", apiEntryPath)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err := cmd.Start()
  if err != nil {
    return nil, err
  }

  return cmd.Process, nil
}
