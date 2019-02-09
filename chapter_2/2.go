package main

import (
	"../util"
	"container/list"
	"fmt"
)

func main() {
	l1 := util.SequenceList()
	l2 := util.SequenceList()

	pattern1(l1.Front(), 4)
	pattern2(l2.Front(), 4)
}

// pattern1
func pattern1(e *list.Element, k int) int {
	if e == nil {
		return 0
	}

	index := pattern1(e.Next(), k) + 1
	if index == k {
		fmt.Printf("%dth to last element is %d\n", k, e.Value)
	}

	return index
}

// pattern2
// takes time O(N), takes space (1).
func pattern2(e *list.Element, k int) {
	if e == nil {
		fmt.Printf("element is nil\n")
		return
	}

	p1 := e
	p2 := e

	for i := 0; i < k; i = i + 1 {
		p1 = p1.Next()
	}

	for p1 != nil {
		p1 = p1.Next()
		p2 = p2.Next()
	}

	fmt.Printf("%dth to last element is %d\n", k, p2.Value)
}
