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

	// User
	// user := api.Group("/users")
	// user.Post("/", handler.CreateUser)                                                           // create user
	// user.Get("/:user_id", middleware.Protected(), handler.GetUser)                               // get user by id
	// user.Post("/:user_id/address", middleware.Protected(), handler.CreateUserAddress)            // create user address
	// user.Get("/:user_id/address", middleware.Protected(), handler.GetPreferredUserAddress)       // get user address
	// user.Post("/:user_id/pets", middleware.Protected(), handler.CreatePetOwn)                    // CreatePet And Register Owner
	// user.Get("/:user_id/pets", middleware.Protected(), handler.GetAllPetsOwns)                   // GetAllPetsOwns by user
	// user.Post("/:user_id/transfer/pets/:pet_id", middleware.Protected(), handler.TransferPet)    // TransferPet to other user
	// user.Get("/:user_id/pets/:pet_id/records", middleware.Protected(), handler.GetAllPetRecords) // GetAllPetRecords
	// user.Post("/:user_id/pets/:pet_id/records", middleware.Protected(), handler.CreatePetRecord) // CreatePetRecord

	// // Country
	// country := api.Group("/countries")
	// country.Get("/", middleware.Protected(), handler.GetAllCountrys)
	// country.Get("/:country_id/departments", middleware.Protected(), handler.GetCountryDepartments) // Get all departments of country

	// // Department
	// department := api.Group("/departments")
	// // department.Get("/", middleware.Protected(), handler.GetAllDepartments)
	// department.Get("/:department_id/citys", middleware.Protected(), handler.GetDepartmentCitys)

	// User
	user := api.Group("/users")
	user.Post("/", handler.CreateUser)
	user.Get("/:id", middleware.Protected(), handler.GetUser)

	// Address
	address := api.Group("/address")
	address.Post("/", middleware.Protected(), handler.CreateUserAddress)
	address.Get("/active", middleware.Protected(), handler.GetActiveUserAddress) // funciona - get para relacion
	// address.Get("/activate", middleware.Protected(), handler.GetPreferredUserAddress) // funciona - get para relacion
	// http://localhost:8000/api/address/search?user_id=1

	// Country
	// country := api.Group("/countrys")
	// country.Get("/", middleware.Protected(), handler.GetAllCountrys)

	// // Department
	// department := api.Group("/departments")
	// department.Get("/", middleware.Protected(), handler.GetAllDepartments)
	// department.Get("/search", middleware.Protected(), handler.GetCountryDepartments)
	// http://localhost:8000/api/departments/search?country_id=1

	// City
	// city := api.Group("/citys")
	// city.Get("/", middleware.Protected(), handler.GetAllCitys)
	// city.Get("/search", middleware.Protected(), handler.GetDepartmentCitys)
	// http://localhost:8000/api/citys/search?department_id=1

	// Owns
	owns := api.Group("/owns")
	owns.Get("/search", middleware.Protected(), handler.GetAllPetsOwns) // GetAllPetsOfUser
	// http://localhost:8000/api/owns/search?user_id=1
	owns.Post("/transfer", middleware.Protected(), handler.TransferPet) // TransferPet to other user
	// http://localhost:8000/api/owns/transfer?user_id=1&pet_id=1

	// Pets
	pets := api.Group("/pets")
	pets.Post("/", middleware.Protected(), handler.CreatePetOwn)        // CreatePet and Own witout image
	pets.Put("/image/:id", middleware.Protected(), handler.PutPetImage) // CreatePet and Own witout image
	pets.Get("/:id", middleware.Protected(), handler.GetPet)            // CreatePet and Own witout image

	// PetRecords
	records := api.Group("/records")
	records.Get("/search", middleware.Protected(), handler.GetAllPetRecords) // GetAllPetRecords
	// http://localhost:8000/api/records/search?pet_id=1
	records.Post("/", middleware.Protected(), handler.CreatePetRecord) // CreatePetRecord

	// general public --> get pet data and user data and address

}
