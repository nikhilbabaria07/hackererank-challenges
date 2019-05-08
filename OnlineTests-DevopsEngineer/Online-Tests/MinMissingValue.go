package main

import (
	"fmt"
	"sort"
)

func maxVal(A []int) int {
	sort.Ints(A)
	return A[len(A)-1]
}

func Solution1(A []int) int {
	// write your code in Go 1.4

	//var flag bool
	var res int
	mV := maxVal(A)
	for i := 1; i < mV; i++ {
		for _, value := range A {
			if value != i {
				res = i
			} else {
				res = mV
			}
		}
	}

	return res
}

func main() {
	fmt.Println([]int{1, 2, 2, 2, 2, 3, 1, 1, 5, 6, 7})
	fmt.Println(Solution([]int{1, 2, 2, 2, 2, 3, 1, 1, 5, 6, 7}))
}
