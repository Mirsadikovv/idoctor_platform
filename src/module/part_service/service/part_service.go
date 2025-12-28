package part_service

import (
	"context"

	part_dto "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/dto"
	part_model "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/model"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"gorm.io/gorm"
)

type PartService interface {
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*part_dto.PartPage, error)
	Find(ctx context.Context, filter pg.Filter) ([]part_dto.Part, error)
	FindOne(ctx context.Context, filter pg.Filter) (*part_dto.Part, error)
	Create(ctx context.Context, partDto *part_dto.PartCreate) (int64, error)
	Update(ctx context.Context, id int64, partDto *part_dto.PartUpdate) error
	DeleteOrRestore(ctx context.Context, id int64) error
}

type partService struct {
	db *gorm.DB
}

func NewPartService(db *gorm.DB) PartService {
	return &partService{
		db: db,
	}
}

func (s *partService) Find(ctx context.Context, filter pg.Filter) ([]part_dto.Part, error) {
	return pg.FindWithScan[part_model.Part, part_dto.Part](s.db, filter)
}

func (s *partService) FindOne(ctx context.Context, filter pg.Filter) (*part_dto.Part, error) {
	return pg.FindOneWithScan[part_model.Part, part_dto.Part](s.db, filter)
}

func (s *partService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*part_dto.PartPage, error) {
	return pg.PageWithScan[part_model.Part, part_dto.Part](s.db.WithContext(ctx), paginate, filter)
}

func (s *partService) Create(ctx context.Context, partDto *part_dto.PartCreate) (int64, error) {
	partModel := &part_model.Part{
		Name:       partDto.Name,
		DeviceId:   partDto.DeviceId,
		SupplierId: partDto.SupplierId,
	}

	if err := pg.Create(s.db.WithContext(ctx), partModel, "id"); err != nil {
		return 0, err
	}

	return partModel.Id, nil
}

func (s *partService) Update(ctx context.Context, id int64, partDto *part_dto.PartUpdate) error {
	partModel := &part_model.Part{
		Name:       partDto.Name,
		DeviceId:   partDto.DeviceId,
		SupplierId: partDto.SupplierId,
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if _, err := pg.Update[part_model.Part](s.db.WithContext(ctx), partModel, filter, "id"); err != nil {
		return err
	}

	return nil
}

func (s *partService) DeleteOrRestore(ctx context.Context, id int64) error {
	var part part_model.Part

	// Check if part exists (including soft deleted)
	if err := s.db.WithContext(ctx).Unscoped().Where("id = ?", id).First(&part).Error; err != nil {
		return err
	}

	// If deleted, restore; otherwise, soft delete
	if part.DeletedAt != nil && (*part.DeletedAt).Valid {
		// Restore
		return s.db.WithContext(ctx).Model(&part_model.Part{}).Where("id = ?", id).Update("deleted_at", nil).Error
	} else {
		// Soft delete
		return s.db.WithContext(ctx).Where("id = ?", id).Delete(&part_model.Part{}).Error
	}
}
