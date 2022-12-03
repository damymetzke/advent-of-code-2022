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

func CalculatePriorityFromAscii(value byte) uint32 {
	if value >= 97 && value < 123 {
		return uint32(value - 96)
	} else if value >= 65 && value < 91 {
		return uint32(value - 38)
	}
	log.Fatalf("Invalid character '%v'", value)
	return 0
}

// This function will calculate the priority of a specific string
func CalculatePriority(input string) uint32 {

	// Loop throug left side
	for i := 0; i < len(input)/2; i++ {
		// Loop throug right side
		for j := len(input) / 2; j < len(input); j++ {
			if input[i] == input[j] {
				// Convert ascii to priority
        return CalculatePriorityFromAscii(input[i])
			}
		}
	}

	log.Fatalf("Valid string, shouldn't be valid '%v'", input)
	return 0
}

func CalculateGroupPriority(first string, second string, third string) uint32 {
	// Loop through first
	for _, firstValue := range first {
		for _, secondValue := range second {
			if firstValue == secondValue {
				for _, thirdvalue := range third {
					if firstValue == thirdvalue {
						return CalculatePriorityFromAscii(byte(firstValue))
					}
				}
				break
			}
		}
	}
	log.Fatalf("No common item '%v' '%v' '%v'", first, second, third)
	return 0
}

func main() {
	var result uint32
	var groupResult uint32

	var waitGroup sync.WaitGroup

	// Go though all rucksacks
	var first string
	var second string

	for i, rucksack := range strings.Split(GetInput(), "\n") {
		// Decided to try using goroutines
		// This went as well as you would expect
		waitGroup.Add(1)
		value := rucksack
		go func() {
			atomic.AddUint32(&result, CalculatePriority(value))
			waitGroup.Done()
		}()

		// Check 3 at the time using modulo logic
		if i%3 == 0 {
			first = rucksack
		} else if i%3 == 1 {
			second = rucksack
		} else {
			waitGroup.Add(1)
      firstValue := first
      secondValue := second
			go func() {
				atomic.AddUint32(&groupResult, CalculateGroupPriority(firstValue, secondValue, value))
				waitGroup.Done()
			}()
		}
	}

	waitGroup.Wait()

	fmt.Println("=-= PART 1 =-=")
	fmt.Println(result)
	fmt.Println("=-= PART 2 =-=")
	fmt.Println(groupResult)
}
