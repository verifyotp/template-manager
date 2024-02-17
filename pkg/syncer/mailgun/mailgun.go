package mailgun

import (
	"context"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v3"
)

type Syncer struct {
}

func ListActiveTemplates(domain, apiKey string) ([]mailgun.Template, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	it := mg.ListTemplates(&mailgun.ListTemplateOptions{Active: true})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var page, result []mailgun.Template
	for it.Next(ctx, &page) {
		result = append(result, page...)
	}

	if it.Err() != nil {
		return nil, it.Err()
	}
	return result, nil
}
