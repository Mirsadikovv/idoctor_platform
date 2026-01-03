import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { IdType, PageDataType } from "@/service";

export type LanguageType = {
	id: number;
	name: string;
	description: string;
	deletedAt?: string;
	value?: string | number;
	label?: string;
};

export type LanguagePartialType = Partial<LanguageType>;
export type LanguagePageData = PageDataType<LanguageType>;

class LanguageService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(query: string = "") {
		const { data } = await api.get<LanguagePageData>(`/language/page?${query}`);
		return data;
	}
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async pageWithDelete(query: string = "") {
		const { data } = await api.get<LanguagePageData>(
			`/language/page?with_deleted=true&${query}`,
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
	async getByID(id: number) {
		const { data } = await api.get<LanguageType>(`/language/${id}`);

		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async first() {
		const { data } = await api.get<LanguageType>(`/language/first`);

		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async search(q: string = "") {
		const { data } = await api.get<LanguageType[]>(`/language/search?q=${q}`);

		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async create(language: LanguageType) {
		const { data } = await api.post<IdType>(`/language`, language);

		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async update(id: number, language: LanguagePartialType) {
		await api.patch(`/language/${id}`, language);
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
		await api.delete(`/language/${id}`);
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
		await api.patch(`/language/${id}/restore`);
		return true;
	}
}

const languageService = new LanguageService();

export { languageService as LanguageService };
