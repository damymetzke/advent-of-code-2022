package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}
type PositionSet map[Position]struct{}

func GetInput() string {
	data, err := os.ReadFile("input/9")
	if err != nil {
		log.Fatal("Could not read file 'input/9':\n  * ", err)
	}

	return string(data)
}

func getDelta(direction byte) (int, int) {
	switch direction {
	case 'L':
		return -1, 0
	case 'R':
		return 1, 0
	case 'U':
		return 0, -1
	case 'D':
		return 0, 1
	default:
		log.Fatalf("Could not understand direction '%v'", direction)
	}
	return 0, 0
}

func inRange(x, y, tX, tY int) bool {
	dX := tX - x
	dY := tY - y

	if dX < 0 {
		dX *= -1
	}

	if dY < 0 {
		dY *= -1
	}

	return dX <= 1 && dY <= 1
}

func followDelta(x, y, tX, tY int) (int, int) {
	var dX, dY int

	if tX < x {
		dX = 1
	} else if tX > x {
		dX = -1
	}

	if tY < y {
		dY = 1
	} else if tY > y {
		dY = -1
	}

	return dX, dY
}

func main() {
	lines := strings.Split(GetInput(), "\n")

	var x, y, tX, tY int

	visited := PositionSet{{
		x: 0,
		y: 0}: {}}

	for _, line := range lines {
		dX, dY := getDelta(line[0])
		amount, err := strconv.ParseInt(line[2:], 10, 64)
		if err != nil {
			log.Fatalf("Cannot parse line '%v':\n%v", line, err)
		}

		// Movement of tail can be calculated for every multiple orthogonal steps
		// The result will be the same
		x += dX * int(amount)
		y += dY * int(amount)

    // Move the tail until in range
		for !inRange(x, y, tX, tY) {
			fX, fY := followDelta(x, y, tX, tY)
			tX += fX
			tY += fY
			visited[Position{x: tX, y: tY}] = struct{}{}
		}
	}

	fmt.Println("=-= PART 1 =-=")
  fmt.Println(len(visited))
	fmt.Println("=-= PART 2 =-=")
}
