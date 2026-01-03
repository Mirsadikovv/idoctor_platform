import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { IdType, PageDataType } from "@/service";

export type ProblemType = {
	id: number;
	name: string;
	price: number;
	created_at: string;
	value?: string | number;
	label?: string;
};

export type ProblemCreateType = {
	name: string;
	price: number;
};

export type ProblemUpdateType = {
	name: string;
	price: number;
};

export type ProblemPartialType = Partial<ProblemType>;
export type ProblemPageData = PageDataType<ProblemType>;

export type ProblemSearchParams = {
	name?: string;
	min_price?: number;
	max_price?: number;
	page?: number;
	perpage?: number;
	limit?: number;
};

class ProblemService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(params: ProblemSearchParams = {}) {
		const searchParams = new URLSearchParams();
		Object.entries(params).forEach(([key, value]) => {
			if (value !== undefined) {
				searchParams.append(key, String(value));
			}
		});

		const { data } = await api.get<ProblemPageData>(`/problem/page?${searchParams}`);
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
		const { data } = await api.get<ProblemType>(`/problem/${id}`);
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
		const { data } = await api.get<ProblemType[]>(`/problem/search?${params}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async create(problem: ProblemCreateType) {
		const { data } = await api.post<IdType>(`/problem`, problem);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async update(id: number, problem: ProblemUpdateType) {
		const { data } = await api.put<IdType>(`/problem/${id}`, problem);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async delete(id: number) {
		await api.delete(`/problem/${id}`);
		return true;
	}
}

const problemService = new ProblemService();

export { problemService as ProblemService };
