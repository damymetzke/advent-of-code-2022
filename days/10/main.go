package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() string {
	data, err := os.ReadFile("input/10")
	if err != nil {
		log.Fatal("Could not read file 'input/10':\n  * ", err)
	}

	return string(data)
}

// The first value is the cycle delta
// The second value is the counter delta
func getDelta(command string) (int, int) {
  if command == "noop" {
    return 1, 0;
  }

  counter, err := strconv.ParseInt(command[5:], 10, 64)
  if err != nil {
    log.Fatalf("Could not parse command '%v'", command)
  }
  return 2, int(counter)
}

func main() {
  commands := strings.Split(GetInput(), "\n")

  var signelStrength int
  cycle, counter := 1, 1

  for _, command := range commands {
    cycleDelta, counterDelta := getDelta(command)

    // I have to loop for each cycle delta,
    // in case the signal exists during an addx command
    for i := 0; i < cycleDelta; i++ {
      cycle++
      if i + 1 == cycleDelta {
        counter += counterDelta
      }

      if (cycle + 20) % 40 == 0 {
        fmt.Println(cycle)
        fmt.Println(command)
        signelStrength += counter * cycle
      }
    }
  }

	fmt.Println("=-= PART 1 =-=")
  fmt.Println(signelStrength)
	fmt.Println("=-= PART 2 =-=")
}
