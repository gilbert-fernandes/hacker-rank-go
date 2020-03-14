package main

import "fmt"

/*
 * Given a 6 x 6 2D Array arr
 *
 * 1 1 1 0 0 0
 * 0 1 0 0 0 0
 * 1 1 1 0 0 0
 * 0 0 0 0 0 0
 * 0 0 0 0 0 0
 * 0 0 0 0 0 0
 *
 * We define an hourglass A to be a subset of values with indices falling in this pattern in arr's
 * graphical representation:
 *
 * a b c
 *   d
 * e f g
 *
 * There are 16 hourglasses in arr, and an hourglass sum is the sum of an hourglass' values
 * Calculate the hourglass sum for every hourglass in arr then print the maximum hourglass sum
 *
 * For example, given the 2D array:
 *
 * -9  -9  -9   1  1  1
 *  0  -9   0   4  3  2
 * -9  -9  -9   1  2  3
 *  0   0   8   6  6  0
 *  0   0   0  -2  0  0
 *  0   0   1   2  4  0
 *
 * We calculate the following 16 hourglass values:
 *
 * -63, -34, -9, 12,
 * -10,   0, 28, 23,
 * -27, -11, -2, 10,
 *   9,  17, 25, 18
 *
 * Our highest hourglass value is 28 from the hourglass:
 *
 * 0 4 3
 *   1
 * 8 6 6
 *
 * Complete the function hourglassSum in the editor below
 * It should return an integer, the maximum hourglass sum in the array.
 *
 */

func hourglassSum(arr [][]int32) int32 {
	var max int32 = -100
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			a := arr[j][i]
			b := arr[j][i+1]
			c := arr[j][i+2]
			d := arr[j+1][i+1]
			e := arr[j+2][i]
			f := arr[j+2][i+1]
			g := arr[j+2][i+2]
			sum := a + b + c + d + e + f + g
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	/*
		1 1 1 0 0 0
		0 1 0 0 0 0
		1 1 1 0 0 0
		0 0 2 4 4 0
		0 0 0 2 0 0
		0 0 1 2 4 0
	*/
	sample0 := [][]int32{{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	fmt.Println(hourglassSum(sample0))
}
