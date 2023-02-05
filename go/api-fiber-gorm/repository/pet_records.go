package repository

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
)

// CreateUserAddress new address
func CreatePetRecord(mdl model.PetRecord) (model.PetRecord, error) {

	db := database.DB
	if err := db.Create(&mdl).Error; err != nil {
		return mdl, err
	}

	return mdl, nil
}

// GetAllPetRecords get preferred useraddress
func GetAllPetRecords(pet_id int) ([]model.PetRecord, error) {
	db := database.DB
	var mdl []model.PetRecord

	result := db.Where(&mdl, model.PetRecord{PetId: pet_id}).Find(&mdl)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, nil
	}

	return mdl, nil
}
