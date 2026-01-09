import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type {
	MasterStatisticsParams,
	MasterStatisticsResponse,
	OrderStatisticsParams,
	OrderStatistics,
	PaymentStatisticsParams,
	PaymentStatisticsResponse,
	RevenueByPeriodParams,
	RevenueByPeriodResponse,
	TopPartsStatisticsParams,
	TopPartsStatisticsResponse,
} from "./types";

class StatisticsService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data?.message || err.message || "Ошибка загрузки статистики мастеров",
			);
		},
	})
	async getMasterStatistics(params: MasterStatisticsParams) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<MasterStatisticsResponse>(
			`/statistics/masters?${searchParams}`,
		);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data?.message || err.message || "Ошибка загрузки статистики заказов",
			);
		},
	})
	async getOrderStatistics(params: OrderStatisticsParams) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<OrderStatistics>(`/statistics/orders?${searchParams}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data?.message || err.message || "Ошибка загрузки статистики платежей",
			);
		},
	})
	async getPaymentStatistics(params: PaymentStatisticsParams) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<PaymentStatisticsResponse>(
			`/statistics/payments?${searchParams}`,
		);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data?.message || err.message || "Ошибка загрузки статистики дохода",
			);
		},
	})
	async getRevenueByPeriod(params: RevenueByPeriodParams) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<RevenueByPeriodResponse>(
			`/statistics/revenue?${searchParams}`,
		);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data?.message ||
					err.message ||
					"Ошибка загрузки статистики популярных деталей",
			);
		},
	})
	async getTopParts(params: TopPartsStatisticsParams) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<TopPartsStatisticsResponse>(
			`/statistics/top-parts?${searchParams}`,
		);
		return data;
	}
}

const statisticsService = new StatisticsService();

export { statisticsService as StatisticsService };
