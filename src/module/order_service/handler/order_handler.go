package order_handler

import (
	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	order_dto "github.com/Mirsadikovv/idoctor_platform/src/module/order_service/dto"
	order_service "github.com/Mirsadikovv/idoctor_platform/src/module/order_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type orderHandler struct {
	db             *gorm.DB
	log            logger.Logger
	authMiddleware *auth_middleware.AuthMiddleware
	orderService   order_service.OrderService
}

func NewOrderHandler(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &orderHandler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		orderService:   order_service.NewOrderService(db),
	}

	orderServiceMiddleware := handler.authMiddleware.BuildMiddleware()
	orderGroup := router.Group("/api/v1/order", orderServiceMiddleware)
	{
		orderGroup.GET("/:id", handler.FindByID)
		orderGroup.GET("/search", handler.Search)
		orderGroup.GET("/page", handler.Page)
		orderGroup.POST("", handler.Create)
		orderGroup.PUT("/:id", handler.Update)
		orderGroup.DELETE("/:id", handler.DeleteOrRestore)
	}
}

// Search godoc
// @Summary      Get all orders with search
// @Description  Get all orders with search
// @Tags 		 order
// @ID           search-order
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        params query order_dto.OrderParams false "params"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} []order_dto.Order "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order/search [get]
func (h *orderHandler) Search(c echo.Context) error {
	req := request.RequestWithData[any](c)

	var params order_dto.OrderParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if params.ClientID != nil {
			tx = tx.Where("orders.client_id = ?", *params.ClientID)
		}

		if params.MasterID != nil {
			tx = tx.Where("orders.master_id = ?", *params.MasterID)
		}

		if params.Status != nil {
			tx = tx.Where("orders.status = ?", *params.Status)
		}

		if params.PaymentType != nil {
			tx = tx.Where("orders.payment_type = ?", *params.PaymentType)
		}

		if params.PaymentStatus != nil {
			tx = tx.Where("orders.payment_status = ?", *params.PaymentStatus)
		}

		if params.MinPrice != nil {
			tx = tx.Where("orders.price >= ?", *params.MinPrice)
		}

		if params.MaxPrice != nil {
			tx = tx.Where("orders.price <= ?", *params.MaxPrice)
		}

		// Handle soft delete filtering
		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("orders.deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Select(
			"orders.id",
			"orders.client_id",
			"orders.master_id",
			"orders.price",
			"orders.status",
			"orders.payment_type",
			"orders.payment_status",
			"orders.created_at",
			"orders.deleted_at",
		).Order("orders.id DESC")
	}

	orders, err := h.orderService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(orders)
}

// Page 			godoc
// @Summary			order-service-page
// @Description		order-service-page
// @Tags 			order
// @Param       	filter query order_dto.OrderParams false "Filter parameters"
// @Param       	page query int false "Page number"
// @Param       	perpage query int false "Items per page"
// @Security     	ApiKeyAuth
// @Success	200 	{object} order_dto.OrderPage "Successful operation"
// @Failure	400 	{object} response.HttpSuccess "Bad request"
// @Failure	500 	{object} response.HttpSuccess "Internal server error"
// @Router			/order/page [get]
func (h *orderHandler) Page(c echo.Context) error {
	req := request.Request(c)

	var params order_dto.OrderParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	tx := func(tx *gorm.DB) *gorm.DB {
		if params.ClientID != nil {
			tx = tx.Where("client_id = ?", *params.ClientID)
		}

		if params.MasterID != nil {
			tx = tx.Where("master_id = ?", *params.MasterID)
		}

		if params.Status != nil {
			tx = tx.Where("status = ?", *params.Status)
		}

		if params.PaymentType != nil {
			tx = tx.Where("payment_type = ?", *params.PaymentType)
		}

		if params.PaymentStatus != nil {
			tx = tx.Where("payment_status = ?", *params.PaymentStatus)
		}

		if params.MinPrice != nil {
			tx = tx.Where("price >= ?", *params.MinPrice)
		}

		if params.MaxPrice != nil {
			tx = tx.Where("price <= ?", *params.MaxPrice)
		}

		// Handle soft delete filtering
		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Select(
			"orders.id",
			"orders.client_id",
			"orders.master_id",
			"orders.price",
			"orders.status",
			"orders.payment_type",
			"orders.payment_status",
			"orders.created_at",
			"orders.deleted_at",
		).Order("orders.id DESC")
	}

	items, err := h.orderService.Page(req.Context(), req.NewPaginate(), tx)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(items)
}

// GetById godoc
// @Summary      Get order by ID
// @Description  Get order by ID with parts and problems
// @Tags 		 order
// @ID           get-order-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "order ID"
// @Success      200 {object} order_dto.Order "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order/{id} [get]
func (h *orderHandler) FindByID(c echo.Context) error {
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

	order, err := h.orderService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(order)
}

// Create godoc
// @Summary      Create order
// @Description  Create new order with parts and problems
// @Tags 		 order
// @ID           create-order
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body order_dto.OrderCreate true "order information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order [POST]
func (h *orderHandler) Create(c echo.Context) error {
	req := request.Request(c)

	var orderDto order_dto.OrderCreate
	{
		if err := req.BindBody(&orderDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := h.orderService.Create(c.Request().Context(), &orderDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(response.NewID(id))
}

// Update godoc
// @Summary      Update order
// @Description  Update existing order with parts and problems
// @Tags 		 order
// @ID           update-order
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "order ID"
// @Param        input body order_dto.OrderUpdate true "order information"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order/{id} [PUT]
func (h *orderHandler) Update(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var orderDto order_dto.OrderUpdate
	{
		if err := req.BindBody(&orderDto); err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.orderService.Update(c.Request().Context(), id, &orderDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Order updated successfully"})
}

// DeleteOrRestore godoc
// @Summary      Delete or restore order
// @Description  Soft delete order if active, restore if already deleted
// @Tags 		 order
// @ID           delete-or-restore-order
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "order ID"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /order/{id} [delete]
func (h *orderHandler) DeleteOrRestore(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.orderService.DeleteOrRestore(req.Context(), id)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Order status toggled successfully"})
}
