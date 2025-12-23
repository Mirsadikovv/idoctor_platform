import type { RouteRecordRaw } from "vue-router";

const orderPageRoute: RouteRecordRaw = {
	path: "orders",
	name: "ORDER_PAGE",
	component: () => import("@module/Order/pages/Page.vue"),
	meta: {
		title: "order_page_title",
		activeLinkGroup: "ORDER_GROUP",
		sidebar: {
			label: "order_page_title",
			icon: "assignment",
			isExpandedGroup: false,
		},
	},
};

const orderCreateRoute: RouteRecordRaw = {
	path: "orders/create",
	name: "ORDER_CREATE",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "order_create_title",
		activeLinkGroup: "ORDER_GROUP",
	},
};

const orderEditRoute: RouteRecordRaw = {
	path: "orders/:id/edit",
	name: "ORDER_EDIT",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "order_edit_title",
		activeLinkGroup: "ORDER_GROUP",
	},
};

const orderViewRoute: RouteRecordRaw = {
	path: "orders/:id",
	name: "ORDER_VIEW",
	props: true,
	component: () => import("@module/Order/pages/View.vue"),
	meta: {
		title: "order_view_title",
		activeLinkGroup: "ORDER_GROUP",
	},
};

export function OrderRoutes(sort: number): RouteRecordRaw[] {
	return [orderPageRoute, orderCreateRoute, orderEditRoute, orderViewRoute].map((route) => {
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