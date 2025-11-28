package user_service

import (
	"context"

	user_dto "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/dto"
	user_model "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/model"
	"github.com/labstack/echo/v4"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService interface {
	Upsert(ctx echo.Context, userDto *user_dto.User) (*user_model.User, error)
	Delete(ctx echo.Context, filter pg.Filter) error
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*user_dto.UserPage, error)
	Find(ctx context.Context, filter pg.Filter) ([]user_dto.User, error)
	FindOne(ctx context.Context, filter pg.Filter) (*user_dto.User, error)
	DeleteOrRestore(ctx echo.Context, filter pg.Filter) error
	Create(userDto *user_dto.UserCreate) (int64, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (s *userService) Delete(ctx echo.Context, filter pg.Filter) error {
	return pg.Delete[user_model.User](s.db.WithContext(ctx.Request().Context()), nil, filter)
}

func (s *userService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*user_dto.UserPage, error) {
	return pg.PageWithScan[user_model.User, user_dto.User](s.db, paginate, filter)
}

func (s *userService) Find(ctx context.Context, filter pg.Filter) ([]user_dto.User, error) {
	users, err := pg.FindWithScan[user_model.User, user_dto.User](s.db, filter)
	{
		if err != nil {
			return nil, err
		}
	}

	if users == nil {
		users = []user_dto.User{}
	}

	return users, nil
}

func (s *userService) FindOne(ctx context.Context, filter pg.Filter) (*user_dto.User, error) {
	return pg.FindOneWithScan[user_model.User, user_dto.User](s.db, filter)
}


func clauseOnConflict() clause.OnConflict {
	return clause.OnConflict{
		Columns: []clause.Column{{Name: "username"}},
		DoUpdates: clause.Assignments(map[string]any{
			"first_name":    gorm.Expr("excluded.first_name"),
			"last_name":     gorm.Expr("excluded.last_name"),
			"middle_name":   gorm.Expr("excluded.middle_name"),
			"date_of_birth": gorm.Expr("excluded.date_of_birth"),
			"last_visit":    gorm.Expr("excluded.last_visit"),
			"gender":        gorm.Expr("COALESCE(NULLIF(users.gender, ''), excluded.gender)"),
			"role_id":       gorm.Expr("excluded.role_id"),
		}),
	}
}

func (s *userService) Upsert(ctx echo.Context, userDto *user_dto.User) (*user_model.User, error) {
	user := &user_model.User{
		Username:    userDto.Username,
		FirstName:   userDto.FirstName,
		LastName:    userDto.LastName,
		MiddleName:  userDto.MiddleName,
		DateOfBirth: userDto.DateOfBirth,
		LastVisit:   userDto.LastVisit,
		Gender:      string(userDto.Gender),
		RoleId:      userDto.RoleId,
	}

	if err := s.db.WithContext(ctx.Request().Context()).Clauses(clauseOnConflict()).
		Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteOrRestore(ctx echo.Context, filter pg.Filter) error {

	restoreData := map[string]any{
		"deleted_at": gorm.Expr("CASE WHEN deleted_at IS NULL THEN now() ELSE null END"),
	}

	if _, err := pg.Update[user_model.User](s.db.WithContext(ctx.Request().Context()).Unscoped(), restoreData, filter); err != nil {
		return err
	}

	return nil
}

func (s *userService) Create(userDto *user_dto.UserCreate) (int64, error) {

	data := map[string]any{
		"username": userDto.Username,
		"password": gorm.Expr("HASH_MAKE(?)", userDto.Password),
		"role_id":  userDto.RoleId,
	}

	// Добавляем опциональные поля, если они указаны
	if userDto.FirstName != nil {
		data["first_name"] = *userDto.FirstName
	}
	if userDto.LastName != nil {
		data["last_name"] = *userDto.LastName
	}
	if userDto.MiddleName != nil {
		data["middle_name"] = *userDto.MiddleName
	}
	if userDto.DateOfBirth != nil {
		data["date_of_birth"] = *userDto.DateOfBirth
	}
	if userDto.Gender != nil {
		data["gender"] = *userDto.Gender
	}

	userModel := &user_model.User{
		Username: userDto.Username,
		RoleId:   userDto.RoleId,
	}

	// Используем Clauses для применения HASH_MAKE к паролю
	if err := s.db.Table("users").Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Create(data).Scan(userModel).Error; err != nil {
		return 0, err
	}

	return userModel.Id, nil
}
