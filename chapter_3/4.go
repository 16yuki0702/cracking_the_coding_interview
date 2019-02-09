package main

import (
	"../util"
	"fmt"
)

func main() {
	q := util.NewDoubleStackQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Printf("peek %d\n", q.Peek())
	fmt.Printf("pop %d\n", q.Dequeue())

	fmt.Printf("peek %d\n", q.Peek())
	fmt.Printf("pop %d\n", q.Dequeue())

	fmt.Printf("peek %d\n", q.Peek())
	fmt.Printf("pop %d\n", q.Dequeue())

	fmt.Printf("peek %d\n", q.Peek())
	fmt.Printf("pop %d\n", q.Dequeue())
}
