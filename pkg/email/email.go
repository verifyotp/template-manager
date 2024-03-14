package email

import (
	"errors"
	"template-manager/internal/entity"
)

var (
	ErrPublicKeyRequired  = errors.New("public key is required")
	ErrPrivateKeyRequired = errors.New("private key is required")
	ErrDomainRequired     = errors.New("domain is required")
)

type Provider interface {
	GetTemplates(input *TemplateQuery) (*TemplateResponse, error)
}

type Template struct {
	ID          int64       `mailjet:"ID" mailgun:"-" json:"id,omitempty"`
	Name        string      `mailjet:"Name" mailgun:"name" json:"name,omitempty"`
	Author      string      `mailjet:"Author" mailgun:"author" json:"author,omitempty"`
	Description string      `mailjet:"Description" mailgun:"description" json:"description,omitempty"`
	Version     interface{} `mailjet:"-" mailgun:"version" json:"version,omitempty"`
}

type TemplateResponse struct {
	HasNext  bool       `json:"has_next,omitempty"`
	Count    int        `json:"count,omitempty"`
	Total    int        `json:"total,omitempty"`
	DataList []Template `json:"data_list,omitempty"`
	Data     Template   `json:"data,omitempty"`
}

type TemplateContent struct {
	HTMLContent string      `mailjet:"HTML-part" json:"html_content,omitempty"`
	TextContent string      `mailjet:"Text-part" json:"text_content,omitempty"`
	MJMLContent interface{} `mailjet:"MJMLContent" json:"mjml_content,omitempty"`
	Headers     entity.Map  `mailjet:"Headers" json:"headers,omitempty"`
}

type TemplateContentResponse struct {
	Count    int               `mailjet:"Count" json:"count,omitempty"`
	Total    int               `mailjet:"Total" json:"total,omitempty"`
	DataList []TemplateContent `mailjet:"Data" json:"data_list,omitempty"`
	Data     TemplateContent   `mailjet:"-" json:"data,omitempty"`
}

type AuthCredential struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	Domain     string `json:"domain"`
}

type TemplateQuery struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	SortBy string `json:"sort_by"`
	Active bool   `json:"active"`
	AuthCredential
}

type TemplateInput struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Author      string     `json:"author"`
	ID          int64      `json:"id"`
	HTMLContent string     `json:"html_content"`
	TextContent string     `json:"text_content"`
	MJMLContent string     `json:"mjml_content"`
	Subject     string     `json:"subject"`
	Headers     entity.Map `json:"headers"`
	Tag         string     `json:"tag"`
	Template    string     `json:"template"`
	Engine      string     `json:"engine"`
	Comment     string     `json:"comment"`
	// Active    bool           `json:"active"`
	AuthCredential
}
