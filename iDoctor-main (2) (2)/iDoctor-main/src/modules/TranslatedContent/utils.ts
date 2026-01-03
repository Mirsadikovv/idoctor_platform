import { type LangContentCreateOrUpdate, type LanguageContentPartialType } from "@/service";
import { useLanguageStore } from "@/store/language-store";

const languageStore = useLanguageStore();

export async function searchLangs(_params: string = "") {
	return languageStore.languages;
}
export async function initalBatch() {
	const content =
		languageStore.languages &&
		languageStore.languages.map((language) => ({
			category: "",
			key: "",
			languageId: language.id,
			value: "",
			langName: language.name,
			LangDescp: language.description,
		}));
	return {
		category: "",
		key: "",
		languageId: languageStore.langID,
		contents: content || [],
	};
}

export function populateModelValues(
	baseModel: Partial<LangContentCreateOrUpdate>,
	valuesArray: LanguageContentPartialType[],
): LangContentCreateOrUpdate {
	// Создаем копию базовой модели для избежания мутации оригинала
	const result: LangContentCreateOrUpdate = JSON.parse(JSON.stringify(baseModel));

	// Создаем карту значений по languageId для быстрого поиска
	const valuesByLangId: Record<number, string> = {};
	let commonKey: string = "";

	valuesArray.forEach((item) => {
		// @ts-ignore
		valuesByLangId[item?.languageId] = item.value;
		if (!commonKey && item.key) {
			commonKey = item.key; // Берем key из первого элемента
		}
	});

	// Заполняем key в основной модели
	if (commonKey) {
		result.key = commonKey;
	}

	// Заполняем значения в contents
	result.contents.forEach((content) => {
		if (commonKey) {
			content.key = commonKey;
		}
		if (valuesByLangId[content.languageId]) {
			content.value = valuesByLangId[content.languageId];
		}
	});

	return result;
}
