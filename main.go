package main

import (
	"fmt"

	"github.com/eliasuran/gomon2/cmd"
)

// ALL ERROR HANDLING HERE FOR NOW
func main() {
  apiEntryPath, err := cmd.Init()
  if err != nil {
    fmt.Printf("Error in init: %v\n", err)
  }

  err = cmd.Run(apiEntryPath)
  if err != nil {
    fmt.Printf("Error in run: %v\n", err)
  }
}
