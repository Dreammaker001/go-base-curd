package v1

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App) *fiber.App {
	api := app.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	return app
}
