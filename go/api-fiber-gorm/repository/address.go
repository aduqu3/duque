package repository

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"

	"gorm.io/gorm"
)

// CreateUserAddress new address
func CreateUserAddress(u_address model.UserAddress) (model.UserAddress, error) {

	db := database.DB
	if err := db.Create(&u_address).Error; err != nil {
		return u_address, err
	}

	return u_address, nil
}

// GetUserAddress get preferred useraddress
func GetActiveUserAddress(id int) (model.UserAddress, error) {
	db := database.DB
	var u_address model.UserAddress

	result := db.Where(model.UserAddress{UserId: id, Preferred: true}).First(&u_address)

	if result.Error != nil || result.RowsAffected == 0 {
		return u_address, gorm.ErrEmptySlice
	}

	return u_address, nil
}
