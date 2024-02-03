package database

import "context"

type Operations[T any] interface {
	Create(ctx context.Context, object *T) error
	Update(ctx context.Context, object *T) error
	Get(ctx context.Context, object T) (*T, error)
	Find(ctx context.Context, object T) ([]*T, error)
	Delete(ctx context.Context, object T) error
}
