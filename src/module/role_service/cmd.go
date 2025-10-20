package role_cmd

import (
	auth_middleware "git.sriss.uz/mehnat/inspector_platform/src/module/auth_service/middleware"
	role_handler "git.sriss.uz/mehnat/inspector_platform/src/module/role_service/handler"

	"git.sriss.uz/shared/shared_service/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Role Service API
// @version 1.0
// @description This is a Role Service API.
// @BasePath  /api/v1
// @Schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Cmd(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	{
		role_handler.NewRoleHandler(router, db, log, authMiddleware)
	}
}
