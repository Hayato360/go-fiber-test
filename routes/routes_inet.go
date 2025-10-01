package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "1234",
		},
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", c.HelloTest)

	v1.Post("/", c.BodyParserTest)

	v1.Get("/user/:name", c.ParamsTest)

	v1.Post("/inet", c.QueryTest)

	v1.Post("/valid", c.ValidTest)

	// Mock data route
	v1.Post("/mock-data", c.CreateMockData)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

	//CRUD companies
	company := v1.Group("/company")
	company.Get("/", c.GetCompanies)
	company.Get("/:id", c.GetCompany)
	company.Get("/search", c.GetCompanyByName)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)

}
