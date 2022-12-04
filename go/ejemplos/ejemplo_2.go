// for and "while"

package main

import (
	"fmt"
)

func main() {
	var i = 0

	// while
	for ; i < 10 ; {
		fmt.Println(i)
		i++
	}

	// do while
	// i = 0
	for {
		fmt.Println(i)
		if i == 10 {
			break
		}
		i++
	}

}
