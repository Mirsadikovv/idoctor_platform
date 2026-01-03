import type { RouteRecordRaw } from "vue-router";

const languagePageRoute: RouteRecordRaw = {
	path: "languages",
	name: "LANGUAGE_PAGE",
	component: () => import("@module/Language/pages/Page.vue"),
	meta: {
		title: "language_page_title",
		activeLinkGroup: "LANGUAGE_GROUP",
		sidebar: {
			label: "language_page_title",
			icon: "label",
			isExpandedGroup: true,
		},
	},
};

const languageCreateRoute: RouteRecordRaw = {
	path: "languages/create",
	name: "LANGUAGE_CREATE",
	props: true,
	component: () => import("@module/Language/pages/Create.vue"),
	meta: {
		title: "language_create_title",
		activeLinkGroup: "LANGUAGE_GROUP",
	},
};

const languageViewRoute: RouteRecordRaw = {
	path: "languages/:id",
	name: "LANGUAGE_VIEW",
	props: true,
	component: () => import("@module/Language/pages/View.vue"),
	meta: {
		title: "language_view_title",
		activeLinkGroup: "LANGUAGE_GROUP",
	},
};

const languageEditRoute: RouteRecordRaw = {
	path: "languages/:id/edit",
	name: "LANGUAGE_EDIT",
	props: true,
	component: () => import("@module/Language/pages/Edit.vue"),
	meta: {
		title: "language_edit_title",
		activeLinkGroup: "LANGUAGE_GROUP",
	},
};

export function LanguageRoutes(sort: number): RouteRecordRaw[] {
	return [languagePageRoute, languageCreateRoute, languageViewRoute, languageEditRoute].map(
		(route) => {
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
		},
	);
}
