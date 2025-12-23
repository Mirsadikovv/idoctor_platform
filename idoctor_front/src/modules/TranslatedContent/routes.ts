import type { RouteRecordRaw } from "vue-router";

const translatedContentRoute: RouteRecordRaw = {
	path: "translated-content",
	name: "TRANSLATED_CONTENT",
	component: () => import("@module/TranslatedContent/pages/Page.vue"),
	meta: {
		title: "translated_content_page",
		activeLinkGroup: "TRANSLATED_GROUP",
		sidebar: {
			isExpandedGroup: true,
			label: "translated_content_page",
			icon: "label",
		},
	},
};
const translateCreateRoute: RouteRecordRaw = {
	path: "",
	name: "TRANSLATE_CREATE",
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "translate_create_page",
		activeLinkGroup: "TRANSLATED_GROUP",
	},
};

const translateUpdateRoute: RouteRecordRaw = {
	path: "",
	name: "TRANSLATE_UPDATE",
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "translate_update_page",
		activeLinkGroup: "TRANSLATED_GROUP",
	},
};

const translateDeleteRoute: RouteRecordRaw = {
	path: "",
	name: "TRANSLATE_DELETE",
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "translate_delete_page",
		activeLinkGroup: "TRANSLATED_GROUP",
	},
};

export function TranslatedContentRoutes(sort: number): RouteRecordRaw[] {
	return [
		translatedContentRoute,
		translateCreateRoute,
		translateUpdateRoute,
		translateDeleteRoute,
	].map((route) => {
		if (route?.meta?.sidebar) {
			return {
				...route,
				meta: {
					...route.meta,
					sort,
				},
			};
		}
		return route;
	});
}
