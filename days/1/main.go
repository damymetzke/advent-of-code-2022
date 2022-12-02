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

	max := 0
	for _, elve := range elves {
		caloryStrings := strings.Split(elve, "\n")
		total := 0
		for _, caloryString := range caloryStrings {
      if caloryString == "" {
        continue
      }

			calories, err := strconv.Atoi(caloryString)
			if err != nil {
				log.Fatalf("Could not parse as integer '%v':\n  * %v", caloryString, err)
			}

			total += calories
		}

		if max < total {
			max = total
		}
	}

	fmt.Println(max)
}
