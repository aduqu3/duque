package handler

import (
	"api-fiber-gorm/model"
	"api-fiber-gorm/repository"

	"github.com/gofiber/fiber/v2"
)

// CreateUserAddress new address
func CreateUserAddress(c *fiber.Ctx) error {
	user_id, err := c.ParamsInt("user_id")

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found id", "data": nil})
	}

	u_a := new(model.UserAddress)
	if err := c.BodyParser(u_a); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if u_id := GetUserIdOfToken(c); u_id != user_id {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}
	// asing user id
	u_a.UserId = user_id

	u_address, err := repository.CreateUserAddress(*u_a)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user address", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user address", "data": u_address})
}

// GetUserAddress get preferred useraddress
func GetPreferredUserAddress(c *fiber.Ctx) error {
	user_id, err := c.ParamsInt("user_id")
	// strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found id", "data": nil})
	}

	if u_id := GetUserIdOfToken(c); u_id != user_id {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	u_addres, err := repository.GetPreferredUserAddress(user_id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No preferred user address found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Preferred User Address found", "data": u_addres})
	// return c.JSON(u_addres) // idea solo retorna data
}

// GetUserAddress get preferred useraddress
// func Test(c *fiber.Ctx) error {
// 	fmt.Println("test 1")
// 	id, err := c.ParamsInt("id")

// 	fmt.Println(id)

// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found id", "data": nil})
// 	}

// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No preferred user address found with ID", "data": nil})
// 	// return c.JSON(u_addres) // idea solo retorna data
// }

// GetUserAddress get preferred useraddress
// func Test2(c *fiber.Ctx) error {
// 	fmt.Println("test 2")
// 	id, err := strconv.Atoi(c.Params("id"))

// 	fmt.Println(id)
// 	fmt.Println(c.Params("user_id"))
// 	fmt.Println(c.Params("pet_id"))

// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found id", "data": nil})
// 	}

// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No preferred user address found with ID", "data": nil})
// 	// return c.JSON(u_addres) // idea solo retorna data
// }

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
