package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// A direction to check in
type LookDirection struct {
  // The amount of trees to check,
  // used for simple optimisation
	amountToCheck int
  // Deltas will be a combination of 0 and 1/-1
  // This will be used to iterate orthogonally through the grid
	deltaX        int
	deltaY        int
}

func GetInput() string {
	data, err := os.ReadFile("input/8")
	if err != nil {
		log.Fatal("Could not read file 'input/8':\n  * ", err)
	}

	return string(data)
}

// Thank you ChatGPT
func convertTo2DByteArray(input string) [][]byte {
	// Split the input string into a slice of strings, where
	// each string represents a row in the grid.
	rows := strings.Split(input, "\n")

	// Create a 2D byte array with the same dimensions as the
	// input grid.
	grid := make([][]byte, len(rows))
	for i := range grid {
		grid[i] = make([]byte, len(rows[i]))
	}

	// Iterate over the rows and convert each string into a
	// slice of bytes.
	for i, row := range rows {
		for j := range row {
			// ChatGPT tried to parse the int, I don't care so I removed the parsing
			// The values should be ordered the same
			grid[i][j] = row[j]
		}
	}

	return grid
}

func incrementForDirection(x *int, y *int, direction LookDirection) {
	*x += direction.deltaX
	*y += direction.deltaY
}

func main() {
	grid := convertTo2DByteArray(GetInput())
	// Width and height, easy optimisation
	width := len(grid[0])
	height := len(grid)

	var result int

	for i, row := range grid {
		for j, tree := range row {
			directions := []LookDirection{
				{
					amountToCheck: j,
					deltaX:        -1,
					deltaY:        0,
				},
				{
					amountToCheck: width - j - 1,
					deltaX:        1,
					deltaY:        0,
				},
				{
					amountToCheck: i,
					deltaX:        0,
					deltaY:        -1,
				},
				{
					amountToCheck: height - i - 1,
					deltaX:        0,
					deltaY:        1,
				},
			}

			sort.Slice(directions, func(i, j int) bool {
				return directions[i].amountToCheck < directions[j].amountToCheck
			})

      // Test all directions
		labelDirection:
			for _, direction := range directions {
				for y, x := i+direction.deltaY, j+direction.deltaX; y >= 0 && y < height && x >= 0 && x < width; incrementForDirection(&x, &y, direction) {
          // In this case a tree is blocking the view
          // So try the next direction
					if grid[y][x] >= tree {
						continue labelDirection
					}
				}

        // If at least 1 loop was successful there is a clear path
        // The tree is visible so add 1 to the result
				result++
				break
			}
		}
	}

	fmt.Println("=-= PART 1 =-=")
	fmt.Println(result)
	fmt.Println("=-= PART 2 =-=")
}
