package bot_handler

import (
	auth_middleware "git.sriss.uz/mehnat/inspector_platform/src/module/auth_service/middleware"
	bot_dto "git.sriss.uz/mehnat/inspector_platform/src/module/bot_service/dto"
	bot_service "git.sriss.uz/mehnat/inspector_platform/src/module/bot_service/service"

	"git.sriss.uz/shared/shared_service/logger"
	"git.sriss.uz/shared/shared_service/request"
	"git.sriss.uz/shared/shared_service/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BotHandler interface {
}

type botHandler struct {
	db             *gorm.DB
	log            logger.Logger
	botService     bot_service.BotService
	authMiddleware *auth_middleware.AuthMiddleware
}

func NewBotHandler(group *echo.Group, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &botHandler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		botService:     bot_service.NewBotService(db),
	}

	botGroup := group.Group("/bot_user")
	{
		botGroup.POST("/create", handler.Create)
		botGroup.PATCH("/:id", handler.Update)
		botGroup.GET("/page", handler.Page)
		botGroup.GET("/search", handler.Search)
		botGroup.GET("/:id", handler.GetById)
	}

}

// Create godoc
// @Summary      bot
// @Description  bot
// @Tags 		 bot_user
// @ID           create-bot
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body bot_dto.BotUserCreate true "bot information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /bot_user/create [post]
func (o *botHandler) Create(ctx echo.Context) error {

	req := request.Request(ctx)

	var botDto bot_dto.BotUserCreate
	{
		if err := req.BindBody(&botDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := o.botService.Create(ctx,&botDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.Created(response.NewID(id))
}

// Update 		 godoc
// @Summary      bot
// @Description  bot
// @Tags 		 bot_user
// @ID           update-bot
// @Accept       json
// @Produce      json
// @Param        id path string true "bot ID"
// @Security     ApiKeyAuth
// @Param        input body bot_dto.BotUserUpdate true "bot information"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /bot_user/{id} [patch]
func (o *botHandler) Update(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var botDto bot_dto.BotUserUpdate
	{
		if err := req.BindBody(&botDto); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("bots.id = ?", id)
	}

	if err := o.botService.Update(ctx, &botDto, filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Page 		 godoc
// @Summary      GetContent all bot with pagination
// @Description  GetContent all bot with pagination
// @Tags 		 bot_user
// @ID           get-all-bot
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        params query bot_dto.BotUserParams false "params"
// @Success      200 {object} bot_dto.BotUserPage "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /bot_user/page [get]
func (o *botHandler) Page(ctx echo.Context) error {

	req := request.Request(ctx)

	var params bot_dto.BotUserParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if params.Name != nil {
			tx = tx.Where("bot_users.name ILIKE ?", "%"+*params.Name+"%")
		}

		if params.PhoneNumber != nil {
			tx = tx.Where("bot_users.phone_number ILIKE ?", "%"+*params.PhoneNumber+"%")
		}

		if params.TelegramId != nil {
			tx = tx.Where("bot_users.telegram_id = ?", *params.TelegramId)
		}

		if params.Username != nil {
			tx = tx.Where("bot_users.username ILIKE ?", "%"+*params.Username+"%")
		}

		return tx.Select("bot_users.*")
	}

	userPage, err := o.botService.Page(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(userPage)
}

// Search godoc
// @Summary      GetContent all bot with pagination
// @Description  GetContent all bot with pagination
// @Tags 		 bot_user
// @ID           search-bot
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Param        params query bot_dto.BotUserParams false "params"
// @Success      200 {object} []bot_dto.BotUserDto "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /bot_user/search [get]
func (o *botHandler) Search(ctx echo.Context) error {

	req := request.Request(ctx)

	var params bot_dto.BotUserParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if params.Name != nil {
			tx = tx.Where("bot_users.name ILIKE ?", "%"+*params.Name+"%")
		}

		if params.PhoneNumber != nil {
			tx = tx.Where("bot_users.phone_number ILIKE ?", "%"+*params.PhoneNumber+"%")
		}

		if params.TelegramId != nil {
			tx = tx.Where("bot_users.telegram_id = ?", *params.TelegramId)
		}

		if params.Username != nil {
			tx = tx.Where("bot_users.username ILIKE ?", "%"+*params.Username+"%")
		}

		return tx.Select("bot_users.*")
	}

	bots, err := o.botService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}

		if bots == nil {
			return req.OK([]bot_dto.BotUserDto{})
		}
	}

	return req.OK(bots)
}

// GetById godoc
// @Summary      GetContent bot by ID
// @Description  GetContent bot by ID
// @Tags 		 bot_user
// @ID           get-bot-user-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "bot ID"
// @Success      200 {object} bot_dto.BotUserDto "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /bot_user/{id} [get]
func (o *botHandler) GetById(ctx echo.Context) error {

	req := request.Request(ctx)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("bot_users.id = ?", id)

		return tx.Select("bot_users.*")
	}

	user, err := o.botService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(user)
}
