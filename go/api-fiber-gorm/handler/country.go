package handler

// import (
// 	"api-fiber-gorm/repository"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// )

// // GetAllCountrys
// func GetAllCountrys(c *fiber.Ctx) error {
// 	countrys, err := repository.GetAllCountrys()

// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
// 	}

// 	return c.JSON(fiber.Map{"status": "success", "message": "All countrys", "data": countrys})
// }

// func GetCountryDepartments(c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Query("country_id"))

// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
// 	}

// 	departments, err := repository.GetCountryDepartments(id)

// 	// var departments []model.Department
// 	// result := db.Find(&departments, id)

// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found resource", "data": nil})
// 	}
// 	// // if product.Title == "" {
// 	// // 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
// 	// // }

// 	return c.JSON(fiber.Map{"status": "success", "message": "Country found", "data": departments})
// }
