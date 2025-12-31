package language_handler

import (
	"fmt"

	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	language_dto "github.com/Mirsadikovv/idoctor_platform/src/module/language_service/dto"
	language_service "github.com/Mirsadikovv/idoctor_platform/src/module/language_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type languageContentHandler struct {
	db                     *gorm.DB
	log                    logger.Logger
	authMiddleware         any
	languageContentService language_service.LanguageContentService
}

func NewLanguageContentHandler(group *echo.Group, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {

	handler := &languageContentHandler{
		db:                     db,
		log:                    log,
		authMiddleware:         authMiddleware,
		languageContentService: language_service.NewLanguageContentService(db),
	}

	languageAuthMiddleware := authMiddleware.BuildMiddleware()
	languageGroup := group.Group("/language-content")
	{
		languageGroup.GET("/by-key/:key", handler.FindByKey)
		languageGroup.PUT("", handler.CreateOrUpdate, languageAuthMiddleware)
		languageGroup.GET("/:language_id", handler.GetByLangId)
		languageGroup.GET("/page/:language_id", handler.PageByLangId)
		languageGroup.DELETE("/:key", handler.DeleteByKey, languageAuthMiddleware)
	}
}

// Create godoc
// @Summary      language-content
// @Description  language-content
// @Tags 		 language-content
// @ID           create-or-update-language-content
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body language_dto.LanguageContentsCreateOrUpdate true "language-content information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language-content [put]
func (l *languageContentHandler) CreateOrUpdate(c echo.Context) error {

	req := request.RequestWithData[any](c)

	var languageContentDto language_dto.LanguageContentsCreateOrUpdate
	{
		if err := req.BindBody(&languageContentDto); err != nil {
			return req.BadRequest(err)
		}
	}
	// TODO: remove if it needs
	languageContentDto.Category = "GLOBAL"

	id, err := l.languageContentService.CreateOrUpdate(c, &languageContentDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.Created(response.NewID(id))
}

// GetById godoc
// @Summary      GetContent language-content by ID
// @Description  GetContent language-content by ID
// @Tags 		 language-content
// @ID           get-language-content-by-key
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        key path string true "language-content ID"
// @Success      200 {object} language_dto.LanguageContent "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language-content/by-key/{key} [get]
func (l *languageContentHandler) FindByKey(c echo.Context) error {

	req := request.RequestWithData[any](c)

	key := req.Param("key")

	filter := func(tx *gorm.DB) *gorm.DB {

		tx = tx.Where("category = 'GLOBAL' AND key = ?", key)

		return tx.
			Select(
				"key",
				"category",

				`JSON_AGG(
					JSON_BUILD_OBJECT(
						'id', id,
						'languageId', language_id, 
						'value', value
					)
				) AS contents`).
			Group("key, category")
	}

	language, err := l.languageContentService.FindByKey(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(language)
}

// GetByLangId godoc
// @Summary      GetContent language-content by ID
// @Description  GetContent language-content by ID
// @Tags 		 language-content
// @ID           find-translation-by-language-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        language_id path string true "language-content ID"
// @Success      200 {object} sharedutil.JsonObject "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language-content/{language_id} [get]
func (l *languageContentHandler) GetByLangId(c echo.Context) error {

	req := request.RequestWithData[any](c)

	languageId, err := req.ParamToInt("language_id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		return tx.
			Select(`JSONB_OBJECT_AGG(key,value) AS json_object`).
			Where("language_id = ?", languageId).
			Where("category = 'GLOBAL'").
			Group("language_id")
	}

	contents, err := l.languageContentService.FindByLanguage(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(contents)
}

// PageByLangId godoc
// @Summary      GetContent language-content by ID
// @Description  GetContent language-content by ID
// @Tags 		 language-content
// @ID           page-translation-by-language-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        language_id path string true "language-content ID"
// @Param        params query language_dto.PageByLangParams false "Params"
// @Success      200 {object} sharedutil.JsonObject "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language-content/page/{language_id} [get]
func (l *languageContentHandler) PageByLangId(c echo.Context) error {

	req := request.RequestWithData[any](c)

	var params language_dto.PageByLangParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	languageId, err := req.ParamToInt("language_id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if params.Key != nil {
			key := fmt.Sprintf("%%%s%%", *params.Key)
			tx = tx.Where("language_contents.key ILIKE ?", key)
		}

		if params.Value != nil {
			value := fmt.Sprintf("%%%s%%", *params.Value)
			tx = tx.Where("language_contents.value ILIKE ?", value)
		}

		return tx.
			Joins("LEFT JOIN languages ON languages.id = language_contents.language_id").
			Select(
				"language_contents.id",
				"language_contents.key",
				"language_contents.value",
				"language_contents.category",
				"language_contents.language_id",

				`JSONB_BUILD_OBJECT(
					'id', languages.id, 
					'name', 	   languages.name,
					'description', languages.description
				) AS language`).
			Where("languages.id = ?", languageId).
			Where("language_contents.category = 'GLOBAL'")
	}

	contents, err := l.languageContentService.PageByLanguage(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(contents)
}

// Create godoc
// @Summary      language-content
// @Description  language-content
// @Tags 		 language-content
// @ID           delete-language-content-by-key
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        key path string true "key"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /language-content/{key} [delete]
func (l *languageContentHandler) DeleteByKey(c echo.Context) error {

	req := request.RequestWithData[any](c)

	key := req.Param("key")
	{
		if key == "" {
			return req.BadRequest("key is required")
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("key = ?", key)
	}

	if err := l.languageContentService.Delete(c.Request().Context(), filter); err != nil {
		return req.BadRequest(err)
	}

	return req.NoContent()
}
