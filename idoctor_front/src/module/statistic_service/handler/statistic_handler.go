package handler

import (
	"strconv"

	auth_middleware "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/middleware"
	statistic_dto "github.com/Mirsadikovv/idoctor_platform/src/module/statistic_service/dto"
	statistic_service "github.com/Mirsadikovv/idoctor_platform/src/module/statistic_service/service"

	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type statisticHandler struct {
	db               *gorm.DB
	log              logger.Logger
	authMiddleware   *auth_middleware.AuthMiddleware
	statisticService statistic_service.StatisticService
}

func NewStatisticHandler(router *echo.Echo, db *gorm.DB, log logger.Logger, authMiddleware *auth_middleware.AuthMiddleware) {
	handler := &statisticHandler{
		db:               db,
		log:              log,
		authMiddleware:   authMiddleware,
		statisticService: statistic_service.NewStatisticService(db),
	}

	statisticServiceMiddleware := handler.authMiddleware.BuildMiddleware()
	statisticGroup := router.Group("/api/v1/statistics", statisticServiceMiddleware)
	{
		statisticGroup.GET("/orders", handler.GetOrderStatistics)
		statisticGroup.GET("/revenue", handler.GetRevenueByPeriod)
		statisticGroup.GET("/masters", handler.GetMasterStatistics)
		statisticGroup.GET("/payments", handler.GetPaymentStatistics)
		statisticGroup.GET("/top-parts", handler.GetTopParts)
	}
}

// GetOrderStatistics godoc
// @Summary      Get order statistics
// @Description  Get overall order statistics including revenue, costs, and profit
// @Tags         statistics
// @ID           get-order-statistics
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        start_date query string false "Start date (YYYY-MM-DD)"
// @Param        end_date query string false "End date (YYYY-MM-DD)"
// @Param        master_id query int false "Master ID"
// @Param        status query string false "Order status"
// @Param        payment_status query string false "Payment status"
// @Param        payment_type query string false "Payment type"
// @Success      200 {object} statistic_dto.OrderStatistics "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /statistics/orders [get]
func (h *statisticHandler) GetOrderStatistics(c echo.Context) error {
	req := request.Request(c)

	var filter statistic_dto.StatisticsFilter
	{
		if err := req.BindQuery(&filter); err != nil {
			return req.BadRequest(err)
		}
	}

	stats, err := h.statisticService.GetOrderStatistics(req.Context(), &filter)
	{
		if err != nil {
			h.log.Error(err)
			return req.BadRequest(err)
		}
	}

	return req.OK(stats)
}

// GetRevenueByPeriod godoc
// @Summary      Get revenue by period
// @Description  Get revenue statistics grouped by period (day, week, month, year)
// @Tags         statistics
// @ID           get-revenue-by-period
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        start_date query string false "Start date (YYYY-MM-DD)"
// @Param        end_date query string false "End date (YYYY-MM-DD)"
// @Param        master_id query int false "Master ID"
// @Param        group_by query string false "Group by period (day, week, month, year)" default(day)
// @Success      200 {array} statistic_dto.RevenueByPeriod "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /statistics/revenue [get]
func (h *statisticHandler) GetRevenueByPeriod(c echo.Context) error {
	req := request.Request(c)

	var filter statistic_dto.StatisticsFilter
	{
		if err := req.BindQuery(&filter); err != nil {
			return req.BadRequest(err)
		}
	}

	results, err := h.statisticService.GetRevenueByPeriod(req.Context(), &filter)
	{
		if err != nil {
			h.log.Error(err)
			return req.BadRequest(err)
		}
	}

	return req.OK(results)
}

// GetMasterStatistics godoc
// @Summary      Get master statistics
// @Description  Get statistics for each master including orders count and revenue
// @Tags         statistics
// @ID           get-master-statistics
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        start_date query string false "Start date (YYYY-MM-DD)"
// @Param        end_date query string false "End date (YYYY-MM-DD)"
// @Success      200 {array} statistic_dto.MasterStatistics "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /statistics/masters [get]
func (h *statisticHandler) GetMasterStatistics(c echo.Context) error {
	req := request.Request(c)

	var filter statistic_dto.StatisticsFilter
	{
		if err := req.BindQuery(&filter); err != nil {
			return req.BadRequest(err)
		}
	}

	results, err := h.statisticService.GetMasterStatistics(req.Context(), &filter)
	{
		if err != nil {
			h.log.Error(err)
			return req.BadRequest(err)
		}
	}

	return req.OK(results)
}

// GetPaymentStatistics godoc
// @Summary      Get payment statistics
// @Description  Get statistics grouped by payment type
// @Tags         statistics
// @ID           get-payment-statistics
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        start_date query string false "Start date (YYYY-MM-DD)"
// @Param        end_date query string false "End date (YYYY-MM-DD)"
// @Param        master_id query int false "Master ID"
// @Success      200 {array} statistic_dto.PaymentStatistics "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /statistics/payments [get]
func (h *statisticHandler) GetPaymentStatistics(c echo.Context) error {
	req := request.Request(c)

	var filter statistic_dto.StatisticsFilter
	{
		if err := req.BindQuery(&filter); err != nil {
			return req.BadRequest(err)
		}
	}

	results, err := h.statisticService.GetPaymentStatistics(req.Context(), &filter)
	{
		if err != nil {
			h.log.Error(err)
			return req.BadRequest(err)
		}
	}

	return req.OK(results)
}

// GetTopParts godoc
// @Summary      Get top parts statistics
// @Description  Get statistics for most used parts
// @Tags         statistics
// @ID           get-top-parts
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        start_date query string false "Start date (YYYY-MM-DD)"
// @Param        end_date query string false "End date (YYYY-MM-DD)"
// @Param        master_id query int false "Master ID"
// @Param        limit query int false "Limit results" default(10)
// @Success      200 {array} statistic_dto.TopPartsStatistics "Successful operation"
// @Failure      400 {object} response.HttpSuccess "Bad request"
// @Failure      500 {object} response.HttpSuccess "Internal server error"
// @Router       /statistics/top-parts [get]
func (h *statisticHandler) GetTopParts(c echo.Context) error {
	req := request.Request(c)

	var filter statistic_dto.StatisticsFilter
	{
		if err := req.BindQuery(&filter); err != nil {
			return req.BadRequest(err)
		}
	}

	limit := 10
	if limitParam := c.QueryParam("limit"); limitParam != "" {
		if val, err := strconv.Atoi(limitParam); err == nil {
			limit = val
		}
	}

	results, err := h.statisticService.GetTopParts(req.Context(), &filter, limit)
	{
		if err != nil {
			h.log.Error(err)
			return req.BadRequest(err)
		}
	}

	return req.OK(results)
}
