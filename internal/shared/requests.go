package shared

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mileusna/useragent"

	"template-manager/internal/entity"
)

type SignUpRequest struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`

	Device entity.Device `json:"-"`
}

type LogoutRequest struct {
	AccountID string `json:"account_id"`
	Token     string `json:"token"`
}

type InitiateResetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type DeleteAccessKeyRequest struct {
	AccountID   string `json:"account_id"`
	AccessKeyID string `json:"access_key_id"`
}

type ListAccessKeysRequest struct {
	AccountID string `json:"account_id"`
}

type CreateAccessKeyRequest struct {
	AccountID     string `json:"account_id"`
	AccessKeyName string `json:"access_key_name"`
}

type Device struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

func (d Device) Transform() entity.Device {
	ua := useragent.Parse(d.UserAgent)
	return entity.Device{
		IP:             d.IP,
		UserAgent:      d.UserAgent,
		OS:             ua.OS,
		OSVersion:      ua.OSVersion,
		Browser:        ua.Name,
		BrowserVersion: ua.Version,
	}
}

type CreateTemplateRequest struct {
	AccountID   string     `json:"account_id"`
	Name        string     `json:"name"`
	ContentType string     `json:"content_type"`
	Location    string     `json:"location"`
	Vars        entity.Map `json:"vars"`
}

func (r CreateTemplateRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.AccountID, validation.Required),
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.ContentType, validation.Required),
		validation.Field(&r.Location, validation.Required, is.URL),
	)
}

type UpdateTemplateRequest struct {
	AccountID  string     `json:"account_id"`
	TemplateID string     `json:"template_id"`
	Location   string     `json:"location"`
	Vars       entity.Map `json:"vars"`
}

func (r UpdateTemplateRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.AccountID, validation.Required),
		validation.Field(&r.TemplateID, validation.Required),
		validation.Field(&r.Location, validation.Required, is.URL),
	)
}

type DeleteTemplateRequest struct {
	AccountID  string `json:"account_id"`
	TemplateID string `json:"template_id"`
	Version    uint64 `json:"version"`
}

func (r DeleteTemplateRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.AccountID, validation.Required),
		validation.Field(&r.TemplateID, validation.Required),
		validation.Field(&r.Version, validation.Required, is.Digit),
	)
}

type ListTemplatesRequest struct {
	AccountID string `json:"account_id"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

func (r ListTemplatesRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.AccountID, validation.Required),
		validation.Field(&r.Page, validation.Required),
		validation.Field(&r.PageSize, validation.Required),
	)
}

type GetTemplateRequest struct {
	AccountID  string `json:"account_id"`
	TemplateID string `json:"template_id"`
}

func (r GetTemplateRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.AccountID, validation.Required),
		validation.Field(&r.TemplateID, validation.Required),
	)
}
