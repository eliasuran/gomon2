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

  _, err = lib.StartServer(apiEntryPath)
  if err != nil {
    return err
  }

  for {
    fmt.Println("listening cuh")

    var stats lib.PathStats
    err := lib.GetDirStats(root, &stats)
    if err != nil {
      return err
    }

    changeFound := false

    for _, item := range stats {
      if changeFound {
        break
      }
      for _, initialItem := range initialStats {
        if item.Path == initialItem.Path && item.Stats.ModTime() != initialItem.Stats.ModTime() {
          fmt.Println("A CHANGE WAS SPOTTED")
          initialStats = stats
          changeFound = true
          break
        }
      }
    }

    time.Sleep(1000 * time.Millisecond)
  }
}
