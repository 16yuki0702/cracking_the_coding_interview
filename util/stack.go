package util

import (
	"container/list"
	"fmt"
)

type LinkedListStack struct {
	Stack *list.List
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		Stack: list.New(),
	}
}

func (s *LinkedListStack) Push(v interface{}) {
	s.Stack.PushBack(v)
}

func (s *LinkedListStack) Pop() interface{} {
	top := s.Stack.Back()
	res := top
	s.Stack.Remove(top)
	return res
}

func (s *LinkedListStack) Peek() interface{} {
	return s.Stack.Back()
}

func DumpLinkedListStack(s *LinkedListStack) {
	for e := s.Stack.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}

type SliceStack struct {
	Stack []int
}

func NewSliceStack() *SliceStack {
	return &SliceStack{
		Stack: make([]int, 0),
	}
}

func (s *SliceStack) Push(v int) {
	s.Stack = append(s.Stack, v)
}

func (s *SliceStack) Pop() int {
	length := len(s.Stack)
	if length == 0 {
		return -1
	}

	res := s.Stack[length-1]
	if length == 1 {
		s.Stack = nil
	} else {
		s.Stack = s.Stack[:length-1]
	}

	return res
}

func (s *SliceStack) Peek() int {
	length := len(s.Stack)
	if length == 0 {
		return -1
	}
	return s.Stack[length-1]
}

func DumpSliceStack(s *SliceStack) {
	i := len(s.Stack)
	for top := i - 1; 0 <= top; top = top - 1 {
		fmt.Println(s.Stack[top])
	}
}
