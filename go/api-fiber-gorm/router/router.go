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
	address.Get("/user_id=:id", middleware.Protected(), handler.GetPreferredUserAddress) // funciona - get para relacion
	// http://localhost:8000/api/address/1
	// address.Get("/:id", middleware.Protected(), handler.) // funciona - solo get id address

	// Country
	country := api.Group("/countrys")
	country.Get("/", middleware.Protected(), handler.GetAllCountrys)

	// Department
	department := api.Group("/departments")
	// department.Get("/", middleware.Protected(), handler.GetAllDepartments)
	department.Get("/country_id=:id", middleware.Protected(), handler.GetCountryDepartments)

	// City
	city := api.Group("/citys")
	// city.Get("/", middleware.Protected(), handler.GetAllCitys)
	city.Get("/deparment_id=:id", middleware.Protected(), handler.GetDepartmentCitys)

	// Owns
	owns := api.Group("/owns")
	owns.Get("/user_id=:id", middleware.Protected(), handler.GetAllPetsOwns)             // GetAllPetsOwns
	owns.Post("/transfer/user_id=:user_id-pet_id=:pet_id", middleware.Protected(), handler.TransferPet) // TransferPet to other user

	// Pets
	pets := api.Group("/pets")
	pets.Post("/", middleware.Protected(), handler.CreatePetOwn) // CreatePet And Register Owner

	// PetRecords
	records := api.Group("/petrecords")
	records.Get("/pet_id=:id", middleware.Protected(), handler.GetAllPetRecords) // GetAllPetRecords
	records.Post("/", middleware.Protected(), handler.CreatePetRecord)           // CreatePetRecord

	// get pet data and user data and address
}
