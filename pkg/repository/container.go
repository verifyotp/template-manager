package repository

import (
	"template-manager/internal/entity"
	"template-manager/pkg/database"
)

type Container struct {
	AuthRepository     AccountRepositoryInterface[entity.Account]
	KeyRepository      KeyRepositoryInterface[entity.Key]
	TemplateRepository TemplateRepositoryInterface[entity.Template]
}

func NewRepositoryContainer(db *database.PostgresClient) *Container {
	return &Container{
		AuthRepository:     NewRepository[entity.Account](db.Client.Table(entity.Account{}.TableName())),
		KeyRepository:      NewRepository[entity.Key](db.Client.Table(entity.Key{}.TableName())),
		TemplateRepository: NewRepository[entity.Template](db.Client.Table(entity.Template{}.TableName())),
	}
}
