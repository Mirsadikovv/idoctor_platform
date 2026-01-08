package order_part_dto

import (
	"encoding/json"
	"time"

	"github.com/Mirsadikovv/shared/response"
)

type OrderPartCreate struct {
	OrderId     int64   `json:"order_id" binding:"required"`
	PartId      int64   `json:"part_id" binding:"required"`
	SupplierId  int64   `json:"supplier_id" binding:"required"`
	IncomePrice float64 `json:"income_price" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

type OrderPartUpdate struct {
	SupplierId  *int64   `json:"supplier_id,omitempty"`
	IncomePrice *float64 `json:"income_price,omitempty"`
	Price       *float64 `json:"price,omitempty"`
}

type OrderPartPage = response.PageData[OrderPartResponse] // @name OrderPartPage

type OrderPartParams struct {
	OrderId        *int64 `query:"order_id" json:"order_id"`
	PartId         *int64 `query:"part_id" json:"part_id"`
	SupplierId     *int64 `query:"supplier_id" json:"supplier_id"`
	IncludeDeleted *bool  `query:"include_deleted" json:"include_deleted"`
	OnlyDeleted    *bool  `query:"only_deleted" json:"only_deleted"`
} // @Name OrderPartParams

type OrderPartResponse struct {
	Id          int64            `json:"id"`
	OrderId     int64            `json:"order_id"`
	PartId      int64            `json:"part_id"`
	SupplierId  int64            `json:"supplier_id"`
	IncomePrice float64          `json:"income_price"`
	Price       float64          `json:"price"`
	CreatedAt   *time.Time       `json:"created_at"`
	Order       json.RawMessage  `json:"order,omitempty"`
	Part        json.RawMessage  `json:"part,omitempty"`
	Supplier    json.RawMessage  `json:"supplier,omitempty"`
}
