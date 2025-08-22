package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"0xroot/internal/config"
)

func main() {

	cfg := config.LoadConfig("config.yaml")
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON("Hello wolrd")
	})

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	app.Listen(addr)
	app.Listen(":3000")
}
