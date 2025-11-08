package supplier_dto

import (
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
	Name              *string `query:"name" json:"name"`
	IncludeDeleted    *bool   `query:"include_deleted" json:"include_deleted"`
	OnlyDeleted       *bool   `query:"only_deleted" json:"only_deleted"`
} // @Name SupplierParams
