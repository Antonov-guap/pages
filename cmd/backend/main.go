package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func main() {
	app := fiber.New()
	app.Get("/time", func(c *fiber.Ctx) error {
		return c.SendString(time.Now().String())
	})
	lo.Must0(app.Listen("0.0.0.0:8080"))
}
