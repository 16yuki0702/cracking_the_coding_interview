package main

import (
	"../util"
	"fmt"
)

func main() {
	testSliceStack()
}

func testSliceStack() {
	s := util.NewSliceStack()

	s.Push(5)
	s.Push(6)
	s.Push(3)
	s.Push(7)

	fmt.Println(s.Peek())
	fmt.Println(s.Min())
	util.DumpSliceStack(s)

	s.Pop()

	fmt.Println(s.Peek())
	fmt.Println(s.Min())
	util.DumpSliceStack(s)

	s.Pop()

	fmt.Println(s.Peek())
	fmt.Println(s.Min())
	util.DumpSliceStack(s)
}
