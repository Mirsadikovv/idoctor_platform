package part_dto

import (
	"time"

	"github.com/Mirsadikovv/shared/response"
)

type PartCreate struct {
	Name       string `json:"name" validate:"required"`
	DeviceID   int64  `json:"device_id" validate:"required"`
	SupplierID int64  `json:"supplier_id" validate:"required"`
} // @Name PartCreate

type PartUpdate struct {
	Name       string `json:"name" validate:"required"`
	DeviceID   int64  `json:"device_id" validate:"required"`
	SupplierID int64  `json:"supplier_id" validate:"required"`
} // @Name PartUpdate

type PartPage = response.PageData[Part] // @name PartPage

type Part struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	DeviceID   int64      `json:"device_id"`
	SupplierID int64      `json:"supplier_id"`
	CreatedAt  *time.Time `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
} // @Name Part

type PartParams struct {
	Name           *string `query:"name" json:"name"`
	DeviceID       *int64  `query:"device_id" json:"device_id"`
	SupplierID     *int64  `query:"supplier_id" json:"supplier_id"`
	IncludeDeleted *bool   `query:"include_deleted" json:"include_deleted"`
	OnlyDeleted    *bool   `query:"only_deleted" json:"only_deleted"`
} // @Name PartParams
