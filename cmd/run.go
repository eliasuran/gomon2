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

  process, err := lib.StartServer(apiEntryPath)
  if err != nil {
    return err
  }

  defer func() {
    if process != nil {
      lib.StopServer(process)   
    }
  }()

  for {
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
        // going of modtime for now
        if item.Path == initialItem.Path && item.Stats.ModTime() != initialItem.Stats.ModTime() {
          fmt.Println("A CHANGE WAS SPOTTED")
          //

          go lib.StopServer(process)

          if process, err = lib.StartServer(apiEntryPath); err != nil {
            return err
          }

          //
          initialStats = stats
          changeFound = true

          break
        }
      }
    }

    time.Sleep(100 * time.Millisecond)
  }
}
