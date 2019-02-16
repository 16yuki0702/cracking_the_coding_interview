package main

import (
	"../util"
	"fmt"
)

type Tower struct {
	Disks *util.SliceStack
	Index int
}

func NewTower(i int) *Tower {
	return &Tower{
		Disks: util.NewSliceStack(),
		Index: i,
	}
}

func (t *Tower) Add(v int) {
	if !t.Disks.IsEmpty() && t.Disks.Peek() <= v {
		fmt.Printf("Error placing disk %d\n", v)
	} else {
		t.Disks.Push(v)
	}
}

func (t *Tower) MoveTopTo(dest *Tower) {
	p := t.Disks.Pop()
	dest.Disks.Push(p)
}

func (t *Tower) MoveDisks(v int, dest *Tower, buffer *Tower) {
	if v <= 0 {
		return
	}
	t.MoveDisks(v-1, buffer, dest)
	t.MoveTopTo(dest)
	buffer.MoveDisks(v-1, dest, t)
}

func main() {
	n := 3
	towers := make([]*Tower, 3)
	for i := 0; i < 3; i = i + 1 {
		towers[i] = NewTower(i)
	}
	for i := n - 1; i >= 0; i = i - 1 {
		towers[0].Add(i)
	}

	fmt.Println(towers[0].Disks.Size())
	fmt.Println(towers[2].Disks.Size())

	towers[0].MoveDisks(n, towers[2], towers[1])

	fmt.Println(towers[0].Disks.Size())
	fmt.Println(towers[2].Disks.Size())
}
