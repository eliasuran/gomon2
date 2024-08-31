package cmd

import (
	"fmt"
	"time"

	"github.com/eliasuran/gomon2/lib"
)

func Run(root string, apiEntryPath string) error {
  fmt.Println("run")

  var initialStats lib.PathStats
  err := lib.GetDirStats(root, &initialStats)
  if err != nil {
    return err
  }
  fmt.Println(initialStats)

  _, err = lib.StartServer(apiEntryPath)
  if err != nil {
    return err
  }

  for {
    fmt.Println("listening cuh")


    time.Sleep(5000 * time.Millisecond)
  }
}
