package Models

type Profesor struct {
	Nombre   string
	Facultad string
}

type Users struct {
	Estudiantes []Estudiante
	Profesores  []Profesor
}

type Curso struct {
	Id     int
	Nombre string
}

type Estudiante struct {
	Nombre string
	Edad   int
	Cursos []Curso
}
