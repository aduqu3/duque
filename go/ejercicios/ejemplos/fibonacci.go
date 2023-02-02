package main

import (
	"fmt"
)

func fibonacci(fibo int) (result int) {
	result = fibo
	if result == 0 {
		return 0
	}
	if result == 1 {
		return 1
	}
	return fibonacci(result-1) + fibonacci(result-2)

}

func main() {
	var indice int
	fmt.Scanln(&indice)
	result := fibonacci(indice)
	println(result)
}
