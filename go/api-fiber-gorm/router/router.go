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
	user := api.Group("/users")
	user.Post("/", handler.CreateUser)                                                           // create user
	user.Get("/:user_id", middleware.Protected(), handler.GetUser)                               // get user by id
	user.Post("/:user_id/address", middleware.Protected(), handler.CreateUserAddress)            // create user address
	user.Get("/:user_id/address", middleware.Protected(), handler.GetPreferredUserAddress)       // get user address
	user.Post("/:user_id/pets", middleware.Protected(), handler.CreatePetOwn)                    // CreatePet And Register Owner
	user.Get("/:user_id/pets", middleware.Protected(), handler.GetAllPetsOwns)                   // GetAllPetsOwns by user
	user.Post("/:user_id/transfer/pets/:pet_id", middleware.Protected(), handler.TransferPet)    // TransferPet to other user
	user.Get("/:user_id/pets/:pet_id/records", middleware.Protected(), handler.GetAllPetRecords) // GetAllPetRecords
	user.Post("/:user_id/pets/:pet_id/records", middleware.Protected(), handler.CreatePetRecord) // CreatePetRecord

	// Country
	country := api.Group("/countries")
	country.Get("/", middleware.Protected(), handler.GetAllCountrys)
	country.Get("/:country_id/departments", middleware.Protected(), handler.GetCountryDepartments) // Get all departments of country

	// Department
	department := api.Group("/departments")
	// department.Get("/", middleware.Protected(), handler.GetAllDepartments)
	department.Get("/:department_id/citys", middleware.Protected(), handler.GetDepartmentCitys)

	// Media
	media := api.Group("/images")
	media.Post("/", middleware.Protected(), handler.Fileupload)

	// City
	// city := api.Group("/citys")
	// city.Get("/", middleware.Protected(), handler.GetAllCitys)

	// Owns
	// owns := api.Group("/owns")
	// owns.Get("/user_id=:id", middleware.Protected(), handler.GetAllPetsOwns)                            // GetAllPetsOwns
	// owns.Post("/transfer/user_id=:user_id-pet_id=:pet_id", middleware.Protected(), handler.TransferPet) // TransferPet to other user

	// Pets
	// pets := api.Group("/pets")
	// pets.Get("/:pet_id/records", middleware.Protected(), handler.GetAllPetRecords) // GetAllPetRecords
	// pets.Post("/:pet_id/records", middleware.Protected(), handler.CreatePetRecord) // CreatePetRecord

	// PetRecords
	// records := api.Group("/petrecords")
	// records.Get("/pet_id=:id", middleware.Protected(), handler.GetAllPetRecords) // GetAllPetRecords
	// records.Post("/", middleware.Protected(), handler.CreatePetRecord)           // CreatePetRecord

	// get pet data and user data and address
}
