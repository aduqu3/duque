package main

import (
	"fmt"
)

func main(){
	arr1 := [4]string{"Richard","pC","Materia","curso"}
	
	for i, ele := range arr1{
		fmt.Printf("%v - %v \n", i, ele)
	}
}