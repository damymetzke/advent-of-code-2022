package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

func GetInput() string {
	data, err := os.ReadFile("input/3")
	if err != nil {
		log.Fatal("Could not read file 'input/3':\n  * ", err)
	}

	return string(data)
}

// This function will calculate the priority of a specific string
func CalculatePriority(input string) uint32 {

  // Loop throug left side
	for i := 0; i < len(input)/2; i++ {
    // Loop throug right side
		for j := len(input) / 2; j < len(input); j++ {
			if input[i] == input[j] {
        // Convert ascii to priority
				if input[i] >= 97 && input[i] < 123 {
					return uint32(input[i] - 96)
				} else if input[i] >= 65 && input[i] < 91 {
					return uint32(input[i] - 38)
				} else {
					log.Fatalf("Invalid character '%v'", input[i])
				}
			}
		}
	}

	log.Fatalf("Valid string, shouldn't be valid '%v'", input)
	return 0
}

func main() {
	var result uint32

	var waitGroup sync.WaitGroup

  // Go though all rucksacks
	for _, rucksack := range strings.Split(GetInput(), "\n") {
		waitGroup.Add(1)
    value := rucksack
    // Decided to try using goroutines
    // This went as well as you would expect
		go func() {
			atomic.AddUint32(&result, CalculatePriority(value))
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println("=-= PART 1 =-=")
	fmt.Println(result)
	fmt.Println("=-= PART 2 =-=")
}
