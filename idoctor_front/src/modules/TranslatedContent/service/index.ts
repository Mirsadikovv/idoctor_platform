import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { LanguagePartialType, PageDataType } from "@/service";

export type SelectableRoute = {
	path: string;
	name: string;
	meta: {
		title: string;
		activeLinkGroup?: string;
		isEmpty?: boolean;
	};
	label: string;
};

export type LanguageContnetType = {
	id: number;
	key: string;
	value: string;
	category: string;
};

type LanguageResult = {
	category: string;
	languageId: number;
};

export interface LanguageContent {
	id: number;
	key: string;
	value: string;
	languageId: number;
	language: LanguagePartialType;
}

export interface LangContentCreateOrUpdate {
	key: string;
	category: string;
	languageId: number;
	contents:
		| {
				category: string;
				key: string;
				languageId: number;
				value: string;
				langName: string;
				LangDescp: string;
		  }[]
		| []; // Allow for an empty array
}

export type LanguageContentPartialType = Partial<LanguageContent>;
export type LanguageContentPageData = PageDataType<LanguageContent>;

class LanguageContentService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(query: string = "", langID: number) {
		const { data } = await api.get<LanguageContentPageData>(
			`/language-content/page/${langID}?${query}`,
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
	async getLangContentsKey(id: number) {
		const { data } = await api.get<LanguageContentPartialType[]>(`/language-content/${id}`);

		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async getLanguage(languageId: number, category: string = "GLOBAL") {
		const { data } = await api.get<Record<string, string>>(
			`/language-content/${languageId}/${category}`,
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
	async getLanguageContents(languageId: number) {
		const { data } = await api.get<Record<string, string>>(`/language-content/${languageId}`);

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
		const { data } = await api.get<LanguageContnetType[]>(`/language/search?q=${q}`);

		return data;
	}
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async getByKey(key: string = "") {
		const { data } = await api.get<LangContentCreateOrUpdate>(
			`/language-content/by-key/${key}`,
		);

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
	async createOrUpdate(languageContentsCreateOrUpdate: LangContentCreateOrUpdate) {
		const { data } = await api.put<LanguageResult>(
			`/language-content`,
			languageContentsCreateOrUpdate,
		);

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
	async update(id: number, language: LanguageContentPartialType) {
		await api.patch(`/language/${id}`, language);
		return true;
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
	async delete(_languageId: number, key: string) {
		await api.delete(`/language-content/${key}`);
		return true;
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
	async restore(id: number) {
		await api.patch(`/language/${id}/restore`);
		return true;
	}
}

const languageContentService = new LanguageContentService();

export { languageContentService as LanguageContentService };
