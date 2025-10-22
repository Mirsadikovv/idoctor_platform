package auth_cmd

import (
	auth_handler "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/handler"
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"

	"github.com/Mirsadikovv/shared/shared_service/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Auth Service API
// @version 1.0
// @description This is a Auth Service API.
// @BasePath  /api/v1
// @Schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Cmd(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {

	routerGroup := router.Group("/api/v1")
	{
		auth_handler.NewAuthHandler(routerGroup, db, log, authMiddleware)
	}
}
