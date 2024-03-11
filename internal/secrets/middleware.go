package secrets

import "github.com/gofiber/fiber/v2"

const NoMiddlewareKsmConfig = "none"

func NewKsmEngineFromConfig() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if cfg := c.Get("KSM_CONFIG", "none"); cfg != NoMiddlewareKsmConfig {
			eng, err := NewKsmEngine(cfg)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status": "error",
					"error":  "failed to initialize KSM engine",
				})
			}
			c.Locals("KSM_ENGINE", eng)
			return c.Next()
		}
		c.Locals("KSM_ENGINE", KsmEngine)

		return c.Next()
	}
}
