package organization_cmd

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	organization_handler "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service/handler"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Organization Service API
// @version 1.0
// @description This is a Organization Service API.
// @BasePath  /api/v1
// @Schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Cmd(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {

	routerGroup := router.Group("/api/v1")
	{
		organization_handler.NewOrganizationHandler(routerGroup, db, log, authMiddleware)
	}
}
