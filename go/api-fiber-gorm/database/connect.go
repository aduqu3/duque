package database

import (
	"api-fiber-gorm/config"
	"api-fiber-gorm/model"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	// DB.AutoMigrate(
	// 	&model.Country{}, &model.Department{}, &model.City{}, &model.Address{},
	// 	&model.UserType{}, &model.User{}, &model.Pet{},
	// 	&model.UserAddress{}, &model.UserPet{}, &model.PetVaccine{},
	// )
	DB.AutoMigrate(
		&model.Country{}, &model.Department{}, &model.City{},
		&model.UserType{}, &model.User{}, &model.UserAddress{},
		&model.Pet{}, &model.UserPet{}, &model.PetVaccine{},
	)

	// DB.Create(&model.UserType{Role: "normal"}) // init database !import!
	fmt.Println("Database Migrated")
}