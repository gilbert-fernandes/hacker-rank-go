package main

import (
	"fmt"
)

func reverseArray(a []int32) []int32 {
	// could use sort.Reverse but uses an int[] instead of an int32[] slice
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - i - 1
		a[i], a[j] = a[j], a[i]
	}
	return a
}

/*
 * An array is a type of data structure that stores elements of the same type in a contiguous block of memory
 * In an array, A, of size N, each memory location has some unique index i (where 0 <= i < N)
 * that can be referenced as A[i]
 * Given an array A of N integers, print each element in reverse order as a single line of space-separated integers
 *
 */
func main() {
	// sample0
	// 1 4 3 2
	sample0 := []int32{1, 4, 3, 2}
	fmt.Println(reverseArray(sample0))
}
