package order_dto

import (
	"time"

	"github.com/Mirsadikovv/shared/response"
)

type OrderCreate struct {
	ClientID      *int64  `json:"client_id"`
	MasterID      *int64  `json:"master_id"`
	Price         float64 `json:"price" validate:"required,min=0"`
	Status        string  `json:"status" validate:"required"`
	PaymentType   string  `json:"payment_type" validate:"required"`
	PaymentStatus string  `json:"payment_status" validate:"required"`
	PartIDs       []int64 `json:"part_ids"`
	ProblemIDs    []int64 `json:"problem_ids"`
} // @Name OrderCreate

type OrderUpdate struct {
	ClientID      *int64  `json:"client_id"`
	MasterID      *int64  `json:"master_id"`
	Price         float64 `json:"price" validate:"required,min=0"`
	Status        string  `json:"status" validate:"required"`
	PaymentType   string  `json:"payment_type" validate:"required"`
	PaymentStatus string  `json:"payment_status" validate:"required"`
	PartIDs       []int64 `json:"part_ids"`
	ProblemIDs    []int64 `json:"problem_ids"`
} // @Name OrderUpdate

type OrderPage = response.PageData[Order] // @name OrderPage

type Order struct {
	ID            int64      `json:"id"`
	ClientID      *int64     `json:"client_id"`
	MasterID      *int64     `json:"master_id"`
	Price         float64    `json:"price"`
	Status        string     `json:"status"`
	PaymentType   string     `json:"payment_type"`
	PaymentStatus string     `json:"payment_status"`
	CreatedAt     *time.Time `json:"created_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	PartIDs       []int64    `json:"part_ids,omitempty"`
	ProblemIDs    []int64    `json:"problem_ids,omitempty"`
} // @Name Order

type OrderParams struct {
	ClientID       *int64   `query:"client_id" json:"client_id"`
	MasterID       *int64   `query:"master_id" json:"master_id"`
	Status         *string  `query:"status" json:"status"`
	PaymentType    *string  `query:"payment_type" json:"payment_type"`
	PaymentStatus  *string  `query:"payment_status" json:"payment_status"`
	MinPrice       *float64 `query:"min_price" json:"min_price"`
	MaxPrice       *float64 `query:"max_price" json:"max_price"`
	IncludeDeleted *bool    `query:"include_deleted" json:"include_deleted"`
	OnlyDeleted    *bool    `query:"only_deleted" json:"only_deleted"`
	StartDate      *string  `query:"start_date" json:"start_date"`
	EndDate        *string  `query:"end_date" json:"end_date"`
} // @Name OrderParams
