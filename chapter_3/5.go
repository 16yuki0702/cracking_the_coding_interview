package main

import (
	"../util"
	"fmt"
)

func main() {
	s := util.NewSliceStack()
	s.Push(1)
	s.Push(3)
	s.Push(8)
	s.Push(7)
	s.Push(10)
	s.Push(5)
	s.Push(12)

	sortStack(s)
}

func sortStack(s *util.SliceStack) {
	t := util.NewSliceStack()
	for s.Peek() != -1 {
		tmp := s.Pop()

		for t.Peek() != -1 && tmp < t.Peek() {
			s.Push(t.Pop())
		}
		t.Push(tmp)
	}

	for t.Peek() != -1 {
		s.Push(t.Pop())
	}

	for e := s.Pop(); e != -1; e = s.Pop() {
		fmt.Println(e)
	}
}
