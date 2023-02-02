package main

import (
	"fmt"
)

func main() {
	var mxn [4][4]string
	// fmt.Println(mxn)

	for i := 0; i < len(mxn); i++ {
		for j := 0; j < len(mxn[0]); j++ {
			if i == 0 {
				mxn[0][j] = "*"
			}

			if i == 3 {
				mxn[3][j] = "*"
			}

			if j == 0 {
				mxn[i][0] = "*"
			}

			if j == 3 {
				mxn[i][3] = "*"
			}
		}
	}

	for i := 0; i < len(mxn); i++ {
		for j := 0; j < len(mxn[0]); j++ {
			if mxn[i][j] == "*" {
				fmt.Printf("|%s|", mxn[i][j])
			} else {
				fmt.Printf("| %s|", mxn[i][j])
			}

		}
		fmt.Println(" ")
	}
}
