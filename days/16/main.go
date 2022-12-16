package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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

func ParseAllLines(lines []string) <-chan Valve {
	result := make(chan Valve, len(lines))
	var wait sync.WaitGroup

	go func() {
		for _, line := range lines {
			current := line
			wait.Add(1)
			go func() {
				result <- ParseLine(current)
				wait.Done()
			}()
		}

		wait.Wait()
		close(result)
	}()

	return result
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

func BuildGraph(valves <-chan Valve) <-chan Graph {
	result := make(chan Graph, 1)

	go func() {
		var graph Graph
		for valve := range valves {
			UpdateGraph(&graph, valve)
		}

		result <- graph
		close(result)
	}()

	return result
}

func GenerateOptimizedGraph(graphChannel <-chan Graph) <-chan Graph {
	result := make(chan Graph, 1)

	go func() {
		fullGraph := <-graphChannel
		var graph Graph
		nodeMap := make(map[string]*GraphNode)

		for _, node := range fullGraph.nodes {
			if node.valve.rate == 0 && node.valve.name != "AA" {
				continue
			}

			var nextEdges []GraphEdge

			visited := []*GraphNode{node}
			backlog := []*GraphNode{node}
			nextBacklog := []*GraphNode{node}

			step := 0

			for len(backlog) != 0 {
				for _, item := range backlog {
					currentValue, currentValueOk := nodeMap[item.valve.name]
					if !currentValueOk {
						currentValue = &GraphNode{
							valve: item.valve,
							edges: []GraphEdge{},
						}
						nodeMap[item.valve.name] = currentValue
            if item.valve.rate > 0 || item.valve.name == "AA" {
              graph.nodes = append(graph.nodes, currentValue)
            }
					}

					if step != 0 && item.valve.rate > 0 {

						nextEdges = append(nextEdges, GraphEdge{
							node:  currentValue,
							steps: step,
						})
					}

				checkEdge:
					for _, edge := range item.edges {
						for _, checkVisisted := range visited {
							if edge.node == checkVisisted {
								continue checkEdge
							}
						}

						nextBacklog = append(nextBacklog, edge.node)
            visited = append(visited, edge.node)
					}
				}

				backlog = nextBacklog
				nextBacklog = []*GraphNode{}
				step++
			}

			nodeMap[node.valve.name].edges = nextEdges
		}

    graph.root = nodeMap["AA"]

		result <- graph
		close(result)
	}()

	return result
}

func FindHeighestScore(node *GraphNode, stepsLeft, losingPressure int, visited []*GraphNode) int {
  max := stepsLeft * losingPressure
  outerLoop:
  for _, edge := range node.edges {
    // Don't visit twice
    for _, checkVisited := range visited {
      if edge.node == checkVisited {
        continue outerLoop
      }
    }

    // Not enough time
    if edge.steps + 1 > stepsLeft {
      continue
    }

    score := FindHeighestScore(edge.node, stepsLeft - (edge.steps + 1), losingPressure + edge.node.valve.rate, append(visited, edge.node)) + (edge.steps + 1) * losingPressure
    if score > max {
      max = score
    }
  }
  return max
}

func SolvePart1(graph <-chan Graph) int {
  value := <-graph
  return FindHeighestScore(value.root, 30, 0, []*GraphNode{value.root})
}

func main() {
	lines := strings.Split(GetInput(), "\n")

	valves := ParseAllLines(lines)

	graph := BuildGraph(valves)

	optimizedGraph := GenerateOptimizedGraph(graph)

	fmt.Println("=-= PART 1 =-=")
  fmt.Println(SolvePart1(optimizedGraph))
	fmt.Println("=-= PART 2 =-=")
}
