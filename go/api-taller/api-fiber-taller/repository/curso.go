package repository

import (
	"api-fiber-taller/database"
	"api-fiber-taller/model"
)

// GetAllCursos query all estudiantes
func GetAllCursos() ([]model.Curso, error) {

	db := database.DB
	var curso []model.Curso
	result := db.Find(&curso)

	if result.Error != nil {
		return curso, result.Error
	}

	return curso, nil
}

// GetCurso query mejor curso by id
func GetCurso(id int) (model.Curso, error) {

	db := database.DB
	var curso model.Curso
	result := db.Find(&curso, id)

	if result.Error != nil {
		return curso, result.Error
	}

	return curso, nil
}

// CreateUserAddress new address
func CreateCourse(u_address model.Curso) (model.Curso, error) {
	db := database.DB
	if err := db.Create(&u_address).Error; err != nil {
		return u_address, err
	}
	return u_address, nil
}
