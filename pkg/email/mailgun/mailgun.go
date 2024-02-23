package mailgun

import (
	"context"
	"template-manager/pkg/email"

	jsoniter "github.com/json-iterator/go"
	mailgun "github.com/mailgun/mailgun-go/v3"
)

type Mailgun struct {
	marshaler jsoniter.API
	ctx       context.Context
}

func New(ctx context.Context) *Mailgun {
	return &Mailgun{
		ctx:       ctx,
		marshaler: jsoniter.Config{TagKey: "mailgun"}.Froze(),
	}
}

func (m *Mailgun) initMailgunClient(input *email.AuthCredential) (*mailgun.MailgunImpl, error) {
	if err := m.validateCredential(input); err != nil {
		return nil, err
	}
	return mailgun.NewMailgun(input.Domain, input.PrivateKey), nil
}

func (m *Mailgun) validateCredential(input *email.AuthCredential) error {
	if input.Domain == "" {
		return email.ErrDomainRequired
	}
	if input.PrivateKey == "" {
		return email.ErrPrivateKeyRequired
	}
	return nil
}

func (m *Mailgun) GetTemplates(input *email.TemplateQuery) (*email.TemplateResponse, error) {
	client, err := m.initMailgunClient(&input.AuthCredential)
	if err != nil {
		return nil, err
	}
	var data []mailgun.Template
	result := client.ListTemplates(&mailgun.ListTemplateOptions{})
	hasNext := result.Next(m.ctx, &data)
	marshaledData, err := m.marshaler.Marshal(data)
	if err != nil {
		return nil, err
	}
	response := &email.TemplateResponse{
		HasNext: hasNext,
	}
	if err := m.marshaler.Unmarshal(marshaledData, &response.DataList); err != nil {
		return nil, err
	}
	return response, nil
}

func (m *Mailgun) GetTemplate(input *email.TemplateQuery) (*email.TemplateResponse, error) {
	client, err := m.initMailgunClient(&input.AuthCredential)
	if err != nil {
		return nil, err
	}
	tmpl, err := client.GetTemplate(m.ctx, input.Name)
	if err != nil {
		return nil, err
	}
	marshaledData, err := m.marshaler.Marshal(tmpl)
	if err != nil {
		return nil, err
	}
	response := &email.TemplateResponse{}
	if err := m.marshaler.Unmarshal(marshaledData, &response.Data); err != nil {
		return nil, err
	}
	return response, nil
}

func (m *Mailgun) AddTemplate(input *email.TemplateInput) (*email.TemplateResponse, error) {
	client, err := m.initMailgunClient(&input.AuthCredential)
	if err != nil {
		return nil, err
	}
	version := mailgun.TemplateVersion{
		Active:   true,
		Tag:      input.Tag,
		Template: input.Template,
		Comment:  input.Comment,
	}
	tmpl := &mailgun.Template{
		Name:        input.Name,
		Description: input.Description,
		Version:     version,
	}
	if err := client.CreateTemplate(m.ctx, tmpl); err != nil {
		return nil, err
	}
	return &email.TemplateResponse{}, nil
}

func (m *Mailgun) UpdateTemplate(input *email.TemplateInput) (*email.TemplateResponse, error) {
	client, err := m.initMailgunClient(&input.AuthCredential)
	if err != nil {
		return nil, err
	}
	tmpl := &mailgun.Template{}
	if input.Name != "" {
		tmpl.Name = input.Name
	}
	if input.Description != "" {
		tmpl.Description = input.Description
	}
	if err := client.UpdateTemplate(m.ctx, tmpl); err != nil {
		return nil, err
	}
	return &email.TemplateResponse{}, nil
}

func (m *Mailgun) DeleteTemplate(input *email.TemplateQuery) error {
	client, err := m.initMailgunClient(&input.AuthCredential)
	if err != nil {
		return err
	}
	return client.DeleteTemplate(m.ctx, input.Name)
}
