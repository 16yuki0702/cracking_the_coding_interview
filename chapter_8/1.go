package main

import (
	"fmt"
)

func countWays(n int, memo []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if memo[n] > 0 {
		return memo[n]
	} else {
		memo[n] = countWays(n-1, memo) + countWays(n-2, memo) + countWays(n-3, memo)
		return memo[n]
	}
}

func main() {
	n := 10
	memo := make([]int, n+1)
	fmt.Println(countWays(n, memo))
}
