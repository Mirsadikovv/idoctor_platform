import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { IdType, PageDataType } from "@/service";

export type PartType = {
	id: number;
	name: string;
	device_id: number;
	supplier_id: number;
	created_at: string;
	deleted_at?: string;
	value?: string | number;
	label?: string;
};

export type PartCreateType = {
	name: string;
	device_id: number;
	supplier_id: number;
};

export type PartUpdateType = {
	name: string;
	device_id: number;
	supplier_id: number;
};

export type PartPartialType = Partial<PartType>;
export type PartPageData = PageDataType<PartType>;

export type PartSearchParams = {
	name?: string;
	device_id?: number;
	supplier_id?: number;
	include_deleted?: boolean;
	only_deleted?: boolean;
	page?: number;
	perpage?: number;
	limit?: number;
};

class PartService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(params: PartSearchParams = {}) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<PartPageData>(`/part/page?${searchParams}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async pageWithDelete(params: PartSearchParams = {}) {
		const searchParams = new URLSearchParams();
		searchParams.append("include_deleted", "true");
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined && key !== "include_deleted") {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<PartPageData>(`/part/page?${searchParams}`);
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
		const { data } = await api.get<PartType>(`/part/${id}`);
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
		const { data } = await api.get<PartType[]>(`/part/search?${params}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async create(part: PartCreateType) {
		const { data } = await api.post<IdType>(`/part`, part);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async update(id: number, part: PartUpdateType) {
		await api.put(`/part/${id}`, part);
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
		await api.delete(`/part/${id}`);
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
		await api.delete(`/part/${id}`);
		return true;
	}
}

const partService = new PartService();

export { partService as PartService };
