package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(apiEntryPath string) error {
  fmt.Println("run")

  cmd := exec.Command("go", "run", apiEntryPath)
  cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

  err := cmd.Run()
  if err != nil {
    return err
  }

  return nil
}
