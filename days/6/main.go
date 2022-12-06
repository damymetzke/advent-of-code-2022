package main

import (
	"fmt"
	"log"
	"os"
)

func GetInput() string {
	data, err := os.ReadFile("input/6")
	if err != nil {
		log.Fatal("Could not read file 'input/6':\n  * ", err)
	}

	return string(data)
}

func areUnique(values [4]byte) bool {
	return values[0] != values[1] &&
		values[0] != values[2] &&
		values[0] != values[3] &&
		values[1] != values[2] &&
		values[1] != values[3] &&
		values[2] != values[3]
}

func main() {

	input := GetInput()

	memory := [4]byte{input[0], input[1], input[2], input[3]}

  var i int
  for i = 4; !areUnique(memory); i++ {
    memory[i % 4] = input[i]
  }

	fmt.Println("=-= PART 1 =-=")
  fmt.Println(i)
	fmt.Println("=-= PART 2 =-=")
}
