import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { IdType, PageDataType } from "@/service";

export type SupplierType = {
	id: number;
	name: string;
	created_at: string;
	deleted_at?: string;
	value?: string | number;
	label?: string;
};

export type SupplierCreateType = {
	name: string;
};

export type SupplierUpdateType = {
	name: string;
};

export type SupplierPartialType = Partial<SupplierType>;
export type SupplierPageData = PageDataType<SupplierType>;

export type SupplierSearchParams = {
	name?: string;
	include_deleted?: boolean;
	only_deleted?: boolean;
	page?: number;
	perpage?: number;
	limit?: number;
};

class SupplierService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(params: SupplierSearchParams = {}) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<SupplierPageData>(`/supplier/page?${searchParams}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async pageWithDelete(params: SupplierSearchParams = {}) {
		const searchParams = new URLSearchParams();
		searchParams.append("include_deleted", "true");
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined && key !== "include_deleted") {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<SupplierPageData>(`/supplier/page?${searchParams}`);
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
		const { data } = await api.get<SupplierType>(`/supplier/${id}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async search(params: string) {
		const { data } = await api.get<SupplierType[]>(`/supplier/search?${params}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async create(supplier: SupplierCreateType) {
		const { data } = await api.post<IdType>(`/supplier`, supplier);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async update(id: number, supplier: SupplierUpdateType) {
		await api.put(`/supplier/${id}`, supplier);
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
		await api.delete(`/supplier/${id}`);
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
		await api.delete(`/supplier/${id}`);
		return true;
	}
}

const supplierService = new SupplierService();

export { supplierService as SupplierService };
