package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type DirData struct {
	fileSizes        int64
	childDirectories []string
}

func GetInput() string {
	data, err := os.ReadFile("input/7")
	if err != nil {
		log.Fatal("Could not read file 'input/7':\n  * ", err)
	}

	return string(data)
}

func setDir(dirs *map[string]DirData, cwd []string, dirName string) {
	key := strings.Join(cwd, "/")
	data, exists := (*dirs)[key]
	if !exists {
		data = DirData{}
	}
	data.childDirectories = append(data.childDirectories, dirName)
	(*dirs)[key] = data
}

func addSize(dirs *map[string]DirData, cwd []string, size int64) {
	key := strings.Join(cwd, "/")
	data, exists := (*dirs)[key]
	if !exists {
		data = DirData{}
	}
  data.fileSizes += size
	(*dirs)[key] = data
}

func EvaluateDir(dirs *map[string]DirData, cwd []string) int64 {
  key := strings.Join(cwd, "/")
  data := (*dirs)[key]
  for _, childDir := range data.childDirectories {
    data.fileSizes += EvaluateDir(dirs, append(cwd, childDir))
  }

  (*dirs)[key] = data

  return data.fileSizes
}

func findSmallestDir(dirs *map[string]DirData, key string, required int64) (bool, int64) {
  data := (*dirs)[key]
  if data.fileSizes < required {
    return false, 0
  }

  minResult := false
  minSize := int64(0)

  for _, childDir := range data.childDirectories {
    result, size := findSmallestDir(dirs, key + "/" + childDir, required)
    if !result {
      continue
    }
    minResult = true
    if size < minSize {
      minSize = size
    }
  }

  if !minResult {
    return true, data.fileSizes
  }

  return true, minSize
}

func main() {
  dirs := map[string]DirData{}

	input := strings.Split(GetInput(), "\n")

	// Keep track of the CWD
	currentWorkingDir := []string{}

	// Run each command
	for _, line := range input {
		// ls is always follwed by either dirs or files.
		// I can safely ignore the command and look at the other patterns
		if line == "$ ls" {
			continue
		}

		// Change directory
		if line[0] == '$' && line[:5] == "$ cd " {
			// Back to root
			if line[5:] == "/" {
				currentWorkingDir = []string{}
			} else if line[5:] == ".." {
				currentWorkingDir = currentWorkingDir[:len(currentWorkingDir)-1]
			} else {
				currentWorkingDir = append(currentWorkingDir, line[5:])
			}
      continue
		}

		// From this point on all the values are output of the ls command

		// This is a dir, so mark it
		if line[:4] == "dir " {
			setDir(&dirs, currentWorkingDir, line[4:])
      continue
		}

    // Add the file size to the total
    size, err := strconv.ParseInt(strings.Split(line, " ")[0], 10, 64)
    if err != nil {
      log.Fatalf("Error:\n%v", err)
    }

    addSize(&dirs, currentWorkingDir, size)
	}

  var total, min int64
  required := EvaluateDir(&dirs, []string{}) - 40000000

  for _, dir := range dirs {
    size := dir.fileSizes
    if size <= 100000 {
      total += size
    }
    if size > required {
      if size < min || min == 0 {
        min = size
      }
    }
  }

	fmt.Println("=-= PART 1 =-=")
  fmt.Println(total)
	fmt.Println("=-= PART 2 =-=")
  fmt.Println(min)
}
