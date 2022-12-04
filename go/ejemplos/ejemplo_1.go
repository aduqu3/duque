// switch case

package main

import (
	"fmt"
	"time"
)

func main(){
	// today = time.Now()
	switch today := time.Now().Day(); {
		
		case today < 10:
			fmt.Println("hoy es primero")
		
		case today < 28:
			fmt.Println("Es par < 15")
			fallthrough

		case today < 30:
			fmt.Println("> 20 y par")
			fallthrough

		case today == 31:
			fmt.Println("> 20 e impar")
			fallthrough

		default:
			fmt.Println(today)
			fmt.Println("Defecto")
	}

}