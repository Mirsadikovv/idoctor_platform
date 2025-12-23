import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { IdType, PageDataType } from "@/service";

export type OrderType = {
	id: number;
	client_id?: number;
	master_id?: number;
	part_ids?: number[];
	problem_ids?: number[];
	price: number;
	status: string;
	payment_status: string;
	payment_type: string;
	created_at: string;
	deleted_at?: string;
	value?: string | number;
	label?: string;
};

export type OrderCreateType = {
	client_id?: number;
	master_id?: number;
	part_ids?: number[];
	problem_ids?: number[];
	price: number;
	status: string;
	payment_status: string;
	payment_type: string;
};

export type OrderUpdateType = {
	client_id?: number;
	master_id?: number;
	part_ids?: number[];
	problem_ids?: number[];
	price: number;
	status: string;
	payment_status: string;
	payment_type: string;
};

export type OrderPartialType = Partial<OrderType>;
export type OrderPageData = PageDataType<OrderType>;

export type OrderSearchParams = {
	client_id?: number;
	master_id?: number;
	min_price?: number;
	max_price?: number;
	status?: string;
	payment_status?: string;
	payment_type?: string;
	include_deleted?: boolean;
	only_deleted?: boolean;
	page?: number;
	perpage?: number;
	limit?: number;
};

class OrderService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(params: OrderSearchParams = {}) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<OrderPageData>(`/order/page?${searchParams}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async pageWithDelete(params: string) {
		const { data } = await api.get<OrderPageData>(`/order/page?${params}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async getByID(id: number) {
		const { data } = await api.get<OrderType>(`/order/${id}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async search(params: OrderSearchParams = {}) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<OrderType[]>(`/order/search?${searchParams}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async create(order: OrderCreateType) {
		const { data } = await api.post<IdType>(`/order`, order);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async update(id: number, order: OrderUpdateType) {
		await api.put(`/order/${id}`, order);
		return true;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async delete(id: number) {
		await api.delete(`/order/${id}`);
		return true;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async restore(id: number) {
		await api.delete(`/order/${id}`);
		return true;
	}
}

const orderService = new OrderService();

export { orderService as OrderService };
