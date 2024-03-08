package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"hexagonal.software/ksm-api/internal/config"
	"hexagonal.software/ksm-api/internal/secrets"
)

type Server struct {
	srv *fiber.App
	cfg *config.Server
}

func NewServer(c *config.Server) *Server {
	return &Server{
		cfg: c,
	}
}

func (s *Server) InitServer() error {
	app := fiber.New(fiber.Config{
		// Override default configuration
		ServerHeader:          "KSM-API",
		StrictRouting:         true,
		CaseSensitive:         true,
		AppName:               "KSM-API",
		CompressedFileSuffix:  ".gz",
		EnablePrintRoutes:     false,
		DisableStartupMessage: true,

		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"status": "error",
			})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())

	s.srv = app

	registerRoutes(s.srv)

	return nil
}

func (s *Server) RunServer() error {
	return s.srv.Listen(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
}

func (s *Server) StopServer() error {
	return s.srv.Shutdown()
}

func registerRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/api/v1/secret/:record/:type/:query", secrets.GetSecretByNotation())
}
