package app

import (
	"fmt"

	"hexagonal.software/ksm-api/internal/config"
	"hexagonal.software/ksm-api/internal/secrets"
)

type Application struct {
	srv *Server
	cfg *config.Config
}

func NewApplication(c *config.Config) *Application {
	return &Application{
		cfg: c,
	}
}

func (a *Application) Bootstrap() error {
	a.srv = NewServer(&a.cfg.Server)

	if err := a.srv.InitServer(); err != nil {
		return err
	}

	registerRoutes(a.srv.srv)

	secrets.BootstrapKsmEngine(&a.cfg.KV)

	return nil
}

func (a *Application) Shutdown() {
	fmt.Println("Shutting down server...")
	if a.srv != nil {
		a.srv.StopServer()
	}
}

func (a *Application) RunServer() error {
	return a.srv.RunServer()
}
