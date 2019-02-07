package main

import (
	"fmt"
)

func main() {
	s1 := []string{"pale", "pales", "pale", "pale"}
	s2 := []string{"ple", "pale", "bale", "bae"}

	fmt.Println("pattern 1")
	for i, v1 := range s1 {
		v2 := s2[i]
		pattern1(v1, v2)
	}

	fmt.Println("pattern 2")
	for i, v1 := range s1 {
		v2 := s2[i]
		pattern2(v1, v2)
	}
}

// pattern1
func pattern1(s1, s2 string) {
	if len(s1) == len(s2) {
		canEdit(s1, s2)
	} else if len(s1) > len(s2) {
		canInsert(s1, s2)
	} else {
		canInsert(s2, s1)
	}
}

func canEdit(s1, s2 string) {
	var found = false
	for i, v1 := range s1 {
		v2 := rune(s2[i])
		if v1 != v2 {
			if found {
				fmt.Printf("can't edit s1 %s, s2 %s\n", s1, s2)
				return
			}
			found = true
		}
	}

	if found {
		fmt.Printf("can edit s1 %s, s2 %s\n", s1, s2)
	} else {
		fmt.Printf("can't edit s1 %s, s2 %s\n", s1, s2)
	}
}

func canInsert(s1, s2 string) {
	var i1, i2 = 0, 0

	for i1 < len(s1) && i2 < len(s2) {
		v1 := s1[i1]
		v2 := s2[i2]

		if v1 != v2 {
			if i1 != i2 {
				fmt.Printf("can't insert s1 %s, s2 %s\n", s1, s2)
				return
			}
			i1 = i1 + 1
		} else {
			i1 = i1 + 1
			i2 = i2 + 1
		}
	}
	fmt.Printf("can insert s1 %s, s2 %s\n", s1, s2)
}

// pattern2
func pattern2(s1, s2 string) {
	len1 := len(s1)
	len2 := len(s2)

	var longer, shorter string

	if len1 > len2 {
		longer = s1
		shorter = s2
	} else {
		longer = s2
		shorter = s1
	}

	i1, i2 := 0, 0
	found := false

	for i1 < len1 && i2 < len2 {
		c1 := longer[i1]
		c2 := shorter[i2]

		if c1 != c2 {
			if found {
				if len1 == len2 {
					fmt.Printf("can't edit s1 %s, s2 %s\n", s1, s2)
				} else {
					fmt.Printf("can't insert s1 %s, s2 %s\n", s1, s2)
				}
				return
			}
			found = true

			if len1 == len2 {
				i2 = i2 + 1
			}
		} else {
			i2 = i2 + 1
		}
		i1 = i1 + 1
	}

	if len1 == len2 {
		fmt.Printf("can edit s1 %s, s2 %s\n", s1, s2)
	} else {
		fmt.Printf("can insert s1 %s, s2 %s\n", s1, s2)
	}
}
