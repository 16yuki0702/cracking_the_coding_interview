package main

import (
	"../util"
	"fmt"
)

func main() {
	testSliceStack()
}

func testLikedListStack() {
	s := util.NewLinkedListStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Peek())
	util.DumpLinkedListStack(s)

	s.Pop()
	s.Pop()

	fmt.Println(s.Peek())

	util.DumpLinkedListStack(s)
}

func testSliceStack() {
	s := util.NewSliceStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Peek())
	util.DumpSliceStack(s)

	s.Pop()
	s.Pop()

	fmt.Println(s.Peek())

	util.DumpSliceStack(s)
}
