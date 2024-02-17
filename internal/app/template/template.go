package template

import (
	"context"
	"errors"
	"log/slog"

	"template-manager/internal/entity"
	"template-manager/internal/shared"
	"template-manager/pkg/config"
	"template-manager/pkg/email"
	"template-manager/pkg/repository"
)

type App struct {
	config *config.Config
	email  email.Provider
	logger *slog.Logger
	db     *repository.Container // TODO: replace with repository

}

func New(config *config.Config, email email.Provider, logger *slog.Logger, db *repository.Container) *App {
	return &App{
		config: config,
		email:  email,
		db:     db,
		logger: logger,
	}
}

func (a *App) Create(ctx context.Context, req shared.SignUpRequest) error {
	var account = entity.Account{
		Email: req.Email,
	}

	// generate password
	randomPassword := entity.GenerateRandomPassword()
	if err := account.SetPassword(randomPassword); err != nil {
		a.logger.ErrorContext(ctx, "failed to set password %+v", err)
		return err
	}

	// find existing account
	if _, err := a.db.AuthRepository.Get(ctx, "email = ?", req.Email); err == nil {
		return errors.New("account already exists")
	}

	if _, err := a.db.AuthRepository.Create(ctx, &account); err != nil {
		a.logger.ErrorContext(ctx, "failed to create account %+v", err)
		return err
	}

	//send email
	vars := map[string]any{
		"to":           req.Email,
		"subject":      "Password Setup",
		"password":     randomPassword,
		"company_name": "Template Manager",
	}
	if err := a.email.Send(ctx, email.TemplateIDSignupVerification, vars); err != nil {
		a.logger.ErrorContext(ctx, "failed to send email %+v", err)
		return err
	}
	return nil
}

const (
	LoginFailed = "login failed. please check your email and password and try again"
)
