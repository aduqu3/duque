package Models

// type Profesor struct {
// 	Nombre   string
// 	Facultad string
// }

// type Users struct {
// 	Estudiantes []Estudiante
// 	Profesores  []Profesor
// }

// type Curso struct {
// 	Id     int
// 	Nombre string
// }

// type Estudiante struct {
// 	Nombre string
// 	Edad   int
// 	Cursos []Curso
// }

type Users2 struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
	LinkedIn string `json:"linkedin"`
}

type Personas struct {
	Personas []Persona `json:"personas"`
}

type Persona struct {
	Nombre string  `json:"nombre"`
	Cursos []Curso `json:"cursos"`
}

type Curso struct {
	Nombre   string `json:"nombre"`
	Facultad string `json:"facultad"`
}
