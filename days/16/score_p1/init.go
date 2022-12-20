package score_p1

import (
	. "damymetzke.com/advent-of-code-2022/d16/shared"
)

type state uint8

const (
	P1 state = iota
	P2
)

func findHeighestScore(
	p1,
	p2 *GraphNode,
	stepsLeft,
	losingPressure,
	altStepsLose,
	altPressureGain int,
	visited []*GraphNode,
	state state) int {

	max := stepsLeft * losingPressure + (stepsLeft - altStepsLose) * altPressureGain

	node := p1
	alt := P2
	if state == P2 {
		node = p2
		alt = P1
	}

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

		n1, n2 := p1, p2
		if state == P1 {
			n1 = edge.Node
		} else {
			n2 = edge.Node
		}
		stepsLose := edge.Steps + 1
		var score int
		if stepsLose > altStepsLose {
			score = findHeighestScore(
				n1,
				n2,
				stepsLeft-altStepsLose,
				losingPressure+altPressureGain,
				stepsLose-altStepsLose,
				edge.Node.Valve.Rate,
				append(visited, edge.Node),
				alt) + altStepsLose*losingPressure
		} else if stepsLose < altStepsLose {
			score = findHeighestScore(
				n1,
				n2,
				stepsLeft-stepsLose,
				losingPressure+edge.Node.Valve.Rate,
				altStepsLose-stepsLose,
				altPressureGain,
				append(visited, edge.Node),
				state) + stepsLose*losingPressure
		} else {
			score = findHeighestScore(
				n1,
				n2,
				stepsLeft-stepsLose,
				losingPressure+edge.Node.Valve.Rate+altPressureGain,
				0,
				0,
				append(visited, edge.Node),
				P1) + stepsLose*losingPressure
		}
		if score > max {
			max = score
		}
	}
	return max
}

func SolvePart2(graph <-chan Graph) int {
	value := <-graph
	return findHeighestScore(value.Root, value.Root, 26, 0, 0, 0, []*GraphNode{value.Root}, P1)
}
