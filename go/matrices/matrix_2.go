package main

import (
	"fmt"
)

func main(){
	
	mxn := make([][]string, 4)
	for i := 0; i < len(mxn); i++ {
		mxn[i] = make([]string, 4)
	}

	fmt.Println(mxn)

	for i:= 0; i < len(mxn); i++{
		for j := 0; j < len(mxn[0]); j++{
			if i == 0 && j == 0 {
				mxn[0][0] = "*"
			}

			if i == 0 && j == len(mxn)-1 {
				mxn[0][len(mxn)-1] = "*"
			}

			if i == len(mxn)-1 && j == len(mxn[0])-1 {
				mxn[len(mxn)-1][len(mxn[0])-1] = "*"
			}

			if i == len(mxn)-1 && j == 0 {
				mxn[len(mxn)-1][0] = "*"
			}

			// if i == len(mxn)-1 && j == 0 {
			// 	mxn[len(mxn)-1][0] = "*"
			// }

			
		}
	}
	
	fmt.Println(mxn)

	for i:= 0; i < len(mxn); i++{
		for j := 0; j < len(mxn[0]); j++{
			if mxn[i][j] == "*"{
				fmt.Printf("|%s|", mxn[i][j])
			}else{
				fmt.Printf("| %s|", mxn[i][j])
			}
		}
		fmt.Println(" ")
	}
}