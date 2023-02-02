// Una función que reciba (pida por reclado) una cadena de texto de una ecuación de segundo
// orden y que halle las raíces usando la fórmula general. Por ejemplo: “x^2+8x+15=0” o
// “8x+15+x^2=0” voy a obtener x1 = -3 y x2 = -5.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func find_sqrt_x(a float64, b float64, c float64) (float64, float64){
	var raiz = math.Sqrt((b*b)-(4*a*c))
	var x_neg = (-b - raiz) / (2*a)
	var x_pos = (-b + raiz) / (2*a)
	return x_pos, x_neg
}

func find_variables(fx string) {
	// a := []rune(fx)
	// fmt.Println(string(a))

	var a float64
	var b float64
	var c float64

	var arr = strings.Split(fx, "+")
	// fmt.Println(arr)	

	for i := 0; i < len(arr); i++ {
		if _, err := strconv.Atoi(arr[i]); err == nil {
			c, _ = strconv.ParseFloat(arr[i], 64)
		} else {
			temp := strings.Split(arr[i], "")
			// fmt.Println(len(temp))
			if len(temp) == 2 {
				b, _ = strconv.ParseFloat(temp[0], 64)
			} else {
				temp2 := strings.Split(arr[i], "=")
				if temp2[0] == "x^2" {
					// a, _ = strconv.ParseFloat(temp2[0], 64)
					// fmt.Println(temp2[0])
					a = 1
				}else {
					// if a == 0 {
					// 	temp4 := strings.Split(temp2[0], "x^2")
					// 	a, _ = strconv.ParseFloat(temp4[0], 64)
					// }

					if b == 0 {
						temp3 := strings.Split(temp2[0], "x")
						b, _ = strconv.ParseFloat(temp3[0], 64)
					}

					if c == 0{
						c, _ = strconv.ParseFloat(temp2[0], 64)
					}
				}
			}

		}
	}

	fmt.Println(fx)
	fmt.Printf("a: %f \t b: %f \t c: %f \n", a, b, c)
	x1, x2 := find_sqrt_x(a,b,c)
	fmt.Printf("Raices: x1 = %f \t x2 = %f \n", x1, x2)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Ingresa una ecuacion de segundo grado: ")
    input, _ := reader.ReadString('\n')

	// var input = "x^2+8x+15=0"
	// var input = "8x+15+x^2=0"
	// var input = "x^2+6+5x=0"
	find_variables(input)
}
