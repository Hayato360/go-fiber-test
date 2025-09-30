package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"john": "doe",
	// 		"admin": "1234",
	// 		"gofiber":"21022566",
	// 	},
	// }))

	// api := app.Group("/api")
	// v1 := api.Group("/v1")

	// v1.Get("/", c.HelloTest)

	// type Person struct {
	// 	Name string `json:"name"`
	// 	Pass string `json:"pass"`
	// }

	// v1.Post("/", c.BodyParserTest)

	// v1.Get("/user/:name", c.ParamsTest)

	// v1.Post("/inet", c.QueryTest)

	// v1.Post("/valid", c.ValidTest)

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
		}

		v3 := api.Group("/v3")
		{
			v3.Get("/tew", c.AsciiConverter)
		}
	}
}
