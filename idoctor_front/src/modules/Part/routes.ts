import type { RouteRecordRaw } from "vue-router";

const partPageRoute: RouteRecordRaw = {
	path: "parts",
	name: "PART_PAGE",
	component: () => import("@module/Part/pages/Page.vue"),
	meta: {
		title: "part_page_title",
		activeLinkGroup: "PART_GROUP",
		sidebar: {
			label: "part_page_title",
			icon: "build",
			isExpandedGroup: false,
		},
	},
};

const partCreateRoute: RouteRecordRaw = {
	path: "parts/create",
	name: "PART_CREATE",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "part_create_title",
		activeLinkGroup: "PART_GROUP",
	},
};

const partEditRoute: RouteRecordRaw = {
	path: "parts/:id/edit",
	name: "PART_EDIT",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "part_edit_title",
		activeLinkGroup: "PART_GROUP",
	},
};

export function PartRoutes(sort: number): RouteRecordRaw[] {
	return [partPageRoute, partCreateRoute, partEditRoute].map((route) => {
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