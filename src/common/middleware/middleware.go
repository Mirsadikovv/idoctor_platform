package middleware

import (
	"net/http"
	"strings"

	auth_dto "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/dto"
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	"github.com/Mirsadikovv/shared/middleware"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
)

func AllowXAuthXMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// req := request.RequestWithData[auth_dto.AuthUser](c)

		// // Check if user is authenticated
		// fmt.Println("AllowXAuthXMiddleware:", req.AuthUser() == nil)
		// if req.AuthUser() != nil {
		// 	return next(c)
		// }

		//TODO: Check if user is authenticated
		token, _ := middleware.AuthorizationToken(c)
		{
			if token != "" {
				//TODO: If token is present, assume user is authenticated
				return next(c)
			}
		}

		if strings.TrimSpace(c.Request().Header.Get("X-Auth")) == "7" {
			return next(c)
		}

		// Deny otherwise
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden: missing or invalid X-Auth header")
	}
}

func TokenParser(authMiddleware *auth_middleware.AuthMiddleware) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			token, err := middleware.AuthorizationToken(ctx)
			{
				if err != nil {
					return response.HTTPError(err).Unauthorized()
				}
			}

			user, err := authMiddleware.ParseTokenWithExpired(token)
			{
				if err != nil {
					return response.HTTPError(err.Error()).Unauthorized()
				}
			}

			req := request.RequestWithData[auth_dto.AuthUser](ctx)
			req.SetUser(user)

			return next(ctx)
		}
	}
}
