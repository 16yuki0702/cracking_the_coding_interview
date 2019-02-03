package main

import (
	"../util"
	"fmt"
)

func main() {
	s := util.ReadLine()

	pattern1(s)
	pattern2(s)
}

func pattern1(s string) {
	m := make(map[rune]bool)

	for _, v := range s {
		r := v - 'a'

		if m[r] {
			fmt.Println("not unique char")
			return
		}
		m[r] = true
	}

	fmt.Println("unique char")
}

func pattern2(s string) {
	// it only be available if str is lowercase.
	checker := 0

	for _, v := range s {
		r := uint(v - 'a')

		if checker&(1<<r) > 0 {
			fmt.Println("not unique char")
			return
		}
		checker |= (1 << r)
	}

	fmt.Println("unique char")
}
