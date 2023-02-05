package router

import (
	"api-fiber-gorm/handler"
	"api-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	// auth.Get("/logout", middleware.Protected(), handler.Logout)

	// User
	user := api.Group("/user")
	user.Post("/", handler.CreateUser)
	user.Get("/:id", middleware.Protected(), handler.GetUser)

	// Address
	address := api.Group("/address")
	address.Post("/", middleware.Protected(), handler.CreateUserAddress)
	// http://localhost:8000/api/address/user_id=1
	address.Get("/user_id=:id", middleware.Protected(), handler.Test2) // funciona - get para relacion
	// http://localhost:8000/api/address/user_id:1
	address.Get("/user_id::id", middleware.Protected(), handler.GetPreferredUserAddress) // funciona - get para relacion
	// http://localhost:8000/api/address/1
	address.Get("/:id", middleware.Protected(), handler.Test) // funciona - solo get id address
	

	// Country
	country := api.Group("/countrys")
	country.Get("/", middleware.Protected(), handler.GetAllCountrys)

	// Department
	department := api.Group("/departments")
	department.Get("/", middleware.Protected(), handler.GetAllDepartments)
	department.Get("/?country_id=:id", middleware.Protected(), handler.GetCountryDepartments)

	// City
	city := api.Group("/citys")
	city.Get("/", middleware.Protected(), handler.GetAllCitys)
	city.Get("/?deparment_id=:id", middleware.Protected(), handler.GetDepartmentCitys)

	// Owns
	owns := api.Group("/owns")
	owns.Get("/?user_id=:id", middleware.Protected(), handler.GetDepartmentCitys)                                 // GetAllPetsOfUser
	owns.Post("/transfer?pet_id=:pet_id&new_owner=:username", middleware.Protected(), handler.GetDepartmentCitys) // TransferPet

	// Pets
	pets := api.Group("/pets")
	pets.Get("/?pet_id=:id", middleware.Protected(), handler.GetDepartmentCitys) // GetAllRecordsOfPet
	pets.Post("/", middleware.Protected(), handler.GetDepartmentCitys)           // CreatePet

	// PetRecords
	records := api.Group("/petrecords")
	records.Get("/?pet_id=:id", middleware.Protected(), handler.GetDepartmentCitys) // GetAllRecordsOfPet
	records.Post("/", middleware.Protected(), handler.GetDepartmentCitys)           // CreatePetRecord

	// user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
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
