package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

type Range struct {
	min int
	max int
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

func GetDistance(lhs, rhs Point) int {
	x := lhs.x - rhs.x
	y := lhs.y - rhs.y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	return x + y
}

func GetRangeY(sensor SensorData, yRow int) Range {
	verticalDifference := sensor.sensorPosition.y - yRow
	if verticalDifference < 0 {
		verticalDifference = -verticalDifference
	}

	difference := verticalDifference - GetDistance(sensor.sensorPosition, sensor.closestBeacon)
	return Range{
		min: sensor.sensorPosition.x + difference,
		max: sensor.sensorPosition.x - difference,
	}
}

func IsInRange(data SensorData, point Point) bool {
	return GetDistance(data.sensorPosition, point) <=
		GetDistance(data.sensorPosition, data.closestBeacon)
}

func FindRangesY(datas []SensorData, y int) []Range {
	ranges := make([]Range, len(datas))

	for i, data := range datas {
		ranges[i] = GetRangeY(data, y)
	}

  sort.Slice(ranges, func(i, j int) bool {
    return ranges[i].min < ranges[j].min
  })


  var result []Range

  for i := 0; i < len(ranges); i++ {
    if ranges[i].max < ranges[i].min {
      continue
    }

    if len(result) == 0 {
      result = []Range{ranges[i]}
    }

    last := len(result) - 1
    if result[last].max < ranges[i].min {
      result = append(result, ranges[i])
      continue
    }

    if ranges[i].max > result[last].max {
      result[last].max = ranges[i].max
    }
  }

  return result
}

func main() {
	lines := strings.Split(GetInput(), "\n")
	datas := make([]SensorData, len(lines))
	maxX := 0
	minX := 0

	for i, line := range lines {
		datas[i] = ParseLine(line)
		if datas[i].sensorPosition.x > maxX {
			maxX = datas[i].sensorPosition.x
		}
		if datas[i].sensorPosition.x-
			GetDistance(datas[i].sensorPosition, datas[i].closestBeacon) < minX {
			minX = datas[i].sensorPosition.x - GetDistance(datas[i].sensorPosition, datas[i].closestBeacon)
		}
	}

  combinedRanges := FindRangesY(datas, 2000000)

  var numCannotContainBeacon int

  for _, value := range combinedRanges {
    numCannotContainBeacon += value.max - value.min
  }

  var result int

  for y := 0; y <= 4000000; y++ {
    combinedRanges := FindRangesY(datas, y)
    i := 0
    for combinedRanges[i].max < 0 {
      i++
    }

    if combinedRanges[i].min < 0 && combinedRanges[i].max > 4000000 {
      continue
    }

    if combinedRanges[i].min == 1 {
      result = 0
    }

    result = (combinedRanges[i].max + 1) * 4000000 + y
    break
  }

	fmt.Println("=-= PART 1 =-=")
  fmt.Println(numCannotContainBeacon)
	fmt.Println("=-= PART 2 =-=")
  fmt.Println(result)
}
