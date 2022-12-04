package main

import (
	"fmt"
)

func spiral(n int) []int {
	left, top, right, bottom := 0, 0, n-1, n-1
	var m = make([]int, n*n)
	var k = 1

	for left <= right {
		// moving left to right, along top
		for c := left; c <= right; c++ {
			m[top*n+c] = k
			k++
		}
		top++

		// moving top -> bottom, right side
		for r := top; r <= bottom; r++ {
			m[r*n+right] = k
			k++
		}
		right--

		// moving right->left, along bottom
		for c := right; c >= left; c-- {
			m[bottom*n+c] = k
			k++
		}
		bottom--

		// work up left side
		for r := bottom; r >= top; r-- {
			m[r*n+left] = k
			k++
		}
		left++
	}

	return m

}

func main() {
	var N = 5
	var matrix = spiral(N)

	fmt.Println(matrix)

	// var limit = 3 * 3

	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[i]); j++ {

	// 		if i == 0 {
	// 			matrix[i][j] = fill
	// 			fill++
	// 		} else if i == 1 {
	// 			matrix[i][len(matrix[0])-1] = fill
	// 			fill++
	// 		}
	// 	}
	// }

	for i := 0; i < len(matrix); i++ {
		// for j := 0; j < len(matrix[i]); j++ {
		fmt.Printf("|%d|", matrix[i])
		if i%N == N-1 {
			fmt.Println("")
		}
		// }

	}
}
