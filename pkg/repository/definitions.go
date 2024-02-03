package repository

import (
	"context"

	"template-manager/internal/entity"
)

type AccountRepositoryInterface[T entity.Account] interface {
	Create(ctx context.Context, t *T) (*T, error)
	Find(ctx context.Context, conds ...interface{}) ([]*T, error)
	Get(ctx context.Context, conds ...interface{}) (*T, error)
	Update(ctx context.Context, column string, value interface{}) (*T, error)
}

type KeyRepositoryInterface[T entity.Key] interface {
	Create(ctx context.Context, t *T) (*T, error)
	Find(ctx context.Context, conds ...interface{}) ([]*T, error)
	Delete(ctx context.Context, t *T) error
}
