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
	Id             int64      `json:"id"`
	Pin            int64      `json:"pin"`
	Username       string     `json:"username"`
	Name           string     `json:"name"`
	Valid          string     `json:"valid"`
	PassportNumber string     `json:"passportnumber"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	MiddleName     string     `json:"middleName"`
	DateOfBirth    string     `json:"dateOfBirth"`
	UserType       string     `json:"userType"`
	Gender         Gender     `json:"gender"`
	Email          string     `json:"email"`
	SessionId      string     `json:"sessionId"`
	Photo          string     `json:"photo"`
	Nationality    string     `json:"nationality"`
	PlaceOfBirth   string     `json:"placeOfBirth"`
	CountryOfBirth string     `json:"countryOfBirth"`
	Citizenship    string     `json:"citizenship"`
	RoleId         int64      `json:"roleId"`
	LastVisit      *time.Time `json:"lastVisit"`
} // @name User

type UserParams struct {
	utils.OrderParams
	Fullname *string `json:"fullname" query:"fullname"`
	Pin      *string `json:"pin" query:"pin"`
	Gender   Gender  `json:"gender" query:"gender"`
} // @name UserParams

type UserCreate struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	RoleId     int64  `json:"roleId"`
	EmployeeId int64  `json:"employeeId"`
}
