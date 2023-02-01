package handler

import (
	"api-fiber-taller/repository"
	"strconv"

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
	gender := c.Params("gender")
	old, err := repository.GetMoreOldEstudiante(gender)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Found old persons", "data": old})
}

func GetCursoAvg(c *fiber.Ctx) error {
	curso_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	avg_c, err := repository.GetCursoAvg(curso_id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Found old persons", "data": avg_c})
}

func GetLowEstudianteCurso(c *fiber.Ctx) error {
	curso_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	estudiantes, err := repository.GetLowEstudianteCurso(curso_id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Found old persons", "data": estudiantes})
}

// GetAllCourse
func GetAllCourse(c *fiber.Ctx) error {

	course, err := repository.GetAllCursos()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "All courses", "data": course})
}
