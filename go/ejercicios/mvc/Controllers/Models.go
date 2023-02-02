package Controllers

import "fmt"

type Persona struct {
	Id       int
	Nombre   string
	Apellido string
	Edad     int
	Amigos   []Persona
}

func (pers *Persona) Saludar() {
	fmt.Printf("Hola soy %s %s tengo %d\n", pers.Nombre, pers.Apellido, pers.Edad)
}

func (pers *Persona) ListAmigos() {
	fmt.Printf("Amigos de %s %s\n", pers.Nombre, pers.Apellido)
	for _, item := range pers.Amigos {
		// fmt.Printf("\t%s %s %d \n",item.Nombre, item.Apellido, item.Edad)
		fmt.Printf("\t%v\n", item)
	}
}

func (pers *Persona) AddAmigo(amigo Persona) {
	pers.Amigos = append(pers.Amigos, amigo)
}

func (pers *Persona) RemoveAmigo(id int) {
	// return append(pers.Amigos[:i], pers.Amigos[i+1:]...)
	for i, amigo := range pers.Amigos {
		if amigo.Id != id {
			return
		} else {
			pers.Amigos = append(pers.Amigos[:i], pers.Amigos[i+1:]...)
			return
		}
	}
}

func (pers *Persona) UpdateAmigo(id int, data Persona) {
	// return append(pers.Amigos[:i], pers.Amigos[i+1:]...)
	for i, amigo := range pers.Amigos {
		if amigo.Id != id {
			return
		} else {
			// amigo.Nombre = data.Nombre
			// amigo.Apellido = data.Apellido
			// amigo.Edad = data.Edad
			// amigo.Amigos = data.Amigos
			pers.Amigos[i] = data
			// fmt.Println(pers.Amigos[i])
			return
		}
	}
}
