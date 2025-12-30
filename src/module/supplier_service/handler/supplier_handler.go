package supplier_handler

import (
	"fmt"

	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	supplier_dto "github.com/Mirsadikovv/idoctor_platform/src/module/supplier_service/dto"
	supplier_service "github.com/Mirsadikovv/idoctor_platform/src/module/supplier_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type supplierHandler struct {
	db              *gorm.DB
	log             logger.Logger
	authMiddleware  *auth_middleware.AuthMiddleware
	supplierService supplier_service.SupplierService
}

func NewSupplierHandler(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &supplierHandler{
		db:              db,
		log:             log,
		authMiddleware:  authMiddleware,
		supplierService: supplier_service.NewSupplierService(db),
	}

	supplierServiceMiddleware := handler.authMiddleware.BuildMiddleware()
	supplierGroup := router.Group("/api/v1/supplier", supplierServiceMiddleware)
	{
		supplierGroup.POST("", handler.Create)
		supplierGroup.PUT("/:id", handler.Update)
		supplierGroup.GET("/:id", handler.FindByID)
		supplierGroup.GET("/page", handler.Page)
		supplierGroup.GET("/search", handler.Search)
		supplierGroup.DELETE("/:id", handler.DeleteOrRestore)
	}
}

// Search godoc
// @Summary      Get all suppliers with search
// @Description  Get all suppliers with search
// @Tags 		 supplier
// @ID           search-supplier
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        params query supplier_dto.SupplierParams false "params"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} []supplier_dto.Supplier "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /supplier/search [get]
func (h *supplierHandler) Search(c echo.Context) error {
	req := request.RequestWithData[any](c)

	var params supplier_dto.SupplierParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if params.Name != nil {
			name := fmt.Sprintf("%%%s%%", *params.Name)
			tx = tx.Where("suppliers.name ILIKE ?", name)
		}

		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("suppliers.deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Select(
			"suppliers.id",
			"suppliers.name",
			"suppliers.created_at",
			"suppliers.deleted_at",
		).Order("suppliers.id DESC")
	}

	suppliers, err := h.supplierService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(suppliers)
}

// Page 			godoc
// @Summary			supplier-service-page
// @Description		supplier-service-page
// @Tags 			supplier
// @Param       	filter query supplier_dto.SupplierParams false "Filter parameters"
// @Param       	page query int false "Page number"
// @Param       	perpage query int false "Items per page"
// @Security     	ApiKeyAuth
// @Success	200 	{object} supplier_dto.SupplierPage "Successful operation"
// @Failure	400 	{object} response.HttpSuccess "Bad request"
// @Failure	500 	{object} response.HttpSuccess "Internal server error"
// @Router			/supplier/page [get]
func (h *supplierHandler) Page(c echo.Context) error {
	req := request.Request(c)

	var params supplier_dto.SupplierParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	tx := func(tx *gorm.DB) *gorm.DB {
		if params.Name != nil {
			tx = tx.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *params.Name))
		}

		// Handle soft delete filtering
		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Select(
			"suppliers.id",
			"suppliers.name",
			"suppliers.created_at",
			"suppliers.deleted_at",
		).Order("suppliers.id DESC")
	}

	items, err := h.supplierService.Page(req.Context(), req.NewPaginate(), tx)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(items)
}

// GetById godoc
// @Summary      Get supplier by ID
// @Description  Get supplier by ID
// @Tags 		 supplier
// @ID           get-supplier-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "supplier ID"
// @Success      200 {object} supplier_dto.Supplier "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /supplier/{id} [get]
func (h *supplierHandler) FindByID(c echo.Context) error {
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

	supplier, err := h.supplierService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(supplier)
}

// Create godoc
// @Summary      Create supplier
// @Description  Create new supplier
// @Tags 		 supplier
// @ID           create-supplier
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body supplier_dto.SupplierCreate true "supplier information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /supplier [POST]
func (h *supplierHandler) Create(c echo.Context) error {
	req := request.Request(c)

	var supplierDto supplier_dto.SupplierCreate
	{
		if err := req.BindBody(&supplierDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := h.supplierService.Create(c.Request().Context(), &supplierDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(response.NewID(id))
}

// Update godoc
// @Summary      Update supplier
// @Description  Update existing supplier
// @Tags 		 supplier
// @ID           update-supplier
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "supplier ID"
// @Param        input body supplier_dto.SupplierUpdate true "supplier information"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /supplier/{id} [PUT]
func (h *supplierHandler) Update(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var supplierDto supplier_dto.SupplierUpdate
	{
		if err := req.BindBody(&supplierDto); err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.supplierService.Update(c.Request().Context(), id, &supplierDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Supplier updated successfully"})
}

// DeleteOrRestore godoc
// @Summary      Delete or restore supplier
// @Description  Soft delete supplier if active, restore if already deleted
// @Tags 		 supplier
// @ID           delete-or-restore-supplier
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "supplier ID"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /supplier/{id} [delete]
func (h *supplierHandler) DeleteOrRestore(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.supplierService.DeleteOrRestore(req.Context(), id)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Supplier status toggled successfully"})
}
