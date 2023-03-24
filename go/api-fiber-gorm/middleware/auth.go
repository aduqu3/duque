package middleware

import (
	"api-fiber-gorm/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Config("SECRET")),
		ErrorHandler: jwtError,
	})
}

func ProtectedByRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the user has a valid JWT token
		err := jwtware.New(jwtware.Config{
			SigningKey:   []byte(config.Config("SECRET")),
			ErrorHandler: jwtError,
		})(c)
		if err != nil {
			return err
		}

		// Get the user from the context
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		// Check if the user has a valid role
		if len(allowedRoles) > 0 {

			if _, ok := claims["role"]; !ok {
				return fiber.NewError(fiber.StatusForbidden, "no 'role' field found in the JWT token")
			}

			role := claims["role"].(string)

			// Check if the user's role is allowed
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					// fmt.Println("entra")
					// fmt.Println(c.Next())
					// return c.Next()
					return nil
				}
			}

			return fiber.NewError(fiber.StatusForbidden, "You don't have permission to access this route.")
		}

		// return fiber.NewError(fiber.StatusForbidden, "You don't have permission to access this route.")
		return c.Next()
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
