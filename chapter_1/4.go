package main

import (
	"fmt"
)

func main() {
	s := "tact coa"

	pattern1(s)
	pattern2(s)
	pattern3(s)
}

// pattern1
func pattern1(s string) {
	max := uint('z' - 'a')
	counter := make([]uint, max)

	for _, v := range s {
		c := uint(v - 'a')
		if c < 0 || max < c {
			continue
		}
		counter[c] = counter[c] + 1
	}

	found := false

	for _, v := range counter {
		if v%2 == 1 {
			if found {
				fmt.Println("not palindromes")
				return
			}
			found = true
		}
	}
	fmt.Println("palindromes")
}

// pattern2
func pattern2(s string) {
	max := uint('z' - 'a')
	counter := make([]uint, max)
	oddCount := 0

	for _, v := range s {
		c := uint(v - 'a')
		if c < 0 || max < c {
			continue
		}

		counter[c] = counter[c] + 1

		if counter[c]%2 == 1 {
			oddCount = oddCount + 1
		} else {
			oddCount = oddCount - 1
		}
	}

	if oddCount <= 1 {
		fmt.Println("palindromes")
	} else {
		fmt.Println("not palindromes")
	}
}

// pattern3
func pattern3(s string) {
	checker := 0

	for _, v := range s {
		c := uint(v - 'a')
		f := 1 << c

		if checker&f > 0 {
			checker ^= f
		} else {
			checker |= f
		}
	}

	if checker == 0 {
		fmt.Println("palindromes")
		return
	}

	if (checker-1)&checker == 0 {
		fmt.Println("palindromes")
		return
	}

	fmt.Println("not palindromes")
}
