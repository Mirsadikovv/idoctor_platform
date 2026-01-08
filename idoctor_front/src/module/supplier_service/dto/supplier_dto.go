package supplier_dto

import (
	supplier_model "github.com/Mirsadikovv/idoctor_platform/src/module/supplier_service/model"
	"time"

	"github.com/Mirsadikovv/shared/response"
)

type SupplierCreate struct {
	Name string `json:"name" validate:"required"`
} // @Name SupplierCreate

type SupplierUpdate struct {
	Name string `json:"name" validate:"required"`
} // @Name SupplierUpdate

type SupplierPage = response.PageData[Supplier] // @name SupplierPage

type Supplier struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
} // @Name Supplier

type SupplierParams struct {
	Name           *string `query:"name" json:"name"`
	IncludeDeleted *bool   `query:"include_deleted" json:"include_deleted"`
	OnlyDeleted    *bool   `query:"only_deleted" json:"only_deleted"`
} // @Name SupplierParams

func ToSupplierDto(supplier *supplier_model.Supplier) *Supplier {
	if supplier == nil {
		return nil
	}

	var deletedAt *time.Time
	if supplier.DeletedAt != nil && supplier.DeletedAt.Valid {
		deletedAt = &supplier.DeletedAt.Time
	}

	var createdAt time.Time
	if supplier.CreatedAt != nil {
		createdAt = *supplier.CreatedAt
	}

	return &Supplier{
		ID:        supplier.ID,
		Name:      supplier.Name,
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
	}
}
