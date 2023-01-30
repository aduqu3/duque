package repository

import (
	"api-fiber-taller/database"
	"api-fiber-taller/model"
)

// GetAllEstudiantes query all estudiantes
func GetAllEstudiantes() ([]model.Estudiante, error) {

	db := database.DB
	var estudiantes []model.Estudiante
	result := db.Find(&estudiantes)

	if result.Error != nil {
		return estudiantes, result.Error
	}

	return estudiantes, nil
}

func AvgCurso(cursos []model.EstudianteCurso) float64 {
	sum := 0.0
	for _, item := range cursos {
		sum += float64(item.Nota)
	}
	avg := sum / (float64(len(cursos)))
	return avg
}

type MejorEstudiante struct {
	Estudiante model.Estudiante        `json:"estudiante"`
	Cursos     []model.EstudianteCurso `json:"cursos"`
}

// GetMejorEstudiante query mejor estudiante
func GetMejorEstudiante() (MejorEstudiante, error) {
	var est []model.Estudiante
	var err error
	est, err = GetAllEstudiantes()

	if err != nil {
		return MejorEstudiante{}, err
	}

	var maximo = 0.0
	var index int
	for i, ele := range est {
		cursos, err := GetEstudianteCursobyEstudianteID(int(ele.ID))
		if err != nil {
			return MejorEstudiante{}, err
		}
		tmp2 := AvgCurso(cursos)
		if tmp2 > maximo {
			maximo = tmp2
			index = i
		}
	}

	best := est[index]
	best_cursos, err := GetEstudianteCursobyEstudianteID(int(best.ID))

	for i, item := range best_cursos {
		best_cursos[i].Curso, err = GetCurso(item.CursoId)
		if err != nil {
			return MejorEstudiante{}, err
		}
	}

	if err != nil {
		return MejorEstudiante{}, err
	}

	response := MejorEstudiante{
		Estudiante: best,
		Cursos:     best_cursos,
	}

	return response, nil
}

func GetMoreOldEstudiante() ([]model.Estudiante, error) {
	var est []model.Estudiante
	var err error
	est, err = GetAllEstudiantes()

	if err != nil {
		return est, err
	}

	var maximo_mujer = 0
	var index_mujer int

	var maximo_hombre = 0
	var index_hombre int

	for i, item := range est {
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

	var old []model.Estudiante
	old = append(old, est[index_hombre])
	old = append(old, est[index_mujer])

	return old, nil
}

func GetCursoAvg() ([]model.EstudianteCurso, error) {
	var cursos []model.Curso
	var err error
	cursos, err = GetAllCursos()
	if err != nil {
		return nil, err
	}

	var avg_curso []model.EstudianteCurso
	for _, item := range cursos {
		e_c, err := GetEstudianteCursoByCursoID(int(item.ID))
		if err != nil {
			return nil, err
		}
		var avg model.EstudianteCurso
		avg.CursoId = int(item.ID)
		avg.Curso.Curso = item.Curso
		avg.Nota = AvgCurso(e_c)
		avg_curso = append(avg_curso, avg)
	}

	return avg_curso, nil
}
