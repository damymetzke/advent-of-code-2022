package output

import (
	"fmt"
	"sync"
	"time"

	. "damymetzke.com/advent-of-code-2022/d17/shared"
)

func queueFrames(input <-chan BoardDisplay) <-chan BoardDisplay {
  output := make(chan BoardDisplay, 5)

  var result []BoardDisplay
  var done bool
  var lock sync.Mutex

  go func(){
    for value := range input {
      lock.Lock()
      result = append(result, value)
      lock.Unlock()
    }

    lock.Lock()
    done = true
    lock.Unlock()
  }()

  go func(){
    var i int
    for {
      lock.Lock()
      if done && i == len(result) {
        lock.Unlock()
        break
      }

      if i >= len(result) {
        lock.Unlock()
        continue
      }
      next := result[i]
      i++
      lock.Unlock()

      output <- next
    }
    close(output)
  }()

  return output
}

// This will output 
func Output(boardInput <-chan BoardDisplay) {
  for i := 0; i < 40; i++ {
    fmt.Println()
  }

	for board := range queueFrames(boardInput) {
    for i := 0; i < 40; i++ {
      fmt.Print("\x1b[1A")
    }
		for _, line := range board {
			for _, value := range line {
				switch value {
        case Empty:
          fmt.Print(".")
        case StandingRock:
          fmt.Print("#")
        case FallingRock:
          fmt.Print("V")
				}
			}
			fmt.Println()
		}

    time.Sleep(1000 * time.Millisecond)
	}
}
