package dto

import "time"

// OrderStatistics общая статистика по заказам
type OrderStatistics struct {
	TotalOrders       int64   `json:"total_orders"`        // Общее количество заказов
	CompletedOrders   int64   `json:"completed_orders"`    // Завершенные заказы
	InProgressOrders  int64   `json:"in_progress_orders"`  // Заказы в работе
	CancelledOrders   int64   `json:"cancelled_orders"`    // Отмененные заказы
	TotalRevenue      float64 `json:"total_revenue"`       // Общая выручка (orders.price)
	TotalCost         float64 `json:"total_cost"`          // Общие расходы (order_parts.income_price)
	TotalProfit       float64 `json:"total_profit"`        // Общая прибыль (revenue - cost)
	AverageOrderPrice float64 `json:"average_order_price"` // Средняя цена заказа
	PaidOrders        int64   `json:"paid_orders"`         // Оплаченные заказы
	UnpaidOrders      int64   `json:"unpaid_orders"`       // Неоплаченные заказы
} // @name OrderStatistics

// RevenueByPeriod статистика доходов по периодам
type RevenueByPeriod struct {
	Period   string  `json:"period"`   // Период (дата или месяц)
	Revenue  float64 `json:"revenue"`  // Выручка за период
	Cost     float64 `json:"cost"`     // Расходы за период
	Profit   float64 `json:"profit"`   // Прибыль за период
	Orders   int64   `json:"orders"`   // Количество заказов
} // @name RevenueByPeriod

// MasterStatistics статистика по мастерам
type MasterStatistics struct {
	MasterId      int64   `json:"master_id"`
	MasterName    string  `json:"master_name"`
	TotalOrders   int64   `json:"total_orders"`
	CompletedOrders int64 `json:"completed_orders"`
	TotalRevenue  float64 `json:"total_revenue"`
	AveragePrice  float64 `json:"average_price"`
} // @name MasterStatistics

// PaymentStatistics статистика по платежам
type PaymentStatistics struct {
	PaymentType   string  `json:"payment_type"`   // Тип оплаты (cash, card, etc.)
	TotalAmount   float64 `json:"total_amount"`   // Общая сумма
	OrderCount    int64   `json:"order_count"`    // Количество заказов
	PaidCount     int64   `json:"paid_count"`     // Количество оплаченных
	UnpaidCount   int64   `json:"unpaid_count"`   // Количество неоплаченных
} // @name PaymentStatistics

// StatisticsFilter фильтр для получения статистики
type StatisticsFilter struct {
	StartDate     *time.Time `json:"start_date" query:"start_date"`
	EndDate       *time.Time `json:"end_date" query:"end_date"`
	MasterId      *int64     `json:"master_id" query:"master_id"`
	Status        *string    `json:"status" query:"status"`
	PaymentStatus *string    `json:"payment_status" query:"payment_status"`
	PaymentType   *string    `json:"payment_type" query:"payment_type"`
	GroupBy       *string    `json:"group_by" query:"group_by"` // day, week, month, year
} // @name StatisticsFilter

// TopPartsStatistics статистика по самым используемым запчастям
type TopPartsStatistics struct {
	PartId       int64   `json:"part_id"`
	PartName     string  `json:"part_name"`
	UsageCount   int64   `json:"usage_count"`   // Количество использований
	TotalRevenue float64 `json:"total_revenue"` // Общая выручка от продаж
	TotalCost    float64 `json:"total_cost"`    // Общие затраты
	TotalProfit  float64 `json:"total_profit"`  // Общая прибыль
} // @name TopPartsStatistics

// DeviceStatistics статистика по устройствам
type DeviceStatistics struct {
	DeviceId     int64  `json:"device_id"`
	DeviceName   string `json:"device_name"`
	RepairCount  int64  `json:"repair_count"`  // Количество ремонтов
	AveragePrice float64 `json:"average_price"` // Средняя стоимость ремонта
} // @name DeviceStatistics
