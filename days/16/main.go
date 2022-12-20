package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"damymetzke.com/advent-of-code-2022/d16/score_p1"
	. "damymetzke.com/advent-of-code-2022/d16/shared"
)

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
		Name:  name,
		Rate:  int(rate),
		Edges: edges,
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
		Valve: valve,
		Edges: []GraphEdge{},
	}

	for _, edge := range valve.Edges {
		for i := range graph.Nodes {
			if graph.Nodes[i].Valve.Name == edge {
				graph.Nodes[i].Edges = append(graph.Nodes[i].Edges, GraphEdge{
					Node:  next,
					Steps: 1,
				})

				next.Edges = append(next.Edges, GraphEdge{
					Node:  graph.Nodes[i],
					Steps: 1,
				})
			}
		}
	}

	graph.Nodes = append(graph.Nodes, next)
	if valve.Name == "AA" {
		graph.Root = next
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

		for _, node := range fullGraph.Nodes {
			if node.Valve.Rate == 0 && node.Valve.Name != "AA" {
				continue
			}

			var nextEdges []GraphEdge

			visited := []*GraphNode{node}
			backlog := []*GraphNode{node}
			nextBacklog := []*GraphNode{node}

			step := 0

			for len(backlog) != 0 {
				for _, item := range backlog {
					currentValue, currentValueOk := nodeMap[item.Valve.Name]
					if !currentValueOk {
						currentValue = &GraphNode{
							Valve: item.Valve,
							Edges: []GraphEdge{},
						}
						nodeMap[item.Valve.Name] = currentValue
						if item.Valve.Rate > 0 || item.Valve.Name == "AA" {
							graph.Nodes = append(graph.Nodes, currentValue)
						}
					}

					if step != 0 && item.Valve.Rate > 0 {

						nextEdges = append(nextEdges, GraphEdge{
							Node:  currentValue,
							Steps: step,
						})
					}

				checkEdge:
					for _, edge := range item.Edges {
						for _, checkVisisted := range visited {
							if edge.Node == checkVisisted {
								continue checkEdge
							}
						}

						nextBacklog = append(nextBacklog, edge.Node)
						visited = append(visited, edge.Node)
					}
				}

				backlog = nextBacklog
				nextBacklog = []*GraphNode{}
				step++
			}

			nodeMap[node.Valve.Name].Edges = nextEdges
		}

		graph.Root = nodeMap["AA"]

		result <- graph
		close(result)
	}()

	return result
}

func FindHeighestScore(node *GraphNode, stepsLeft, losingPressure int, visited []*GraphNode) int {
	max := stepsLeft * losingPressure
outerLoop:
	for _, edge := range node.Edges {
		// Don't visit twice
		for _, checkVisited := range visited {
			if edge.Node == checkVisited {
				continue outerLoop
			}
		}

		// Not enough time
		if edge.Steps+1 > stepsLeft {
			continue
		}

		score := FindHeighestScore(edge.Node, stepsLeft-(edge.Steps+1), losingPressure+edge.Node.Valve.Rate, append(visited, edge.Node)) + (edge.Steps+1)*losingPressure
		if score > max {
			max = score
		}
	}
	return max
}

func SolvePart1(graph <-chan Graph) int {
	value := <-graph
	return FindHeighestScore(value.Root, 30, 0, []*GraphNode{value.Root})
}

func ForkGraph(graph <-chan Graph) (<-chan Graph, <-chan Graph) {
  left := make(chan Graph, 1)
  right := make(chan Graph, 1)

  go func(){
    result := <-graph
    left <- result
    right <-result
    close(left)
    close(right)
  }()

  return left, right
}

func main() {
	lines := strings.Split(GetInput(), "\n")

	valves := ParseAllLines(lines)

	graph := BuildGraph(valves)

	graph1, graph2 := ForkGraph(GenerateOptimizedGraph(graph))

	fmt.Println("=-= PART 1 =-=")
	fmt.Println(SolvePart1(graph1))
	fmt.Println("=-= PART 2 =-=")
  fmt.Println(score_p1.SolvePart2(graph2))
}
