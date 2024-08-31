package lib

import (
	"io/fs"
	"os"
	"strings"
)

type PathStats []PathStat

type PathStat struct {
  Path string
  Stats fs.FileInfo
}

// GET STATS TO LOOK FOR CHANGES
func GetDirStats(root string, pathStats *PathStats) error {
  entries, err := os.ReadDir(root)
  if err != nil {
    return err
  }

  for _, entry := range entries {
    name := entry.Name()
    path := root + "/" + name

    if isHidden(name) {
      continue
    }
    if entry.IsDir() {
      GetDirStats(path, pathStats)
    }
    if !isGoFile(path) {
      continue
    }

    stats, err := os.Stat(path)
    if err != nil {
      return err
    }
    pathStat := PathStat{ Path: path, Stats: stats }
    *pathStats = append(*pathStats, pathStat)
  }

  return nil
}

// FIND MAIN FUNCTION
func FindApiEntry(root string) (string, error) {
  var apiEntry string

  entries, err := os.ReadDir(root)
  if err != nil {
    return "", err
  }

  for _, entry := range entries {
    name := entry.Name()
    path := root + "/" + name

    if isHidden(name) {
      continue
    }
    if entry.IsDir() {
      FindApiEntry(path)
    }
    if !isGoFile(path) {
      continue
    }

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
