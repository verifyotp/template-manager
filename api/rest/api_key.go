package rest

import (
	fiber "github.com/gofiber/fiber/v2"

	"template-manager/internal/shared"
)

func (s *server) AddKey(c *fiber.Ctx) error {
	ctx := c.Context()

	var request shared.CreateAccessKeyRequest
	err := c.BodyParser(&request)
	if err != nil {
		return HandleBadRequest(c, err)
	}
	request.AccountID = c.Locals("account_id").(string)
	err = s.authApp.CreateAccessKey(ctx, request)
	if err != nil {
		return HandleError(c, err)
	}

	return HandleSuccess(c, "successfully created key", nil)
}

func (s *server) ListAccessKeys(c *fiber.Ctx) error {
	ctx := c.Context()
	var request = shared.ListAccessKeysRequest{
		AccountID: c.Locals("account_id").(string),
	}
	keys, err := s.authApp.ListAccessKeys(ctx, request)
	if err != nil {
		return HandleError(c, err)
	}

	return HandleSuccess(c, "successfully retrieved key", keys)
}

func (s *server) DeleteKey(c *fiber.Ctx) error {
	ctx := c.Context()
	ID := c.Params("id")

	request := shared.DeleteAccessKeyRequest{
		AccessKeyID: ID,
	}
	err := s.authApp.DeleteAccessKey(ctx, request)
	if err != nil {
		return HandleError(c, err)
	}

	return HandleSuccess(c, "successfully deleted key", nil)
}
