package service

import (
	"context"
	"fmt"

	statistic_dto "github.com/Mirsadikovv/idoctor_platform/src/module/statistic_service/dto"
	"gorm.io/gorm"
)

type StatisticService interface {
	GetOrderStatistics(ctx context.Context, filter *statistic_dto.StatisticsFilter) (*statistic_dto.OrderStatistics, error)
	GetRevenueByPeriod(ctx context.Context, filter *statistic_dto.StatisticsFilter) ([]statistic_dto.RevenueByPeriod, error)
	GetMasterStatistics(ctx context.Context, filter *statistic_dto.StatisticsFilter) ([]statistic_dto.MasterStatistics, error)
	GetPaymentStatistics(ctx context.Context, filter *statistic_dto.StatisticsFilter) ([]statistic_dto.PaymentStatistics, error)
	GetTopParts(ctx context.Context, filter *statistic_dto.StatisticsFilter, limit int) ([]statistic_dto.TopPartsStatistics, error)
}

type statisticService struct {
	db *gorm.DB
}

func NewStatisticService(db *gorm.DB) StatisticService {
	return &statisticService{db: db}
}

// GetOrderStatistics получить общую статистику по заказам
func (s *statisticService) GetOrderStatistics(ctx context.Context, filter *statistic_dto.StatisticsFilter) (*statistic_dto.OrderStatistics, error) {
	var stats statistic_dto.OrderStatistics

	query := s.db.WithContext(ctx).Table("orders")
	query = s.applyFilters(query, filter)

	// Основной запрос со всеми агрегациями
	err := query.
		Select(`
			COUNT(*) as total_orders,
			COUNT(CASE WHEN status = 'completed' THEN 1 END) as completed_orders,
			COUNT(CASE WHEN status IN ('in_progress', 'pending') THEN 1 END) as in_progress_orders,
			COUNT(CASE WHEN status = 'cancelled' THEN 1 END) as cancelled_orders,
			COALESCE(SUM(price), 0) as total_revenue,
			COALESCE(AVG(price), 0) as average_order_price,
			COUNT(CASE WHEN payment_status = 'paid' THEN 1 END) as paid_orders,
			COUNT(CASE WHEN payment_status = 'unpaid' THEN 1 END) as unpaid_orders
		`).
		Scan(&stats).Error

	if err != nil {
		return nil, err
	}

	// Получаем общие затраты отдельным запросом
	var totalCost float64

	// Создаем подзапрос с теми же фильтрами что и для основного запроса
	orderIdsSubquery := s.db.Table("orders").Select("id")
	orderIdsSubquery = s.applyFilters(orderIdsSubquery, filter)

	costQuery := s.db.WithContext(ctx).Table("order_parts").
		Select("COALESCE(SUM(income_price), 0)").
		Where("order_id IN (?)", orderIdsSubquery).
		Where("deleted_at IS NULL")

	costQuery.Scan(&totalCost)

	stats.TotalCost = totalCost
	stats.TotalProfit = stats.TotalRevenue - stats.TotalCost

	return &stats, nil
}

// GetRevenueByPeriod получить статистику доходов по периодам
func (s *statisticService) GetRevenueByPeriod(ctx context.Context, filter *statistic_dto.StatisticsFilter) ([]statistic_dto.RevenueByPeriod, error) {
	var results []statistic_dto.RevenueByPeriod

	groupBy := "day"
	if filter != nil && filter.GroupBy != nil {
		groupBy = *filter.GroupBy
	}

	var dateFormat string
	switch groupBy {
	case "day":
		dateFormat = "DATE(created_at)"
	case "week":
		dateFormat = "DATE_TRUNC('week', created_at)"
	case "month":
		dateFormat = "DATE_TRUNC('month', created_at)"
	case "year":
		dateFormat = "DATE_TRUNC('year', created_at)"
	default:
		dateFormat = "DATE(created_at)"
	}

	query := s.db.WithContext(ctx).Table("orders")
	query = s.applyFilters(query, filter)

	// Подзапрос для получения затрат по каждому заказу
	costSubquery := s.db.Table("order_parts").
		Select("order_id, SUM(income_price) as cost").
		Where("deleted_at IS NULL").
		Group("order_id")

	err := query.
		Select(fmt.Sprintf(`
			%s as period,
			COALESCE(SUM(orders.price), 0) as revenue,
			COUNT(*) as orders
		`, dateFormat)).
		Group("period").
		Order("period ASC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// Для каждого периода получаем затраты
	for i := range results {
		var cost float64
		costQuery := s.db.WithContext(ctx).Table("orders").
			Joins("LEFT JOIN (?) as op ON orders.id = op.order_id", costSubquery).
			Where(fmt.Sprintf("%s = ?", dateFormat), results[i].Period)

		// Применяем те же фильтры что и для orders
		costQuery = s.applyFilters(costQuery, filter)

		costQuery.Select("COALESCE(SUM(op.cost), 0)").Scan(&cost)

		results[i].Cost = cost
		results[i].Profit = results[i].Revenue - results[i].Cost
	}

	return results, nil
}

// GetMasterStatistics получить статистику по мастерам
func (s *statisticService) GetMasterStatistics(ctx context.Context, filter *statistic_dto.StatisticsFilter) ([]statistic_dto.MasterStatistics, error) {
	var results []statistic_dto.MasterStatistics

	query := s.db.WithContext(ctx).Table("orders")
	query = s.applyFilters(query, filter)

	err := query.
		Joins("LEFT JOIN users ON orders.master_id = users.id").
		Select(`
			orders.master_id,
			COALESCE(users.username, 'Unknown') as master_name,
			COUNT(*) as total_orders,
			COUNT(CASE WHEN orders.status = 'completed' THEN 1 END) as completed_orders,
			COALESCE(SUM(orders.price), 0) as total_revenue,
			COALESCE(AVG(orders.price), 0) as average_price
		`).
		Where("orders.master_id IS NOT NULL").
		Group("orders.master_id, users.username").
		Order("total_revenue DESC").
		Scan(&results).Error

	return results, err
}

// GetPaymentStatistics получить статистику по платежам
func (s *statisticService) GetPaymentStatistics(ctx context.Context, filter *statistic_dto.StatisticsFilter) ([]statistic_dto.PaymentStatistics, error) {
	var results []statistic_dto.PaymentStatistics

	query := s.db.WithContext(ctx).Table("orders")
	query = s.applyFilters(query, filter)

	err := query.
		Select(`
			payment_type,
			COALESCE(SUM(price), 0) as total_amount,
			COUNT(*) as order_count,
			COUNT(CASE WHEN payment_status = 'paid' THEN 1 END) as paid_count,
			COUNT(CASE WHEN payment_status = 'unpaid' THEN 1 END) as unpaid_count
		`).
		Group("payment_type").
		Order("total_amount DESC").
		Scan(&results).Error

	return results, err
}

// GetTopParts получить статистику по самым используемым запчастям
func (s *statisticService) GetTopParts(ctx context.Context, filter *statistic_dto.StatisticsFilter, limit int) ([]statistic_dto.TopPartsStatistics, error) {
	var results []statistic_dto.TopPartsStatistics

	if limit <= 0 {
		limit = 10
	}

	query := s.db.WithContext(ctx).Table("order_parts")

	// Применяем фильтры через join с orders
	if filter != nil {
		if filter.StartDate != nil && filter.EndDate != nil {
			query = query.Where("order_id IN (?)",
				s.db.Table("orders").
					Select("id").
					Where("created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate),
			)
		}
		if filter.MasterId != nil {
			query = query.Where("order_id IN (?)",
				s.db.Table("orders").
					Select("id").
					Where("master_id = ?", *filter.MasterId),
			)
		}
	}

	err := query.
		Joins("LEFT JOIN parts ON order_parts.part_id = parts.id").
		Select(`
			order_parts.part_id,
			parts.name as part_name,
			COUNT(*) as usage_count,
			COALESCE(SUM(order_parts.price), 0) as total_revenue,
			COALESCE(SUM(order_parts.income_price), 0) as total_cost,
			COALESCE(SUM(order_parts.price - order_parts.income_price), 0) as total_profit
		`).
		Where("order_parts.deleted_at IS NULL").
		Group("order_parts.part_id, parts.name").
		Order("usage_count DESC").
		Limit(limit).
		Scan(&results).Error

	return results, err
}

// applyFilters применяет фильтры к запросу
func (s *statisticService) applyFilters(query *gorm.DB, filter *statistic_dto.StatisticsFilter) *gorm.DB {
	if filter == nil {
		return query.Where("deleted_at IS NULL")
	}

	query = query.Where("deleted_at IS NULL")

	if filter.StartDate != nil && filter.EndDate != nil {
		query = query.Where("created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	if filter.MasterId != nil {
		query = query.Where("master_id = ?", *filter.MasterId)
	}

	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}

	if filter.PaymentStatus != nil {
		query = query.Where("payment_status = ?", *filter.PaymentStatus)
	}

	if filter.PaymentType != nil {
		query = query.Where("payment_type = ?", *filter.PaymentType)
	}

	return query
}
