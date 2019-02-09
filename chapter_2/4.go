package main

import (
	"../util"
	"container/list"
)

func main() {
	l := util.PartitionList()
	pattern1(l, 5)
}

func pattern1(l *list.List, partition int) {
	beforeList := list.New()
	afterList := list.New()

	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)

		if v < partition {
			beforeList.PushBack(v)
		} else {
			afterList.PushBack(v)
		}
	}

	beforeList.PushBackList(afterList)

	util.DumpList(beforeList)
}
