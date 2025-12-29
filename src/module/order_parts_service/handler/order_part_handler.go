package order_part_handler

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	order_part_dto "github.com/Mirsadikovv/idoctor_platform/src/module/order_parts_service/dto"
	order_part_service "github.com/Mirsadikovv/idoctor_platform/src/module/order_parts_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type orderPartHandler struct {
	db               *gorm.DB
	log              logger.Logger
	authMiddleware   *auth_middleware.AuthMiddleware
	orderPartService order_part_service.OrderPartService
}

func NewOrderPartHandler(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &orderPartHandler{
		db:               db,
		log:              log,
		authMiddleware:   authMiddleware,
		orderPartService: order_part_service.NewOrderPartService(db),
	}

	orderPartServiceMiddleware := handler.authMiddleware.BuildMiddleware()
	orderPartGroup := router.Group("/api/v1/order-part", orderPartServiceMiddleware)
	{
		orderPartGroup.GET("/:id", handler.FindByID)
		orderPartGroup.GET("/search", handler.Search)
		orderPartGroup.GET("/page", handler.Page)
		orderPartGroup.POST("", handler.Create)
		orderPartGroup.PUT("/:id", handler.Update)
		orderPartGroup.DELETE("/:id", handler.DeleteOrRestore)
	}
}

// Search godoc
// @Summary      Get all order parts with search
// @Description  Get all order parts with search
// @Tags 		 order-part
// @ID           search-order-part
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        params query order_part_dto.OrderPartParams false "params"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} []order_part_dto.OrderPartResponse "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order-part/search [get]
func (h *orderPartHandler) Search(c echo.Context) error {
	req := request.RequestWithData[any](c)

	var params order_part_dto.OrderPartParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if params.OrderId != nil {
			tx = tx.Where("order_parts.order_id = ?", *params.OrderId)
		}

		if params.PartId != nil {
			tx = tx.Where("order_parts.part_id = ?", *params.PartId)
		}

		if params.SupplierId != nil {
			tx = tx.Where("order_parts.supplier_id = ?", *params.SupplierId)
		}

		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("order_parts.deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Order("order_parts.id DESC")
	}

	orderParts, err := h.orderPartService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(orderParts)
}

// Page 			godoc
// @Summary			order-part-service-page
// @Description		order-part-service-page
// @Tags 			order-part
// @Param       	filter query order_part_dto.OrderPartParams false "Filter parameters"
// @Param       	page query int false "Page number"
// @Param       	perpage query int false "Items per page"
// @Security     	ApiKeyAuth
// @Success	200 	{object} order_part_dto.OrderPartPage "Successful operation"
// @Failure	400 	{object} response.HttpSuccess "Bad request"
// @Failure	500 	{object} response.HttpSuccess "Internal server error"
// @Router			/order-part/page [get]
func (h *orderPartHandler) Page(c echo.Context) error {
	req := request.Request(c)

	var params order_part_dto.OrderPartParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	tx := func(tx *gorm.DB) *gorm.DB {
		if params.OrderId != nil {
			tx = tx.Where("order_parts.order_id = ?", *params.OrderId)
		}

		if params.PartId != nil {
			tx = tx.Where("order_parts.part_id = ?", *params.PartId)
		}

		if params.SupplierId != nil {
			tx = tx.Where("order_parts.supplier_id = ?", *params.SupplierId)
		}

		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("order_parts.deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Order("order_parts.id DESC")
	}

	items, err := h.orderPartService.Page(req.Context(), req.NewPaginate(), tx)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(items)
}

// GetById godoc
// @Summary      Get order part by ID
// @Description  Get order part by ID
// @Tags 		 order-part
// @ID           get-order-part-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "order part ID"
// @Success      200 {object} order_part_dto.OrderPartResponse "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order-part/{id} [get]
func (h *orderPartHandler) FindByID(c echo.Context) error {
	req := request.RequestWithData[any](c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("order_parts.id = ?", id)
	}

	orderPart, err := h.orderPartService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(orderPart)
}

// Create godoc
// @Summary      Create order part
// @Description  Create new order part
// @Tags 		 order-part
// @ID           create-order-part
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body order_part_dto.OrderPartCreate true "order part information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order-part [POST]
func (h *orderPartHandler) Create(c echo.Context) error {
	req := request.Request(c)

	var orderPartDto order_part_dto.OrderPartCreate
	{
		if err := req.BindBody(&orderPartDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := h.orderPartService.Create(c.Request().Context(), &orderPartDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(response.NewID(id))
}

// Update godoc
// @Summary      Update order part
// @Description  Update existing order part
// @Tags 		 order-part
// @ID           update-order-part
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "order part ID"
// @Param        input body order_part_dto.OrderPartUpdate true "order part information"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order-part/{id} [PUT]
func (h *orderPartHandler) Update(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var orderPartDto order_part_dto.OrderPartUpdate
	{
		if err := req.BindBody(&orderPartDto); err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.orderPartService.Update(c.Request().Context(), id, &orderPartDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Order part updated successfully"})
}

// DeleteOrRestore godoc
// @Summary      Delete or restore order part
// @Description  Soft delete order part if active, restore if already deleted
// @Tags 		 order-part
// @ID           delete-or-restore-order-part
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "order part ID"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order-part/{id} [delete]
func (h *orderPartHandler) DeleteOrRestore(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.orderPartService.DeleteOrRestore(req.Context(), id)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Order part status toggled successfully"})
}
