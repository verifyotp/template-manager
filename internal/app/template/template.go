package template

import (
	"context"
	"errors"
	"log/slog"

	"template-manager/internal/entity"
	"template-manager/internal/shared"
	"template-manager/pkg/config"
	"template-manager/pkg/repository"
)

type App struct {
	config *config.Config
	logger *slog.Logger
	db     *repository.Container // TODO: replace with repository

}

func New(config *config.Config, logger *slog.Logger, db *repository.Container) *App {
	return &App{
		config: config,
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

	return nil
}

func (a *App) Update(ctx context.Context, req shared.SignUpRequest) error {
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

	return nil
}

func (a *App) Delete(ctx context.Context, req shared.SignUpRequest) error {
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

	return nil
}
