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

func GetInput() string {
	data, err := os.ReadFile("input/4")
	if err != nil {
		log.Fatal("Could not read file 'input/4':\n  * ", err)
	}

	return string(data)
}

func IsOverlapping(ranges string) bool {
	// Split up the string, replace works here because format is predictable
	splits := strings.Split(strings.Replace(ranges, ",", "-", 1), "-")
	leftMin, _ := strconv.ParseInt(splits[0], 10, 16)
	leftMax, _ := strconv.ParseInt(splits[1], 10, 16)
	rightMin, _ := strconv.ParseInt(splits[2], 10, 16)
	rightMax, _ := strconv.ParseInt(splits[3], 10, 16)

	// Check if any of the ranges is fully contained in the other
	return leftMin >= rightMin && leftMax <= rightMax || leftMin <= rightMin && leftMax >= rightMax
}

func main() {
	var waitGroup sync.WaitGroup
	var totalOverlapping uint32

	for _, ranges := range strings.Split(GetInput(), "\n") {
		waitGroup.Add(1)
		rangesValue := ranges
		// More goroutines, getting the hang of this now :)
		go func() {
			if IsOverlapping(rangesValue) {
				atomic.AddUint32(&totalOverlapping, 1)
			}
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println("=-= PART 1 =-=")
	fmt.Println(totalOverlapping)
	fmt.Println("=-= PART 2 =-=")
}
