package rest

import (
	"template-manager/internal/shared"

	fiber "github.com/gofiber/fiber/v2"
)

func (s *server) AddCredential(c *fiber.Ctx) error {
	req := new(shared.CredentialInput)
	if err := c.BodyParser(req); err != nil {
		return HandleBadRequest(c, err)
	}

	accountID := c.Locals("account_id").(string)
	if err := s.credentialApp.Create(c.Context(), accountID, req); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "credential created successfully", nil)
}

func (s *server) UpdateCredential(c *fiber.Ctx) error {
	req := new(shared.CredentialInput)
	if err := c.BodyParser(req); err != nil {
		return HandleBadRequest(c, err)
	}

	if err := s.credentialApp.Update(c.Context(), req); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "credential updated successfully", nil)
}

func (s *server) GetCredentials(c *fiber.Ctx) error {
	accountID := c.Locals("account_id").(string)
	credentials, err := s.credentialApp.GetByAccountID(c.Context(), accountID)
	if err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "credentials retrieved successfully", credentials)
}

func (s *server) DeleteCredential(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := s.credentialApp.Delete(c.Context(), id); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "credential deleted successfully", nil)
}

func (s *server) GetCredential(c *fiber.Ctx) error {
	id := c.Params("id")
	credential, err := s.credentialApp.GetByID(c.Context(), id)
	if err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "credential retrieved successfully", credential)
}
