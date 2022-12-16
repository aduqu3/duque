package main

import (
	"mvc/Controllers"
)

// type Person struct {
// 	name string
// 	age int
// 	job string
// 	salary int
// }

func main() {
	// fmt.Println(Controllers.Persona)

	// var pers Controllers.Persona
	// pers.Nombre = "Andres"
	// pers.Apellido = "Duque"
	// pers.Edad = 23

	// fmt.Printf("%+v\n", pers)
	pers := Controllers.Persona{Id: 1, Nombre: "Andres", Apellido: "Duque", Edad: 23}
	pers.Saludar()

	// pers.ListAmigos()

	amigo1 := Controllers.Persona{Id: 2, Nombre: "Santiago", Apellido: "Cruz", Edad: 23}
	pers.AddAmigo(amigo1)

	pers.ListAmigos()

	// pers.RemoveAmigo(amigo1.Id)
	// fmt.Printf("%+v\n", pers)

	amigo1.Nombre = "Felipe"
	amigo1.Apellido = "Santiago"
	amigo1.AddAmigo(Controllers.Persona{Id: 3, Nombre: "Pepito", Apellido: "Perez", Edad: 23})
	pers.UpdateAmigo(amigo1.Id, amigo1)

	pers.ListAmigos()
}
