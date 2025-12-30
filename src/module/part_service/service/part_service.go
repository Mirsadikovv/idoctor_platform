package part_service

import (
	"context"
	"time"

	device_model "github.com/Mirsadikovv/idoctor_platform/src/module/device_service/model"
	part_dto "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/dto"
	part_model "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/model"
	supplier_model "github.com/Mirsadikovv/idoctor_platform/src/module/supplier_service/model"
	user_model "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/model"

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
	var parts []part_dto.Part

	tx := s.db.WithContext(ctx)
	if filter != nil {
		tx = filter(tx)
	}

	if err := tx.Preload("Device").Preload("Supplier").Preload("Master").Find(&parts).Error; err != nil {
		return nil, err
	}

	return parts, nil
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

// func convertPartToDTO(part *part_model.Part) part_dto.Part {
// 	return part_dto.Part{
// 		Id:          part.Id,
// 		Name:        part.Name,
// 		DeviceId:    part.DeviceId,
// 		SupplierId:  part.SupplierId,
// 		MasterId:    part.MasterId,
// 		Device:      convertDeviceToInfo(part.Device),
// 		Supplier:    convertSupplierToInfo(part.Supplier),
// 		Master:      convertMasterToInfo(part.Master),
// 		IncomePrice: part.IncomePrice,
// 		Price:       part.Price,
// 		CreatedAt:   part.CreatedAt,
// 		DeletedAt:   convertDeletedAt(part.DeletedAt),
// 	}
// }

func convertDeviceToInfo(device *device_model.Device) *part_dto.DeviceInfo {
	if device == nil {
		return nil
	}
	return &part_dto.DeviceInfo{
		Id:        device.ID,
		Name:      device.Name,
		BrandName: device.BrandName,
	}
}

func convertSupplierToInfo(supplier *supplier_model.Supplier) *part_dto.SupplierInfo {
	if supplier == nil {
		return nil
	}
	return &part_dto.SupplierInfo{
		Id:   supplier.ID,
		Name: supplier.Name,
	}
}

func convertMasterToInfo(master *user_model.User) *part_dto.MasterInfo {
	if master == nil {
		return nil
	}
	return &part_dto.MasterInfo{
		Id:          master.Id,
		FirstName:   master.FirstName,
		LastName:    master.LastName,
		MiddleName:  master.MiddleName,
		Username:    master.Username,
		PhoneNumber: master.PhoneNumber,
	}
}

func convertDeletedAt(deletedAt *gorm.DeletedAt) *time.Time {
	if deletedAt != nil && deletedAt.Valid {
		return &deletedAt.Time
	}
	return nil
}
