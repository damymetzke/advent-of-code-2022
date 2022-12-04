package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type Overlap byte

const (
	NoOverlap   Overlap = 0
	SomeOverlap         = 1
	FullOverlap         = 2
)

func GetInput() string {
	data, err := os.ReadFile("input/4")
	if err != nil {
		log.Fatal("Could not read file 'input/4':\n  * ", err)
	}

	return string(data)
}

func GetOverlapType(ranges string) Overlap {
	// Split up the string, replace works here because format is predictable
	splits := strings.Split(strings.Replace(ranges, ",", "-", 1), "-")
	leftMin, _ := strconv.ParseInt(splits[0], 10, 16)
	leftMax, _ := strconv.ParseInt(splits[1], 10, 16)
	rightMin, _ := strconv.ParseInt(splits[2], 10, 16)
	rightMax, _ := strconv.ParseInt(splits[3], 10, 16)

	// Check if the ranges are fully separated
	if leftMin > rightMax || rightMin > leftMax {
		return NoOverlap
	}

	// Check if any of the ranges is fully contained in the other
	if leftMin >= rightMin && leftMax <= rightMax || leftMin <= rightMin && leftMax >= rightMax {
		return FullOverlap
	}
	return SomeOverlap
}

func main() {
	var waitGroup sync.WaitGroup
	var totalFullOverlapping uint32
	var totalSomeOverlapping uint32

	for _, ranges := range strings.Split(GetInput(), "\n") {
		waitGroup.Add(1)
		rangesValue := ranges
		// More goroutines, getting the hang of this now :)
		go func() {
			result := GetOverlapType(rangesValue)
			if result == NoOverlap {
				goto complete
			}

			if result == FullOverlap {
				atomic.AddUint32(&totalFullOverlapping, 1)
			}

			atomic.AddUint32(&totalSomeOverlapping, 1)

		complete:
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println("=-= PART 1 =-=")
	fmt.Println(totalFullOverlapping)
	fmt.Println("=-= PART 2 =-=")
	fmt.Println(totalSomeOverlapping)
}
