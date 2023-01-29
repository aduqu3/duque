package repository

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
	"errors"
)

// GetAllCountrys query all countrys
func GetAllCountrys() ([]model.Country, error) {
	db := database.DB
	var countrys []model.Country
	result := db.Find(&countrys)

	if result.Error != nil {
		return countrys, result.Error
	}
	return countrys, nil
}

// GetCountry get country by id
func GetCountry(id int) (model.Country, error) {
	db := database.DB
	var country model.Country

	err := db.Find(&country, id).Error
	if err != nil {
		return country, err
	}
	
	return country, nil
}

// GetCountryDepartments query country
func GetCountryDepartments(id int) ([]model.Department, error) {
	departments, err := GetAllDepartments()

	if err != nil {
		return departments, err
	}

	var new_departments []model.Department
	for _, ele := range departments {
		if ele.CountryId == id {
			new_departments = append(new_departments, ele)
		}
	}

	if len(new_departments) == 0 {
		return new_departments, errors.New("Not found resource for ID")
	}

	return new_departments, nil
}
