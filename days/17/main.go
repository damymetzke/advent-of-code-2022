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
	fmt.Println(<-GetInput())
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

  input := make(chan BoardDisplay, 3)
  input <- a
  input <- b
  input <- c
  close(input)

  output.Output(input)

}
