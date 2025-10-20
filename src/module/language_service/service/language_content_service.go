package language_service

import (
	"context"
	"fmt"

	language_dto "git.sriss.uz/mehnat/inspector_platform/src/module/language_service/dto"
	language_model "git.sriss.uz/mehnat/inspector_platform/src/module/language_service/model"
	log_service "git.sriss.uz/mehnat/inspector_platform/src/module/log_service/service"
	"github.com/labstack/echo/v4"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"git.sriss.uz/shared/shared_service/sharedutil"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LanguageContentService interface {
	Find(ctx context.Context, filter pg.Filter) ([]language_dto.LanguageContent, error)
	FindOne(ctx context.Context, filter pg.Filter) (*language_dto.LanguageContent, error)
	FindByKey(ctx context.Context, filter pg.Filter) (*language_dto.LanguageContentByKeyResp, error)
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*language_dto.LanguageContentPage, error)
	Create(ctx context.Context, languageDto *language_dto.LanguageContentCreate) (int64, error)
	CreateOrUpdate(ctx echo.Context, languageDto *language_dto.LanguageContentsCreateOrUpdate) (*language_dto.LanguageResult, error)
	FindByLanguage(ctx context.Context, filter pg.Filter) (sharedutil.JsonObject, error)
	PageByLanguage(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*language_dto.ContentPage, error)
	SearchByLanguage(ctx context.Context, filter pg.Filter) ([]language_dto.ContentDto, error)
	Update(ctx context.Context, languageDto *language_dto.LanguageContentUpdate, filter pg.Filter) error
	Delete(ctx context.Context, filter pg.Filter) error
}

type languageContentService struct {
	db *gorm.DB
}

func NewLanguageContentService(db *gorm.DB) LanguageContentService {
	return &languageContentService{
		db: db,
	}
}

func (l *languageContentService) Find(ctx context.Context, filter pg.Filter) ([]language_dto.LanguageContent, error) {
	contents, err := pg.FindWithScan[language_model.LanguageContent, language_dto.LanguageContent](l.db, filter)
	{
		if err != nil {
			return nil, err
		}
	}

	if contents == nil {
		contents = []language_dto.LanguageContent{}
	}

	return contents, nil
}

func (l *languageContentService) FindOne(ctx context.Context, filter pg.Filter) (*language_dto.LanguageContent, error) {
	return pg.FindOneWithScan[language_model.LanguageContent, language_dto.LanguageContent](l.db, filter)
}

func (l *languageContentService) FindByKey(ctx context.Context, filter pg.Filter) (*language_dto.LanguageContentByKeyResp, error) {
	return pg.FindOneWithScan[language_model.LanguageContent, language_dto.LanguageContentByKeyResp](l.db, filter)
}

func (l *languageContentService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*language_dto.LanguageContentPage, error) {
	return pg.PageWithScan[language_model.LanguageContent, language_dto.LanguageContent](l.db, paginate, filter)
}

func (l *languageContentService) PageByLanguage(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*language_dto.ContentPage, error) {
	return pg.PageWithScan[language_model.LanguageContent, language_dto.ContentDto](l.db, paginate, filter)
}

func (l *languageContentService) SearchByLanguage(ctx context.Context, filter pg.Filter) ([]language_dto.ContentDto, error) {
	langContents, err := pg.FindWithScan[language_model.LanguageContent, language_dto.ContentDto](l.db, filter)
	{
		if err != nil {
			return nil, err
		}
	}
	return langContents, nil
}

func (l *languageContentService) FindByLanguage(ctx context.Context, filter pg.Filter) (sharedutil.JsonObject, error) {

	var contents sharedutil.JsonObject
	{
		result := l.db.Model(&language_model.LanguageContent{}).
			Scopes(filter).Scan(&contents)

		if err := result.Error; err != nil {
			return nil, err
		}
	}

	if contents == nil {
		contents = make(sharedutil.JsonObject)
	}

	return contents, nil
}

func (l *languageContentService) CreateOrUpdate(ctx echo.Context, languageDto *language_dto.LanguageContentsCreateOrUpdate) (*language_dto.LanguageResult, error) {

	var contents []*language_model.LanguageContent
	{
		for _, content := range languageDto.Contents {
			contents = append(contents, &language_model.LanguageContent{
				Key:        content.Key,
				Value:      content.Value,
				Category:   languageDto.Category,
				LanguageId: content.LanguageId,
			})
		}
	}

	// 1. Oldindan mavjud recordlarni olish
	var existingRecords []*language_model.LanguageContent
	oldDataMap := make(map[string]map[string]any) // key: "key_category_langId", value: old data

	{
		keys := make([]string, len(contents))
		for i, c := range contents {
			keys[i] = c.Key
		}

		err := l.db.Where("key IN ? AND category = ? AND language_id = ?", keys, languageDto.Category, languageDto.LanguageId).
			Find(&existingRecords).Error
		if err != nil {
			return nil, err
		}

		// Old data map ni to'ldirish
		for _, record := range existingRecords {
			compositeKey := fmt.Sprintf("%s_%s_%d", record.Key, record.Category, record.LanguageId)
			oldDataMap[compositeKey] = map[string]any{
				"key":         record.Key,
				"value":       record.Value,
				"category":    record.Category,
				"language_id": record.LanguageId,
			}
		}
	}

	// 2. Create or Update
	clauseOnConflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "key"}, {Name: "category"}, {Name: "language_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"value", // faqat value update bo'ladi
		}),
	}

	result := l.db.WithContext(ctx.Request().Context()).Clauses(clauseOnConflict).Create(contents)
	{
		if err := result.Error; err != nil {
			return nil, err
		}
	}

	// 3. Log qilish
	for _, c := range contents {
		compositeKey := fmt.Sprintf("%s_%s_%d", c.Key, c.Category, c.LanguageId)

		data := map[string]any{
			"id":   c.Id,
			"data": c,
		}

		// Agar old data mavjud bo'lsa - update, aks holda - create
		if oldData, exists := oldDataMap[compositeKey]; exists {
			data["old_data"] = oldData
		}

		if _, errLog := log_service.TableCrud(l.db, ctx, "language_contents", data); errLog != nil {
			fmt.Println("log error:", errLog)
		}
	}

	languageResult := &language_dto.LanguageResult{
		LanguageId: languageDto.LanguageId,
		Category:   languageDto.Category,
	}

	return languageResult, nil
}

func (l *languageContentService) Create(ctx context.Context, languageDto *language_dto.LanguageContentCreate) (int64, error) {
	languageContentModel := &language_model.LanguageContent{
		Key:        languageDto.Key,
		Value:      languageDto.Value,
		Category:   languageDto.Category,
		LanguageId: languageDto.LanguageId,
	}

	if err := pg.Create(l.db.WithContext(ctx), languageContentModel, "id"); err != nil {
		return 0, err
	}

	return languageContentModel.Id, nil
}

func (l *languageContentService) Update(ctx context.Context, languageDto *language_dto.LanguageContentUpdate, filter pg.Filter) error {

	if _, err := pg.Update[language_model.LanguageContent](l.db.WithContext(ctx), languageDto, filter); err != nil {
		return err
	}

	return nil
}

func (l *languageContentService) Delete(ctx context.Context, filter pg.Filter) error {

	if err := pg.Delete[language_model.LanguageContent](l.db.WithContext(ctx).Debug(), nil, filter); err != nil {
		return err
	}

	return nil
}
