package main

import (
	"../util"
	"container/list"
	"fmt"
)

func main() {
	g := util.NewGraph()
	n1 := util.NewNode("node 1")
	n2 := util.NewNode("node 2")

	n1.AddNode(n2)

	g.AddGraph(n1)

	search(g, n1, n2)
}

func search(g *util.Graph, start, end *util.Node) {
	if start == end {
		fmt.Printf("found node %s, %s\n", start.Name, end.Name)
		return
	}

	// using LinkedList as a queue
	q := list.New()

	for _, v := range g.Nodes {
		v.State = util.Unvisited
	}

	start.State = util.Visiting
	q.PushBack(start)

	for q.Len() != 0 {
		e := q.Front()

		if e != nil {
			q.Remove(e)
			tmp := e.Value.(*util.Node)

			for _, child := range tmp.Children {
				if child.State == util.Unvisited {
					if child == end {
						fmt.Printf("found node %s, %s\n", start.Name, end.Name)
						return
					} else {
						child.State = util.Visiting
						q.PushBack(child)
					}
				}
			}

			tmp.State = util.Visited
		}
	}

	fmt.Printf("not found node %s, %s\n", start.Name, end.Name)
}
