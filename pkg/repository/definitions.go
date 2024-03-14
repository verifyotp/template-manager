package repository

import (
	"context"

	"template-manager/internal/entity"
	"template-manager/pkg/repository/util"
)

type AccountRepositoryInterface[T entity.Account] interface {
	Create(ctx context.Context, t *T) error
	Find(ctx context.Context, conds ...interface{}) ([]T, error)
	Get(ctx context.Context, conds ...interface{}) (*T, error)
	Update(ctx context.Context, E *T) error
}

type KeyRepositoryInterface[T entity.Key] interface {
	Create(ctx context.Context, t *T) error
	Find(ctx context.Context, conds ...interface{}) ([]T, error)
	FindManyWithOptions(ctx context.Context, query any, opts ...Opt) ([]T, error)
	Delete(ctx context.Context, t *T) error
}

type TemplateRepositoryInterface[T entity.Template] interface {
	Create(ctx context.Context, t *T) error
	Find(ctx context.Context, conds ...interface{}) ([]T, error)
	Get(ctx context.Context, conds ...interface{}) (*T, error)
	Update(ctx context.Context, E *T) error
	Delete(ctx context.Context, t *T) error
	FindWithPagination(ctx context.Context, query any, opts ...Opt) (*util.PaginationT[[]T], error)
}

type CredentialRepositoryInterface[T entity.Credential] interface {
	Create(ctx context.Context, t *T) error
	Find(ctx context.Context, conds ...interface{}) ([]T, error)
	Get(ctx context.Context, conds ...interface{}) (*T, error)
	Update(ctx context.Context, E *T) error
	Delete(ctx context.Context, t *T) error
	FindWithPagination(ctx context.Context, query any, opts ...Opt) (*util.PaginationT[[]T], error)
}
