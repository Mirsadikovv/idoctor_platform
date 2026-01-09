// Enums
export enum OrderStatus {
	PENDING = 'pending',
	IN_PROGRESS = 'in_progress',
	COMPLETED = 'completed',
	CANCELLED = 'cancelled',
}

export enum PaymentStatus {
	PAID = 'paid',
	UNPAID = 'unpaid',
	PARTIAL = 'partial',
}

export enum PaymentType {
	CASH = 'cash',
	CARD = 'card',
	TRANSFER = 'transfer',
	ONLINE = 'online',
}

export enum GroupByPeriod {
	DAY = 'day',
	WEEK = 'week',
	MONTH = 'month',
	YEAR = 'year',
}

// Request Parameters
export interface BaseDateRangeParams {
	start_date: string; // YYYY-MM-DD
	end_date: string;
}

export interface MasterStatisticsParams extends BaseDateRangeParams {}

export interface OrderStatisticsParams extends BaseDateRangeParams {
	master_id?: number;
	status?: OrderStatus;
	payment_status?: PaymentStatus;
	payment_type?: PaymentType;
}

export interface PaymentStatisticsParams extends BaseDateRangeParams {
	master_id?: number;
}

export interface RevenueByPeriodParams extends BaseDateRangeParams {
	master_id?: number;
	group_by?: GroupByPeriod;
}

export interface TopPartsStatisticsParams extends BaseDateRangeParams {
	master_id?: number;
	limit?: number;
}

// Response DTOs
export interface MasterStatistics {
	master_id: number;
	master_name: string;
	total_orders: number;
	completed_orders: number;
	total_revenue: number;
	average_price: number;
}

export interface OrderStatistics {
	total_orders: number;
	completed_orders: number;
	in_progress_orders: number;
	cancelled_orders: number;
	paid_orders: number;
	unpaid_orders: number;
	total_revenue: number;
	total_cost: number;
	total_profit: number;
	average_order_price: number;
}

export interface PaymentStatistics {
	payment_type: string;
	order_count: number;
	paid_count: number;
	unpaid_count: number;
	total_amount: number;
}

export interface RevenueByPeriod {
	period: string;
	orders: number;
	revenue: number;
	cost: number;
	profit: number;
}

export interface TopPartsStatistics {
	part_id: number;
	part_name: string;
	usage_count: number;
	total_revenue: number;
	total_cost: number;
	total_profit: number;
}

export type MasterStatisticsResponse = MasterStatistics[];
export type PaymentStatisticsResponse = PaymentStatistics[];
export type RevenueByPeriodResponse = RevenueByPeriod[];
export type TopPartsStatisticsResponse = TopPartsStatistics[];
