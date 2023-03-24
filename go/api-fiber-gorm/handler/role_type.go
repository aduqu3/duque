package handler

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllRoles
func GetAllRoles(c *fiber.Ctx) error {
	db := database.DB
	var roles []model.UserType
	result := db.Find(&roles)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "All countrys", "data": roles})
}
