package language_service

import (
	"context"

	language_dto "git.sriss.uz/mehnat/inspector_platform/src/module/language_service/dto"
	language_model "git.sriss.uz/mehnat/inspector_platform/src/module/language_service/model"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"gorm.io/gorm"
)

type LanguageService interface {
	Find(ctx context.Context, filter pg.Filter) ([]language_dto.Language, error)
	FindOne(ctx context.Context, filter pg.Filter) (*language_dto.Language, error)
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*language_dto.LanguagePage, error)
	Create(ctx context.Context, languageDto *language_dto.LanguageCreate) (int64, error)
	Update(ctx context.Context, languageDto *language_dto.LanguageUpdate, filter pg.Filter) error
	Delete(ctx context.Context, filter pg.Filter) error
	DeleteOrRestore(ctx context.Context, filter pg.Filter) error
}

type languageService struct {
	db *gorm.DB
}

func NewLanguageService(db *gorm.DB) LanguageService {
	return &languageService{
		db: db,
	}
}

func (l *languageService) Find(ctx context.Context, filter pg.Filter) ([]language_dto.Language, error) {
	languages, err := pg.FindWithScan[language_model.Language, language_dto.Language](l.db, filter)
	{
		if err != nil {
			return nil, err
		}
	}

	if languages == nil {
		languages = []language_dto.Language{}
	}

	return languages, nil
}

func (l *languageService) FindOne(ctx context.Context, filter pg.Filter) (*language_dto.Language, error) {
	return pg.FindOneWithScan[language_model.Language, language_dto.Language](l.db, filter)
}

func (l *languageService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*language_dto.LanguagePage, error) {
	return pg.PageWithScan[language_model.Language, language_dto.Language](l.db, paginate, filter)
}

func (l *languageService) Create(ctx context.Context, languageDto *language_dto.LanguageCreate) (int64, error) {
	languageModel := &language_model.Language{
		Name:        languageDto.Name,
		Description: languageDto.Description,
	}

	if err := pg.Create(l.db.WithContext(ctx), languageModel, "id"); err != nil {
		return 0, err
	}

	return languageModel.Id, nil
}

func (l *languageService) Update(ctx context.Context, languageDto *language_dto.LanguageUpdate, filter pg.Filter) error {

	if _, err := pg.Update[language_model.Language](l.db.WithContext(ctx), languageDto, filter); err != nil {
		return err
	}

	return nil
}

func (l *languageService) Delete(ctx context.Context, filter pg.Filter) error {
	if err := pg.Delete[language_model.Language](l.db.WithContext(ctx), nil, filter); err != nil {
		return err
	}

	return nil
}

func (l *languageService) DeleteOrRestore(ctx context.Context, filter pg.Filter) error {

	restoreData := map[string]any{
		"deleted_at": gorm.Expr("CASE WHEN deleted_at IS NULL THEN now() ELSE null END"),
	}

	if _, err := pg.Update[language_model.Language](l.db.WithContext(ctx).Unscoped(), restoreData, filter); err != nil {
		return err
	}

	return nil
}
