package output

import (
	"fmt"
	"time"

	. "damymetzke.com/advent-of-code-2022/d17/shared"
)

func Output(boardInput <-chan BoardDisplay) {
	for board := range boardInput {
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

    fmt.Println()

    time.Sleep(1000 * time.Millisecond)
	}
}
