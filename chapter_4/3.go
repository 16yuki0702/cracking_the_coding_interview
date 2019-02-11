package main

import (
	"../util"
	"container/list"
	"fmt"
)

func main() {
	seed := []int{1, 2, 3, 4, 5, 6, 7}
	root := util.CreateMinimalBST(seed)
	CreateLevelLinkedList(root)

}

func CreateLevelLinkedList(root *util.TreeNode) []*list.List {
	lists := []*list.List{}
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
