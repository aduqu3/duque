package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"web2/Models"
)

func Index(writer http.ResponseWriter, resp *http.Request) {
	//fmt.Fprint(writer, "<h1>Hola mundo</h1>")

	template, err := template.ParseFiles("Templates/index.html")

	// funciona Users
	// load json file
	// fileContent, err := os.Open("./Static/users.json")

	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Println("The File is opened successfully...")

	// defer fileContent.Close()

	// byteResult, _ := ioutil.ReadAll(fileContent)

	// var JUsers Models.Users2
	// json.Unmarshal(byteResult, &JUsers)
	// funciona

	fileContent, err := os.Open("./Static/persons.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var JUsers Models.Personas
	json.Unmarshal(byteResult, &JUsers)

	// fmt.Println(JUsers.Users)

	// for i := 0; i < len(JUsers.Users); i++ {
	// 	fmt.Println("User Name: " + JUsers.Users[i].Name)
	// 	fmt.Println("Facebook Url: " + JUsers.Users[i].Social.LinkedIn)
	// }

	// JUsers.Users --> arreglo

	// var res map[string]interface{}
	// json.Unmarshal([]byte(byteResult), &res)

	// fmt.Println(res["users"])

	// Curso1Est1 := Models.Curso{Id: 1, Nombre: "Promgramación en Go"}
	// Curso2Est1 := Models.Curso{Id: 2, Nombre: "Promgramación en Java"}
	// Curso3Est1 := Models.Curso{Id: 3, Nombre: "Promgramación en Python"}
	// Curso4Est1 := Models.Curso{Id: 4, Nombre: "Promgramación en C#"}

	// sliceCursos := []Models.Curso{Curso1Est1, Curso2Est1, Curso3Est1, Curso4Est1}

	// est1 := Models.Estudiante{Nombre: "Andres Duque", Edad: 23, Cursos: sliceCursos}

	// Curso1Est2 := Models.Curso{Id: 5, Nombre: "Promgramación en Ruby"}
	// Curso2Est2 := Models.Curso{Id: 6, Nombre: "Promgramación en R"}
	// Curso3Est2 := Models.Curso{Id: 7, Nombre: "Promgramación en Haskell"}
	// Curso4Est2 := Models.Curso{Id: 8, Nombre: "Promgramación en Javascript"}

	// sliceCursos2 := []Models.Curso{Curso1Est2, Curso2Est2, Curso3Est2, Curso4Est2}

	// est2 := Models.Estudiante{Nombre: "Juan Benito", Edad: 23, Cursos: sliceCursos2}

	// prof1 := Models.Profesor{Nombre: "Angel Cruz", Facultad: "Ingeniería"}
	// prof2 := Models.Profesor{Nombre: "Roger Calderon", Facultad: "Ingeniería"}
	// prof3 := Models.Profesor{Nombre: "Carlos Navarro", Facultad: "Biología"}

	// datosEst := []Models.Estudiante{est1, est2}
	// datosProf := []Models.Profesor{prof1, prof2, prof3}

	// datos := Models.Users{Estudiantes: datosEst, Profesores: datosProf}

	if err != nil {
		fmt.Fprint(writer, "<h1>Página no encontrada<h1/>")
		panic(err)
	} else {
		template.Execute(writer, JUsers) //renderizar html
	}
}

func main() {
	var ADDRESS = "localhost:8000"

	http.HandleFunc("/", Index)
	fmt.Println("Starting server...")
	fmt.Printf("Listening in http://%s\n", ADDRESS)
	http.ListenAndServe(ADDRESS, nil)
}
