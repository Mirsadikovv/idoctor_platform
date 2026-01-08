package language_cmd

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	language_handler "github.com/Mirsadikovv/idoctor_platform/src/module/language_service/handler"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Language Service API
// @version 1.0
// @description This is a Language Service API.
// @BasePath  /api/v1
// @Schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Cmd(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {

	routerGroup := router.Group("/api/v1")
	{
		language_handler.NewLanguageHandler(routerGroup, db, log, authMiddleware)
		language_handler.NewLanguageContentHandler(routerGroup, db, log, authMiddleware)
	}
}
