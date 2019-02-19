package main

import (
	"container/list"
	"fmt"
)

func main() {
	l1 := list.New()
	l1.PushFront(6)
	l1.PushFront(1)
	l1.PushFront(7)

	l2 := list.New()
	l2.PushFront(2)
	l2.PushFront(9)
	l2.PushFront(5)

	pattern1(l1, l2)

	l3 := list.New()
	l3.PushFront(4)
	l3.PushFront(3)
	l3.PushFront(2)
	l3.PushFront(1)

	l4 := list.New()
	l4.PushFront(7)
	l4.PushFront(6)
	l4.PushFront(5)

	pattern2(l3, l4)
}

func pattern1(l1, l2 *list.List) {
	result := list.New()
	addList(result, l1.Front(), l2.Front(), 0)

	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

func addList(result *list.List, e1, e2 *list.Element, carry int) {
	if e1 == nil && e2 == nil && carry == 0 {
		return
	}

	v := carry

	if e1 != nil {
		v = v + e1.Value.(int)
	}
	if e2 != nil {
		v = v + e2.Value.(int)
	}

	result.PushFront(v % 10)

	if e1 != nil || e2 != nil {
		nextCarry := 0
		if v >= 10 {
			nextCarry = 1
		}
		var next1 *list.Element
		if e1.Next() != nil {
			next1 = e1.Next()
		}
		var next2 *list.Element
		if e2.Next() != nil {
			next2 = e2.Next()
		}
		addList(result, next1, next2, nextCarry)
	}
}

// pattern2
type Sum struct {
	sum   *list.List
	carry int
}

func NewSum() *Sum {
	return &Sum{
		sum:   list.New(),
		carry: 0,
	}
}

func pattern2(l1, l2 *list.List) {
	len1 := l1.Len()
	len2 := l2.Len()

	if len1 < len2 {
		padList(l1, len2-len1)
	} else if len2 < len1 {
		padList(l2, len1-len2)
	}

	s := NewSum()

	addSum(l1.Front(), l2.Front(), s)

	if s.carry != 0 {
		s.sum.PushFront(s.carry)
	}

	for e := s.sum.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

func addSum(e1, e2 *list.Element, s *Sum) {
	if e1 == nil && e2 == nil {
		return
	}

	addSum(e1.Next(), e2.Next(), s)

	v := e1.Value.(int) + e2.Value.(int) + s.carry

	s.sum.PushFront(v % 10)
	s.carry = v / 10
}

func padList(l *list.List, padNum int) {
	for i := 0; i < padNum; i = i + 1 {
		l.PushFront(0)
	}
}
