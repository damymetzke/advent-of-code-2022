package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Position struct {
	x int
	y int
}

func GetInput() string {
	data, err := os.ReadFile("input/12")
	if err != nil {
		log.Fatal("Could not read file 'input/12':\n  * ", err)
	}

	return string(data)
}

func GetPossibleDirections(position Position) [4]Position {
	return [4]Position{
		{
			x: position.x - 1,
			y: position.y,
		},
		{
			x: position.x + 1,
			y: position.y,
		},
		{
			x: position.x,
			y: position.y - 1,
		},
		{
			x: position.x,
			y: position.y + 1,
		},
	}
}

func VisualizeGrid(grid [][]byte, visited [][]bool, clear bool) {
	if clear {
		for range grid {
			fmt.Print("\x1b[1A\x1b[K")
		}
	}

	for y, row := range grid {
		for x, value := range row {
			colorValue := 255 - ((value - 97) * 94)
			if visited[y][x] {
				fmt.Printf("\x1b[38;2;%v;255;%vm#\x1b[0m", colorValue, colorValue)
			} else {
				fmt.Printf("\x1b[38;2;255;%v;%vm.\x1b[0m", colorValue, colorValue)
			}
		}
		fmt.Println()
	}

	time.Sleep(5 * time.Millisecond)

}

// For part 2 I changed the direction
func FindShortestPath(start, end Position, grid [][]byte, visualize bool, findLowestInstead bool) int {
	var steps int
	currentPositions := []Position{end}
	var nextPositions []Position

	visited := make([][]bool, len(grid))
	gridWidth := len(grid[0])
	for i := range visited {
		visited[i] = make([]bool, gridWidth)
	}

	if visualize {
		VisualizeGrid(grid, visited, false)
	}

	visited[end.y][end.x] = true

	for {
		if len(currentPositions) == 0 {
			break
		}

		steps++
		for _, next := range currentPositions {
			height := grid[next.y][next.x]
			for _, possible := range GetPossibleDirections(next) {
				// Check bounds
				if possible.x < 0 || possible.x >= gridWidth || possible.y < 0 || possible.y >= len(grid) {
					continue
				}

				// Check if already visited
				if visited[possible.y][possible.x] == true {
					continue
				}

				// Check if unreachable
				if grid[possible.y][possible.x] < height-1 {
					continue
				}

        // Check for lowest
        if findLowestInstead && grid[possible.y][possible.x] == 'a' {
          return steps
        }

				// Check for end
				if possible == start {
					return steps
				}

				// Add next position
				visited[possible.y][possible.x] = true
				nextPositions = append(nextPositions, possible)
			}
		}

		if visualize {
			VisualizeGrid(grid, visited, true)
		}
		currentPositions = nextPositions
		nextPositions = []Position{}
	}

	return 0
}

func main() {
	lines := strings.Split(GetInput(), "\n")
	grid := make([][]byte, len(lines))

	for i, line := range lines {
		grid[i] = []byte(line)
	}

	var start, end Position

	for y, row := range grid {
		for x, value := range row {
			if value == 'S' {
				grid[y][x] = 'a'
				start = Position{
					x,
					y,
				}
			} else if value == 'E' {
				grid[y][x] = 'z'
				end = Position{
					x,
					y,
				}
			}
		}
	}

	fmt.Println("=-= PART 1 =-=")
  fmt.Println(FindShortestPath(start, end, grid, false, false))
	fmt.Println("=-= PART 2 =-=")
  fmt.Println(FindShortestPath(start, end, grid, false, true))

	fmt.Println("=-= Visual 1 =-=")
	FindShortestPath(start, end, grid, true, false)
  fmt.Println("=-= Visual 2 =-=")
	FindShortestPath(start, end, grid, true, true)
}
