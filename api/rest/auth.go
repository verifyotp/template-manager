package rest

import (
	fiber "github.com/gofiber/fiber/v2"

	"template-manager/internal/shared"
)

func (s *server) Signup(c *fiber.Ctx) error {
	ctx := c.Context()

	var request shared.SignUpRequest
	err := c.BodyParser(&request)
	if err != nil {
		return HandleError(c, err)
	}

	err = s.app.Signup(ctx, request)
	if err != nil {
		return HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "check your email to continue sign up",
	})
}

func (s *server) InitiateResetPassword(c *fiber.Ctx) error {
	var request shared.InitiateResetPasswordRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = s.app.InitiateResetPassword(c.Context(), request)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "check your email to continue reset password",
	})
}

func (s *server) getDevice(c *fiber.Ctx) shared.Device {
	return shared.Device{
		IP:        c.IP(),
		UserAgent: c.Get("User-Agent"),
	}
}

func (s *server) Login(c *fiber.Ctx) error {
	var request shared.LoginRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	//get device info
	request.Device = s.getDevice(c).Transform()
	response, err := s.app.Login(c.Context(), request)
	if err != nil {
		return HandleError(c, err)
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "welcome back",
		"body":    response,
	})
}

func (s *server) Logout(c *fiber.Ctx) error {
	var request shared.LogoutRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = s.app.Logout(c.Context(), request)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "successfully logged out",
	})
}
