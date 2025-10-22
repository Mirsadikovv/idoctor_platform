package language_handler

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	language_dto "github.com/Mirsadikovv/idoctor_platform/src/module/language_service/dto"
	language_service "github.com/Mirsadikovv/idoctor_platform/src/module/language_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type languageHandler struct {
	db              *gorm.DB
	log             logger.Logger
	authMiddleware  any
	languageService language_service.LanguageService
}

func NewLanguageHandler(group *echo.Group, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &languageHandler{
		db:              db,
		log:             log,
		authMiddleware:  authMiddleware,
		languageService: language_service.NewLanguageService(db),
	}

	languageAuthMiddleware := authMiddleware.BuildMiddleware()
	languageGroup := group.Group("/language")
	{
		languageGroup.GET("/first", handler.First)
		languageGroup.GET("/:id", handler.FindByID)
		languageGroup.GET("/search", handler.Find)
		languageGroup.GET("/page", handler.Page)

		languageGroup.POST("", handler.Create, languageAuthMiddleware)
		languageGroup.PATCH("/:id", handler.Update, languageAuthMiddleware)
		languageGroup.PATCH("/:id/restore", handler.Restore, languageAuthMiddleware)
		languageGroup.DELETE("/:id", handler.Delete, languageAuthMiddleware)
	}
}

// godoc
// @Summary      Create Language
// @Description  Create Language
// @Tags         language
// @Accept       json
// @Produce      json
// @Param        body  body      language_dto.LanguageCreate  true  "Language Create"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language [post]
func (l *languageHandler) Create(c echo.Context) error {

	req := request.RequestWithData[any](c)

	var languageDto language_dto.LanguageCreate
	{
		if err := req.BindBody(&languageDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := l.languageService.Create(c.Request().Context(), &languageDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.Created(response.NewID(id))
}

// Create godoc
// @Summary      language
// @Description  language
// @Tags 		 language
// @ID           update-language
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "language ID"
// @Param        input body language_dto.LanguageUpdate true "language information"
// @Success      204 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language/{id} [patch]
func (l *languageHandler) Update(c echo.Context) error {

	req := request.RequestWithData[any](c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var (
		languageDto language_dto.LanguageUpdate
	)
	{
		if err := req.BindBody(&languageDto); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if err := l.languageService.Update(c.Request().Context(), &languageDto, filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Create godoc
// @Summary      language
// @Description  language
// @Tags 		 language
// @ID           delete-language
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "language ID"
// @Success      204 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language/{id} [delete]
func (l *languageHandler) Delete(c echo.Context) error {

	req := request.RequestWithData[any](c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if err := l.languageService.Delete(c.Request().Context(), filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Create godoc
// @Summary      language
// @Description  language
// @Tags 		 language
// @ID           restore-language
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "language ID"
// @Success      204 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language/{id}/restore [patch]
func (l *languageHandler) Restore(c echo.Context) error {

	req := request.RequestWithData[any](c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	if err := l.languageService.DeleteOrRestore(c.Request().Context(), filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}

// Create godoc
// @Summary      language
// @Description  language
// @Tags 		 language
// @ID           language-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "language ID"
// @Success      200 {object} language_dto.Language "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language/{id} [get]
func (l *languageHandler) FindByID(c echo.Context) error {

	req := request.RequestWithData[any](c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}

	language, err := l.languageService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(language)
}

// Create godoc
// @Summary      language
// @Description  language
// @Tags 		 language
// @ID           language-first
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200 {object} language_dto.Language "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language/first [get]
func (l *languageHandler) First(c echo.Context) error {

	req := request.RequestWithData[any](c)

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Order("id ASC").Limit(1)
	}

	language, err := l.languageService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(language)
}

// Create godoc
// @Summary      language
// @Description  language
// @Tags 		 language
// @ID           language-find
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        name query string false "name"
// @Param        description query string false "description"
// @Success      200 {object} []language_dto.Language "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language/search [get]
func (l *languageHandler) Find(c echo.Context) error {

	req := request.RequestWithData[any](c)

	filter := func(tx *gorm.DB) *gorm.DB {

		if name := req.Query("name"); name != "" {
			tx = tx.Where("name = ?", name)
		}

		if description := req.Query("description"); description != "" {
			tx = tx.Where("description = ?", description)
		}

		return tx
	}

	languages, err := l.languageService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(languages)
}

// Create godoc
// @Summary      language
// @Description  language
// @Tags 		 language
// @ID           language-page
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        with_deleted query string false "with_deleted"
// @Param        name query string false "name"
// @Param        description query string false "description"
// @Success      200 {object} []language_dto.Language "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language/page [get]
func (l *languageHandler) Page(c echo.Context) error {

	req := request.RequestWithData[any](c)

	filter := func(tx *gorm.DB) *gorm.DB {

		if withDeleted := req.Query("with_deleted"); withDeleted == "true" {
			tx = tx.Unscoped()
		}

		if name := req.Query("name"); name != "" {
			tx = tx.Where("name = ?", name)
		}

		if description := req.Query("description"); description != "" {
			tx = tx.Where("description = ?", description)
		}

		return tx.Order("id desc")
	}

	languages, err := l.languageService.Page(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(languages)
}
