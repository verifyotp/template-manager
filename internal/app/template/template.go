package template

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"template-manager/internal/entity"
	"template-manager/internal/shared"
	"template-manager/pkg/config"
	"template-manager/pkg/repository"
	"template-manager/pkg/repository/util"
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

func (a *App) GetUploadURL(ctx context.Context, req shared.GetUploadURLRequest) (*shared.UploadURLResponse, error) {
	return &shared.UploadURLResponse{
		URL:         "https://s3.amazonaws.com/your-bucket-name/" + shared.GenerateSlug(req.Name),
		ExpireAt:    time.Now().Add(time.Hour),
		AccountID:   req.AccountID,
		ContentType: req.ContentType,
	}, nil
}

func (a *App) Create(ctx context.Context, req shared.CreateTemplateRequest) error {
	var template = entity.Template{
		AccountID:   req.AccountID,
		Name:        req.Name,
		Slug:        shared.GenerateSlug(req.Name),
		Version:     1,
		ContentType: req.ContentType,
		Location:    req.Location,
		Vars:        req.Vars,
		Active:      true,
	}
	return a.db.TemplateRepository.Create(ctx, &template)
}

func (a *App) Update(ctx context.Context, req shared.UpdateTemplateRequest) error {
	existing, err := a.db.TemplateRepository.Get(ctx, "id = ? AND account_id = ?", req.TemplateID, req.AccountID)
	if err != nil {
		return err
	}
	newVersion := existing.Version + 1
	return a.db.TemplateRepository.Create(ctx, &entity.Template{
		AccountID:   req.AccountID,
		Name:        fmt.Sprintf("%s-v%d", existing.Name, newVersion),
		Slug:        existing.Slug,
		Version:     newVersion,
		ContentType: existing.ContentType,
		Location:    req.Location,
		Vars:        req.Vars,
		Active:      existing.Active,
	})
}

func (a *App) Delete(ctx context.Context, req shared.DeleteTemplateRequest) error {
	if err := a.db.TemplateRepository.Delete(ctx, &entity.Template{
		ID:        req.TemplateID,
		AccountID: req.AccountID,
		Version:   req.Version,
	}); err != nil {
		a.logger.ErrorContext(ctx, "failed to create account %+v", err)
		return err
	}
	return nil
}

func (a *App) Get(ctx context.Context, req shared.GetTemplateRequest) (*entity.Template, error) {
	return a.db.TemplateRepository.Get(ctx, "id = ? AND account_id = ?", req.TemplateID, req.AccountID)
}

func (a *App) List(ctx context.Context, req shared.ListTemplatesRequest) (*util.PaginationT[[]entity.Template], error) {
	return a.db.TemplateRepository.FindWithPagination(
		ctx,
		util.Eq("account_id", req.AccountID),
		repository.WithPagination(req.Page, req.PageSize),
	)
}
