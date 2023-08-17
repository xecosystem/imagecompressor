package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type ServerOptions struct {
	Path    string
	Handler func(*fiber.Ctx) error
}

func NewServer(options ServerOptions) *fiber.App {
	app := fiber.New()
	app.Use(compress.New())
	app.Use(cors.New())
	app.Get(options.Path, options.Handler)
	return app
}
