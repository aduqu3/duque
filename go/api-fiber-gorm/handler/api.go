package handler

import "github.com/gofiber/fiber/v2"

// Hello hanlde api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello world!", "data": nil})
}

func HelloAdmin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello world Admin!", "data": nil})
}

func HelloAcademico(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello world Academico!", "data": nil})
}

func HelloEmpresario(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello world Empresario!", "data": nil})
}
