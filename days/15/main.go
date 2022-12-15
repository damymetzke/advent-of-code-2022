package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type SensorData struct {
	sensorPosition Point
	closestBeacon  Point
}

func GetInput() string {
	data, err := os.ReadFile("input/15")
	if err != nil {
		log.Fatal("Could not read file 'input/15':\n  * ", err)
	}

	return string(data)
}

func ParsePoint(coords string) Point {
	parts := strings.Split(coords, ", ")

  x, xErr := strconv.ParseInt(parts[0][2:], 10, 64)
  y, yErr := strconv.ParseInt(parts[1][2:], 10, 64)

  if xErr != nil || yErr != nil {
    log.Fatalf("Could not parse coordinates '%v'", coords)
  }

	return Point{
		x: int(x),
		y: int(y),
	}
}

func ParseLine(line string) SensorData {
	parts := strings.Split(line, ": closest beacon is at ")

	sensor := ParsePoint(parts[0][10:])
	beacon := ParsePoint(parts[1])

	return SensorData{
		sensorPosition: sensor,
		closestBeacon:  beacon,
	}
}

func main() {
	lines := strings.Split(GetInput(), "\n")
	datas := make([]SensorData, len(lines))

	for i, line := range lines {
		datas[i] = ParseLine(line)
	}

	for _, data := range datas {
		fmt.Println(data)
	}

	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")
}
