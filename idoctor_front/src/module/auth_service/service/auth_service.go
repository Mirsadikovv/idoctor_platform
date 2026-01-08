package auth_service

import (
	"context"
	"log"

	auth_dto "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/dto"
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	user_dto "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/dto"
	user_model "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/model"
	user_service "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/service"
	"github.com/Mirsadikovv/shared/pg"

	"gorm.io/gorm"
)

type AuthService interface {
	SignIn(signIn *auth_dto.SignIn) (*auth_dto.Token, error)
	SignUp(signUp *auth_dto.SingUp) error
	Me(ctx context.Context, token string) (*user_dto.User, error)
	SignInTelegram(ctx context.Context, telegramId *int64) (*auth_dto.TelegramRole, error)
}

type authService struct {
	db             *gorm.DB
	userService    user_service.UserService
	authMiddleware *auth_middleware.AuthMiddleware
}

func NewAuthService(db *gorm.DB, authMiddleware *auth_middleware.AuthMiddleware) AuthService {
	return &authService{
		db:             db,
		authMiddleware: authMiddleware,
		userService:    user_service.NewUserService(db),
	}
}

func (a *authService) SignIn(signIn *auth_dto.SignIn) (*auth_dto.Token, error) {

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("username = ?", signIn.Username).
			Where("HASH_CHECK(?,password)", signIn.Password).
			Where("blocked_at IS NULL")
	}

	lastVisit := map[string]any{"last_visit": gorm.Expr("CURRENT_TIMESTAMP")}

	user, err := pg.Update[user_model.User](a.db, lastVisit, filter, "id", "role_id")
	{
		if err != nil {
			return nil, err
		}
	}

	authUser := auth_dto.AuthUser{
		Id:     user.Id,
		RoleId: user.RoleId,
	}

	token, err := a.authMiddleware.Token(&authUser)
	{
		if err != nil {
			return nil, err
		}
	}

	responseToken := auth_dto.Token{
		Token: token,
	}

	return &responseToken, nil
}

func (a *authService) SignUp(signUp *auth_dto.SingUp) error {

	userDto := user_dto.UserCreate{
		Username:    signUp.Username,
		Password:    signUp.Password,
		FirstName:   signUp.FirstName,
		LastName:    signUp.LastName,
		MiddleName:  signUp.MiddleName,
		DateOfBirth: signUp.DateOfBirth,
		Gender:      signUp.Gender,
		RoleId:      signUp.RoleId,
	}

	if _, err := a.userService.Create(&userDto); err != nil {
		return err
	}

	return nil
}

func (a *authService) Me(ctx context.Context, token string) (*user_dto.User, error) {

	authUser, err := a.authMiddleware.ParseTokenWithExpired(token)
	{
		if err != nil {
			return nil, err
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Select(
			"users.id",
			"users.username",
			"users.role_id",
			"roles.name as role",
			"users.first_name",
			"users.last_name",
			"users.middle_name",
			"users.date_of_birth",
			"users.gender",
			"roles.pages",
			"users.last_visit",
			"users.created_at",
			"users.blocked_at",
		).Joins("INNER JOIN roles ON roles.id = users.role_id").
			Where("users.id = ?", authUser.Id)
	}

	return a.userService.FindOne(ctx, filter)
}

// TODO: check possible mistakes
func (a *authService) SignInTelegram(ctx context.Context, telegramId *int64) (*auth_dto.TelegramRole, error) {

	if telegramId == nil {
		authUser := auth_dto.AuthUser{
			Id:     0,
			RoleId: 0,
		}

		token, err := a.authMiddleware.Token(&authUser)
		if err != nil {
			return nil, err
		}

		return &auth_dto.TelegramRole{
			Role:  "user",
			Token: token,
		}, nil
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Select(
			"users.id",
			"users.role_id",
			"roles.name as role",
		).Joins("LEFT JOIN roles ON roles.id = users.role_id").
			Where("users.telegram_id = ?", *telegramId).
			Where("users.blocked_at IS NULL").
			Order("users.last_visit DESC")
	}

	var result struct {
		Id     int64  `gorm:"column:id"`
		RoleId int64  `gorm:"column:role_id"`
		Role   string `gorm:"column:role"`
	}

	err := a.db.
		Table("users").
		Scopes(filter).
		Take(&result).Error

	if err != nil {
		return nil, err
	}

	log.Println("result", result)
	log.Println(result.RoleId)
	if err == gorm.ErrRecordNotFound || result.Role == "" || result.RoleId == 0 {
		authUser := auth_dto.AuthUser{
			Id:     result.Id,
			RoleId: 0,
		}

		token, err := a.authMiddleware.Token(&authUser)
		if err != nil {
			return nil, err
		}

		return &auth_dto.TelegramRole{
			Role:  "user",
			Token: token,
		}, nil
	}

	// Обновляем last_visit
	lastVisit := map[string]any{"last_visit": gorm.Expr("CURRENT_TIMESTAMP")}
	a.db.Table("users").Where("id = ?", result.Id).Updates(lastVisit)

	// Генерируем токен
	authUser := auth_dto.AuthUser{
		Id:     result.Id,
		RoleId: result.RoleId,
	}

	token, err := a.authMiddleware.Token(&authUser)
	if err != nil {
		return nil, err
	}

	return &auth_dto.TelegramRole{
		Role:  result.Role,
		Token: token,
	}, nil
}
