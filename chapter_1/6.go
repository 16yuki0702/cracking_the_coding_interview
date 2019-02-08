package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "aabcccccaaa"
	pattern1(str)
	pattern2(str)
	pattern3(str)
}

// pattern1
func pattern1(str string) {
	var result string
	count := 0

	for i, _ := range str {
		count = count + 1

		if len(str) <= i+1 || str[i] != str[i+1] {
			result = result + fmt.Sprintf("%c%d", str[i], count)
			count = 0
		}
	}

	if len(str) < len(result) {
		result = str
	}

	fmt.Printf("result %s\n", result)
}

// pattern2
func pattern2(str string) {
	var result string
	buf := make([]byte, 0, len(str))
	count := 0

	for i, _ := range str {
		count = count + 1

		if len(str) <= i+1 || str[i] != str[i+1] {
			buf = append(buf, fmt.Sprintf("%c%d", str[i], count)...)
			count = 0
		}
	}

	result = string(buf)
	if len(str) < len(result) {
		result = str
	}

	fmt.Printf("result %s\n", result)
}

// pattern3
func pattern3(str string) {
	compressedCount := countCompressed(str)

	if len(str) < compressedCount {
		fmt.Printf("result %s\n", str)
		return
	}

	buf := make([]byte, compressedCount)
	count := 0

	for i, _ := range str {
		count = count + 1

		if len(str) <= i+1 || str[i] != str[i+1] {
			buf = append(buf, fmt.Sprintf("%c%d", str[i], count)...)
			count = 0
		}
	}

	result := string(buf)

	fmt.Printf("result %s\n", result)
}

func countCompressed(str string) int {
	compressedCount := 0
	count := 0

	for i, _ := range str {
		count = count + 1

		if len(str) <= i+1 || str[i] != str[i+1] {
			countStr := strconv.Itoa(count)
			compressedCount = compressedCount + (1 + len(countStr))
			count = 0
		}
	}

	return compressedCount
}
