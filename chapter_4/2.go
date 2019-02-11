package main

import (
	"../util"
	"fmt"
)

func main() {
	seed := []int{1, 2, 3, 4, 5, 6, 7}

	n := createMinimalBST(seed, 0, len(seed)-1)

	fmt.Println(n)
}

func createMinimalBST(seed []int, start, end int) *util.TreeNode {
	if end < start {
		return nil
	}
	mid := (start + end) / 2
	n := util.NewTreeNode(mid)
	n.Left = createMinimalBST(seed, start, mid-1)
	n.Right = createMinimalBST(seed, mid+1, end)

	return n
}
