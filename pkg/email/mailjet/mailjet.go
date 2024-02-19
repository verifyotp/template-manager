package mailjet

import (
	// "context"
	// "errors"
	// "fmt"
	// "log"

	jsoniter "github.com/json-iterator/go"
	"template-manager/pkg/email"

	mailjet "github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
	// "template-manager/internal/pkg/email"
)

type Mailjet struct {
	marshaler jsoniter.API
}

func New() *Mailjet {
	return &Mailjet{
		marshaler: jsoniter.Config{TagKey: "mailjet"}.Froze(),
	}
}

func (m *Mailjet) validateCredential(credential *email.TemplateInput) error {
	if credential.PublicKey == "" {
		return email.ErrPublicKeyRequired
	}
	if credential.PrivateKey == "" {
		return email.ErrPrivateKeyRequired
	}
	return nil
}

func (m *Mailjet) GetTemplates(credential *email.TemplateInput) (*email.TemplateList, error) {
	if err := m.validateCredential(credential); err != nil {
		return nil, err
	}
	client := mailjet.NewMailjetClient(credential.PublicKey, credential.PrivateKey)
	var data []resources.Template
	count, total, err := client.List("template", &data)
	if err != nil {
		return nil, err
	}
	marshaledData, err := m.marshaler.Marshal(data)
	if err != nil {
		return nil, err
	}
	response := &email.TemplateList{
		Count: count,
		Total: total,
	}
	if err := m.marshaler.Unmarshal(marshaledData, response.Data); err != nil {
		return nil, err
	}
	return response, nil
}

func (m *Mailjet) GetTemplate(credential *email.TemplateInput) (*email.TemplateList, error) {
	if err := m.validateCredential(credential); err != nil {
		return nil, err
	}
	client := mailjet.NewMailjetClient(credential.PublicKey, credential.PrivateKey)
	var data []resources.Template
	mr := &mailjet.Request{
		Resource: "template",
		ID:       credential.ID,
	}
	if err := client.Get(mr, &data); err != nil {
		return nil, err
	}
	marshaledData, err := m.marshaler.Marshal(data)
	if err != nil {
		return nil, err
	}
	response := &email.TemplateList{}
	if err := m.marshaler.Unmarshal(marshaledData, response.Data); err != nil {
		return nil, err
	}
	return response, nil
}

func (m *Mailjet) GetTemplateContent(credential *email.TemplateInput) (*email.TemplateContent, error) {
	if err := m.validateCredential(credential); err != nil {
		return nil, err
	}
	client := mailjet.NewMailjetClient(credential.PublicKey, credential.PrivateKey)
	var data []resources.TemplateDetailcontent
	mr := &mailjet.Request{
		Resource: "template",
		ID:       credential.ID,
		Action:   "detailcontent",
	}
	if err := client.Get(mr, &data); err != nil {
		return nil, err
	}
	marshaledData, err := m.marshaler.Marshal(data)
	if err != nil {
		return nil, err
	}
	response := &email.TemplateContent{}
	if err := m.marshaler.Unmarshal(marshaledData, response); err != nil {
		return nil, err
	}
	return response, nil
}
