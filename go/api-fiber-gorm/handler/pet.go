package handler

import (
	"api-fiber-gorm/config"
	"api-fiber-gorm/model"
	"api-fiber-gorm/repository"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreatePet and Own
func CreatePetOwn(c *fiber.Ctx) error {
	user_id, err := c.ParamsInt("user_id")

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	pet := new(model.Pet)
	if err := c.BodyParser(pet); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if u_id := GetUserIdOfToken(c); u_id != user_id {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Forbidden action", "data": nil})
	}

	// get image of pet
	file, err := c.FormFile("image")

	// if get image dont present error, save image to folder
	if err == nil {
		log.Println("image upload error --> ", err)
		// return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

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

		if err == nil {
			fmt.Println(image)
			pet.Image = image
			log.Println("image saved and passed to pet")
			// log.Println("image save error --> ", err)
			// return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}
	}

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
