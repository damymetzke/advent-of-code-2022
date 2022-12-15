package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Point struct {
	x int
	y int
}

type Path []Point

type Bounds struct {
	left   int
	right  int
	bottom int
}

type Grid [][]bool

func GetInput() string {
	data, err := os.ReadFile("input/14")
	if err != nil {
		log.Fatal("Could not read file 'input/14':\n  * ", err)
	}

	return string(data)
}

func ParsePath(line string) Path {
	pointStrings := strings.Split(line, " -> ")
	result := make(Path, len(pointStrings))
	for i, point := range pointStrings {
		coords := strings.Split(point, ",")
		left, leftErr := strconv.ParseInt(coords[0], 10, 64)
		right, rightErr := strconv.ParseInt(coords[1], 10, 64)

		if leftErr != nil || rightErr != nil {
			log.Fatalf("Could not parse string '%v'", line)
		}
		result[i] = Point{
			x: int(left),
			y: int(right),
		}
	}

	return result
}

func GetBounds(path Path) Bounds {
	result := Bounds{
		left:   path[0].x,
		right:  path[0].x,
		bottom: path[0].y,
	}

	for _, point := range path {
		if point.x < result.left {
			result.left = point.x
		}
		if point.x > result.right {
			result.right = point.x
		}
		if point.y > result.bottom {
			result.bottom = point.y
		}
	}

	return result
}

func ShouldUpdateBounds(total, current Bounds) bool {
	return total.bottom == -1 ||
		current.left < total.left ||
		current.right > total.right ||
		current.bottom > total.bottom
}

func UpdateBounds(total, current Bounds) Bounds {
	if total.bottom == -1 {
		return current
	}

	if current.left < total.left {
		total.left = current.left
	}

	if current.right > total.right {
		total.right = current.right
	}

	if current.bottom > total.bottom {
		total.bottom = current.bottom
	}

	return total
}

func GridToString(grid Grid) string {
	var result strings.Builder

	for _, row := range grid {
		for _, value := range row {
			if value {
				result.WriteRune('#')
			} else {
				result.WriteRune('.')
			}
		}
		result.WriteRune('\n')
	}

	return result.String()
}

func GetPathDelta(from, to Point) Point {
	if from.x == to.x {
		if from.y < to.y {
			return Point{
				x: 0,
				y: 1,
			}
		}

		return Point{
			x: 0,
			y: -1,
		}
	}

	if from.x < to.x {
		return Point{
			x: 1,
			y: 0,
		}
	}

	return Point{
		x: -1,
		y: 0,
	}
}

func main() {
	input := strings.Split(GetInput(), "\n")

	totalBounds := Bounds{0, 0, -1}
	paths := make([]Path, len(input))

	var wait sync.WaitGroup
	var lock sync.Mutex

	for i, line := range input {
		wait.Add(1)
		num := i
		currentLine := line
		go func() {
			path := ParsePath(currentLine)
			paths[num] = path
			bounds := GetBounds(path)

			if ShouldUpdateBounds(totalBounds, bounds) {
				lock.Lock()
				totalBounds = UpdateBounds(totalBounds, bounds)
				lock.Unlock()
			}
			wait.Done()
		}()
	}

	wait.Wait()

	width := totalBounds.right - totalBounds.left
	depth := totalBounds.bottom
	left := totalBounds.left

	grid := make(Grid, depth + 1)
	for i := range grid {
		grid[i] = make([]bool, width + 1)
	}

	for _, path := range paths {
		wait.Add(1)
		currentPath := path

		go func() {
			for j := 0; j < len(currentPath)-1; j++ {
        delta := GetPathDelta(currentPath[j], currentPath[j+1])
        follow := currentPath[j]

        for follow != currentPath[j+1] {
          grid[follow.y][follow.x - left] = true
          follow.x += delta.x
          follow.y += delta.y
        }
			}
      end := currentPath[len(currentPath) - 1]
      grid[end.y][end.x - left] = true
			wait.Done()
		}()
	}

	wait.Wait()


	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")
  fmt.Println("=-= Visual =-=")

	fmt.Println(GridToString(grid))
}
