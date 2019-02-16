package main

import (
	"../util"
	"fmt"
	"math"
)

func main() {
	seed := []int{0, 1, 2, 3, 4, 5, 6}
	root := util.CreateMinimalBST(seed)
	pattern1(root)
	pattern2(root)
}

// pattern1
func pattern1(root *util.TreeNode) {
	if root == nil {
		fmt.Println("root is nil")
		return
	}

	diff := getHeight(root.Left) - getHeight(root.Right)
	if abs(diff) > 1 {
		fmt.Println("root is not balanced.")
		return
	} else {
		fmt.Println("root is balanced.")
		return
	}
}

// pattern2
func pattern2(root *util.TreeNode) int64 {
	if root == nil {
		return -1
	}

	leftHeight := pattern2(root.Left)
	if leftHeight == math.MinInt64 {
		return math.MinInt64
	}

	rightHeight := pattern2(root.Right)
	if rightHeight == math.MinInt64 {
		return math.MinInt64
	}

	diff := leftHeight - rightHeight
	if abs64(diff) > 1 {
		return math.MinInt64
	} else {
		m := max64(leftHeight, rightHeight) + 1
		return m
	}
}

func getHeight(root *util.TreeNode) int {
	if root == nil {
		return -1
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func max64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
