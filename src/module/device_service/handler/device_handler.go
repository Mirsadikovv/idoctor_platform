package device_handler

import (
	"fmt"

	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	device_dto "github.com/Mirsadikovv/idoctor_platform/src/module/device_service/dto"
	device_service "github.com/Mirsadikovv/idoctor_platform/src/module/device_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type deviceHandler struct {
	db             *gorm.DB
	log            logger.Logger
	authMiddleware *auth_middleware.AuthMiddleware
	deviceService  device_service.DeviceService
}

func NewDeviceHandler(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &deviceHandler{
		db:             db,
		log:            log,
		authMiddleware: authMiddleware,
		deviceService:  device_service.NewDeviceService(db),
	}

	deviceServiceMiddleware := handler.authMiddleware.BuildMiddleware()
	deviceGroup := router.Group("/api/v1/device", deviceServiceMiddleware)
	{
		deviceGroup.GET("/:id", handler.FindByID)
		deviceGroup.GET("/search", handler.Search)
		deviceGroup.GET("/page", handler.Page)
		deviceGroup.POST("", handler.Create)
		deviceGroup.PUT("/:id", handler.Update)
		deviceGroup.DELETE("/:id", handler.DeleteOrRestore)
	}
}

// Search godoc
// @Summary      Get all devices with search
// @Description  Get all devices with search
// @Tags 		 device
// @ID           search-device
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        params query device_dto.DeviceParams false "params"
// @Param        limit  query int    false "Limit the number of results" default(20)
// @Success      200 {object} []device_dto.Device "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /device/search [get]
func (h *deviceHandler) Search(c echo.Context) error {
	req := request.RequestWithData[any](c)

	var params device_dto.DeviceParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if params.Name != nil {
			name := fmt.Sprintf("%%%s%%", *params.Name)
			tx = tx.Where("devices.name ILIKE ?", name)
		}

		if params.BrandName != nil {
			brandName := fmt.Sprintf("%%%s%%", *params.BrandName)
			tx = tx.Where("devices.brand_name ILIKE ?", brandName)
		}

		// Handle soft delete filtering
		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("devices.deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Select(
			"devices.id",
			"devices.name",
			"devices.brand_name",
			"devices.created_at",
			"devices.deleted_at",
		).Order("devices.id DESC")
	}

	devices, err := h.deviceService.Find(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(devices)
}

// Page 			godoc
// @Summary			device-service-page
// @Description		device-service-page
// @Tags 			device
// @Param       	filter query device_dto.DeviceParams false "Filter parameters"
// @Param       	page query int false "Page number"
// @Param       	perpage query int false "Items per page"
// @Security     	ApiKeyAuth
// @Success	200 	{object} device_dto.DevicePage "Successful operation"
// @Failure	400 	{object} response.HttpSuccess "Bad request"
// @Failure	500 	{object} response.HttpSuccess "Internal server error"
// @Router			/device/page [get]
func (h *deviceHandler) Page(c echo.Context) error {
	req := request.Request(c)

	var params device_dto.DeviceParams
	{
		if err := req.BindQuery(&params); err != nil {
			return req.BadRequest(err)
		}
	}

	tx := func(tx *gorm.DB) *gorm.DB {
		if params.Name != nil {
			tx = tx.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *params.Name))
		}

		if params.BrandName != nil {
			tx = tx.Where("brand_name ILIKE ?", fmt.Sprintf("%%%s%%", *params.BrandName))
		}

		// Handle soft delete filtering
		if params.OnlyDeleted != nil && *params.OnlyDeleted {
			tx = tx.Unscoped().Where("deleted_at IS NOT NULL")
		} else if params.IncludeDeleted != nil && *params.IncludeDeleted {
			tx = tx.Unscoped()
		}

		return tx.Select(
			"devices.id",
			"devices.name",
			"devices.brand_name",
			"devices.created_at",
			"devices.deleted_at",
		).Order("devices.id DESC")
	}

	items, err := h.deviceService.Page(req.Context(), req.NewPaginate(), tx)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(items)
}

// GetById godoc
// @Summary      Get device by ID
// @Description  Get device by ID
// @Tags 		 device
// @ID           get-device-by-id
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "device ID"
// @Success      200 {object} device_dto.Device "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /device/{id} [get]
func (h *deviceHandler) FindByID(c echo.Context) error {
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

	device, err := h.deviceService.FindOne(req.Context(), filter)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(device)
}

// Create godoc
// @Summary      Create device
// @Description  Create new device
// @Tags 		 device
// @ID           create-device
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        input body device_dto.DeviceCreate true "device information"
// @Success      201 {object} response.ID64 "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /device [POST]
func (h *deviceHandler) Create(c echo.Context) error {
	req := request.Request(c)

	var deviceDto device_dto.DeviceCreate
	{
		if err := req.BindBody(&deviceDto); err != nil {
			return req.BadRequest(err)
		}
	}

	id, err := h.deviceService.Create(c.Request().Context(), &deviceDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(response.NewID(id))
}

// Update godoc
// @Summary      Update device
// @Description  Update existing device
// @Tags 		 device
// @ID           update-device
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "device ID"
// @Param        input body device_dto.DeviceUpdate true "device information"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /device/{id} [PUT]
func (h *deviceHandler) Update(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	var deviceDto device_dto.DeviceUpdate
	{
		if err := req.BindBody(&deviceDto); err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.deviceService.Update(c.Request().Context(), id, &deviceDto)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Device updated successfully"})
}

// DeleteOrRestore godoc
// @Summary      Delete or restore device
// @Description  Soft delete device if active, restore if already deleted
// @Tags 		 device
// @ID           delete-or-restore-device
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "device ID"
// @Success      200 {object} response.HttpSuccess "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /device/{id} [delete]
func (h *deviceHandler) DeleteOrRestore(c echo.Context) error {
	req := request.Request(c)

	id, err := req.ParamToInt("id")
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	err = h.deviceService.DeleteOrRestore(req.Context(), id)
	{
		if err != nil {
			return req.BadRequest(err)
		}
	}

	return req.OK(map[string]string{"message": "Device status toggled successfully"})
}
