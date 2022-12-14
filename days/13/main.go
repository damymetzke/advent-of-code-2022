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

type Data struct {
	isArray bool
	value   int
	array   []Data
}

func GetInput() string {
	data, err := os.ReadFile("input/13")
	if err != nil {
		log.Fatal("Could not read file 'input/13':\n  * ", err)
	}

	return string(data)
}

func ParseNumber(line string, offset int) (int, Data) {
	end := offset
	for line[end] != ',' && line[end] != ']' {
		end++
	}
	parsed, err := strconv.ParseInt(line[offset:end], 10, 64)

	if err != nil {
		log.Fatalf("Could not parse integer from '%v'\n  with offset '%v'", line, offset)
	}

	return end, Data{
		isArray: false,
		value:   int(parsed),
		array:   []Data{},
	}
}

func ParseArray(line string, offset int) (int, Data) {
	result := []Data{}

	for line[offset] != ']' {
		offset++
		if line[offset] == ']' {
			break
		}
		var value Data
		if line[offset] == '[' {
			offset, value = ParseArray(line, offset)

		} else {
			offset, value = ParseNumber(line, offset)
		}

		result = append(result, value)
	}
	return offset + 1, Data{
		isArray: true,
		value:   0,
		array:   result,
	}
}

func ParseLine(line string) Data {
	_, result := ParseArray(line, 0)
	return result
}

func IsInRightOrder(left, right Data) int {
  for i := 0; i < len(left.array) && i < len(right.array); i++ {
    leftValue := left.array[i]
    rightValue := right.array[i]

    // Compare numbers
    if !leftValue.isArray && !rightValue.isArray {
      if leftValue.value == rightValue.value {
        continue
      } else if leftValue.value < rightValue.value {
        return 1
      }
      return -1
    }

    // Convert left value if not array
    if !leftValue.isArray {
      leftValue = Data{
        isArray: true,
        value: 0,
        array: []Data{leftValue},
      }
    }
    // Convert right value if not array
    if !rightValue.isArray {
      rightValue = Data{
        isArray: true,
        value: 0,
        array: []Data{rightValue},
      }
    }

    // Compare arrays
    result := IsInRightOrder(leftValue, rightValue)
    if result == 0 {
      continue
    }
    return result
  }

  // Which one ran out?
  if len(left.array) == len(right.array) {
    return 0
  } else if len(left.array) < len(right.array) {
    return 1
  }
  return -1
}

func main() {
	input := strings.Split(GetInput(), "\n")

	var wait sync.WaitGroup

	left := make([]Data, len(input)/3+1)
	right := make([]Data, len(input)/3+1)

	for i := 0; i < len(input)/3+1; i++ {
		wait.Add(1)
		num := i

		go func() {
			left[num] = ParseLine(input[3*num])
			right[num] = ParseLine(input[3*num+1])
			wait.Done()
		}()
	}

	wait.Wait()

	var result int64

	for i := range left {
    index := int64(i + 1)
		leftValue := left[i]
		rightValue := right[i]
		wait.Add(1)

		go func() {
			if IsInRightOrder(leftValue, rightValue) == 1 {
				atomic.AddInt64(&result, index)
			} else if IsInRightOrder(leftValue, rightValue) == 0 {
      }
			wait.Done()
		}()
	}

  wait.Wait()
	fmt.Println("=-= PART 1 =-=")
  fmt.Println(result)
	fmt.Println("=-= PART 2 =-=")
}
