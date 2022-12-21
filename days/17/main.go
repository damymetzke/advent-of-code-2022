package main

import (
	"fmt"
	"log"
	"os"

	"damymetzke.com/advent-of-code-2022/d17/output"
	. "damymetzke.com/advent-of-code-2022/d17/shared"
)

func GetInput() <-chan string {
	result := make(chan string, 1)
	go func() {
		data, err := os.ReadFile("input/17")
		if err != nil {
			log.Fatal("Could not read file 'input/17':\n  * ", err)
		}

		result <- string(data)
		close(result)
	}()
	return result
}

func main() {
	// fmt.Println(<-GetInput())
	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")

	fmt.Println("=-= Visual =-=")

	var a, b, c BoardDisplay

	for i := 0; i < 40; i++ {
		a[i] = [7]DisplayType{
			StandingRock, Empty, Empty, FallingRock, Empty, Empty, StandingRock,
		}

		b[i] = [7]DisplayType{
			StandingRock, Empty, FallingRock, FallingRock, FallingRock, Empty, StandingRock,
		}

		c[i] = [7]DisplayType{
			StandingRock, Empty, FallingRock, Empty, FallingRock, Empty, StandingRock,
		}
	}

	input := make(chan StateCollection, 3)

	var board Board

	for i := 0; i < 35; i++ {
		var next [7]bool
		next[i%7] = true
		board = append(board, next)
	}

	input <- StateCollection{
		Board: board,
	}

  for i := 0; i < 10; i++ {
		var next [7]bool
		next[i%14 / 2] = true
    board = append(board, next)
  }
  input <- StateCollection{
    Board: board,
  }
  close(input)

	output.Output(output.StateToDisplay(input))

}
