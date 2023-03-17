package repository

// import (
// 	"api-fiber-gorm/database"
// 	"api-fiber-gorm/model"
// 	"errors"
// )

// // GetAllDepartments query all departments
// func GetAllDepartments() ([]model.Department, error) {

// 	db := database.DB
// 	var departments []model.Department
// 	result := db.Find(&departments)

// 	if result.Error != nil {
// 		return departments, result.Error
// 	}

// 	return departments, nil
// }

// // GetDepartment get department by id
// func GetDepartment(id int) (model.Department, error) {
// 	db := database.DB
// 	var department model.Department

// 	err := db.Find(&department, id).Error
// 	if err != nil {
// 		return department, err
// 	}

// 	return department, nil
// }

// // GetCountryDepartments query country
// func GetDepartmentCitys(id int) ([]model.City, error) {
// 	citys, err := GetAllCitys()

// 	if err != nil {
// 		return citys, err
// 	}

// 	var new_citys []model.City
// 	for _, ele := range citys {
// 		if ele.DepartmentId == id {
// 			new_citys = append(new_citys, ele)
// 		}
// 	}

// 	if len(new_citys) == 0 {
// 		return new_citys, errors.New("Not found resource for ID")
// 	}

// 	return new_citys, nil
// }
