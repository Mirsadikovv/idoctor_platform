package helpers

import (
	"strings"

	auth_dto "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/dto"
	"github.com/Mirsadikovv/shared/jwt"
	"github.com/labstack/echo/v4"
)

func SoftAuthUser(c echo.Context, jwtConfig *jwt.JwtConfig) *auth_dto.AuthUser {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return nil
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return nil
	}

	token := parts[1]

	user, err := jwt.NewJwtService[*auth_dto.AuthUser](jwtConfig).ParseTokenWithExpired(token)
	if err != nil {
		return nil
	}

	return user
}
