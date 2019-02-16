package main

import (
	"fmt"
)

func main() {
	l := []int{-40, -20, -1, 1, 2, 3, 5, 7, 9, 12, 13}

	fmt.Println(pattern1(l, 0, len(l)-1))

	l2 := []int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}

	fmt.Println(pattern2(l2, 0, len(l2)-1))
}

// if the values are all distinct.
func pattern1(l []int, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if l[mid] == mid {
		return mid
	} else if l[mid] > mid {
		return pattern1(l, start, mid-1)
	} else {
		return pattern1(l, mid+1, end)
	}
}

// if the values are not distinct.
func pattern2(l []int, start, end int) int {
	if end < start {
		return -1
	}

	mid := (start + end) / 2
	midValue := l[mid]

	if mid == midValue {
		return mid
	}

	leftIndex := min(mid-1, midValue)

	left := pattern2(l, start, leftIndex)
	if left >= 0 {
		return left
	}

	rightIndex := max(mid+1, midValue)
	return pattern2(l, rightIndex, end)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
