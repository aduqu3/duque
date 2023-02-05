package handler

import (
	"api-fiber-gorm/model"
	"api-fiber-gorm/repository"

	"github.com/gofiber/fiber/v2"
)

// CreatePet and Own
func CreatePetOwn(c *fiber.Ctx) error {

	pet := new(model.Pet)
	if err := c.BodyParser(pet); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	user_id := GetUserIdOfToken(c)

	pet_data, err := repository.CreatePet(*pet)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user address", "data": err})
	}

	mdl_own := model.Own{
		PetId:  int(pet_data.ID),
		UserId: user_id,
	}

	own_pet, err := repository.CreateOwnPet(mdl_own)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user address", "data": err})
	}

	own_pet.Pet = pet_data

	return c.JSON(fiber.Map{"status": "success", "message": "Created user address", "data": own_pet})
}
