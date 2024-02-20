package mailjet

import (
	jsoniter "github.com/json-iterator/go"
	"template-manager/pkg/email"

	mailjet "github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
)

type Mailjet struct {
	marshaler jsoniter.API
}

func New() *Mailjet {
	return &Mailjet{
		marshaler: jsoniter.Config{TagKey: "mailjet"}.Froze(),
	}
}

func (m *Mailjet) initMailjetClient(input *email.AuthCredential) (*mailjet.Client, error) {
	if err := m.validateCredential(input); err != nil {
		return nil, err
	}
	return mailjet.NewMailjetClient(input.PublicKey, input.PrivateKey), nil
}

func (m *Mailjet) validateCredential(input *email.AuthCredential) error {
	if input.PublicKey == "" {
		return email.ErrPublicKeyRequired
	}
	if input.PrivateKey == "" {
		return email.ErrPrivateKeyRequired
	}
	return nil
}

func (m *Mailjet) GetTemplates(input *email.TemplateQuery) (*email.TemplateList, error) {
	client, err := m.initMailjetClient(&input.AuthCredential)
	if err != nil {
		return nil, err
	}
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

func (m *Mailjet) GetTemplate(input *email.TemplateQuery) (*email.TemplateList, error) {
	client, err := m.initMailjetClient(&input.AuthCredential)
	if err != nil {
		return nil, err
	}
	var data []resources.Template
	mr := &mailjet.Request{
		Resource: "template",
		ID:       input.ID,
	}
	if err := client.Get(mr, &data); err != nil {
		return nil, err
	}
	marshaledData, err := m.marshaler.Marshal(data)
	if err != nil {
		return nil, err
	}
	response := new(email.TemplateList)
	if err := m.marshaler.Unmarshal(marshaledData, response.Data); err != nil {
		return nil, err
	}
	return response, nil
}

func (m *Mailjet) AddTemplate(input *email.TemplateInput) (*email.TemplateList, error) {
	client, err := m.initMailjetClient(&input.AuthCredential)
	if err != nil {
		return nil, err
	}
	var data []resources.Template
	mr := &mailjet.Request{
		Resource: "template",
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: resources.Template{
			Name:        input.Name,
			Description: input.Description,
			Author:      input.Author,
		},
	}
	if err := client.Post(fmr, &data); err != nil {
		return nil, err
	}
	marshaledData, err := m.marshaler.Marshal(data)
	if err != nil {
		return nil, err
	}
	response := new(email.TemplateList)
	if err := m.marshaler.Unmarshal(marshaledData, response.Data); err != nil {
		return nil, err
	}
	return response, nil
}

func (m *Mailjet) UpdateTemplate(input *email.TemplateInput) error {
	client, err := m.initMailjetClient(&input.AuthCredential)
	if err != nil {
		return err
	}
	fields := make([]string, 0)
	payload := resources.Template{}
	if input.Name != "" {
		payload.Name = input.Name
		fields = append(fields, "Name")
	}
	if input.Description != "" {
		payload.Description = input.Description
		fields = append(fields, "Description")
	}
	if input.Author != "" {
		payload.Author = input.Author
		fields = append(fields, "Author")
	}
	mr := &mailjet.Request{
		Resource: "template",
		ID:       input.ID,
	}
	fmr := &mailjet.FullRequest{
		Info:    mr,
		Payload: payload,
	}
	if err := client.Put(fmr, fields); err != nil {
		return err
	}
	return nil
}

func (m *Mailjet) DeleteTemplate(input *email.TemplateQuery) error {
	client, err := m.initMailjetClient(&input.AuthCredential)
	if err != nil {
		return err
	}
	mr := &mailjet.Request{
		Resource: "template",
		ID:       input.ID,
	}
	if err := client.Delete(mr); err != nil {
		return err
	}
	return nil
}

func (m *Mailjet) GetTemplateContent(input *email.TemplateQuery) (*email.TemplateContentList, error) {
	client, err := m.initMailjetClient(&input.AuthCredential)
	if err != nil {
		return nil, err
	}
	var data []resources.TemplateDetailcontent
	mr := &mailjet.Request{
		Resource: "template",
		ID:       input.ID,
		Action:   "detailcontent",
	}
	if err := client.Get(mr, &data); err != nil {
		return nil, err
	}
	marshaledData, err := m.marshaler.Marshal(data)
	if err != nil {
		return nil, err
	}
	response := new(email.TemplateContentList)
	if err := m.marshaler.Unmarshal(marshaledData, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (m *Mailjet) AddTemplateContent(input *email.TemplateInput) (*email.TemplateContentList, error) {
	client, err := m.initMailjetClient(&input.AuthCredential)
	if err != nil {
		return nil, err
	}
	var data []resources.TemplateDetailcontent
	mr := &mailjet.Request{
		Resource: "template",
		ID:       input.ID,
		Action:   "detailcontent",
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: resources.TemplateDetailcontent{
			Headers:  input.Headers,
			HtmlPart: input.HTMLContent,
			TextPart: input.TextContent,
		},
	}
	if err := client.Post(fmr, &data); err != nil {
		return nil, err
	}
	marshaledData, err := m.marshaler.Marshal(data)
	if err != nil {
		return nil, err
	}
	response := new(email.TemplateContentList)
	if err := m.marshaler.Unmarshal(marshaledData, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (m *Mailjet) UpdateTemplateContent(input *email.TemplateInput) error {
	client, err := m.initMailjetClient(&input.AuthCredential)
	if err != nil {
		return err
	}
	fields := make([]string, 0)
	payload := resources.TemplateDetailcontent{}
	if input.HTMLContent != "" {
		payload.HtmlPart = input.HTMLContent
		fields = append(fields, "HtmlPart")
	}
	if input.TextContent != "" {
		payload.TextPart = input.TextContent
		fields = append(fields, "TextPart")
	}
	if input.Headers != nil {
		payload.Headers = input.Headers
		fields = append(fields, "Headers")
	}
	mr := &mailjet.Request{
		Resource: "template",
		ID:       input.ID,
		Action:   "detailcontent",
	}
	fmr := &mailjet.FullRequest{
		Info:    mr,
		Payload: payload,
	}
	if err := client.Put(fmr, fields); err != nil {
		return err
	}
	return nil
}
