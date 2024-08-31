package main

import (
	"fmt"

	"github.com/eliasuran/gomon2/cmd"
)

// ALL ERROR HANDLING HERE FOR NOW
func main() {
  root := "."
  apiEntryPath, err := cmd.Init(root)
  if err != nil {
    fmt.Printf("Error in init: %v\n", err)
  }

  err = cmd.Run(root, apiEntryPath)
  if err != nil {
    fmt.Printf("Error in run: %v\n", err)
  }
}
