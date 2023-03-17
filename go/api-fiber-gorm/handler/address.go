package handler

import (
	"api-fiber-gorm/model"
	"api-fiber-gorm/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreateUserAddress new address for user
func CreateUserAddress(c *fiber.Ctx) error {
	// user_id, err := c.ParamsInt("user_id")

	u_a := new(model.UserAddress)
	if err := c.BodyParser(u_a); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if u_id := GetUserIdOfToken(c); u_id != u_a.UserId {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	u_address, err := repository.CreateUserAddress(*u_a)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user address", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user address", "data": u_address})
}

// GetUserAddress get preferred useraddress
func GetActiveUserAddress(c *fiber.Ctx) error {
	user_id, err := strconv.Atoi(c.Query("user_id"))
	// strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	if u_id := GetUserIdOfToken(c); u_id != user_id {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	u_addres, err := repository.GetActiveUserAddress(user_id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No preferred user address found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Preferred User Address found", "data": u_addres})
	// return c.JSON(u_addres) // idea solo retorna data
}
