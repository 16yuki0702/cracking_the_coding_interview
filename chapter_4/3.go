package main

import (
	"../util"
	"container/list"
	"fmt"
)

func main() {
	seed := []int{0, 1, 2, 3, 4, 5, 6}
	root := util.CreateMinimalBST(seed)
	pattern1(root)
	pattern2(root)
}

// pattern1
func pattern1(root *util.TreeNode) []*list.List {
	lists := make([]*list.List, 0)
	createLevelLinkedList(root, lists, 0)
	return lists
}

func createLevelLinkedList(root *util.TreeNode, lists []*list.List, level int) {
	if root == nil {
		return
	}

	var l *list.List

	if len(lists) == level {
		l = list.New()
		lists = append(lists, l)
	} else {
		l = lists[level]
	}

	l.PushBack(root)
	createLevelLinkedList(root.Left, lists, level+1)
	createLevelLinkedList(root.Right, lists, level+1)
}

func pattern2(root *util.TreeNode) {
	result := make([]*list.List, 0)

	current := list.New()
	if root != nil {
		current.PushBack(root)
	}

	for current.Len() > 0 {
		result = append(result, current)
		parent := current
		current = list.New()

		for e := parent.Front(); e != nil; e = e.Next() {
			r := e.Value.(*util.TreeNode)
			if r.Left != nil {
				current.PushBack(r.Left)
			}
			if r.Right != nil {
				current.PushBack(r.Right)
			}
		}
	}

	for i, v := range result {
		for e := v.Front(); e != nil; e = e.Next() {
			r := e.Value.(*util.TreeNode)
			fmt.Printf("Level %d : Value %d\t", i, r.Number)
		}
		fmt.Printf("\n")
	}
}
