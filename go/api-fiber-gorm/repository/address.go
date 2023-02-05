package repository

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"

	"gorm.io/gorm"
)

// GetUserAddress get preferred useraddress
func GetPreferredUserAddress(id int) (model.UserAddress, error) {
	db := database.DB
	var u_address model.UserAddress

	result := db.Where(&u_address, model.UserAddress{UserId: id, Preferred: true}).Find(&u_address)

	if result.Error != nil || result.RowsAffected == 0 {
		return u_address, gorm.ErrEmptySlice
	}

	var err error
	u_address.City, err = GetCity(u_address.CityId)
	if err != nil {
		return u_address, err
	}

	u_address.City.Department, err = GetDepartment(u_address.City.DepartmentId)
	if err != nil {
		return u_address, err
	}

	u_address.City.Department.Country, err = GetCountry(u_address.City.Department.CountryId)
	if err != nil {
		return u_address, err
	}

	return u_address, nil
}

// CreateUserAddress new address
func CreateUserAddress(u_address model.UserAddress) (model.UserAddress, error) {


	db := database.DB
	if err := db.Create(&u_address).Error; err != nil {
		return u_address, err
	}

	var err error
	u_address.City, err = GetCity(u_address.CityId)
	if err != nil {
		return u_address, err
	}

	u_address.City.Department, err = GetDepartment(u_address.City.DepartmentId)
	if err != nil {
		return u_address, err
	}

	u_address.City.Department.Country, err = GetCountry(u_address.City.Department.CountryId)
	if err != nil {
		return u_address, err
	}

	return u_address, nil
}

// UpdateUser update user
// func UpdateUser(c *fiber.Ctx) error {
// 	type UpdateUserInput struct {
// 		Names string `json:"names"`
// 	}
// 	var uui UpdateUserInput
// 	if err := c.BodyParser(&uui); err != nil {
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
// 	}
// 	id := c.Params("id")
// 	token := c.Locals("user").(*jwt.Token)

// 	if !validToken(token, id) {
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
// 	}

// 	db := database.DB
// 	var user model.User

// 	db.First(&user, id)
// 	user.Names = uui.Names
// 	db.Save(&user)

// 	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
// }

// // DeleteUser delete user
// func DeleteUser(c *fiber.Ctx) error {
// 	type PasswordInput struct {
// 		Password string `json:"password"`
// 	}
// 	var pi PasswordInput
// 	if err := c.BodyParser(&pi); err != nil {
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
// 	}
// 	id := c.Params("id")
// 	token := c.Locals("user").(*jwt.Token)

// 	if !validToken(token, id) {
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})

// 	}

// 	if !validUser(id, pi.Password) {
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})

// 	}

// 	db := database.DB
// 	var user model.User

// 	db.First(&user, id)

// 	db.Delete(&user)
// 	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
// }
