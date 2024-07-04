package main

import (
	"log"

	"github.com/elskow/links.helmyl.com/internal/server"
)

func main() {
	server := server.New()

	server.Static("/", "./public")
	server.RegisterFiberRoutes()

	err := server.Listen(":3000")
	if err != nil {
		log.Fatalf("server.Listen: %v", err)
	}
}
