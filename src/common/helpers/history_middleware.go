package helpers

import (
	"context"
	"fmt"

	"git.sriss.uz/shared/shared_service/jwt"
	"github.com/labstack/echo/v4"
)

func HistoryMiddleware(jwtConfig *jwt.JwtConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user := SoftAuthUser(c, jwtConfig)

			var userID int64
			if user != nil {
				userID = user.Id
			} else {
				userID = 0
			}

			ip := c.RealIP()
			path := c.Path()
			method := c.Request().Method

			ctx := context.WithValue(c.Request().Context(), "history_context", map[string]string{
				"user_id": fmt.Sprintf("%d", userID),
				"ip":      ip,
				"path":    path,
				"method":  method,
			})

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
