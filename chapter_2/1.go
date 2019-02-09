package main

import (
	"../util"
	"container/list"
	"fmt"
)

func main() {
	l1 := util.DuplicateList()
	l2 := util.DuplicateList()

	pattern1(l1)
	pattern2(l2)
}

// pattern1
func pattern1(l *list.List) {
	fmt.Println("pattern1")

	util.DumpList(l)

	m := make(map[int]bool)

	for e := l.Front(); e != nil; {
		v := e.Value.(int)
		if m[v] {
			tmp := e
			e = e.Next()
			l.Remove(tmp)
		} else {
			m[v] = true
			e = e.Next()
		}
	}

	util.DumpList(l)
}

// pattern2
func pattern2(l *list.List) {
	fmt.Println("pattern2")

	util.DumpList(l)

	// it is efficient for memory allocation.
	// but cost is O(n^2)
	for e1 := l.Front(); e1 != nil; e1 = e1.Next() {
		v1 := e1.Value.(int)

		for e2 := e1.Next(); e2 != nil; {
			v2 := e2.Value.(int)
			if v1 == v2 {
				tmp := e2
				e2 = e2.Next()
				l.Remove(tmp)
			} else {
				e2 = e2.Next()
			}
		}
	}

	util.DumpList(l)
}
