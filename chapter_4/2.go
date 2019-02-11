package main

import (
	"../util"
	"fmt"
)

func main() {
	seed := []int{1, 2, 3, 4, 5, 6, 7}

	n := util.CreateMinimalBST(seed)

	fmt.Println(n)
}
