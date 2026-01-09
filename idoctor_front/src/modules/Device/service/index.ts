import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { IdType, PageDataType } from "@/service";

export type DeviceType = {
	id: number;
	name: string;
	brand_name?: string;
	created_at: string;
	deleted_at?: string;
	value?: string | number;
	label?: string;
};

export type DeviceCreateType = {
	name: string;
	brand_name?: string;
};

export type DeviceUpdateType = {
	name: string;
	brand_name?: string;
};

export type DevicePartialType = Partial<DeviceType>;
export type DevicePageData = PageDataType<DeviceType>;

export type DeviceSearchParams = {
	name?: string;
	brand_name?: string;
	include_deleted?: boolean;
	only_deleted?: boolean;
	page?: number;
	perpage?: number;
	limit?: number;
};

class DeviceService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(params: DeviceSearchParams = {}) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<DevicePageData>(`/device/page?${searchParams}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async pageWithDelete(params: DeviceSearchParams = {}) {
		const searchParams = new URLSearchParams();
		searchParams.append("include_deleted", "true");
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined && key !== "include_deleted") {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<DevicePageData>(`/device/page?${searchParams}`);
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
		const { data } = await api.get<DeviceType>(`/device/${id}`);
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
		const { data } = await api.get<DeviceType[]>(`/device/search?${params}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async create(device: DeviceCreateType) {
		const { data } = await api.post<IdType>(`/device`, device);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async update(id: number, device: DeviceUpdateType) {
		await api.put(`/device/${id}`, device);
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
		await api.delete(`/device/${id}`);
		return true;
	}
}

const deviceService = new DeviceService();

export { deviceService as DeviceService };
