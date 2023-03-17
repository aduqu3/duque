package handler

// import (
// 	"api-fiber-gorm/repository"

// 	"github.com/gofiber/fiber/v2"
// )

// // GetAllCitys
// func GetAllCitys(c *fiber.Ctx) error {
// 	citys, err := repository.GetAllCitys()

// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Nof found resource", "data": nil})
// 	}

// 	return c.JSON(fiber.Map{"status": "success", "message": "All citys", "data": citys})
// }
