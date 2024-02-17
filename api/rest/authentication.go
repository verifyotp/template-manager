package rest

import (
	"strings"

	fiber "github.com/gofiber/fiber/v2"

	"template-manager/internal/shared"
)

func (s *server) Signup(c *fiber.Ctx) error {
	ctx := c.Context()

	var request shared.SignUpRequest
	err := c.BodyParser(&request)
	if err != nil {
		return HandleBadRequest(c, err)
	}

	err = s.authApp.Signup(ctx, request)
	if err != nil {
		return HandleError(c, err)
	}

	return HandleSuccess(c, "check your email to continue sign up", nil)
}

func (s *server) Login(c *fiber.Ctx) error {
	var request shared.LoginRequest
	err := c.BodyParser(&request)
	if err != nil {
		return HandleBadRequest(c, err)
	}

	//get device info
	request.Device = getDevice(c).Transform()
	response, err := s.authApp.Login(c.Context(), request)
	if err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "welcome back", response)
}

func getDevice(c *fiber.Ctx) shared.Device {
	return shared.Device{
		IP:        c.IP(),
		UserAgent: c.Get("User-Agent"),
	}
}

func (s *server) InitiateResetPassword(c *fiber.Ctx) error {
	var request shared.InitiateResetPasswordRequest
	err := c.BodyParser(&request)
	if err != nil {
		return HandleBadRequest(c, err)
	}

	err = s.authApp.InitiateResetPassword(c.Context(), request)
	if err != nil {
		return HandleError(c, err)
	}

	return HandleSuccess(c, "check your email to continue password reset", nil)
}

func (s *server) Logout(c *fiber.Ctx) error {
	var request shared.LogoutRequest
	err := c.BodyParser(&request)
	if err != nil {
		return HandleBadRequest(c, err)
	}

	request.Token = strings.ReplaceAll(request.Token, "Bearer ", "")
	err = s.authApp.Logout(c.Context(), request)
	if err != nil {
		return HandleError(c, err)
	}

	return HandleSuccess(c, "successfully logged out", nil)
}
