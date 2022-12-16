package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	name  string
	rate  int
	edges []string
}

func GetInput() string {
	data, err := os.ReadFile("input/16")
	if err != nil {
		log.Fatal("Could not read file 'input/16':\n  * ", err)
	}

	return string(data)
}

func ParseLine(line string) Valve {
  parts := strings.Split(line, "; ")
  name := parts[0][6:8]
  rate, rateErr := strconv.ParseInt(parts[0][23:], 10, 64)

  end := strings.Split(parts[1], "valve")[1]
  if end[0] == 's' {
    end = end[2:]
  } else {
    end =end[1:]
  }

  edges := strings.Split(end, ", ")

  if rateErr != nil {
    log.Fatalf("Could not parse line '%v'", line)
  }
  return Valve {
    name: name,
    rate: int(rate),
    edges: edges,
  }
}

func main() {
  lines := strings.Split(GetInput(), "\n")

  valveChannel := make(chan Valve)

  for _, line := range lines {
    value := line
    go func(){
      valveChannel <- ParseLine(value)
    }()
  }

  valves := make([]Valve, len(lines))

  for i := 0; i < len(valves); i++ {
    valves[i] = <- valveChannel
    fmt.Println(valves[i])
  }



	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")
}
