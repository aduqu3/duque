package main

import (
	"fmt"
)

func spiral(n int) [][] int{
	// Defining the boundaries of the matrix.
	left, top, right, bottom := 0, 0, n-1, n-1

	// create matrix (n x n)
	var m = make([][]int, n)
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, n)
	}

	// start point, initial number
	var k = 1
	
	for left <= right {
		// moving left to right, along top
		for c := left; c <= right; c++ {
            m[top][c] = k
            k++
        }
		// Since we have traversed the whole first
      	// row, move down to the next row.
		top++

		// moving top -> bottom, right side
		for r := top; r <= bottom; r++ {
            m[r][right] = k
            k++
        }
		// Since we have traversed the whole last
      	// column, move left to the previous column.
        right--

		// moving right -> left, along bottom
        for c := right; c >= left; c-- {
            m[bottom][c] = k
            k++
        }
		// Since we have traversed the whole last
      	// row, move up to the previous row.
        bottom--

        // work up left side
        for r := bottom; r >= top; r-- {
            m[r][left] = k
            k++
        }
		// Since we have traversed the whole first
      	// col, move right to the next column.
        left++
	}

	return m

}

func main() {
	var N = 4
	var matrix = spiral(N)
	
	// fmt.Println(matrix)
	
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("|%d|", matrix[i][j])
		}
		fmt.Println("")
	}

}
