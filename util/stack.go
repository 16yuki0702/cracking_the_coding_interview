package util

import (
	"container/list"
	"fmt"
)

type LinkedListStack struct {
	stack *list.List
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		stack: list.New(),
	}
}

func (s *LinkedListStack) Push(v interface{}) {
	s.stack.PushBack(v)
}

func (s *LinkedListStack) Pop() interface{} {
	top := s.stack.Back()
	res := top
	s.stack.Remove(top)
	return res
}

func (s *LinkedListStack) Peek() interface{} {
	return s.stack.Back()
}

func DumpLinkedListStack(s *LinkedListStack) {
	for e := s.stack.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}

type SliceStack struct {
	stack    []int
	minStack []int
}

func NewSliceStack() *SliceStack {
	return &SliceStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (s *SliceStack) Push(v int) {
	s.stack = append(s.stack, v)
	if v <= s.Min() {
		s.minStack = append(s.minStack, v)
	}
}

func (s *SliceStack) Pop() int {
	length := len(s.stack)
	if length == 0 {
		return -1
	}

	res := s.stack[length-1]
	s.stack = s.stack[:length-1]

	if res == s.Min() {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}

	return res
}

func (s *SliceStack) Peek() int {
	length := len(s.stack)
	if length == 0 {
		return -1
	}
	return s.stack[length-1]
}

func (s *SliceStack) Min() int {
	min := 0
	length := len(s.minStack)

	if length == 0 {
		min = int(^uint(0) >> 1)
	} else {
		min = s.minStack[length-1]
	}

	return min
}

func (s *SliceStack) IsEmpty() bool {
	length := len(s.stack)
	if length == 0 {
		return true
	}
	return false
}

func (s *SliceStack) Size() int {
	return len(s.stack)
}

func DumpSliceStack(s *SliceStack) {
	i := len(s.stack)
	for top := i - 1; 0 <= top; top = top - 1 {
		fmt.Println(s.stack[top])
	}
}
