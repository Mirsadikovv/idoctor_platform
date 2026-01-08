package auth_middleware

import (
	auth_dto "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/dto"

	"github.com/Mirsadikovv/shared/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthMiddleware = middleware.AuthEchoMiddleware[*auth_dto.AuthUser, echo.Context, *gorm.DB, struct{}]
