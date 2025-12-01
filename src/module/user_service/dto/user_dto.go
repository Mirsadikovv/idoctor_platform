package user_dto

import (
	"context"
	"errors"
	"time"

	"github.com/Mirsadikovv/idoctor_platform/src/common/utils"
	"github.com/Mirsadikovv/shared/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	UserPage = response.PageData[User]
	Password string
	Gender   string
)

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (g Gender) String() string {
	return string(g)
}

func (g Gender) Valid() bool {
	return g == Male || g == Female
}

func (g *Gender) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return errors.New("empty gender")
	}

	switch string(b) {
	case `"male"`:
		*g = Male
	case `"female"`:
		*g = Female
	case `"1"`:
		*g = Male
	case `"2"`:
		*g = Female
	default:
		return errors.New("invalid gender")
	}

	return nil
}

func (p Password) GormValue(context.Context, *gorm.DB) clause.Expr {
	if p == "" {
		return gorm.Expr("password")
	}

	return gorm.Expr("HASH_MAKE(?)", string(p))
}

type User struct {
	Id               int64      `json:"id"`
	Username         string     `json:"username"`
	FirstName        string     `json:"firstName"`
	LastName         string     `json:"lastName"`
	MiddleName       string     `json:"middleName"`
	DateOfBirth      string     `json:"dateOfBirth"`
	Gender           Gender     `json:"gender"`
	RoleId           int64      `json:"roleId"`
	TelegramId       *int64     `json:"telegramId,omitempty"`
	TelegramUsername *string    `json:"telegramUsername,omitempty"`
	PhoneNumber      *string    `json:"phoneNumber,omitempty"`
	LanguageCode     *string    `json:"languageCode,omitempty"`
	LastVisit        *time.Time `json:"lastVisit"`
} // @name User

type UserParams struct {
	utils.OrderParams
	Fullname *string `json:"fullname" query:"fullname"`
	Gender   Gender  `json:"gender" query:"gender"`
} // @name UserParams

type UserCreate struct {
	Username         string  `json:"username" validate:"required"`
	Password         string  `json:"password" validate:"required"`
	FirstName        *string `json:"firstName,omitempty"`
	LastName         *string `json:"lastName,omitempty"`
	MiddleName       *string `json:"middleName,omitempty"`
	DateOfBirth      *string `json:"dateOfBirth,omitempty"`
	Gender           *string `json:"gender,omitempty"`
	RoleId           int64   `json:"roleId" validate:"required"`
	TelegramId       *int64  `json:"telegramId,omitempty"`
	TelegramUsername *string `json:"telegramUsername,omitempty"`
	PhoneNumber      *string `json:"phoneNumber,omitempty"`
	LanguageCode     *string `json:"languageCode,omitempty"`
}
