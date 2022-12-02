package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetInput() string {
	data, err := os.ReadFile("input/2")
	if err != nil {
		log.Fatal("Could not read file 'input/2':\n  * ", err)
	}

	return string(data)
}

func CalculateScore(round string) (int, int) {
	opponentMoveStr := round[0]
	myMoveStr := round[2]

	myMove := 0
	opponentMove := 0
  myMoveAlt := 0

	result := 0
  resultAlt := 0

	switch opponentMoveStr {
	case 'A':
		opponentMove = 0
	case 'B':
		opponentMove = 1
	case 'C':
		opponentMove = 2
	default:
		log.Fatalf("Cound not understand opponentMove '%v'", myMoveStr)
	}

	switch myMoveStr {
	case 'X':
		myMove = 0
    myMoveAlt = (opponentMove + 2) % 3
	case 'Y':
		myMove = 1
    myMoveAlt = opponentMove
	case 'Z':
		myMove = 2
    myMoveAlt = (opponentMove + 1) % 3
	default:
		log.Fatalf("Cound not understand myMove '%v'", myMoveStr)
	}

  result += myMove + 1
  resultAlt += myMoveAlt + 1

  if opponentMove == myMove {
    result += 3
  }
  if (opponentMove + 1) % 3 == myMove {
    result += 6
  }

  if opponentMove == myMoveAlt {
    resultAlt += 3
  }
  if (opponentMove + 1) % 3 == myMoveAlt {
    resultAlt += 6
  }

	return result, resultAlt
}

func main() {

  input := strings.Split(GetInput(), "\n")

  total := 0
  alternativeTotal := 0

  for _, round := range input {
    normal, alternative := CalculateScore(round)
    total += normal
    alternativeTotal += alternative
  }

	fmt.Println("=-= PART 1 =-=")

  fmt.Println(total)

	fmt.Println("=-= PART 2 =-=")

  fmt.Println(alternativeTotal)
}
