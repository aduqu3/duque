package handler

import (
	"api-fiber-taller/model"
	"api-fiber-taller/repository"

	"github.com/gofiber/fiber/v2"
)

// CreateUserAddress new address
func CreateCourse(c *fiber.Ctx) error {

	u_a := new(model.Curso)
	if err := c.BodyParser(u_a); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	u_address, err := repository.CreateCourse(*u_a)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create course", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created course", "data": u_address})
}
