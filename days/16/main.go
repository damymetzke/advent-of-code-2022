package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	name  string
	rate  int
	edges []string
}

type GraphEdge struct {
	node  *GraphNode
	steps int
}

type GraphNode struct {
	valve Valve
	edges []GraphEdge
}

type Graph struct {
	root  *GraphNode
	nodes []*GraphNode
}

func GetInput() string {
	data, err := os.ReadFile("input/16")
	if err != nil {
		log.Fatal("Could not read file 'input/16':\n  * ", err)
	}

	return string(data)
}

func ParseLine(line string) Valve {
	parts := strings.Split(line, "; ")
	name := parts[0][6:8]
	rate, rateErr := strconv.ParseInt(parts[0][23:], 10, 64)

	end := strings.Split(parts[1], "valve")[1]
	if end[0] == 's' {
		end = end[2:]
	} else {
		end = end[1:]
	}

	edges := strings.Split(end, ", ")

	if rateErr != nil {
		log.Fatalf("Could not parse line '%v'", line)
	}
	return Valve{
		name:  name,
		rate:  int(rate),
		edges: edges,
	}
}

func UpdateGraph(graph *Graph, valve Valve) {
	next := &GraphNode{
		valve: valve,
		edges: []GraphEdge{},
	}

	for _, edge := range valve.edges {
		for i := range graph.nodes {
			if graph.nodes[i].valve.name == edge {
				graph.nodes[i].edges = append(graph.nodes[i].edges, GraphEdge{
					node:  next,
					steps: 1,
				})

				next.edges = append(next.edges, GraphEdge{
					node:  graph.nodes[i],
					steps: 1,
				})
			}
		}
	}

	graph.nodes = append(graph.nodes, next)
  if valve.name == "AA" {
    graph.root = next
  }
}

func main() {
	lines := strings.Split(GetInput(), "\n")

	valveChannel := make(chan Valve, len(lines))

	for _, line := range lines {
		value := line
		go func() {
			valveChannel <- ParseLine(value)
		}()
	}

	valves := make([]Valve, len(lines))
  var graph Graph

	for i := 0; i < len(valves); i++ {
		valves[i] = <-valveChannel
    UpdateGraph(&graph, valves[i])
	}

  fmt.Println(*graph.root)
  for _, node := range graph.nodes {
    fmt.Println(*node)
  }

	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")
}
