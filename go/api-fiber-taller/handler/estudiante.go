package handler

import (
	"api-fiber-taller/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllEstudiantes
func GetAllEstudiantes(c *fiber.Ctx) error {

	estudiantes, err := repository.GetAllEstudiantes()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "All departments", "data": estudiantes})
}

func GetMejorEstudiante(c *fiber.Ctx) error {

	best, err := repository.GetMejorEstudiante()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Found best student", "data": best})
}

func GetMoreOldEstudiante(c *fiber.Ctx) error {

	old, err := repository.GetMoreOldEstudiante()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Found old persons", "data": old})
}

func GetCursoAvg(c *fiber.Ctx) error {

	avg_c, err := repository.GetCursoAvg()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Found old persons", "data": avg_c})
}
