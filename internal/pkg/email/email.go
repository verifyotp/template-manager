package email

import "context"

type TemplateID string

const (
	TemplateIDSignupVerification TemplateID = "signup_verification"
)

type Provider interface {
	Send(ctx context.Context, id TemplateID, vars map[string]any) error
}
