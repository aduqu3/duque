package main

import (
	"fmt"
)

func main(){
	n:= 10
	f := make([]int, n+1, n+2)
	fmt.Println(f)
	fmt.Println(len(f))
}