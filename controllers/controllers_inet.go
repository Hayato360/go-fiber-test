package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

type Person struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}

func BodyParserTest(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		log.Println(p.Name)
		log.Println(p.Pass)
		str := p.Name + p.Pass

		return c.JSON(str)
	}