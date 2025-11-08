package device_service

import (
	"context"

	device_dto "github.com/Mirsadikovv/idoctor_platform/src/module/device_service/dto"
	device_model "github.com/Mirsadikovv/idoctor_platform/src/module/device_service/model"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"gorm.io/gorm"
)

type DeviceService interface {
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*device_dto.DevicePage, error)
	Find(ctx context.Context, filter pg.Filter) ([]device_dto.Device, error)
	FindOne(ctx context.Context, filter pg.Filter) (*device_dto.Device, error)
	Create(ctx context.Context, deviceDto *device_dto.DeviceCreate) (int64, error)
	Update(ctx context.Context, id int64, deviceDto *device_dto.DeviceUpdate) error
	DeleteOrRestore(ctx context.Context, id int64) error
}

type deviceService struct {
	db *gorm.DB
}

func NewDeviceService(db *gorm.DB) DeviceService {
	return &deviceService{
		db: db,
	}
}

func (s *deviceService) Find(ctx context.Context, filter pg.Filter) ([]device_dto.Device, error) {
	return pg.FindWithScan[device_model.Device, device_dto.Device](s.db, filter)
}

func (s *deviceService) FindOne(ctx context.Context, filter pg.Filter) (*device_dto.Device, error) {
	return pg.FindOneWithScan[device_model.Device, device_dto.Device](s.db, filter)
}

func (s *deviceService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*device_dto.DevicePage, error) {
	return pg.PageWithScan[device_model.Device, device_dto.Device](s.db.WithContext(ctx), paginate, filter)
}

func (s *deviceService) Create(ctx context.Context, deviceDto *device_dto.DeviceCreate) (int64, error) {
	deviceModel := &device_model.Device{
		Name:      deviceDto.Name,
		BrandName: deviceDto.BrandName,
	}

	if err := pg.Create(s.db.WithContext(ctx), deviceModel, "id"); err != nil {
		return 0, err
	}

	return deviceModel.ID, nil
}

func (s *deviceService) Update(ctx context.Context, id int64, deviceDto *device_dto.DeviceUpdate) error {
	deviceModel := &device_model.Device{
		Name:      deviceDto.Name,
		BrandName: deviceDto.BrandName,
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if _, err := pg.Update[device_model.Device](s.db.WithContext(ctx), deviceModel, filter, "id"); err != nil {
		return err
	}

	return nil
}

func (s *deviceService) DeleteOrRestore(ctx context.Context, id int64) error {
	var device device_model.Device

	// Check if device exists (including soft deleted)
	if err := s.db.WithContext(ctx).Unscoped().Where("id = ?", id).First(&device).Error; err != nil {
		return err
	}

	// If deleted, restore; otherwise, soft delete
	if device.DeletedAt != nil && (*device.DeletedAt).Valid {
		// Restore
		return s.db.WithContext(ctx).Model(&device_model.Device{}).Where("id = ?", id).Update("deleted_at", nil).Error
	} else {
		// Soft delete
		return s.db.WithContext(ctx).Where("id = ?", id).Delete(&device_model.Device{}).Error
	}
}
