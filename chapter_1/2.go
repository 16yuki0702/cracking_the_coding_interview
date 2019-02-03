package main

import (
	"../util"
	"fmt"
)

func main() {
	s1 := util.ReadLine()
	s2 := util.ReadLine()

	pattern1(s1, s2)
	pattern2(s1, s2)
}

func pattern1(s1, s2 string) {
	s1 = util.SortString(s1)
	s2 = util.SortString(s2)

	if s1 != s2 {
		fmt.Println("not permutation")
		return
	}
	fmt.Println("permutation")
}

func pattern2(s1, s2 string) {
	if len(s1) != len(s2) {
		fmt.Println("not permutation")
		return
	}

	m := [256]int{}

	for _, v := range s1 {
		i := uint(v - 'a')
		m[i] = m[i] + 1
	}

	for _, v := range s2 {
		i := uint(v - 'a')
		m[i] = m[i] - 1
		if m[i] < 0 {
			fmt.Println("not permutation")
			return
		}
	}
	fmt.Println("permutation")
}
