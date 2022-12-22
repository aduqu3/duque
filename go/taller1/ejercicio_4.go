// Declarar un arreglo de 10 posiciones y con inicialización directa (no usar para este taller ciclos)
// llenarla con números aleatorios enteros entre -10 y 10 y si el número aleatorio es par se
// agregará cero en su lugar. Si el número es impar y es múltiplo de 3 entonces se guardará como
// 999 en su lugar. Para esto último, se debe declarar una función que genere un aleatorio y
// retorne el valor correspondiente según las restricciones dadas

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genrand() int {
	rand.Seed(time.Now().UnixNano())
	var min = -10
	var max = 10
	var number = rand.Intn(max-min) + min

	if number % 2 == 0 { //par
		number = 0
	} else if number % 3 == 0 { // ya se sabe que es impar, solo falta saber si es multiplo de 3
		number = 999
	}
	return number
}

func main() {
	var arr [10]int
	arr[0] = genrand()
	arr[1] = genrand()
	arr[2] = genrand()
	arr[3] = genrand()
	arr[4] = genrand()
	arr[5] = genrand()
	arr[6] = genrand()
	arr[7] = genrand()
	arr[8] = genrand()
	arr[9] = genrand()
	fmt.Println(arr)
}
