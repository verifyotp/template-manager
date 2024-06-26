package rest

import (
	"template-manager/internal/shared"

	fiber "github.com/gofiber/fiber/v2"
)

func (s *server) GetUploadURL(c *fiber.Ctx) error {
	var req shared.GetUploadURLRequest
	if err := c.BodyParser(&req); err != nil {
		return HandleBadRequest(c, err)
	}

	if err := req.Validate(); err != nil {
		return HandleBadRequest(c, err)
	}

	uploadURL, err := s.templateApp.GetUploadURL(c.Context(), req)
	if err != nil {
		return HandleError(c, err)
	}

	return HandleSuccess(c, "upload url retrieved successfully", uploadURL)
}

func (s *server) AddTemplate(c *fiber.Ctx) error {
	var req shared.CreateTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return HandleBadRequest(c, err)
	}

	req.AccountID = c.Locals("account_id").(string)
	if err := req.Validate(); err != nil {
		return HandleBadRequest(c, err)
	}

	if err := s.templateApp.Create(c.Context(), req); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "template created successfully", nil)
}

func (s *server) GetTemplate(c *fiber.Ctx) error {

	var req = shared.GetTemplateRequest{
		AccountID:  c.Locals("account_id").(string),
		TemplateID: c.Params("id"),
	}

	if err := req.Validate(); err != nil {
		return HandleBadRequest(c, err)
	}

	template, err := s.templateApp.Get(c.Context(), req)
	if err != nil {
		return HandleError(c, err)
	}

	return HandleSuccess(c, "template retrieved successfully", template)
}

func (s *server) ListTemplates(c *fiber.Ctx) error {
	var req = shared.ListTemplatesRequest{
		AccountID: c.Locals("account_id").(string),
		Page:      c.QueryInt("page", 1),
		PageSize:  c.QueryInt("page_size", 10),
	}

	if err := req.Validate(); err != nil {
		return HandleBadRequest(c, err)
	}

	templates, err := s.templateApp.List(c.Context(), req)
	if err != nil {
		return HandleError(c, err)
	}

	return HandleSuccess(c, "templates retrieved successfully", templates)
}

func (s *server) UpdateTemplate(c *fiber.Ctx) error {
	var req shared.UpdateTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return HandleBadRequest(c, err)
	}
	req.AccountID = c.Locals("account_id").(string)
	req.TemplateID = c.Params("id")

	if err := req.Validate(); err != nil {
		return HandleBadRequest(c, err)
	}

	if err := s.templateApp.Update(c.Context(), req); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "template updated successfully", nil)
}

func (s *server) EditTemplate(c *fiber.Ctx) error {
	var req shared.UpdateTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return HandleBadRequest(c, err)
	}
	req.AccountID = c.Locals("account_id").(string)
	req.TemplateID = c.Params("id")

	if err := req.Validate(); err != nil {
		return HandleBadRequest(c, err)
	}

	if err := s.templateApp.Edit(c.Context(), req); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "template updated successfully", nil)
}

func (s *server) DeleteTemplate(c *fiber.Ctx) error {
	var req shared.DeleteTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return HandleBadRequest(c, err)
	}

	req.AccountID = c.Locals("account_id").(string)
	req.TemplateID = c.Params("id")

	if err := req.Validate(); err != nil {
		return HandleBadRequest(c, err)
	}

	if err := s.templateApp.Delete(c.Context(), req); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "template deleted successfully", nil)
}

func (s *server) ImportTemplate(c *fiber.Ctx) error {
	var req shared.ImportTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return HandleBadRequest(c, err)
	}
	req.AccountID = c.Locals("account_id").(string)

	if err := req.Validate(); err != nil {
		return HandleBadRequest(c, err)
	}

	if err := s.templateApp.Import(c.Context(), req); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "template imported successfully", nil)
}

func (s *server) ExportTemplate(c *fiber.Ctx) error {
	var req shared.ExportTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return HandleBadRequest(c, err)
	}
	req.AccountID = c.Locals("account_id").(string)

	if err := req.Validate(); err != nil {
		return HandleBadRequest(c, err)
	}

	if err := s.templateApp.Export(c.Context(), req); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, "template exported successfully", nil)
}
