import { api } from "@/plugins/axios.plugin";
import { Try } from "@/common";
import { type IdType, type PageDataType } from "@/service";
import type { keySignType } from "../utils/utils";

export type MediaType = {
	category: keySignType;
	createdAt: string;
	dirName: string;
	filename: string;
	name: string;
	mimeType: string;
	owner: string;
	primary: boolean;
	id: number;
	size: number;
	sign: string;
};

export type FileQueryType = {
	limit?: number;
	page?: number;
	employeeId?: number;
	sign?: string;
	category?: keySignType;
	owner?: number | string;
};

export type MediaPartialType = Partial<MediaType>;
export type MediaPageDataType = PageDataType<MediaType>;

class MediaService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async getFileById(id: number) {
		const { data } = await api.get<any>(`/file/download/${id}`);

		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(query: FileQueryType) {
		const { data } = await api.get<MediaPageDataType>(`/file/page?`, {
			params: query,
		});

		return data;
	}
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async pageDocx(query: string, objQuery: string) {
		const { data } = await api.get<MediaPageDataType>(`/file/page?${query}&${objQuery}`);

		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async getMedia(category: string, owner: string) {
		const { data } = await api.get<MediaType>(`/file/${category}/owner/${owner}`);

		return data;
	}

	@Try({
		async onSuccess(result) {
			(await import("@/common/Notify")).SuccesNotify(result.statusText);
		},
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async create(category: keySignType, owner: number, formdata: FormData) {
		formdata.append("category", category);
		formdata.append("owner", owner.toString());
		const { data } = await api.post<IdType>(`/file/create`, formdata);

		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async getById(category: string, sign: string, id: number) {
		const { data } = await api.get(`/file/${category}/${sign}/${id}`);
		return data;
	}

	@Try({
		async onSuccess(result) {
			(await import("@/common/Notify")).SuccesNotify(result.statusText);
		},
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async delete(id: number) {
		const res = await api.delete(`/file/${id}`);

		return res;
	}
	// replace
	@Try({
		async onSuccess(result) {
			(await import("@/common/Notify")).SuccesNotify(result.statusText);
		},
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async update(id: number, formdata: FormData) {
		const res = await api.patch<MediaPartialType>(`/file/${id}/replace`, formdata);

		return res;
	}
}

const mediaService = new MediaService();

export { mediaService as MediaService };
