package email

import (
	"errors"
	"template-manager/internal/entity"
)

var (
	ErrPublicKeyRequired  = errors.New("public key is required")
	ErrPrivateKeyRequired = errors.New("private key is required")
)

type Provider interface {
	GetTemplates(input *TemplateQuery) (*TemplateList, error)
}

type Template struct {
	ID          int64  `mailjet:"ID" json:"id"`
	Name        string `mailjet:"Name" json:"name"`
	Author      string `mailjet:"Author" json:"author"`
	Description string `mailjet:"Description" json:"description"`
}

type TemplateList struct {
	Count int        `mailjet:"Count" json:"count"`
	Total int        `mailjet:"Total" json:"total"`
	Data  []Template `mailjet:"Data" json:"data"`
}

type TemplateContent struct {
	HTMLContent string      `mailjet:"HTML-part" json:"html_content"`
	TextContent string      `mailjet:"Text-part" json:"text_content"`
	MJMLContent interface{} `mailjet:"MJMLContent" json:"mjml_content"`
	Headers     entity.Map  `mailjet:"Headers" json:"headers"`
}

type TemplateContentList struct {
	Count int               `mailjet:"Count" json:"count"`
	Total int               `mailjet:"Total" json:"total"`
	Data  []TemplateContent `mailjet:"Data" json:"data"`
}

type AuthCredential struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type TemplateQuery struct {
	ID     int64  `json:"id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	SortBy string `json:"sort_by"`
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
	AuthCredential
}
