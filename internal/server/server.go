package server

import (
	"github.com/elskow/links.helmyl.com/config"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	fiberConfig := config.NewFiberConfig()

	server := &FiberServer{
		App: fiber.New(fiberConfig),
	}

	return server
}
