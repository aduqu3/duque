package handler

import (
	"api-fiber-gorm/model"
	"api-fiber-gorm/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// // GetUserAddress get preferred useraddress
// func GetPreferredUserAddress(c *fiber.Ctx) error {
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
func CreateUserAddress(c *fiber.Ctx) error {

	fmt.Println("asdw HHHAHAAAAAAAAAAAAHHHHHHHHHH")
	// db := database.DB
	u_a := new(model.UserAddress)

	// fmt.Println(c.BodyParser(u_address))

	if err := c.BodyParser(u_a); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	u_address, err := repository.CreateUserAddress(u_a)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user address", "data": u_address})
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