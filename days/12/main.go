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
  direction byte
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
      direction: '>',
		},
		{
			x: position.x + 1,
			y: position.y,
      direction: '<',
		},
		{
			x: position.x,
			y: position.y - 1,
      direction: 'V',
		},
		{
			x: position.x,
			y: position.y + 1,
      direction: '^',
		},
	}
}

func VisualizeGrid(grid [][]byte, visited [][]bool, display [][]byte, clear bool) {
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
        result.WriteString("\x1b[38;2;" + colorValue + ";255;" + colorValue +  "m" + string(display[y][x]) + "\x1b[0m")
			} else {
        result.WriteString("\x1b[38;2;255;" + colorValue + ";" + colorValue +  "m#\x1b[0m")
			}
		}
    result.WriteString("\n")
	}
  fmt.Print(result.String())

	time.Sleep(5 * time.Millisecond)

}

func VisualizePath(grid [][]byte, visited [][]bool, display [][]byte, start Position) {
	incrementalDisplay := make([][]byte, len(grid))
	for i := range incrementalDisplay {
		incrementalDisplay[i] = make([]byte, len(grid[0]))
    for j := range incrementalDisplay[i] {
      incrementalDisplay[i][j] = '#'
    }
  }
  currentPosition := start
  incrementalDisplay[currentPosition.y][currentPosition.x] = display[currentPosition.y][currentPosition.x]
  for(display[currentPosition.y][currentPosition.x]) != 'E' {
    VisualizeGrid(grid, visited, incrementalDisplay, true)

    switch display[currentPosition.y][currentPosition.x] {
    case '>':
      currentPosition.x += 1
    case '<':
      currentPosition.x -= 1
    case 'V':
      currentPosition.y += 1
    case '^':
      currentPosition.y -= 1
    }
    incrementalDisplay[currentPosition.y][currentPosition.x] = display[currentPosition.y][currentPosition.x]
  }
}

// For part 2 I changed the direction
func FindShortestPath(start, end Position, grid [][]byte, visualize bool) (int, int) {
	var steps int
	currentPositions := []Position{end}
	var nextPositions []Position
  var firstLowest int

	visited := make([][]bool, len(grid))
	display := make([][]byte, len(grid))
	staticDisplay := make([][]byte, len(grid))
	gridWidth := len(grid[0])
	for i := range visited {
		visited[i] = make([]bool, gridWidth)
    if !visualize {
      continue
    }

		display[i] = make([]byte, gridWidth)
		staticDisplay[i] = make([]byte, gridWidth)
    for j := range staticDisplay[i] {
      display[i][j] = '#'
      staticDisplay[i][j] = '#'
    }
	}


	if visualize {
    display[end.y][end.x] = 'E'
    display[start.y][start.x] = 'S'
		VisualizeGrid(grid, visited, staticDisplay, false)
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
				if possible.x == start.x && possible.y == start.y {
          if visualize {
            display[possible.y][possible.x] = possible.direction
            VisualizePath(grid, visited, display, start)
          }
          if firstLowest == 0 {
            return steps, steps
          }
					return steps, firstLowest
				}

				// Add next position
				visited[possible.y][possible.x] = true
				nextPositions = append(nextPositions, possible)
        if !visualize {
          continue
        }

        display[possible.y][possible.x] = possible.direction
			}
		}

		if visualize {
			VisualizeGrid(grid, visited, staticDisplay, true)
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
          'S',
				}
			} else if value == 'E' {
				grid[y][x] = 'z'
				end = Position{
					x,
					y,
          'E',
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
