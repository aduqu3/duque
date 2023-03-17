package handler

import (
	"api-fiber-gorm/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetOwns GetAllPetsOwn
func GetAllPetsOwns(c *fiber.Ctx) error {
	user_id, err := strconv.Atoi(c.Query("user_id"))

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found id", "data": nil})
	}

	if u_id := GetUserIdOfToken(c); u_id != user_id {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	mdl, err := repository.GetAllPetsOwns(user_id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No preferred user address found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Preferred User Address found", "data": mdl})
	// return c.JSON(u_addres) // idea solo retorna data
}

// GetAllPetRecords
func TransferPet(c *fiber.Ctx) error {

	user_id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Error in convert data", "data": nil})
	}

	pet_id, err := strconv.Atoi(c.Query("pet_id"))

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Error in convert data", "data": nil})
	}

	if u_id := GetUserIdOfToken(c); u_id != user_id {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	if repository.ValidateOwnerPet(user_id, pet_id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	type NewOwner struct {
		Username string `json:"username"`
	}

	var uui NewOwner
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	mdl, err := repository.TransferPet(user_id, pet_id, uui.Username)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Unable to transfer Pet", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Transfer successfully", "data": mdl})
	// return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
	// return c.JSON(u_addres) // idea solo retorna data
}
