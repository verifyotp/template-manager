package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

func (r *Repository[T]) Create(ctx context.Context, t *T) (*T, error) {
	if err := r.db.WithContext(ctx).Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (r *Repository[T]) Find(ctx context.Context, conds ...interface{}) ([]*T, error) {
	var (
		dest []*T
	)

	if err := r.db.WithContext(ctx).Find(dest, conds...).Error; err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *Repository[T]) Get(ctx context.Context, conds ...interface{}) (*T, error) {
	var (
		dest *T
	)
	if err := r.db.WithContext(ctx).First(dest, conds...).Error; err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *Repository[T]) Update(ctx context.Context, column string, value interface{}) (*T, error) {
	var (
		dest *T
	)
	if err := r.db.WithContext(ctx).Update(column, value).Find(dest).Error; err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *Repository[T]) Delete(ctx context.Context, t *T) error {
	if err := r.db.WithContext(ctx).Delete(t).Error; err != nil {
		return err
	}
	return nil
}
