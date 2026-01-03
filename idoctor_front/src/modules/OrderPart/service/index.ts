import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { IdType, PageDataType } from "@/service";

export type OrderPartType = {
	id: number;
	income_price: number;
	order_id: number;
	part_id: number;
	price: number;
	supplier_id: number;
};

export type OrderPartCreateType = Omit<OrderPartType, "id">;

export type OrderPartUpdateType = Omit<OrderPartType, "id">;

export type OrderPartPartialType = Partial<OrderPartType>;
export type OrderPartPageData = PageDataType<OrderPartType>;

export type OrderPartSearchParams = {
	order_id?: number;
	part_id?: number;
	supplier_id?: number;
	include_deleted?: boolean;
	only_deleted?: boolean;
	page?: number;
	perpage?: number;
	limit?: number;
};

class OrderPartService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(params: string, orderId: string | number) {
		const { data } = await api.get<OrderPartPageData>(
			`/order-part/page?order_id=${orderId}&${params}`,
		);
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
		const { data } = await api.get<OrderPartPageData>(`/order-part/page?${params}`);
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
		const { data } = await api.get<OrderPartType>(`/order-part/${id}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async search(params: OrderPartSearchParams = {}) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<OrderPartType[]>(`/order-part/search?${searchParams}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async create(order: OrderPartCreateType) {
		const { data } = await api.post<IdType>(`/order-part`, order);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async update(id: number, order: OrderPartUpdateType) {
		await api.put(`/order-part/${id}`, order);
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
		await api.delete(`/order-part/${id}`);
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
		await api.delete(`/order-part/${id}`);
		return true;
	}
}

const orderService = new OrderPartService();

export { orderService as OrderPartService };
