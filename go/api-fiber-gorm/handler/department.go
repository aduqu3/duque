package handler

import (
	"api-fiber-gorm/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetAllDepartments
func GetAllDepartments(c *fiber.Ctx) error {

	departments, err := repository.GetAllDepartments()

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Nof found resource", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "All departments", "data": departments})
}

func GetDepartmentCitys(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}
	citys, err := repository.GetDepartmentCitys(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Country found", "data": citys})
}
