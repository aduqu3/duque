package repository

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
)

// CreatePet new pet
func CreatePet(mdl model.Pet) (model.Pet, error) {
	db := database.DB
	if err := db.Create(&mdl).Error; err != nil {
		return mdl, err
	}
	return mdl, nil
}

// GetPet get pet by id
func GetPet(pet_id int) (model.Pet, error) {
	db := database.DB
	var mdl model.Pet
	err := db.Find(&mdl, pet_id).Error
	if err != nil {
		return mdl, err
	}
	return mdl, nil
}
