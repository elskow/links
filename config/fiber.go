package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type FiberConfig struct {
	ServerHeader string
	AppName      string
	Views        *html.Engine
}

func NewFiberConfig() fiber.Config {
	engine := html.New("./template", ".html")
	return fiber.Config{
		ServerHeader: "Links",
		AppName:      "Links",
		Views:        engine,
	}
}
