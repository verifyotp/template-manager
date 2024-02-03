package repository

import (
	"template-manager/internal/entity"
	"template-manager/pkg/database"
)

type Container struct {
	AuthRepository AccountRepositoryInterface[entity.Account]
	KeyRepository  KeyRepositoryInterface[entity.Key]
}

func NewRepositoryContainer(db *database.PostgresClient) *Container {
	return &Container{
		AuthRepository: NewRepository[entity.Account](db.Client.Table("accounts")),
		KeyRepository:  NewRepository[entity.Key](db.Client.Table("keys")),
	}
}
