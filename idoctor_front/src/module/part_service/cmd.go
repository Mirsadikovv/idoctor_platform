package part_cmd

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	part_handler "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/handler"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Part Service API
// @version 1.0
// @description This is a Part Service API.
// @BasePath  /api/v1
// @Schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Cmd(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	{
		part_handler.NewPartHandler(router, db, log, authMiddleware)
	}
}
