package main

import (
	"fmt"
	"log"
	"os"
)

type DirData struct {
	fileSizes        int64
	childDirectories []string
}

func GetInput() string {
	data, err := os.ReadFile("input/8")
	if err != nil {
		log.Fatal("Could not read file 'input/8':\n  * ", err)
	}

	return string(data)
}

func main() {
  fmt.Println(GetInput())
	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")
}
