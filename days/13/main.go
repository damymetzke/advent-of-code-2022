package main

import (
	"fmt"
	"log"
	"os"
)

func GetInput() string {
	data, err := os.ReadFile("input/13")
	if err != nil {
		log.Fatal("Could not read file 'input/13':\n  * ", err)
	}

	return string(data)
}

func main() {
  fmt.Println(GetInput())
	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")
}
