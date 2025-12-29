package auth_dto

import (
	"time"

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
	// Get the route pattern (e.g., /api/v1/role/:id) instead of actual path
	routePattern := ctx.Path()
	method := ctx.Request().Method

	// First, get user and role information
	var result struct {
		UserId       int64
		Username     string
		RoleId       int64
		RoleName     string
		LastVisit    *time.Time
		CreatedAt    *time.Time
		BlockedAt    *time.Time
		Permissions  map[string][]string
	}

	err := db.Table("users").
		Select(
			"users.id as user_id",
			"users.username",
			"users.role_id",
			"roles.name as role_name",
			"users.last_visit",
			"users.created_at",
			"users.blocked_at",
			"roles.permissions",
		).
		Joins("INNER JOIN roles ON roles.id = users.role_id").
		Where("users.id = ?", u.Id).
		Where("users.blocked_at IS NULL").
		Scan(&result).Error

	if err != nil {
		return true, err
	}

	// Create user object
	user := &user_model.User{
		Id:        result.UserId,
		Username:  result.Username,
		RoleId:    result.RoleId,
		LastVisit: result.LastVisit,
		CreatedAt: result.CreatedAt,
		BlockedAt: result.BlockedAt,
	}

	// If admin role, skip permission check
	if result.RoleName == "admin" {
		req := request.RequestWithData[user_model.User](ctx)
		req.SetUser(user)
		return false, nil
	}

	// For non-admin users, check if route pattern exists in permissions
	// Permissions structure: { "/api/v1/role/:id": ["GET", "POST"], ... }
	if methods, exists := result.Permissions[routePattern]; exists {
		// Check if current method is allowed
		for _, allowedMethod := range methods {
			if allowedMethod == method {
				req := request.RequestWithData[user_model.User](ctx)
				req.SetUser(user)
				return false, nil
			}
		}
	}

	// Permission denied
	return true, gorm.ErrRecordNotFound
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
	Role  string `json:"role"`
	Token string `json:"token"`
}

type UserRole struct {
	Role string `gorm:"column:role"`
}
