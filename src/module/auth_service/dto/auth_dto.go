package auth_dto

import (
	"fmt"

	user_model "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/model"

	"github.com/Mirsadikovv/shared/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthUser struct {
	Id     int64 `json:"id"`
	RoleId int64 `json:"roleId"`
}

func (u *AuthUser) ID() int64 {
	return u.Id
}

func (u *AuthUser) Pre(ctx echo.Context, db *gorm.DB, _ ...struct{}) (bool, error) {

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Joins("INNER JOIN roles ON roles.id = users.role_id").
			Where("users.id = ?", u.Id).
			Where("users.blocked_at IS NULL").
			Where(fmt.Sprintf(`roles.permissions -> '%s' ? '%s'`, ctx.Path(), ctx.Request().Method)).
			Limit(1)
	}

	var user user_model.User

	result := db.Table("users").
		Scopes(filter).Select(
		"users.id",
		"users.username",
		"users.role_id",
		"users.last_visit",
		"users.created_at",
		"users.blocked_at",
	).Scan(&user)

	if err := result.Error; err != nil {
		return true, err
	}

	if result.RowsAffected == 0 {
		return true, gorm.ErrRecordNotFound
	}

	req := request.RequestWithData[user_model.User](ctx)

	req.SetUser(&user)

	return false, nil
}

type SingUp struct {
	Username    string  `json:"username" validate:"required"`
	Password    string  `json:"password" validate:"required"`
	FirstName   *string `json:"firstName,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	MiddleName  *string `json:"middleName,omitempty"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	Gender      *string `json:"gender,omitempty"`
	RoleId      int64   `json:"roleId" validate:"required"`
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type TelegramRole struct {
	Role string `json:"role"`
}

type UserRole struct {
	Role string `gorm:"column:role"`
}
