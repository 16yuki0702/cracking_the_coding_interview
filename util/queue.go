package util

type DoubleStackQueue struct {
	inOrder *SliceStack
	reverse *SliceStack
}

func NewDoubleStackQueue() *DoubleStackQueue {
	return &DoubleStackQueue{
		inOrder: NewSliceStack(),
		reverse: NewSliceStack(),
	}
}

func (q *DoubleStackQueue) Enqueue(v int) {
	q.inOrder.Push(v)
}

func (q *DoubleStackQueue) Dequeue() int {
	q.shift()
	return q.reverse.Pop()
}

func (q *DoubleStackQueue) Peek() int {
	q.shift()
	return q.reverse.Peek()
}

func (q *DoubleStackQueue) shift() {
	if q.reverse.Peek() != -1 {
		return
	}
	for e := q.inOrder.Pop(); e != -1; e = q.inOrder.Pop() {
		q.reverse.Push(e)
	}
}
