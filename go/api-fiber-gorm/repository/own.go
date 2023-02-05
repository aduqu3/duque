package repository

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
	"fmt"
)

// CreateOwnPet add pet to user
func CreateOwnPet(mdl model.Own) (model.Own, error) {
	db := database.DB
	if err := db.Create(&mdl).Error; err != nil {
		return mdl, err
	}
	return mdl, nil
}

// GetAllPetsOwns get preferred useraddress
func GetAllPetsOwns(user_id int) ([]model.Pet, error) {
	db := database.DB
	var mdl_own []model.Own

	result := db.Where(&mdl_own, model.Own{UserId: user_id}).Find(&mdl_own)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, nil
	}

	// var err error
	var pets []model.Pet
	for _, item := range mdl_own {
		// mdl_own[i].Pet, _ = GetPet(item.PetId)
		pet, _ := GetPet(item.PetId)
		pets = append(pets, pet)
	}

	// fmt.Println(mdl_own[:])

	// if err != nil {
	// 	return mdl_own, err
	// }

	return pets, nil
}

// GetUserAddress get preferred useraddress
func ValidateOwnerPet(user_id int, pet_id int) bool {
	db := database.DB
	var mdl_own model.Own

	result := db.Where(&mdl_own, model.Own{UserId: user_id, PetId: pet_id}).Find(&mdl_own)

	if result.Error != nil || result.RowsAffected == 0 {
		return true
	}

	fmt.Println(result)

	// fmt.Println(result.Statement.Vars...)
	// fmt.Println(result.Statement.TableExpr.Vars...)

	// var err error
	// for i, item := range mdl_own {
	// 	mdl_own[i].Pet, _ = GetPet(item.PetId)
	// }

	// if err != nil {
	// 	return mdl_own, err
	// }

	return false
}

func TransferPet(user_id int, pet_id int, username string) (model.Own, error) {
	fmt.Println("llega al repository")

	db := database.DB
	var mdl_own model.Own
	result := db.Where(&mdl_own, model.Own{UserId: user_id, PetId: pet_id}).Find(&mdl_own)

	if result.Error != nil || result.RowsAffected == 0 {
		return model.Own{}, nil
	}

	fmt.Println(mdl_own.UserId)
	fmt.Println(mdl_own.PetId)

	// change user id before create new record own
	var mdl_user model.User

	if err := db.First(&mdl_user, model.User{Username: username}).Error; err != nil {
		return model.Own{}, err
	}

	// delete old record own
	if err := db.Delete(&mdl_own).Error; err != nil {
		return model.Own{}, err
	}

	// ready for create new record
	new_mdl_own := model.Own{
		UserId: int(mdl_user.ID),
		PetId:  mdl_own.PetId,
	}

	// mdl_own.UserId = int(mdl_user.ID)

	// create new record own with new owner of pet
	if err := db.Create(&new_mdl_own).Error; err != nil {
		return model.Own{}, err
	}

	return mdl_own, nil
}
