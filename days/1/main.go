package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() string {
	data, err := os.ReadFile("input/1")
	if err != nil {
		log.Fatal("Could not read file 'input/1':\n  * ", err)
	}

	return string(data)
}

func main() {
	input := GetInput()

	elves := strings.Split(input, "\n\n")

	maxElves := [3]int{0, 0, 0}
	for _, elve := range elves {
		caloryStrings := strings.Split(elve, "\n")
		elveCalories := 0
		for _, caloryString := range caloryStrings {
      if caloryString == "" {
        continue
      }

			calories, err := strconv.Atoi(caloryString)
			if err != nil {
				log.Fatalf("Could not parse as integer '%v':\n  * %v", caloryString, err)
			}

			elveCalories += calories
		}

    for i := 0; i < 3; i++ {
      if maxElves[i] < elveCalories {
        swap := maxElves[i]
        maxElves[i] = elveCalories
        elveCalories = swap
      }
    }
	}

  fmt.Println("=-= PART 1 =-=")
	fmt.Println(maxElves[0])
  fmt.Println("=-= PART 2 =-=")
  fmt.Println(maxElves[0] + maxElves[1] + maxElves[2])
}
