package lib

import "flag"

func DefineFlags() ([]*bool) {
  helpFlag := flag.Bool("help", false, "Run gomon in the root dir of your go http server")
  // TODO: add flag to allow user to provide a file path to run gomon on 
  // Ex: gomon -f project/api/
  flag.Parse()

  return []*bool{helpFlag}
}
