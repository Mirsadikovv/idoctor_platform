package supplier_service

import (
	"context"

	supplier_dto "github.com/Mirsadikovv/idoctor_platform/src/module/supplier_service/dto"
	supplier_model "github.com/Mirsadikovv/idoctor_platform/src/module/supplier_service/model"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"gorm.io/gorm"
)

type SupplierService interface {
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*supplier_dto.SupplierPage, error)
	Find(ctx context.Context, filter pg.Filter) ([]supplier_dto.Supplier, error)
	FindOne(ctx context.Context, filter pg.Filter) (*supplier_dto.Supplier, error)
	Create(ctx context.Context, supplierDto *supplier_dto.SupplierCreate) (int64, error)
	Update(ctx context.Context, id int64, supplierDto *supplier_dto.SupplierUpdate) error
	DeleteOrRestore(ctx context.Context, id int64) error
}

type supplierService struct {
	db *gorm.DB
}

func NewSupplierService(db *gorm.DB) SupplierService {
	return &supplierService{
		db: db,
	}
}

func (s *supplierService) Find(ctx context.Context, filter pg.Filter) ([]supplier_dto.Supplier, error) {
	return pg.FindWithScan[supplier_model.Supplier, supplier_dto.Supplier](s.db, filter)
}

func (s *supplierService) FindOne(ctx context.Context, filter pg.Filter) (*supplier_dto.Supplier, error) {
	return pg.FindOneWithScan[supplier_model.Supplier, supplier_dto.Supplier](s.db, filter)
}

func (s *supplierService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*supplier_dto.SupplierPage, error) {
	return pg.PageWithScan[supplier_model.Supplier, supplier_dto.Supplier](s.db.WithContext(ctx), paginate, filter)
}

func (s *supplierService) Create(ctx context.Context, supplierDto *supplier_dto.SupplierCreate) (int64, error) {
	supplierModel := &supplier_model.Supplier{
		Name: supplierDto.Name,
	}

	if err := pg.Create(s.db.WithContext(ctx), supplierModel, "id"); err != nil {
		return 0, err
	}

	return supplierModel.ID, nil
}

func (s *supplierService) Update(ctx context.Context, id int64, supplierDto *supplier_dto.SupplierUpdate) error {
	supplierModel := &supplier_model.Supplier{
		Name: supplierDto.Name,
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if _, err := pg.Update[supplier_model.Supplier](s.db.WithContext(ctx), supplierModel, filter, "id"); err != nil {
		return err
	}

	return nil
}

func (s *supplierService) DeleteOrRestore(ctx context.Context, id int64) error {
	var supplier supplier_model.Supplier

	// Check if supplier exists (including soft deleted)
	if err := s.db.WithContext(ctx).Unscoped().Where("id = ?", id).First(&supplier).Error; err != nil {
		return err
	}

	// If deleted, restore; otherwise, soft delete
	if supplier.DeletedAt.Valid {
		// Restore
		return s.db.WithContext(ctx).Model(&supplier_model.Supplier{}).Where("id = ?", id).Update("deleted_at", nil).Error
	} else {
		// Soft delete
		return s.db.WithContext(ctx).Where("id = ?", id).Delete(&supplier_model.Supplier{}).Error
	}
}
