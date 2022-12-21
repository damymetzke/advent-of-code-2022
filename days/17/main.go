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

	input := make(chan StateCollection, 3)

	var board Board

	for i := 0; i < 35; i++ {
		var next [7]bool
		next[i%7] = true
		board = append(board, next)
	}

	input <- StateCollection{
		Board: board,
    PieceType: 1,
    PiecePosition: Position {
      X: 2,
      Y: 37,
    },
	}

  for i := 0; i < 10; i++ {
		var next [7]bool
		next[i%14 / 2] = true
    board = append(board, next)
  }
  input <- StateCollection{
    Board: board,
    PieceType: 1,
    PiecePosition: Position {
      X: 2,
      Y: 47,
    },
  }
  close(input)

  visualChannel := output.StateToDisplay(input)

	// fmt.Println(<-GetInput())
	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")

	fmt.Println("=-= Visual =-=")

	output.Output(visualChannel)

}
