package credential

import (
	"context"
	"template-manager/internal/entity"
	"template-manager/internal/shared"
	"template-manager/pkg/encoding"
	"template-manager/pkg/repository"
)

type Credential struct {
	db repository.Container
}

func New(db repository.Container) *Credential {
	return &Credential{
		db: db,
	}
}

func (c *Credential) Create(ctx context.Context, accountID string, input *shared.CredentialInput) error {
	if err := input.ValidateCreate(); err != nil {
		return err
	}
	cred := new(entity.Credential)
	if err := encoding.UnPackJSON(input, cred); err != nil {
		return err
	}
	cred.AccountID = accountID
	return c.db.CredentialRepository.Create(ctx, cred)
}

func (c *Credential) GetByAccountID(ctx context.Context, accountID string) ([]entity.Credential, error) {
	return c.db.CredentialRepository.Find(ctx, "account_id = ?", accountID)
}

func (c *Credential) Update(ctx context.Context, input *shared.CredentialInput) error {
	if err := input.ValidateUpdate(); err != nil {
		return err
	}
	cred := new(entity.Credential)
	if err := encoding.UnPackJSON(input, cred); err != nil {
		return err
	}
	return c.db.CredentialRepository.Update(ctx, cred)
}

func (c *Credential) Delete(ctx context.Context, id string) error {
	cred, err := c.db.CredentialRepository.Get(ctx, "id = ?", id)
	if err != nil {
		return err
	}
	return c.db.CredentialRepository.Delete(ctx, cred)
}

func (c *Credential) GetByID(ctx context.Context, id string) (*entity.Credential, error) {
	return c.db.CredentialRepository.Get(ctx, "id = ?", id)
}
