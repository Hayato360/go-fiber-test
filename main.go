package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-test/routes"

)

func main() {
	app := fiber.New()
	routes.InetRoutes(app)
	

	app.Listen(":3000")

}
