package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"parcial2/Models"
	"sort"
)

func Promedio_Curso(arr []Models.Curso) float64 {
	sum := 0.0
	for _, item := range arr {
		sum += float64(item.Nota)
	}
	avg := (float64(sum)) / (float64(len(arr)))

	return avg
}

func Promedio_Curso_Student(arr []Models.Student_Curso) float64 {
	sum := 0.0
	for _, item := range arr {
		sum += float64(item.Curso.Nota)
	}
	avg := (float64(sum)) / (float64(len(arr)))

	return avg
}

func Index(writer http.ResponseWriter, resp *http.Request) {
	//fmt.Fprint(writer, "<h1>Hola mundo</h1>")

	template, err := template.ParseFiles("Templates/index.html")

	var response = "Hola mundo"

	if err != nil {
		fmt.Fprint(writer, "<h1>Página no encontrada<h1/>")
		panic(err)
	} else {
		template.Execute(writer, response) //renderizar html
	}
}

func Algebra(writer http.ResponseWriter, resp *http.Request) {
	//fmt.Fprint(writer, "<h1>Hola mundo</h1>")

	template, err := template.ParseFiles("Templates/algebra.html")

	fileContent, err := os.Open("./Static/generated.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var JDatos Models.Datos
	json.Unmarshal(byteResult, &JDatos)

	// fmt.Println(JDatos.Datos)

	// enviar datos
	type JResponse struct {
		Curso       string
		Promedio    float64
		Estudiantes []Models.Student_Curso
	}

	var Algebra_arr []Models.Student_Curso
	var Calculo_arr []Models.Student_Curso
	var Poo_arr []Models.Student_Curso
	var Ctd_arr []Models.Student_Curso

	// recorrer datos del json
	for _, item := range JDatos.Datos {

		// recorrer cursos para obtener estudiantes de ese curso
		for _, item2 := range item.Cursos {

			if item2.Curso == "Algebra lineal" {
				tmp := Models.Student_Curso{Estudiante: item, Curso: item2}
				Algebra_arr = append(Algebra_arr, tmp)

			} else if item2.Curso == "Calculo diferencial" {
				tmp := Models.Student_Curso{Estudiante: item, Curso: item2}
				Calculo_arr = append(Calculo_arr, tmp)

			} else if item2.Curso == "POO" {
				tmp := Models.Student_Curso{Estudiante: item, Curso: item2}
				Poo_arr = append(Poo_arr, tmp)

			} else if item2.Curso == "CTD" {
				tmp := Models.Student_Curso{Estudiante: item, Curso: item2}
				Ctd_arr = append(Ctd_arr, tmp)
			}
		}
	}

	// 10 estudiantes con notas mas bajas por curso
	sort.SliceStable(Algebra_arr, func(i, j int) bool {
		return Algebra_arr[i].Curso.Nota < Algebra_arr[j].Curso.Nota
	})

	sort.SliceStable(Calculo_arr, func(i, j int) bool {
		return Calculo_arr[i].Curso.Nota < Calculo_arr[j].Curso.Nota
	})

	sort.SliceStable(Poo_arr, func(i, j int) bool {
		return Poo_arr[i].Curso.Nota < Poo_arr[j].Curso.Nota
	})

	sort.SliceStable(Ctd_arr, func(i, j int) bool {
		return Ctd_arr[i].Curso.Nota < Ctd_arr[j].Curso.Nota
	})

	var curso_1_avg = Promedio_Curso_Student(Algebra_arr)
	var curso_2_avg = Promedio_Curso_Student(Calculo_arr)
	var curso_3_avg = Promedio_Curso_Student(Poo_arr)
	var curso_4_avg = Promedio_Curso_Student(Ctd_arr)

	var response []JResponse
	response = append(response, JResponse{Curso: Algebra_arr[0].Curso.Curso, Promedio: curso_1_avg, Estudiantes: Algebra_arr[:10]})
	response = append(response, JResponse{Curso: Calculo_arr[0].Curso.Curso, Promedio: curso_2_avg, Estudiantes: Calculo_arr[:10]})
	response = append(response, JResponse{Curso: Poo_arr[0].Curso.Curso, Promedio: curso_3_avg, Estudiantes: Poo_arr[:10]})
	response = append(response, JResponse{Curso: Ctd_arr[0].Curso.Curso, Promedio: curso_4_avg, Estudiantes: Ctd_arr[:10]})

	if err != nil {
		fmt.Fprint(writer, "<h1>Página no encontrada<h1/>")
		panic(err)
	} else {
		template.Execute(writer, response) //renderizar html
	}
}

func Best(writer http.ResponseWriter, resp *http.Request) {

	template, err := template.ParseFiles("Templates/best.html")

	fileContent, err := os.Open("./Static/generated.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var JDatos Models.Datos
	json.Unmarshal(byteResult, &JDatos)

	var maximo = 0.0
	var index int
	for i, item := range JDatos.Datos {
		tmp2 := Promedio_Curso(item.Cursos)
		if tmp2 > maximo {
			maximo = tmp2
			index = i
		}
	}

	if err != nil {
		fmt.Fprint(writer, "<h1>Página no encontrada<h1/>")
		panic(err)
	} else {
		template.Execute(writer, JDatos.Datos[index]) //renderizar html
	}
}

func Sexo(writer http.ResponseWriter, resp *http.Request) {
	//fmt.Fprint(writer, "<h1>Hola mundo</h1>")

	template, err := template.ParseFiles("Templates/sexo.html")

	fileContent, err := os.Open("./Static/generated.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var JDatos Models.Datos
	json.Unmarshal(byteResult, &JDatos)

	var maximo_mujer = 0
	var index_mujer int

	var maximo_hombre = 0
	var index_hombre int

	for i, item := range JDatos.Datos {
		if item.Gender == "female" {
			tmp2 := item.Edad
			if tmp2 > maximo_mujer {
				maximo_mujer = tmp2
				index_mujer = i
			}
		} else {
			tmp2 := item.Edad
			if tmp2 > maximo_hombre {
				maximo_hombre = tmp2
				index_hombre = i
			}
		}
	}

	var mayores []Models.Dato
	mayores = append(mayores, JDatos.Datos[index_mujer])
	mayores = append(mayores, JDatos.Datos[index_hombre])

	// fmt.Println(JDatos.Datos)

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
		template.Execute(writer, mayores) //renderizar html
	}
}

func main() {
	var ADDRESS = "localhost:8000"

	http.HandleFunc("/", Index)
	http.HandleFunc("/Algebra", Algebra)
	http.HandleFunc("/best", Best)
	http.HandleFunc("/Sexo", Sexo)
	fmt.Println("Starting server...")
	fmt.Printf("Listening in http://%s\n", ADDRESS)
	http.ListenAndServe(ADDRESS, nil)
}
