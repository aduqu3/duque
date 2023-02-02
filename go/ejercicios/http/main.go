package main

import (
	"fmt"
	"html/template"
	"net/http"
)


type Estudiante struct{
	Nombre string
	Edad int
	Curso string
}


func Index(response http.ResponseWriter, request *http.Request) {

	est1 := Estudiante{Nombre: "Andres Duque", Edad: 23, Curso: "Golang"}

	template, err := template.ParseFiles("Templates/index.html")

	if err != nil {
		fmt.Fprint(response, "<h1>RECURSO NO ENCONTRADO</h1>")
	} else {
		template.Execute(response, est1)
	}
}

func main() {
	var ADDRESS = "localhost:8000"

	http.HandleFunc("/", Index)

	fmt.Printf("Listening in http://%s\n", ADDRESS)
	http.ListenAndServe(ADDRESS, nil)
}
