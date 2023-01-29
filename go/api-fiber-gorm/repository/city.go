package repository

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
)

// GetAllCitys query all citys
func GetAllCitys() ([]model.City, error) {
	db := database.DB
	var citys []model.City
	result := db.Find(&citys)

	if result.Error != nil {
		return citys, result.Error
	}
	return citys, nil
}
