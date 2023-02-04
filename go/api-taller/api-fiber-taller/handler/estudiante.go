package handler

import (
	"api-fiber-taller/database"
	"api-fiber-taller/model"
	"api-fiber-taller/repository"
	"fmt"
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

// CreateUserAddress new address
func CreateEstudiante(c *fiber.Ctx) error {

	u_a := new(model.Estudiante)
	if err := c.BodyParser(u_a); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	u_address, err := repository.CreateEstudiante(*u_a)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create estudiante", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created estudiante", "data": u_address})
}

// CreateUserAddress new address
func CreateEstudianteCurso(c *fiber.Ctx) error {

	u_a := new(model.EstudianteCurso)
	if err := c.BodyParser(u_a); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	u_address, err := repository.CreateEstudianteCurso(*u_a)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create estudiante", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created estudiante", "data": u_address})
}

func DeleteEstudiante(c *fiber.Ctx) error {
	db := database.DB
	var user model.Estudiante

	id := c.Params("id")
	fmt.Println(id)
	db.First(&user, id)
	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "Estudiante successfully deleted", "data": nil})
}

// UpdateUser update user
func UpdateEstudiante(c *fiber.Ctx) error {
	var uui model.Estudiante
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// fmt.Println(uui.ID)

	// id := c.Params("id")
	// fmt.Println(id)

	db := database.DB
	var user model.Estudiante

	db.First(&user, uui.ID)
	user.Nombre = uui.Nombre
	user.Apellido = uui.Apellido
	user.About = uui.About
	user.Gender = uui.Gender
	user.Address = uui.Address
	user.Email = uui.Email
	user.Edad = uui.Edad
	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "Estudiante successfully updated", "data": user})
}
