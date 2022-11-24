// Crear una función que reciba 4 números como parámetros y retorne: valor mínimo, valor
// máximo, promedio, varianza y desviación estándar

package main

import (
	"fmt"
	"math"
)

func max_arr(arr []int) int {
	var max = 0
	for i, tmp := range arr {
		if i == 0 || tmp > max {
			max = tmp
		}
	}
	return max
}

func avg_arr(arr []int) int {
	var avg int
	for i := range arr {
		avg = avg + arr[i]
	}
	return avg / len(arr)
}

func varianza_arr(arr []int, avg int) float64 {
	var sum float64
	var n = len(arr)

	for i := range arr {
		tmp := arr[i] - avg
		tmp2 := math.Pow(float64(tmp), 2)
		sum = sum + tmp2
	}

	return (float64(1) / float64(n)) * sum
}

func desviacion_arr(avg int) float64 {
	return math.Sqrt(float64(avg))
}

func main() {
	var arr [4]int
	arr[0] = 5
	arr[1] = 6
	arr[2] = 7
	arr[3] = 9

	var max = max_arr(arr[:])
	var avg = avg_arr(arr[:])
	var varianza = varianza_arr(arr[:], avg)
	var desviacion = desviacion_arr(int(varianza))

	fmt.Printf("maximo: %d \n", max)
	fmt.Printf("media: %d \n", avg)
	fmt.Printf("varianza: %f \n", varianza)
	fmt.Printf("desviacion: %f \n", desviacion)
}
