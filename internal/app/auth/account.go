package auth

import (
	"context"
	"errors"
	"log/slog"

	"template-manager/internal/entity"
	"template-manager/internal/pkg/email"
	"template-manager/internal/shared"
	"template-manager/pkg/config"
	"template-manager/pkg/repository"
)

type SessionManager interface {
	//create a session
	Create(ctx context.Context, accountID string, device entity.Device) (*entity.Session, error)
	//verify a session
	Verify(ctx context.Context, token string) (*entity.Session, error)
	//delete a session
	Expire(ctx context.Context, token string) error
	//delete all sessions for an account
	Delete(ctx context.Context, accountID string) error
}

type App struct {
	config *config.Config
	email  email.Provider
	logger *slog.Logger
	db     *repository.Container // TODO: replace with repository
	sess   SessionManager
}

func New(config *config.Config, email email.Provider, logger *slog.Logger, db *repository.Container, sessionManager SessionManager) *App {
	return &App{
		config: config,
		email:  email,
		db:     db,
		logger: logger,
		sess:   sessionManager,
	}
}

func (a *App) Signup(ctx context.Context, req shared.SignUpRequest) error {
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

	if err := a.db.AuthRepository.Create(ctx, &account); err != nil {
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

func (a *App) Login(ctx context.Context, req shared.LoginRequest) (*shared.LoginResponse, error) {

	fetchedAccount, err := a.db.AuthRepository.Get(ctx, "email = ?", req.Email)
	if err != nil {
		a.logger.InfoContext(ctx, "failed to find account %+v", err)
		return nil, errors.New(LoginFailed)
	}
	// check password
	if !fetchedAccount.ComparePassword(req.Password) {
		return nil, errors.New(LoginFailed)
	}

	// delete existing sessions
	if err := a.sess.Delete(ctx, fetchedAccount.ID); err != nil {
		a.logger.ErrorContext(ctx, "failed to delete session %+v", err)
		return nil, err
	}

	// create session
	sess, err := a.sess.Create(ctx, fetchedAccount.ID, req.Device)
	if err != nil {
		a.logger.ErrorContext(ctx, "failed to create session %+v", err)
		return nil, err
	}

	// return token
	return &shared.LoginResponse{
		Account: fetchedAccount,
		Session: sess,
	}, nil
}

func (a *App) InitiateResetPassword(ctx context.Context, req shared.InitiateResetPasswordRequest) error {
	var (
		acc *entity.Account
		err error
	)

	if acc, err = a.db.AuthRepository.Get(ctx, "email = ?", req.Email); err != nil {
		return errors.New("account does not exist")
	}

	// generate password
	randomPassword := entity.GenerateRandomPassword()
	if err := acc.SetPassword(randomPassword); err != nil {
		a.logger.ErrorContext(ctx, "failed to set password %+v", err)
		return err
	}

	// update account
	if err := a.db.AuthRepository.Update(ctx, acc); err != nil {
		a.logger.ErrorContext(ctx, "failed to update account %+v", err)
		return err
	}

	//send email
	vars := map[string]any{
		"to":           req.Email,
		"subject":      "Password Reset",
		"password":     randomPassword,
		"company_name": "Template Manager",
	}
	if err := a.email.Send(ctx, email.TemplateIDSignupVerification, vars); err != nil {
		a.logger.ErrorContext(ctx, "failed to send email %+v", err)
		return err
	}

	return nil
}

func (a *App) Logout(ctx context.Context, req shared.LogoutRequest) error {
	return a.sess.Expire(ctx, req.Token)
}
