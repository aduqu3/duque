package repository

// import (
// 	"api-fiber-gorm/database"
// 	"api-fiber-gorm/model"
// )

// // GetAllCitys query all citys
// func GetAllCitys() ([]model.City, error) {
// 	db := database.DB
// 	var citys []model.City
// 	result := db.Find(&citys)

// 	if result.Error != nil {
// 		return citys, result.Error
// 	}
// 	return citys, nil
// }

// // GetCity get city by id
// func GetCity(id int) (model.City, error) {
// 	db := database.DB
// 	var city model.City

// 	err := db.Find(&city, id).Error
// 	if err != nil {
// 		return city, err
// 	}

// 	return city, nil
// }
