package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func GetInput() string {
	data, err := os.ReadFile("input/9")
	if err != nil {
		log.Fatal("Could not read file 'input/9':\n  * ", err)
	}

	return string(data)
}

func main() {
  fmt.Println(GetInput())
	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")
}
