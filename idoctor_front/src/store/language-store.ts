import { type LanguageType } from "@/service";
import { defineStore } from "pinia";

export type LanguageStoreType = ReturnType<typeof useLanguageStore>;

export const useLanguageStore = defineStore("language", {
	state: () => ({
		local: {} as Record<string, string>,
		global: {} as Record<string, string>,
		_currentLang: getCurrentLang("lang") as LanguageType | null,
		_languages: [] as LanguageType[],
	}),
	actions: {
		tl(key?: string) {
			if (!key) {
				return key;
			}

			const globalItem = this.global[key];
			{
				if (globalItem) {
					return globalItem;
				}
			}

			return key;
		},
		tg(key?: string) {
			return this.tl(key);
		},
		setLanguages(languages: LanguageType[]) {
			this._languages = languages;

			if (!this._currentLang) {
				this._currentLang = getCurrentLang("lang") || languages[0];
			}
		},
		setCurrentLang(language: LanguageType) {
			this._currentLang = language;
			localStorage.setItem("lang", JSON.stringify(language));
		},
		setGlobalLang(content: Record<string, string>) {
			this.global = content || {};
		},
		setLocalLang(content: Record<string, string>) {
			this.local = content || {};
		},
		getName() {
			return this._currentLang?.name;
		},
	},
	getters: {
		lang: (state) => state._currentLang,
		langID: (state) => state._currentLang?.id ?? 0,
		languages: (state) => state._languages,
		currentLang: (state) => state._currentLang,
		hasLang: (state) => !!state._currentLang,
	},
});

function getCurrentLang(key: string) {
	const lang = localStorage.getItem(key);

	if (!lang) {
		return null;
	}

	return JSON.parse(lang);
}
