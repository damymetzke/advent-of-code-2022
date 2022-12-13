package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
  var result strings.Builder
	if clear {
		for range grid {
			result.WriteString("\x1b[1A")
		}
	}


	for y, row := range grid {
		for x, value := range row {
			colorValue := strconv.FormatInt(int64(200 - ((value - 97) * 5)), 10)
			if visited[y][x] {
        result.WriteString("\x1b[38;2;" + colorValue + ";255;" + colorValue +  "m#\x1b[0m")
			} else {
        result.WriteString("\x1b[38;2;255;" + colorValue + ";" + colorValue +  "m#\x1b[0m")
			}
		}
    result.WriteString("\n")
	}
  fmt.Print(result.String())

	time.Sleep(5 * time.Millisecond)

}

// For part 2 I changed the direction
func FindShortestPath(start, end Position, grid [][]byte, visualize bool) (int, int) {
	var steps int
	currentPositions := []Position{end}
	var nextPositions []Position
  var firstLowest int

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
        if firstLowest == 0 && grid[possible.y][possible.x] == 'a' {
          firstLowest = steps
        }

				// Check for end
				if possible == start {
          if firstLowest == 0 {
            return steps, steps
          }
					return steps, firstLowest
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

	return 0, 0
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

  result1, result2 := FindShortestPath(start, end, grid, false)
	fmt.Println("=-= PART 1 =-=")
  fmt.Println(result1)
	fmt.Println("=-= PART 2 =-=")
  fmt.Println(result2)

	fmt.Println("=-= Visual =-=")
	FindShortestPath(start, end, grid, true)
}
