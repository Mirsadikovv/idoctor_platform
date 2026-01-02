package order_dto

import (
	"encoding/json"
	"time"

	"github.com/Mirsadikovv/shared/response"
)

type UserInfo struct {
	Id         int64   `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName string  `json:"middle_name"`
	Username   string  `json:"username"`
	PhoneNumber *string `json:"phone_number,omitempty"`
} // @Name UserInfo

type OrderCreate struct {
	ClientId      *int64  `json:"client_id"`
	ClientPhone   string  `json:"client_phone"`
	ClientName    string  `json:"client_name"`
	PhonePassword string  `json:"phone_password"`
	MasterId      *int64  `json:"master_id"`
	Price         float64 `json:"price" validate:"required,min=0"`
	Status        string  `json:"status" validate:"required"`
	PaymentType   string  `json:"payment_type" validate:"required"`
	PaymentStatus string  `json:"payment_status" validate:"required"`
	PartIds       []int64 `json:"part_ids"`
	ProblemIds    []int64 `json:"problem_ids"`
} // @Name OrderCreate

type OrderUpdate struct {
	ClientId      *int64  `json:"client_id"`
	ClientPhone   string  `json:"client_phone"`
	ClientName    string  `json:"client_name"`
	PhonePassword string  `json:"phone_password"`
	MasterId      *int64  `json:"master_id"`
	Price         float64 `json:"price" validate:"required,min=0"`
	Status        string  `json:"status" validate:"required"`
	PaymentType   string  `json:"payment_type" validate:"required"`
	PaymentStatus string  `json:"payment_status" validate:"required"`
	PartIds       []int64 `json:"part_ids"`
	ProblemIds    []int64 `json:"problem_ids"`
} // @Name OrderUpdate

type OrderPage = response.PageData[Order] // @name OrderPage

type Order struct {
	Id            int64            `json:"id"`
	ClientId      *int64           `json:"client_id"`
	ClientPhone   string           `json:"client_phone"`
	ClientName    string           `json:"client_name"`
	PhonePassword string           `json:"phone_password"`
	MasterId      *int64           `json:"master_id"`
	Client        *json.RawMessage `json:"client,omitempty"`
	Master        *json.RawMessage `json:"master,omitempty"`
	Price         float64          `json:"price"`
	Status        string           `json:"status"`
	PaymentType   string           `json:"payment_type"`
	PaymentStatus string           `json:"payment_status"`
	CreatedAt     *time.Time       `json:"created_at"`
	DeletedAt     *time.Time       `json:"deleted_at,omitempty"`
	PartIds       []int64          `json:"part_ids,omitempty"`
	ProblemIds    []int64          `json:"problem_ids,omitempty"`
} // @Name Order

type OrderParams struct {
	ClientId       *int64   `query:"client_id" json:"client_id"`
	MasterId       *int64   `query:"master_id" json:"master_id"`
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
