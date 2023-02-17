package handler

import (
	"api-fiber-gorm/config"
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
	"api-fiber-gorm/repository"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreatePet and Own
func CreatePetOwn(c *fiber.Ctx) error {
	// user_id, err := c.ParamsInt("id")

	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	// }

	pet := new(model.Pet)
	if err := c.BodyParser(pet); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	user_id := GetUserIdOfToken(c)

	// if u_id := GetUserIdOfToken(c); u_id != user_id {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	// }

	// get image of pet
	// file, err := c.FormFile("image")

	// // if get image dont present error, save image to folder
	// if err == nil {
	// 	log.Println("saving image --> ", err)

	// 	// generate new uuid for image name
	// 	uniqueId := uuid.New()

	// 	// remove "- from imageName"
	// 	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// 	// extract image extension from original file filename
	// 	fileExt := strings.Split(file.Filename, ".")[1]

	// 	// generate image from filename and extension
	// 	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// 	// save image to ./images dir
	// 	err = c.SaveFile(file, fmt.Sprintf(config.Config("MEDIA_PATH")+"%s", image))

	// 	if err == nil {
	// 		pet.Image = image
	// 		log.Println("image saved and saving in pet")
	// 	}
	// } else {
	// 	log.Println("error in saving image --> ", err)
	// }

	pet_data, err := repository.CreatePet(*pet)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create pet", "data": err})
	}

	mdl_own := model.Own{
		PetId:  int(pet_data.ID),
		UserId: user_id,
	}

	own_pet, err := repository.CreateOwnPet(mdl_own)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user own", "data": err})
	}

	own_pet.Pet = pet_data

	return c.JSON(fiber.Map{"status": "success", "message": "Created pet and user own pet", "data": own_pet})
}

// PutImage updates only column image
func PutPetImage(c *fiber.Ctx) error {
	pet_id, err := c.ParamsInt("id")

	user_id := GetUserIdOfToken(c)

	if repository.ValidateOwnerPet(user_id, pet_id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// get image of pet
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update image", "data": err})
	}

	// if get image dont present error, save image to folder
	// log.Println("saving image --> ", err)

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"
	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename
	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf(config.Config("MEDIA_PATH")+"%s", image))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update image", "data": err})
	}

	// pet_id
	db := database.DB
	// if err := db.Create(&mdl).Error; err != nil {
	// 	return mdl, err
	// }
	// return mdl, nil

	// pet_ := model.Pet{
	// 	ID: pet_id,
	// 	Name: ,
	// }

	// newUser := NewUser{
	// 	Email:      user.Email,
	// 	Username:   user.Username,
	// 	FirstName:  user.FirstName,
	// 	LastName:   user.LastName,
	// 	UserTypeId: user.UserTypeId,
	// }
	// 	db := database.DB

	var pet_ model.Pet
	if err = db.Find(&pet_, pet_id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't find pet", "data": err})
	}

	db.Model(&pet_).Updates(model.Pet{Image: image})

	// pet.Image = image
	// log.Println("image saved and saving in pet")
	return c.JSON(fiber.Map{"status": "success", "message": "Updated pet image", "data": "own_pet"})
}

// GetUser get userdata of user authenticate
func GetPet(c *fiber.Ctx) error {
	pet_id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found id", "data": nil})
	}

	user_id := GetUserIdOfToken(c)

	if repository.ValidateOwnerPet(user_id, pet_id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	mdl_pet, err := repository.GetPet(pet_id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Pet found", "data": mdl_pet})
}
