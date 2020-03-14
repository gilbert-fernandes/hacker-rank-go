package main

import "fmt"

/*
 * Complete the 'dynamicArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func dynamicArray(n int32, queries [][]int32) []int32 {
	// so first we are asked to dynamically create an array when supplied size n
	sequences := make([][]int32, n)
	// parse queries. seems the first number is what is asked
	// followed by x and y to calculate
	for _, line := range queries {
		query := line[0]
		x := line[1]
		y := line[2]
		if query == 1 {
			// add y to sequences ? no idea
		} else if query == 2 {

		} else {
			fmt.Println("error : query value not in [1,2] range")
		}
	}

	var toto []int32
	return toto
}

/*
 * the problem is worded so poorly that after reading 5 times the problem
 * i have no idea what is asked for. in a professional environment, anyone daring
 * to write specs that way would not only receive a "no" but would probably be asked
 * to go learn how to word things properly or find another job
 * ikbalkazar just go fuck yourself
 *
 */
func main() {
	sample0 := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}
	fmt.Println(dynamicArray(5, sample0))
}
