package router

import (
	"api-fiber-taller/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	student := api.Group("/student")
	student.Post("/", handler.CreateEstudiante)
	student.Delete("/:id", handler.DeleteEstudiante)
	student.Put("/", handler.UpdateEstudiante)
	student.Post("/course/", handler.CreateEstudianteCurso)
	student.Get("/mejor/", handler.GetMejorEstudiante)
	student.Get("/old/:gender", handler.GetMoreOldEstudiante)
	student.Get("/low/:id", handler.GetLowEstudianteCurso)

	course := api.Group("/course")
	course.Get("/", handler.GetAllCourse)
	course.Post("/", handler.CreateCourse)
	course.Get("/avg/:id", handler.GetCursoAvg)

	// auth := api.Group("/auth")
	// auth.Post("/login", handler.Login)
	// auth.Get("/logout", middleware.Protected(), handler.Logout)

	// // User
	// user := api.Group("/user")
	// user.Post("/", handler.CreateUser)
	// user.Get("/:id", middleware.Protected(), handler.GetUser)
	// // user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// // user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
	// user.Get("/address/:id", middleware.Protected(), handler.GetPreferredUserAddress)
	// user.Post("/address/", middleware.Protected(), handler.CreateUserAddress)

	// // Country
	// country := api.Group("/country")
	// country.Get("/", middleware.Protected(), handler.GetAllCountrys)
	// country.Get("/deparments/:id", middleware.Protected(), handler.GetCountryDepartments)

	// // Department
	// department := api.Group("/department")
	// department.Get("/", middleware.Protected(), handler.GetAllDepartments)
	// department.Get("/citys/:id", middleware.Protected(), handler.GetDepartmentCitys)

	// // City
	// city := api.Group("/city")
	// city.Get("/", middleware.Protected(), handler.GetAllCitys)
	// city.Get("/:id", handler.GetProduct)

	// Product
	// product := api.Group("/product")
	// product.Get("/", handler.GetAllProducts)
	// product.Get("/:id", handler.GetProduct)
	// product.Post("/", middleware.Protected(), handler.CreateProduct)
	// product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
