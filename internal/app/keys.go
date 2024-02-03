package app

import (
	"context"
	"errors"

	"template-manager/internal/entity"
	"template-manager/internal/shared"
)

func (a *App) CreateAccessKey(ctx context.Context, req shared.CreateAccessKeyRequest) error {
	var key = entity.Key{
		AccountID: req.AccountID,
		Name:      req.AccessKeyName,
	}

	if err := key.GenerateKey(); err != nil {
		return err
	}
	if _, err := a.db.KeyRepository.Create(ctx, &key); err != nil {
		a.logger.ErrorContext(ctx, "failed to create account %+v", err)
		return err
	}

	return nil
}

func (a *App) ListAccessKeys(ctx context.Context, req shared.ListAccessKeysRequest) ([]*entity.Key, error) {
	var (
		keys []*entity.Key
		err  error
	)
	if keys, err = a.db.KeyRepository.Find(ctx, "id = ?", req.AccountID); err != nil {
		return nil, errors.New("couldn't find matching keys: " + err.Error())
	}
	return keys, nil
}

func (a *App) DeleteAccessKey(ctx context.Context, req shared.DeleteAccessKeyRequest) error {
	var key = entity.Key{
		AccountID: req.AccountID,
		ID:        req.AccessKeyID,
	}
	if err := a.db.KeyRepository.Delete(ctx, &key); err != nil {
		return errors.New("problem deleting key: " + err.Error())
	}

	return nil
}
