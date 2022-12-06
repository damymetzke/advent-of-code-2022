package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

func areUniqueSorted(values []byte) bool {
	for i := 1; i < len(values); i++ {
		if values[i-1] == values[i] {
			return false
		}
	}
	return true
}

func main() {

	input := GetInput()

	memory := [4]byte{input[0], input[1], input[2], input[3]}
	memoryMany := [14]byte{
		input[0],
		input[1],
		input[2],
		input[3],
		input[4],
		input[5],
		input[6],
		input[7],
		input[8],
		input[9],
		input[10],
		input[11],
		input[12],
		input[13]}

	var i int
	for i = 4; !areUnique(memory); i++ {
		memory[i%4] = input[i]
		memoryMany[i%14] = input[i]
	}

	startPacket := i

	sortedMemory := make([]byte, 14)
	copy(sortedMemory, memoryMany[:])
	sort.Slice(sortedMemory, func(i, j int) bool {
		return sortedMemory[i] < sortedMemory[j]
	})

  // This is overly optinized
  // What I'm doing here is dynamically updating the sorted array
  // The circular buffer memoryMany can be used to find the value that must be added,
  // and the value that must be removed
  // The elements in between the inserted and removed value will be shifted
  // Don't solve problems in this way (at least not at first),
  // This took way too long to implement
	for ; !areUniqueSorted(sortedMemory); i++ {
		last := memoryMany[i%14]
		next := input[i]
		memoryMany[i%14] = next

		// In this case nothing should change
		if last == next {
			continue
		}

		if last < next {
			var replacing bool
			for j := 0; j < 14; j++ {
				if replacing {
					if j == 13 || next < sortedMemory[j+1] {
						sortedMemory[j] = next
						break
					}
          sortedMemory[j] = sortedMemory[j+1]
				} else {
					if sortedMemory[j] == last {
						if j == 13 || next < sortedMemory[j+1] {
							sortedMemory[j] = next
							break
						}
						replacing = true
						sortedMemory[j] = sortedMemory[j+1]
					}
				}
			}

			continue
		}

		var shifting bool
		var shiftMemory byte
		for j := 0; j < 14; j++ {
			if shifting {
				if shiftMemory == last {
					break
				}

				sortedMemory[j], shiftMemory = shiftMemory, sortedMemory[j]

			} else {
				if next < sortedMemory[j] {
					shiftMemory = sortedMemory[j]
					sortedMemory[j] = next
          shifting = true
				}
			}
		}
	}

	fmt.Println("=-= PART 1 =-=")
	fmt.Println(startPacket)
	fmt.Println("=-= PART 2 =-=")
	fmt.Println(i)
}
