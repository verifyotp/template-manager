package shared

import (
	"time"

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
	AccessKeyName string `json:"name"`
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

type GetUploadURLRequest struct {
	AccountID   string `json:"account_id"`
	ContentType string `json:"content_type"`
	Name        string `json:"name"`
}

func (r GetUploadURLRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.AccountID, validation.Required),
		validation.Field(&r.ContentType, validation.Required),
		validation.Field(&r.Name, validation.Required),
	)
}

type UploadURLResponse struct {
	AccountID   string    `json:"account_id"`
	ContentType string    `json:"content_type"`
	URL         string    `json:"url"`
	ExpireAt    time.Time `json:"expire_at"`
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
	AccountID string
	Page      int
	PageSize  int
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

type ImportTemplateRequest struct {
	AccountID          string     `json:"account_id"`
	Provider           string     `json:"provider"`
	ProviderTemplateID string     `json:"provider_template_id"`
	Credentials        entity.Map `json:"credentials"`
}

func (r ImportTemplateRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.AccountID, validation.Required),
		validation.Field(&r.Provider, validation.Required),
		validation.Field(&r.ProviderTemplateID, validation.Required),
		validation.Field(&r.Credentials, validation.Required),
	)
}

type ExportTemplateRequest struct {
	AccountID   string     `json:"account_id"`
	TemplateID  string     `json:"template_id"`
	Provider    string     `json:"provider"`
	Credentials entity.Map `json:"credentials"`
}

func (r ExportTemplateRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.AccountID, validation.Required),
		validation.Field(&r.TemplateID, validation.Required),
		validation.Field(&r.Provider, validation.Required),
		validation.Field(&r.Credentials, validation.Required),
	)
}

type CredentialInput struct {
	ID string `json:"id"`
	// AccountID string `json:"account_id"`
	Platform entity.Platform     `json:"platform"`
	Type     entity.PlatformType `json:"type"`
	Meta     entity.Map          `json:"meta"`
}

func (c *CredentialInput) ValidateCreate() error {
	return validation.ValidateStruct(c,
		// validation.Field(&c.AccountID, validation.Required),
		validation.Field(&c.Platform, validation.Required, validation.In(entity.MAILJET, entity.MAILGUN)),
		validation.Field(&c.Type, validation.Required, validation.In(entity.EMAIL, entity.SMS, entity.PUSH)),
		validation.Field(&c.Meta, validation.Required),
	)
}

func (c *CredentialInput) ValidateUpdate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ID, validation.Required),
		// validation.Field(&c.AccountID, validation.Empty),
		validation.Field(&c.Platform, validation.In(entity.MAILJET, entity.MAILGUN)),
		validation.Field(&c.Type, validation.In(entity.EMAIL, entity.SMS, entity.PUSH)),
	)
}
