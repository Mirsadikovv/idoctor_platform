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
	GetPINFLByUserId(userId int64) (int64, error)
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

func (s *userService) GetPINFLByUserId(userId int64) (int64, error) {

	var pinfl int64
	{
		if err := s.db.Model(&user_model.User{}).
			Where("id = ?", userId).
			Select("pin::BIGINT").
			Take(&pinfl).Error; err != nil {
			return 0, err
		}
	}

	return pinfl, nil
}

func clauseOnConflict() clause.OnConflict {
	return clause.OnConflict{
		Columns: []clause.Column{{Name: "pin"}},
		// DoUpdates: clause.AssignmentColumns([]string{
		// 	"username",
		// 	"valid",
		// 	"passport_number",
		// 	"first_name",
		// 	"last_name",
		// 	"middle_name",
		// 	"date_of_birth",
		// 	"user_type",
		// 	"last_visit",
		// 	"gender",
		// 	"email",
		// 	"session_id",
		// 	"photo",
		// 	"nationality",
		// 	"place_of_birth",
		// 	"country_of_birth",
		// 	"citizenship",
		// }),
		DoUpdates: clause.Assignments(map[string]interface{}{
			"username":        gorm.Expr("excluded.username"),
			"valid":           gorm.Expr("excluded.valid"),
			"passport_number": gorm.Expr("excluded.passport_number"),
			"first_name":      gorm.Expr("excluded.first_name"),
			"last_name":       gorm.Expr("excluded.last_name"),
			"middle_name":     gorm.Expr("excluded.middle_name"),
			"date_of_birth":   gorm.Expr("excluded.date_of_birth"),
			"user_type":       gorm.Expr("excluded.user_type"),
			"last_visit":      gorm.Expr("excluded.last_visit"),
			// Gender faqat mavjud bo'lmasa update qiladi
			"gender":           gorm.Expr("COALESCE(NULLIF(users.gender, ''), excluded.gender)"),
			"email":            gorm.Expr("excluded.email"),
			"session_id":       gorm.Expr("excluded.session_id"),
			"photo":            gorm.Expr("excluded.photo"),
			"nationality":      gorm.Expr("excluded.nationality"),
			"place_of_birth":   gorm.Expr("excluded.place_of_birth"),
			"country_of_birth": gorm.Expr("excluded.country_of_birth"),
			"citizenship":      gorm.Expr("excluded.citizenship"),
		}),
	}
}

func (s *userService) Upsert(ctx echo.Context, userDto *user_dto.User) (*user_model.User, error) {
	user := &user_model.User{
		Pin:            userDto.Pin,
		Username:       userDto.Username,
		Valid:          userDto.Valid,
		PassportNumber: userDto.PassportNumber,
		FirstName:      userDto.FirstName,
		LastName:       userDto.LastName,
		MiddleName:     userDto.MiddleName,
		DateOfBirth:    userDto.DateOfBirth,
		UserType:       userDto.UserType,
		LastVisit:      userDto.LastVisit,
		Gender:         string(userDto.Gender),
		Email:          userDto.Email,
		SessionId:      userDto.SessionId,
		Photo:          userDto.Photo,
		Nationality:    userDto.Nationality,
		PlaceOfBirth:   userDto.PlaceOfBirth,
		CountryOfBirth: userDto.CountryOfBirth,
		Citizenship:    userDto.Citizenship,
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

	userModel := &user_model.User{
		Username: userDto.Username,
		Name:     userDto.Name,
		RoleId:   userDto.RoleId,
	}

	// Используем Clauses для применения HASH_MAKE к паролю
	if err := s.db.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Create(map[string]interface{}{
			"username": userDto.Username,
			"password": gorm.Expr("HASH_MAKE(?)", userDto.Password),
			"name":     userDto.Name,
			"role_id":  userDto.RoleId,
		}).Scan(userModel).Error; err != nil {
		return 0, err
	}

	return userModel.Id, nil
}
