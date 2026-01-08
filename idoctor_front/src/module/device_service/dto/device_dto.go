package device_dto

import (
	"time"

	"github.com/Mirsadikovv/shared/response"
)

type DeviceCreate struct {
	Name      string `json:"name" validate:"required"`
	BrandName string `json:"brand_name"`
} // @Name DeviceCreate

type DeviceUpdate struct {
	Name      string `json:"name" validate:"required"`
	BrandName string `json:"brand_name"`
} // @Name DeviceUpdate

type DevicePage = response.PageData[Device] // @name DevicePage

type Device struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	BrandName string     `json:"brand_name"`
	CreatedAt *time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
} // @Name Device

type DeviceParams struct {
	Name           *string `query:"name" json:"name"`
	BrandName      *string `query:"brand_name" json:"brand_name"`
	IncludeDeleted *bool   `query:"include_deleted" json:"include_deleted"`
	OnlyDeleted    *bool   `query:"only_deleted" json:"only_deleted"`
} // @Name DeviceParams
