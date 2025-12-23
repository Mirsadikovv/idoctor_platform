import type { RouteRecordRaw } from "vue-router";
import { LanguageRoutes } from "./Language/routes";
import { TranslatedContentRoutes } from "./TranslatedContent/routes";

const langGroup: RouteRecordRaw = {
	path: "lang-groups",
	name: "",
	component: () => import("@/components/layouts/EmptyLayout.vue"),
	meta: {
		name: "",
		title: "lang_group_item",
		accsess: async () => {
			return true;
		},
		expanded: {
			label: "lang_and_translations",
			icon: "g_translate",
		},
	},
	children: [
		// Language sort 1
		...LanguageRoutes(1),

		// Translated Content sort 2
		...TranslatedContentRoutes(2),
	],
};

export function LanguageGroupRoutes(sort: number): RouteRecordRaw {
	return {
		...langGroup,
		meta: {
			...langGroup.meta,
			sort,
			title: "lang_group_item",
		},
	};
}

export default langGroup;
