package app

import (
	"log/slog"
	"template-manager/internal/app/auth"
	"template-manager/internal/app/session"
	"template-manager/internal/app/template"
	"template-manager/internal/pkg/email"
	"template-manager/pkg/config"
	"template-manager/pkg/repository"
)

type App struct {
	TemplateApp *template.App
	AuthApp     *auth.App
}

func NewApp(conf *config.Config, mails email.Provider, logger *slog.Logger, repo *repository.Container, sessionManager *session.Session) *App {
	return &App{
		TemplateApp: template.New(conf, logger, repo),
		AuthApp:     auth.New(conf, mails, logger, repo, sessionManager),
	}
}
