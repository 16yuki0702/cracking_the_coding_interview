package util

import (
	"container/list"
	"fmt"
)

func DuplicateList() *list.List {
	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(5)
	l.PushBack(5)
	l.PushBack(6)

	return l
}

func SequenceList() *list.List {
	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)

	return l
}

func DumpList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
