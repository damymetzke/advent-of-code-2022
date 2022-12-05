package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() string {
	data, err := os.ReadFile("input/5")
	if err != nil {
		log.Fatal("Could not read file 'input/5':\n  * ", err)
	}

	return string(data)
}

func main() {
	lines := strings.Split(GetInput(), "\n")
	var offset uint
	var topLines []string

	for i, line := range lines {
		if line == "" {
			offset = uint(i + 1)
			break
		}
		topLines = append(topLines, line)
	}

  width := (len(topLines[0])+1)/4
	top := make([][]rune, width)



  for i := len(topLines) - 2; i >= 0; i-- {
    line := topLines[i]
    for j := 0; j < width; j++ {
      item := line[4 * j + 1]
      if item == ' ' {
        continue
      }

      top[j] = append(top[j], rune(item))
    }
  }

  fmt.Println(offset)

	for _, line := range lines[offset:] {
    if line == "" {
      continue
    }
    parts := strings.Split(line, " ")

    amount, amountErr := strconv.ParseInt(parts[1], 10, 8)
    from, fromErr := strconv.ParseInt(parts[3], 10, 8)
    to, toErr := strconv.ParseInt(parts[5], 10, 8)

    if amountErr != nil || fromErr != nil || toErr != nil {
      log.Panicf("Parse error\n%v\n%v\n%v", amountErr, fromErr, toErr)
    }

    from -= 1
    to -= 1

    //fmt.Printf("(%v), %v -> %v\n", amount, from, to)
    for i:= int64(0); i < amount; i++ {
      top[to] = append(top[to], top[from][len(top[from]) - 1 - int(i)])
    }

    top[from] = top[from][:len(top[from]) - int(amount)]
	}

	fmt.Println("=-= PART 1 =-=")
  for _, stack := range top {
    fmt.Print(string(stack[len(stack) - 1]))
  }
  fmt.Println()
	fmt.Println("=-= PART 2 =-=")
}
