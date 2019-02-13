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

type TreeNode struct {
	Number int
	Right  *TreeNode
	Left   *TreeNode
}

func NewTreeNode(number int) *TreeNode {
	return &TreeNode{
		Number: number,
		Right:  nil,
		Left:   nil,
	}
}

type Tree struct {
	Root *TreeNode
}

func NewTree(root *TreeNode) *Tree {
	return &Tree{
		Root: root,
	}
}

func CreateMinimalBST(seed []int) *TreeNode {
	return createMinimalBST(seed, 0, len(seed))
}

func createMinimalBST(seed []int, start, end int) *TreeNode {
	if end < start {
		return nil
	}

	mid := (start + end) / 2
	if len(seed) <= mid {
		return nil
	}

	n := NewTreeNode(seed[mid])
	n.Left = createMinimalBST(seed, start, mid-1)
	n.Right = createMinimalBST(seed, mid+1, end)

	return n
}
