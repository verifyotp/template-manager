package mailjet

import (
	"context"
	"errors"
	"fmt"

	mailjet "github.com/mailjet/mailjet-apiv3-go/v4"

	"template-manager/internal/pkg/email"
)

type Mailjet struct {
	from       string
	fromName   string
	mj         *mailjet.Client
	publicKey  string
	privateKey string
}

var _ email.Provider = (*Mailjet)(nil)

type Option func(m *Mailjet)

// Set the company name that will appear in the email
func WithName(name string) Option {
	return func(m *Mailjet) {
		m.fromName = name
	}
}

func New(publicKey, privateKey, from string, opts ...Option) *Mailjet {
	mailjetClient := mailjet.NewMailjetClient(publicKey, privateKey)
	client := &Mailjet{
		from:       from,
		mj:         mailjetClient,
		publicKey:  publicKey,
		privateKey: privateKey,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

var templateIDMap = map[email.TemplateID]int{
	email.TemplateIDSignupVerification: 5456918,
}

func (m *Mailjet) Send(ctx context.Context, id email.TemplateID, vars map[string]any) error {
	if err := validateVars(vars); err != nil {
		return err
	}
	to := vars["to"].(string)
	subject := vars["subject"].(string)
	vars["company_email"] = m.from
	vars["logo"] = "https://www.templafy.com/wp-content/uploads/2020/02/corporate-management-templafy.png"
	return sendTemplateEmail(ctx, m.mj, templateIDMap[id], m.from, to, subject, vars)
}

func validateVars(vars map[string]any) error {
	if _, ok := vars["to"]; !ok {
		return errors.New("missing to")
	}
	if _, ok := vars["subject"]; !ok {
		return errors.New("missing subject")
	}
	return nil
}

func sendTemplateEmail(ctx context.Context, mailjetClient *mailjet.Client, templateID int, from, to, subject string, variables map[string]any) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: from,
				Name:  from,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
					Name:  to,
				},
			},
			Subject:          subject,
			TemplateID:       templateID,
			Variables:        variables,
			TemplateLanguage: true,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return err
	}
	fmt.Printf("Data: %+v\n", res)
	return nil
}
