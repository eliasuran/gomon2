package lib

import (
	"os"
	"strings"
)

// FIND MAIN FUNCTION
func FindApiEntry(root string) (string, error) {
  var apiEntry string

  entries, err := os.ReadDir(root)
  if err != nil {
    return "", err
  }

  for _, entry := range entries {
    name := entry.Name()
    if isHidden(name) {
      continue
    }
    if entry.IsDir() {
      FindApiEntry(name)
    }
    if !isGoFile(name) {
      continue
    }
    path := root + "/" + name
    contents, err := os.ReadFile(path)
    if err != nil {
      return "", err
    }
    if strings.Contains(string(contents), "func main()") {
      apiEntry = path
    }
  }

  return apiEntry, nil
}

func isGoFile(name string) bool {
  // TODO: probably a better way to do this
  if strings.Contains(name, ".go") {
    return true
  }
  return false
}

func isHidden(name string) bool {
  if string(name[0]) == "." {
    return true
  }
  return false
}
