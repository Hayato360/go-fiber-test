package controllers

import (
	"errors"
	m "go-fiber-test/models"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)



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
		"asciiValues": asciiValues,
	})
}


func Register(c *fiber.Ctx) error {
	var request m.RegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	validate := validator.New()

	reUsername := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,20}$`)
	reLineID := regexp.MustCompile(`^[a-zA-Z0-9._-]{4,20}$`)
	reSubdomain := regexp.MustCompile(`^[a-z0-9.]{2,30}$`)

	validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		return reUsername.MatchString(fl.Field().String())
	})
	validate.RegisterValidation("line_id", func(fl validator.FieldLevel) bool {
		return reLineID.MatchString(fl.Field().String())
	})
	
	validate.RegisterValidation("bussiness_type", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		allowedBussinessTypes := map[string]bool{
			"retail":    true,
			"service":   true,
			"it":        true,
			"finance":   true,
			"other":     true,
		}
		return allowedBussinessTypes[val]
	})

	validate.RegisterValidation("subdomain", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()

		if !reSubdomain.MatchString(fl.Field().String()) {
			return false
		}

		allowedSuffixes := []string{
			".com",
			".net",
			".org",
			".co",
			".io",
		}

		for _, suffix := range allowedSuffixes {
			if len(value) > len(suffix) && value[len(value)-len(suffix):] == suffix {
				return true
			}
		}
		return false
	})

	if err := validate.Struct(request); err != nil {
		for _,e := range err.(validator.ValidationErrors) {
			errTxt := "Error on field '" + e.Field() + "', condition: " + e.ActualTag()
			return c.Status(fiber.StatusBadRequest).SendString(
				errors.New(errTxt).Error(),
			)
		}
	}

	return c.Status(fiber.StatusOK).JSON(request)
}