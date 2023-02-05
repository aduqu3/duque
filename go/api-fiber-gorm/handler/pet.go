package handler

import (
	"api-fiber-gorm/model"
	"api-fiber-gorm/repository"

	"github.com/gofiber/fiber/v2"
)

// CreatePet and Own
func CreatePetOwn(c *fiber.Ctx) error {
	user_id, err := c.ParamsInt("user_id")

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	pet := new(model.Pet)
	if err := c.BodyParser(pet); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if u_id := GetUserIdOfToken(c); u_id != user_id {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	pet_data, err := repository.CreatePet(*pet)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create pet", "data": err})
	}

	mdl_own := model.Own{
		PetId:  int(pet_data.ID),
		UserId: user_id,
	}

	own_pet, err := repository.CreateOwnPet(mdl_own)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user own", "data": err})
	}

	own_pet.Pet = pet_data

	return c.JSON(fiber.Map{"status": "success", "message": "Created pet and user own pet", "data": own_pet})
}
