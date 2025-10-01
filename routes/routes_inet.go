package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {


	middleware := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	})

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.Get("/fact/:num", middleware, c.FactorialEnd)
			v1.Post("/register", middleware, c.Register)
		}

		v3 := api.Group("/v3")
		{
			v3.Get("/tew", c.AsciiConverter)
		}
	}
}
