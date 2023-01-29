package repository

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
	"fmt"
)

// GetUserAddress get preferred useraddress
// func GetPreferredUserAddress(id) error {
// 	// user_id :=  // user id
// 	user_id, err := strconv.Atoi(c.Params("id"))

// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found id", "data": nil})
// 	}

// 	db := database.DB
// 	var user_address model.UserAddress
// 	result := db.Where(&user_address, model.UserAddress{UserId: user_id, Preferred: true}).Find(&user_address)

// 	fmt.Println(result)
// 	fmt.Println(user_address.User.FirstName)

// 	if result.Error != nil {
// 		fmt.Println("entra al error")
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No preferred user address found with ID", "data": nil})
// 	}
// 	// user.Password = "****"
// 	return c.JSON(fiber.Map{"status": "success", "message": "Preferred User Address found", "data": user_address})
// }

// CreateUserAddress new address
func CreateUserAddress(u_address *model.UserAddress) (map[string]interface{}, error) {

	type NewUserAddress struct {
		UserId       int
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Name         string `json:"name"`
		Country      string `json:"county"`
		Department   string `json:"department"`
		City         string `json:"city"`
		ZipCode      string `json:"zip_code"`
		Details      string `json:"details"`
		OtherDetails string `json:"other_details"`
		Phone        string `json:"phone"`
		Preferred    bool   `json:"preferred"`
	}

	// fmt.Println(u_address)

	db := database.DB
	if err := db.Create(&u_address).Error; err != nil {
		return nil, nil
	}

	var city model.City
	city.ID = uint(u_address.CityId)
	db.Find(&city, city.ID)
	u_address.City = city

	fmt.Println(u_address.City)
	fmt.Println("CPAS")
	// get city
	// get department
	// get country
	// get user

	

	// hash, err := hashPassword(user.Password)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	// }

	// user.Password = hash
	// if err := db.Create(&user).Error; err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	// }

	newUserAddress := map[string]interface{}{
		"UserId":       u_address.UserId,
		"Name":         "",
		"Country":      "1",
		"Department":   "2",
		"City":         "2",
		"ZipCode":      "",
		"Details":      "",
		"OtherDetails": "",
		"Phone":        "",
		"Preferred":    true,
	}

	return newUserAddress, nil
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
