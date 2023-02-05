package handler

import (
	"api-fiber-gorm/config"
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func getUserByUsername(u string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Login get user and password
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	identity := input.Identity
	pass := input.Password
	user := new(model.User)
	email := new(model.User)
	err := *new(error)

	if identity == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
	}

	if valid(identity) {
		email, err = getUserByEmail(identity)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
		}
		user = nil
	} else {
		user, err = getUserByUsername(identity)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
		}
		email = nil
	}

	if email == nil && user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	}

	if email != nil {
		ud = UserData{
			ID:       email.ID,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}

	}
	if user != nil {
		ud = UserData{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}
	}

	if !CheckPasswordHash(pass, ud.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password or username", "data": nil})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"username": ud.Username,
		"user_id":  ud.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	}

	// fmt.Println("Vamos a imprimir identity")
	// fmt.Println(identity)

	// claims["username"] = ud.Username
	// claims["user_id"] = ud.ID
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	fmt.Println(user)

	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	name := claims["username"].(string)

	fmt.Println(user.Raw) // token

	return c.SendString("Welcome " + name)
	// return c.JSON(fiber.Map{"status": "success", "message": "Success logout"})
}

func GetUserIdOfToken(c *fiber.Ctx) int {
	user := c.Locals("user").(*jwt.Token)
	fmt.Println(user)

	claims := user.Claims.(jwt.MapClaims)
	// fmt.Println(claims)
	// name := claims["username"].(string)

	// fmt.Println(user.Raw) // token

	// return c.SendString("Welcome " + name)

	user_id := int(claims["user_id"].(float64))

	// fmt.Println(user_id)                 // token
	// fmt.Println(reflect.TypeOf(user_id)) // token
	return user_id

	// return c.JSON(fiber.Map{"status": "success", "message": "Success logout"})
}
