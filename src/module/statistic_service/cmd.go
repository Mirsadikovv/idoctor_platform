package statistic_cmd

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	statistic_handler "github.com/Mirsadikovv/idoctor_platform/src/module/statistic_service/handler"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Statistic Service API
// @version 1.0
// @description This is a Statistic Service API for order analytics.
// @BasePath  /api/v1
// @Schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Cmd(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	{
		statistic_handler.NewStatisticHandler(router, db, log, authMiddleware)
	}
}
