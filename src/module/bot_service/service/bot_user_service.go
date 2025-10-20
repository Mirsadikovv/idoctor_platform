package bot_service

import (
	"context"

	bot_dto "git.sriss.uz/mehnat/inspector_platform/src/module/bot_service/dto"
	bot_model "git.sriss.uz/mehnat/inspector_platform/src/module/bot_service/model"
	"github.com/labstack/echo/v4"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"gorm.io/gorm"
)

type BotService interface {
	Create(ctx echo.Context, botDto *bot_dto.BotUserCreate) (int64, error)
	Update(ctx echo.Context, botDto *bot_dto.BotUserUpdate, filter pg.Filter) error
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*bot_dto.BotUserPage, error)
	Find(ctx context.Context, filter pg.Filter) ([]bot_dto.BotUserDto, error)
	FindOne(ctx context.Context, filter pg.Filter) (*bot_dto.BotUserDto, error)
}

type botService struct {
	db *gorm.DB
}

func NewBotService(db *gorm.DB) BotService {
	return &botService{
		db: db,
	}
}

func (s *botService) Create(ctx echo.Context, input *bot_dto.BotUserCreate) (int64, error) {
	var model = &bot_model.BotUser{
		Name:         input.Name,
		TelegramId:   input.TelegramId,
		Username:     input.Username,
		PhoneNumber:  input.PhoneNumber,
		LanguageCode: input.LanguageCode,
	}

	if err := pg.Create(s.db.WithContext(ctx.Request().Context()), model, "id"); err != nil {
		return 0, err
	}

	return model.Id, nil
}

func (s *botService) Update(ctx echo.Context, input *bot_dto.BotUserUpdate, filter pg.Filter) error {
	if _, err := pg.Update[bot_model.BotUser](s.db.WithContext(ctx.Request().Context()), input, filter); err != nil {
		return err
	}

	return nil
}

func (s *botService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*bot_dto.BotUserPage, error) {
	return pg.PageWithScan[bot_model.BotUser, bot_dto.BotUserDto](s.db, paginate, filter)
}

func (s *botService) Find(ctx context.Context, filter pg.Filter) ([]bot_dto.BotUserDto, error) {
	users, err := pg.FindWithScan[bot_model.BotUser, bot_dto.BotUserDto](s.db, filter)
	{
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func (s *botService) FindOne(ctx context.Context, filter pg.Filter) (*bot_dto.BotUserDto, error) {
	return pg.FindOneWithScan[bot_model.BotUser, bot_dto.BotUserDto](s.db, filter)
}
