import type { RouteRecordRaw } from "vue-router";

export const notfoundPage: RouteRecordRaw = {
	// path: "/:pathMatch(.*)*", // ← важный момент
	path: "/not-found", // ← важный момент
	name: "404_PAGE",
	component: () => import("@module/Error/pages/404.vue"),
	meta: {
		name: "404_PAGE", // add this property
		title: "notfound_item",
	},
};
export const forbiddenPage: RouteRecordRaw = {
	// path: "/:pathMatch(.*)*", // ← важный момент
	path: "/forbidden", // ← важный момент
	name: "PAGE_FORBIDDEN",
	component: () => import("@module/Error/pages/404.vue"),
	meta: {
		name: "PAGE_FORBIDDEN", // add this property
		title: "forbidden_item",
	},
};
