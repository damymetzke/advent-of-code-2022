package main

import (
	"fmt"
	"log"
	"os"
)

func GetInput() string {
	data, err := os.ReadFile("input/4")
	if err != nil {
		log.Fatal("Could not read file 'input/4':\n  * ", err)
	}

	return string(data)
}

func main() {
	fmt.Println(GetInput())
	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")
}
