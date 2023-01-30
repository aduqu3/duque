package repository

import (
	"api-fiber-taller/database"
	"api-fiber-taller/model"
)

// GetAllEstudiantesCurso query all estudiantes_curso
func GetAllEstudianteCurso() ([]model.EstudianteCurso, error) {

	db := database.DB
	var e_c []model.EstudianteCurso
	result := db.Find(&e_c)

	if result.Error != nil {
		return e_c, result.Error
	}

	return e_c, nil
}

// GetEstudianteCursobyEstudianteID query estudiante curso
func GetEstudianteCursobyEstudianteID(id int) ([]model.EstudianteCurso, error) {

	db := database.DB

	var e_c []model.EstudianteCurso

	result := db.Where(&e_c, model.EstudianteCurso{EstudianteId: id}).Find(&e_c)

	if result.Error != nil {
		return e_c, result.Error
	}

	return e_c, nil
}

// GetEstudianteCursobyEstudianteID query estudiante curso
func GetEstudianteCursoByCursoID(id int) ([]model.EstudianteCurso, error) {

	db := database.DB

	var e_c []model.EstudianteCurso

	result := db.Where(&e_c, model.EstudianteCurso{CursoId: id}).Find(&e_c)

	if result.Error != nil {
		return e_c, result.Error
	}

	return e_c, nil
}
