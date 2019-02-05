package main

import (
	"fmt"
	"strings"
)

var input = "Mr John Smith    "

func main() {
	s := strings.Split(input, "")
	replaceStr(s, 13)
}

func replaceStr(s []string, trueLength int) {
	var blankCount = 0

	for i := 0; i < trueLength; i = i + 1 {
		if s[i] == " " {
			blankCount = blankCount + 1
		}
	}

	index := trueLength + blankCount*2

	for i := trueLength - 1; i > 0; i = i - 1 {
		if s[i] == " " {
			s[index-1] = "0"
			s[index-2] = "2"
			s[index-3] = "%"
			index = index - 3
		} else {
			s[index-1] = s[i]
			index = index - 1
		}
	}

	fmt.Println(s)
}
