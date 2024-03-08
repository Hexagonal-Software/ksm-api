package secrets

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrRetrieveSecret  = errors.New("failed to retrieve secret")
	ErrInvalidNotation = errors.New("invalid notation parameters")
)

func GetSecretByNotation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		notation := fmt.Sprintf("keeper://%s/%s/%s", c.Params("record"), c.Params("type"), c.Params("query"))
		secret, err := GetKsmEngine().GetNotation(notation)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "error",
				"error":  ErrRetrieveSecret.Error(),
			})
		}

		return renderSecret(c, secret)
	}
}

func renderSecret(c *fiber.Ctx, secret []interface{}) error {
	if c.Get("Accept", "text/plain") == "application/json" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"data":   secret,
		})
	}
	c.Request().Header.Set("Accept", "text/plain")

	return c.Status(fiber.StatusOK).SendString(fmt.Sprint(secret[0]))
}
