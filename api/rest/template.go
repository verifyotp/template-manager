package rest

import fiber "github.com/gofiber/fiber/v2"

func (s *server) AddTemplate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func (s *server) GetTemplate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func (s *server) UpdateTemplate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func (s *server) DeleteTemplate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}
