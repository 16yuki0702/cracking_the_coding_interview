package main

import (
	"container/list"
	"fmt"
)

func main() {
	l1 := list.New()
	l1.PushFront(6)
	l1.PushFront(1)
	l1.PushFront(7)

	l2 := list.New()
	l2.PushFront(2)
	l2.PushFront(9)
	l2.PushFront(5)

	pattern1(l1, l2)
}

func pattern1(l1, l2 *list.List) {
	result := list.New()
	addList(result, l1.Front(), l2.Front(), 0)

	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

func addList(result *list.List, e1, e2 *list.Element, carry int) {
	if e1 == nil && e2 == nil && carry == 0 {
		return
	}

	v := carry

	if e1 != nil {
		v = v + e1.Value.(int)
	}
	if e2 != nil {
		v = v + e2.Value.(int)
	}

	result.PushFront(v % 10)

	if e1 != nil || e2 != nil {
		nextCarry := 0
		if v >= 10 {
			nextCarry = 1
		}
		var next1 *list.Element
		if e1.Next() != nil {
			next1 = e1.Next()
		}
		var next2 *list.Element
		if e2.Next() != nil {
			next2 = e2.Next()
		}
		addList(result, next1, next2, nextCarry)
	}
}
