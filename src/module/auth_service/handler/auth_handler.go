package auth_handler

import (
	"fmt"

	auth_dto "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/dto"
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	auth_service "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type authHandler struct {
	db             *gorm.DB
	log            logger.Logger
	authMiddleware *auth_middleware.AuthMiddleware
	authService    auth_service.AuthService
}

func NewAuthHandler(group *echo.Group, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {

	handler := authHandler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		authService:    auth_service.NewAuthService(db, authMiddleware),
	}

	authGroup := group.Group("/auth")
	{
		authGroup.POST("/sign-in", handler.SignIn)
		authGroup.POST("/sign-up", handler.SignUp)
		authGroup.POST("/sign-out", handler.SignOut)
		authGroup.POST("/me", handler.Me)
		authGroup.POST("/sign-in-telegram", handler.SignInTelegram)
	}
}

// Create godoc
// @Summary      sign in
// @Description  sign in
// @Tags 		 auth
// @ID           sign-in
// @Accept       json
// @Produce      json
// @Param        input body auth_dto.SignIn true "sign in information"
// @Success      201 {object} auth_dto.Token "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /auth/sign-in [post]
func (a *authHandler) SignIn(ctx echo.Context) error {

	req := request.Request(ctx)
	a.log.Info("SignIn\n")
	var signIn auth_dto.SignIn
	{
		if err := req.BindBody(&signIn); err != nil {
			a.log.Error(err)
			return req.BadRequest(err)
		}
	}

	token, err := a.authService.SignIn(&signIn)
	{
		if err != nil {
			a.log.Error(err)
			return req.BadRequest(err)
		}
	}

	return req.OK(token)
}

// Create godoc
// @Summary      sign up
// @Description  sign up
// @Tags 		 auth
// @ID           sign-up
// @Accept       json
// @Produce      json
// @Param        input body auth_dto.SingUp true "sign in information"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /auth/sign-up [post]
func (a *authHandler) SignUp(ctx echo.Context) error {

	req := request.Request(ctx)

	var signUp auth_dto.SingUp
	{
		if err := req.BindBody(&signUp); err != nil {
			return req.BadRequest(err)
		}
	}

	if err := a.authService.SignUp(&signUp); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Create godoc
// @Summary      sign out
// @Description  sign out
// @Tags 		 auth
// @ID           sign-out
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      204 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /auth/sign-out [post]
func (a *authHandler) SignOut(ctx echo.Context) error {
	req := request.Request(ctx)

	return req.NoContent()
}

// Create godoc
// @Summary      sign in
// @Description  sign in
// @Tags 		 auth
// @ID           me
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      201 {object} user_dto.User "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /auth/me [post]
func (a *authHandler) Me(ctx echo.Context) error {

	req := request.Request(ctx)

	token, err := req.AuthorizationTokenWithBearer()
	{
		if err != nil {
			return req.Unauthorized(err)
		}
	}

	user, err := a.authService.Me(req.Context(), token)
	{
		if err != nil {
			return req.Unauthorized(err)
		}
	}

	return req.OK(user)
}

// Create godoc
// @Summary      sign in with telegram
// @Description  sign in with telegram
// @Tags 		 auth
// @ID           sign-in-telegram
// @Accept       json
// @Produce      json
// @Param        telegram_id header int false "Telegram ID"
// @Success      200 {object} auth_dto.TelegramRole "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /auth/sign-in-telegram [post]
func (a *authHandler) SignInTelegram(ctx echo.Context) error {

	req := request.Request(ctx)
	a.log.Info("SignInTelegram\n")

	// Получаем telegram_id из header
	telegramIdHeader := ctx.Request().Header.Get("telegram_id")
	var telegramId *int64
	if telegramIdHeader != "" {
		// Парсим telegram_id из строки в int64
		var id int64
		_, err := fmt.Sscanf(telegramIdHeader, "%d", &id)
		if err != nil {
			a.log.Error(err)
			return req.BadRequest(err)
		}
		telegramId = &id
	}
	role, err := a.authService.SignInTelegram(req.Context(), telegramId)
	{
		if err != nil {
			a.log.Error(err)
			return req.BadRequest(err)
		}
	}

	return req.OK(role)
}
