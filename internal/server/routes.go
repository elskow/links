package server

import (
	"github.com/elskow/links.helmyl.com/config"
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.Get("/", s.Home)
}

func (s *FiberServer) Home(c *fiber.Ctx) error {
	urls := config.NewDefinedUrls()

	return c.Render("index", fiber.Map{
		"Urls": urls,
	})
}
