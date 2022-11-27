// Crear una función por cada área y perímetro de las siguientes figuras geométricas donde reciba
// como parámetro la(s) variable(s) requerida(s): https://www.proferecursos.com/areas-y-perimetros/

package main

import (
	"fmt"
	"math"
)

func a_cuadrado(L float64){
	var area = L*L
	var perimetro = 4*L
	fmt.Printf("Cuadrado \t area: %f cm² \t perimetro: %f cm \n", area, perimetro)
}

func a_rectangulo(a float64, b float64){
	var area = a*b
	var perimetro = 2*a + 2*b
	fmt.Printf("Rectandulo \t area: %f cm² \t perimetro: %f cm \n", area, perimetro)
}

func a_triangulo(b float64, h float64){
	var area = (b*h)/2
	var perimetro = 3*b
	fmt.Printf("Triangulo \t area: %f cm² \t perimetro: %f cm \n", area, perimetro)
}

func a_rombo(D float64, d float64, L float64){
	var area = (D*d)/2
	var perimetro = 4*L
	fmt.Printf("Rombo \t area: %f cm² \t perimetro: %f cm \n", area, perimetro)
}

func a_pentagono(P float64, a float64, L float64){
	var area = (P*a)/2
	var perimetro = 5*L
	fmt.Printf("Pentagono \t area: %f cm² \t perimetro: %f cm \n", area, perimetro)
}

func a_hexagono(P float64, a float64, L float64){
	var area = (P*a)/2
	var perimetro = 6*L
	fmt.Printf("Hexagono \t area: %f cm² \t perimetro: %f cm \n", area, perimetro)
}

func a_circulo(r float64){
	var area = math.Pi * (math.Pow(r, 2))
	var perimetro = 2*math.Pi*r
	fmt.Printf("Circulo \t area: %f cm² \t perimetro: %f cm \n", area, perimetro)
}

func a_trapecio(B float64, b float64, h float64, L float64){
	var area = ((B+b)/2)*h
	var perimetro = B+b+(2*L)
	fmt.Printf("Trapecio \t area: %f cm² \t perimetro: %f cm \n", area, perimetro)
}

func a_paralelogramo(B float64, h float64, a float64){
	var area = B*h
	var perimetro = 2*(B+a)
	fmt.Printf("Paralelogramo \t area: %f cm² \t perimetro: %f cm \n", area, perimetro)
}

func main() {
	fmt.Println("Calculador de areas")
	a_paralelogramo(8,6,7)
}