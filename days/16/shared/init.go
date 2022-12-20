package shared

type Valve struct {
	Name  string
	Rate  int
	Edges []string
}

type GraphEdge struct {
	Node  *GraphNode
	Steps int
}

type GraphNode struct {
	Valve Valve
	Edges []GraphEdge
}

type Graph struct {
	Root  *GraphNode
	Nodes []*GraphNode
}
