package bot_handler

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	bot_dto "github.com/Mirsadikovv/idoctor_platform/src/module/bot_service/dto"
	user_dto "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/dto"
	user_model "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/model"
	user_service "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BotHandler interface {
}

type botHandler struct {
	db             *gorm.DB
	log            logger.Logger
	userService    user_service.UserService
	authMiddleware *auth_middleware.AuthMiddleware
}

func NewBotHandler(group *echo.Group, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &botHandler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		userService:    user_service.NewUserService(db),
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

	// Create user with telegram data
	userCreate := &user_dto.UserCreate{
		Username:         botDto.Username,
		Password:         "default_password", // Or generate random password
		FirstName:        &botDto.Name,
		TelegramId:       &botDto.TelegramId,
		TelegramUsername: &botDto.Username,
		PhoneNumber:      &botDto.PhoneNumber,
		LanguageCode:     &botDto.LanguageCode,
		RoleId:           1, // Default role for bot users
	}

	id, err := o.userService.Create(userCreate)
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

	// Update user telegram data
	updates := make(map[string]interface{})

	if botDto.Name != nil {
		updates["first_name"] = *botDto.Name
	}
	if botDto.TelegramId != nil {
		updates["telegram_id"] = *botDto.TelegramId
	}
	if botDto.Username != nil {
		updates["telegram_username"] = *botDto.Username
	}
	if botDto.PhoneNumber != nil {
		updates["phone_number"] = *botDto.PhoneNumber
	}
	if botDto.LanguageCode != nil {
		updates["language_code"] = *botDto.LanguageCode
	}

	if len(updates) > 0 {
		if err := o.db.Model(&user_model.User{}).Where("id = ?", id).Updates(updates).Error; err != nil {
			return req.BadRequest(err)
		}
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

	type BotUserDto struct {
		Id           int64  `json:"id"`
		Name         string `json:"name"`
		TelegramId   int64  `json:"telegramId"`
		Username     string `json:"username"`
		PhoneNumber  string `json:"phoneNumber"`
		LanguageCode string `json:"languageCode"`
	}

	tx := o.db.Table("users").Where("telegram_id IS NOT NULL")

	if params.Name != nil {
		tx = tx.Where("first_name ILIKE ?", "%"+*params.Name+"%")
	}

	if params.PhoneNumber != nil {
		tx = tx.Where("phone_number ILIKE ?", "%"+*params.PhoneNumber+"%")
	}

	if params.TelegramId != nil {
		tx = tx.Where("telegram_id = ?", *params.TelegramId)
	}

	if params.Username != nil {
		tx = tx.Where("telegram_username ILIKE ?", "%"+*params.Username+"%")
	}

	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return req.BadRequest(err)
	}

	paginate := req.NewPaginate()
	tx = tx.Offset(paginate.Offset()).Limit(paginate.Limit())

	var items []BotUserDto
	if err := tx.Select(
		"id",
		"first_name as name",
		"telegram_id",
		"telegram_username as username",
		"phone_number",
		"language_code",
	).Find(&items).Error; err != nil {
		return req.BadRequest(err)
	}

	result := map[string]interface{}{
		"data":  items,
		"total": total,
	}

	return req.OK(result)
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

	type BotUserDto struct {
		Id           int64  `json:"id"`
		Name         string `json:"name"`
		TelegramId   int64  `json:"telegramId"`
		Username     string `json:"username"`
		PhoneNumber  string `json:"phoneNumber"`
		LanguageCode string `json:"languageCode"`
	}

	tx := o.db.Table("users").Where("telegram_id IS NOT NULL")

	if params.Name != nil {
		tx = tx.Where("first_name ILIKE ?", "%"+*params.Name+"%")
	}

	if params.PhoneNumber != nil {
		tx = tx.Where("phone_number ILIKE ?", "%"+*params.PhoneNumber+"%")
	}

	if params.TelegramId != nil {
		tx = tx.Where("telegram_id = ?", *params.TelegramId)
	}

	if params.Username != nil {
		tx = tx.Where("telegram_username ILIKE ?", "%"+*params.Username+"%")
	}

	var items []BotUserDto
	if err := tx.Select(
		"id",
		"first_name as name",
		"telegram_id",
		"telegram_username as username",
		"phone_number",
		"language_code",
	).Limit(20).Find(&items).Error; err != nil {
		return req.BadRequest(err)
	}

	if items == nil {
		items = []BotUserDto{}
	}

	return req.OK(items)
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

	type BotUserDto struct {
		Id           int64  `json:"id"`
		Name         string `json:"name"`
		TelegramId   int64  `json:"telegramId"`
		Username     string `json:"username"`
		PhoneNumber  string `json:"phoneNumber"`
		LanguageCode string `json:"languageCode"`
	}

	var user BotUserDto
	if err := o.db.Table("users").
		Where("id = ? AND telegram_id IS NOT NULL", id).
		Select(
			"id",
			"first_name as name",
			"telegram_id",
			"telegram_username as username",
			"phone_number",
			"language_code",
		).First(&user).Error; err != nil {
		return req.BadRequest(err)
	}

	return req.OK(user)
}
