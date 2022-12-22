// Crear una función que reciba 4 números como parámetros y me retorne un arreglo con los datos
// ordenados. (no se pueden usar ciclos).

package main

import (
	"fmt"
	"sort"
)

func ordenar(a int, b int, c int, d int) []int {
	S := []int{a, b, c, d}
	sort.Ints(S)

	return S
}

func main() {
	fmt.Println(ordenar(3, 4, 8, 1))
}