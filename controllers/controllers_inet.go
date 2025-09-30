package controllers

import (
	m "go-fiber-test/models"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func BodyParserTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name)
	log.Println(p.Pass)
	str := p.Name + p.Pass

	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {
	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	c.Query("search")

	a := c.Query("search")
	str := "my search is " + a
	return c.JSON(str)
}

func ValidTest(c *fiber.Ctx) error {

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	return c.JSON(user)
}

func FactorialEnd(c *fiber.Ctx) error {
	numStr := c.Params("num")
	num, err := strconv.Atoi(numStr)
	if err != nil || num < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid number"})
	}
	fact := 1
	for i := 2; i <= num; i++ {
		fact *= i
	}
	return c.JSON(fiber.Map{
		"number":    num,
		"factorial": fact,
	})
}

func AsciiConverter(c *fiber.Ctx) error {
	taxId := c.Query("tax_id")
	if taxId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "tax_id query parameter is required",
		})
	}

	var asciiValues []int
	for _, char := range taxId {
		asciiValues = append(asciiValues, int(char))
	}

	return c.JSON(fiber.Map{
		"input":        taxId,
		"ascii_values": asciiValues,
	})
}
