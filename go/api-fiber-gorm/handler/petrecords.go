package handler

import (
	"api-fiber-gorm/model"
	"api-fiber-gorm/repository"

	"github.com/gofiber/fiber/v2"
)

// CreateUserAddress new address
func CreatePetRecord(c *fiber.Ctx) error {

	mdl := new(model.PetRecord)
	if err := c.BodyParser(mdl); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	user_id := GetUserIdOfToken(c)

	// validate if pet is own by user authenticate
	if repository.ValidateOwnerPet(user_id, mdl.PetId) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	res_mdl, err := repository.CreatePetRecord(*mdl)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user address", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user address", "data": res_mdl})
}

// GetAllPetRecords
func GetAllPetRecords(c *fiber.Ctx) error {
	pet_id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found id", "data": nil})
	}

	u_id := GetUserIdOfToken(c)

	if repository.ValidateOwnerPet(u_id, pet_id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	mdl, err := repository.GetAllPetsOwns(pet_id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No preferred user address found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Preferred User Address found", "data": mdl})
	// return c.JSON(u_addres) // idea solo retorna data
}
