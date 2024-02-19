package email

import "errors"

var (
	ErrPublicKeyRequired  = errors.New("public key is required")
	ErrPrivateKeyRequired = errors.New("private key is required")
)

type Provider interface {
	GetTemplates(credential *TemplateInput) (*TemplateList, error)
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
	HTML        string      `mailjet:"HTML-part" json:"html"`
	Text        string      `mailjet:"Text-part" json:"text"`
	MJMLContent interface{} `mailjet:"MJMLContent" json:"mjml_content"`
	Heasders    interface{} `mailjet:"Headers" json:"headers"`
}

type TemplateInput struct {
	ID         int64  `json:"id"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	SortBy     string `json:"sort_by"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}
