package cmd

import (
	"fmt"

	"github.com/eliasuran/gomon2/lib"
)

func Init() (string, error) {
  fmt.Println("init")
  fmt.Print(`
 ██████╗  ██████╗ ███╗   ███╗ ██████╗ ███╗   ██╗
██╔════╝ ██╔═══██╗████╗ ████║██╔═══██╗████╗  ██║
██║  ███╗██║   ██║██╔████╔██║██║   ██║██╔██╗ ██║
██║   ██║██║   ██║██║╚██╔╝██║██║   ██║██║╚██╗██║
╚██████╔╝╚██████╔╝██║ ╚═╝ ██║╚██████╔╝██║ ╚████║
 ╚═════╝  ╚═════╝ ╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝ 2

`)

  flags := lib.DefineFlags()
  for _, flag := range flags {
    if *flag {
      // TODO: read out the flag value if the flag was provided
      fmt.Println(flag)
    }
  }

  // get root from user if flag
  root := "."
  entry, err := lib.FindApiEntry(root)
  if err != nil {
    return "", err
  }
  fmt.Println(entry)
  return entry, nil
}
