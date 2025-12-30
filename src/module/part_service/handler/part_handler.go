package part_handler

import (
	"fmt"

	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	part_dto "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/dto"
	part_service "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type partHandler struct {
	db             *gorm.DB
	log            logger.Logger
	authMiddleware *auth_middleware.AuthMiddleware
	partService    part_service.PartService
}

func NewPartHandler(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &partHandler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		partService:    part_service.NewPartService(db),
	}

	// partServiceMiddleware := handler.authMiddleware.BuildMiddleware()
	partGroup := router.Group("/api/v1/part")
	{
		partGroup.GET("/:id", handler.FindById)
		partGroup.GET("/search", handler.Search)
		partGroup.GET("/page", handler.Page)
		partGroup.POST("", handler.Create)
		partGroup.PUT("/:id", handler.Update)
		partGroup.DELETE("/:id", handler.DeleteOrRestore)
	}

}

// Search godoc
// @Summary      Get all parts with search
// @Description  Get all parts with search
// @Tags 		 part
// @Id           search-part
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        params query part_dto.PartParams false "params"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} []part_dto.Part "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /part/search [get]
func (h *partHandler) Search(c echo.Context) error {
	req := request.RequestWithData[any](c)

	var params part_dto.PartParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if params.Name != nil {
			name := fmt.Sprintf("%%%s%%", *params.Name)
			tx = tx.Where("parts.name ILIKE ?", name)
		}

		if params.DeviceId != nil {
			tx = tx.Where("parts.device_id = ?", *params.DeviceId)
		}

		if params.SupplierId != nil {
			tx = tx.Where("parts.supplier_id = ?", *params.SupplierId)
		}

		// Handle soft delete filtering
		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("parts.deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Select(
			"parts.id",
			"parts.name",
			"parts.device_id",
			"parts.supplier_id",
			"parts.created_at",
			"parts.deleted_at",
		).Order("parts.id DESC")
	}

	parts, err := h.partService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(parts)
}

// Page 			godoc
// @Summary			part-service-page
// @Description		part-service-page
// @Tags 			part
// @Param       	filter query part_dto.PartParams false "Filter parameters"
// @Param       	page query int false "Page number"
// @Param       	perpage query int false "Items per page"
// @Security     	ApiKeyAuth
// @Success	200 	{object} part_dto.PartPage "Successful operation"
// @Failure	400 	{object} response.HttpSuccess "Bad request"
// @Failure	500 	{object} response.HttpSuccess "Internal server error"
// @Router			/part/page [get]
func (h *partHandler) Page(c echo.Context) error {
	req := request.Request(c)

	var params part_dto.PartParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	tx := func(tx *gorm.DB) *gorm.DB {
		if params.Name != nil {
			tx = tx.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *params.Name))
		}

		if params.DeviceId != nil {
			tx = tx.Where("device_id = ?", *params.DeviceId)
		}

		if params.SupplierId != nil {
			tx = tx.Where("supplier_id = ?", *params.SupplierId)
		}

		// Handle soft delete filtering
		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Select(
			"parts.id",
			"parts.name",
			"parts.device_id",
			"parts.supplier_id",
			"parts.created_at",
			"parts.deleted_at",
		).Order("parts.id DESC")
	}

	items, err := h.partService.Page(req.Context(), req.NewPaginate(), tx)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(items)
}

// GetById godoc
// @Summary      Get part by Id
// @Description  Get part by Id
// @Tags 		 part
// @Id           get-part-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "part Id"
// @Success      200 {object} part_dto.Part "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /part/{id} [get]
func (h *partHandler) FindById(c echo.Context) error {
	req := request.RequestWithData[any](c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("Device").
			Preload("Supplier").
			Preload("Master").
			Where("id = ?", id)
	}

	part, err := h.partService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(part)
}

// Create godoc
// @Summary      Create part
// @Description  Create new part
// @Tags 		 part
// @Id           create-part
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body part_dto.PartCreate true "part information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /part [POST]
func (h *partHandler) Create(c echo.Context) error {
	req := request.Request(c)

	var partDto part_dto.PartCreate
	{
		if err := req.BindBody(&partDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := h.partService.Create(c.Request().Context(), &partDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(response.NewID(id))
}

// Update godoc
// @Summary      Update part
// @Description  Update existing part
// @Tags 		 part
// @Id           update-part
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "part Id"
// @Param        input body part_dto.PartUpdate true "part information"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /part/{id} [PUT]
func (h *partHandler) Update(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var partDto part_dto.PartUpdate
	{
		if err := req.BindBody(&partDto); err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.partService.Update(c.Request().Context(), id, &partDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Part updated successfully"})
}

// DeleteOrRestore godoc
// @Summary      Delete or restore part
// @Description  Soft delete part if active, restore if already deleted
// @Tags 		 part
// @Id           delete-or-restore-part
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "part Id"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /part/{id} [delete]
func (h *partHandler) DeleteOrRestore(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.partService.DeleteOrRestore(req.Context(), id)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Part status toggled successfully"})
}
