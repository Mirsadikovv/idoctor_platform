package part_dto

import (
	"time"

	"github.com/Mirsadikovv/shared/response"
)

type DeviceInfo struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	BrandName string `json:"brand_name"`
} // @Name DeviceInfo

type SupplierInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
} // @Name SupplierInfo

type MasterInfo struct {
	Id         int64   `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName string  `json:"middle_name"`
	Username   string  `json:"username"`
	PhoneNumber *string `json:"phone_number,omitempty"`
} // @Name MasterInfo

type PartCreate struct {
	Name       string `json:"name" validate:"required"`
	DeviceId   int64  `json:"device_id" validate:"required"`
	SupplierId int64  `json:"supplier_id" validate:"required"`
} // @Name PartCreate

type PartUpdate struct {
	Name       string `json:"name" validate:"required"`
	DeviceId   int64  `json:"device_id" validate:"required"`
	SupplierId int64  `json:"supplier_id" validate:"required"`
} // @Name PartUpdate

type PartPage = response.PageData[Part] // @name PartPage

type Part struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	DeviceId    int64         `json:"device_id"`
	SupplierId  int64         `json:"supplier_id"`
	MasterId    int64         `json:"master_id"`
	Device      *DeviceInfo   `json:"device,omitempty"`
	Supplier    *SupplierInfo `json:"supplier,omitempty"`
	Master      *MasterInfo   `json:"master,omitempty"`
	IncomePrice float64       `json:"income_price"`
	Price       float64       `json:"price"`
	CreatedAt   *time.Time    `json:"created_at"`
	DeletedAt   *time.Time    `json:"deleted_at,omitempty"`
} // @Name Part

type PartParams struct {
	Name           *string `query:"name" json:"name"`
	DeviceId       *int64  `query:"device_id" json:"device_id"`
	SupplierId     *int64  `query:"supplier_id" json:"supplier_id"`
	IncludeDeleted *bool   `query:"include_deleted" json:"include_deleted"`
	OnlyDeleted    *bool   `query:"only_deleted" json:"only_deleted"`
} // @Name PartParams
