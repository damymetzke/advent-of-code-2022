package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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

func GridToString(grid, sand Grid, columns int) string {
	var result strings.Builder

	width := len(grid[0])
	height := len(grid)
	offset := height / columns

	for i := 0; i < len(grid)/columns; i++ {
		for j := 0; j < columns; j++ {
			row := grid[i+j*offset]
			for k, value := range row {
        sand := sand[i+j*offset][k]
				if value {
					result.WriteString("\x1b[48;2;68;64;60m#\x1b[0m")
				} else if sand {
          //fa cc 15
					result.WriteString("\x1b[38;2;250;204;21m\x1b[48;2;68;64;60mO\x1b[0m")
        } else {
					result.WriteString("\x1b[48;2;87;83;78m.\x1b[0m")
				}
			}
			result.WriteRune(' ')
		}
		result.WriteRune('\n')
	}

	for i := 0; i < height%3; i++ {
		for j := 0; j < (width+1)*(columns-1); j++ {
			result.WriteRune(' ')
		}

		row := grid[len(grid)-height%3+i]
		for _, value := range row {
			if value {
				result.WriteString("\x1b[48;2;68;64;60m#\x1b[0m")
			} else {
				result.WriteString("\x1b[48;2;87;83;78m.\x1b[0m")
			}
		}
		result.WriteRune(' ')
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

func FindNextSandPosition(stone, sand Grid, left, bottom int) Point {
	next := Point{
		x: 500,
		y: 0,
	}

	for {
    if next.y >= bottom {
      return Point{
        0,
        -1,
      }
    }
		if !stone[next.y+1][next.x - left] && !sand[next.y+1][next.x - left] {
			next.y++
			continue
		}

		if !stone[next.y+1][next.x-1 - left] && !sand[next.y+1][next.x-1 - left] {
			next.x--
			next.y++
			continue
		}
		if !stone[next.y+1][next.x+1 - left] && !sand[next.y+1][next.x+1 - left] {
			next.x++
			next.y++
			continue
		}
		break
	}

	return next
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

	grid := make(Grid, depth+1)
	for i := range grid {
		grid[i] = make([]bool, width+1)
	}

	for _, path := range paths {
		wait.Add(1)
		currentPath := path

		go func() {
			for j := 0; j < len(currentPath)-1; j++ {
				delta := GetPathDelta(currentPath[j], currentPath[j+1])
				follow := currentPath[j]

				for follow != currentPath[j+1] {
					grid[follow.y][follow.x-left] = true
					follow.x += delta.x
					follow.y += delta.y
				}
			}
			end := currentPath[len(currentPath)-1]
			grid[end.y][end.x-left] = true
			wait.Done()
		}()
	}

	wait.Wait()

	sand := make(Grid, len(grid))
	for i := range grid {
		sand[i] = make([]bool, len(grid[0]))
	}

  var restingSand int

  for {
    nextSand := FindNextSandPosition(grid, sand, left, depth)
    if nextSand.y == -1 {
      break
    }
    sand[nextSand.y][nextSand.x - left] = true
    restingSand++
  }

	fmt.Println("=-= PART 1 =-=")
  fmt.Println(restingSand)
	fmt.Println("=-= PART 2 =-=")
	fmt.Println("=-= Visual =-=")

	sandVisual := make(Grid, len(grid))
	for i := range grid {
		sandVisual[i] = make([]bool, len(grid[0]))
	}

	fmt.Print(GridToString(grid, sandVisual, 3))

	numLines := len(grid)/3 + len(grid)%3
	var clearBuild strings.Builder
	for i := 0; i < numLines; i++ {
		clearBuild.WriteString("\x1b[1A")
	}

	clear := clearBuild.String()

  for {
    nextSand := FindNextSandPosition(grid, sandVisual, left, depth)
    if nextSand.y == -1 {
      break
    }
    sandVisual[nextSand.y][nextSand.x - left] = true
    time.Sleep(5 * time.Millisecond)
    fmt.Println(clear)
    fmt.Print(GridToString(grid, sandVisual, 3))
  }
}
