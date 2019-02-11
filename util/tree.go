package util

type NodeState int

const (
	Unvisited NodeState = iota + 1
	Visited
	Visiting
)

type Node struct {
	State    NodeState
	Name     string
	Children []*Node
}

func NewNode(name string) *Node {
	return &Node{
		State:    Unvisited,
		Name:     name,
		Children: make([]*Node, 0),
	}
}

func (n *Node) AddNode(child *Node) {
	n.Children = append(n.Children, child)
}

type Tree struct {
	Root Node
}

func NewTree() *Tree {
	return &Tree{}
}

type Graph struct {
	Nodes []*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make([]*Node, 0),
	}
}

func (g *Graph) AddGraph(n *Node) {
	g.Nodes = append(g.Nodes, n)
}
