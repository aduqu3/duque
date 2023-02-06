package handler

import (
	"api-fiber-gorm/config"
	"api-fiber-gorm/model"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Fileupload(c *fiber.Ctx) error {

	// parse incomming image file
	normal := new(model.Pet)
	c.BodyParser(&normal)
	
	fmt.Println(normal.Name)

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

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
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN

	// imageUrl := fmt.Sprintf("http://localhost:8000/media/%s", image)

	// create meta data and send to client

	// data := map[string]interface{}{

	// 	"imageName": image,
	// 	"imageUrl":  imageUrl,
	// 	"header":    file.Header,
	// 	"size":      file.Size,
	// }

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": nil})
}
