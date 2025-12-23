import type { RouteRecordRaw } from "vue-router";

const devicePageRoute: RouteRecordRaw = {
	path: "devices",
	name: "DEVICE_PAGE",
	component: () => import("@module/Device/pages/Page.vue"),
	meta: {
		title: "device_page_title",
		activeLinkGroup: "DEVICE_GROUP",
		sidebar: {
			label: "device_page_title",
			icon: "devices",
			isExpandedGroup: false,
		},
	},
};

const deviceCreateRoute: RouteRecordRaw = {
	path: "devices/create",
	name: "DEVICE_CREATE",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "device_create_title",
		activeLinkGroup: "DEVICE_GROUP",
	},
};

const deviceEditRoute: RouteRecordRaw = {
	path: "devices/:id/edit",
	name: "DEVICE_EDIT",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "device_edit_title",
		activeLinkGroup: "DEVICE_GROUP",
	},
};

export function DeviceRoutes(sort: number): RouteRecordRaw[] {
	return [devicePageRoute, deviceCreateRoute, deviceEditRoute].map((route) => {
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
