package device_cmd

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	device_handler "github.com/Mirsadikovv/idoctor_platform/src/module/device_service/handler"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Device Service API
// @version 1.0
// @description This is a Device Service API.
// @BasePath  /api/v1
// @Schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Cmd(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	{
		device_handler.NewDeviceHandler(router, db, log, authMiddleware)
	}
}
