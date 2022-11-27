// Crear una función que reciba una cadena de texto y retorne un true si es palíndromo y un false
// caso contrario

// taller Andres Duque 160003812, Fredy segura 160003830

package main

import (
	"fmt"
)

func reverse(arr []rune) bool {
	
	end := len(arr) - 1
	// fmt.Printf("%d end \n", len(arr))
	
	var count int
	for i := 0; i < len(arr); i++ {
		if arr[end] == arr[i] {
			count++
		}
		end--
	}

	// fmt.Println(count)

	if count == len(arr) {
		fmt.Println("palindroma")
		return true
	}

	fmt.Println("No es palindroma")
	return false

}

func main() {
	var palabra = "casa"
	output := []rune(palabra)
	reverse(output)
}
