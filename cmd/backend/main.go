package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/samber/lo"
)

func main() {
	app := fiber.New()

	app.Static("", ".")

	app.Use(cors.New())
	app.Get("/time", func(c *fiber.Ctx) error {
		return c.SendString(time.Now().Format(time.RFC3339))
	})

	lo.Must0(app.ListenTLS(
		"0.0.0.0:8080",
		"../../../local/htmx/cert/conf/live/htmx.space/fullchain.pem",
		"../../../local/htmx/cert/conf/live/htmx.space/privkey.pem",
	))
}
