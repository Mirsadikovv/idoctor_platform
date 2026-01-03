import type { RouteRecordRaw } from "vue-router";

const orderPartPageRoute: RouteRecordRaw = {
	path: "order-parts/:orderId",
	name: "ORDER_PART_PAGE",
	props: true,
	component: () => import("@module/OrderPart/pages/Page.vue"),
	meta: {
		title: "order_part_page_title",
		activeLinkGroup: "ORDER_PART_GROUP",
	},
};

const orderPartCreateRoute: RouteRecordRaw = {
	path: "order-parts/create/:orderId",
	name: "ORDER_PART_CREATE",
	props: true,
	component: () => import("@module/OrderPart/pages/Create.vue"),
	meta: {
		title: "order_part_create_title",
		activeLinkGroup: "ORDER_PART_GROUP",
	},
};

const orderPartEditRoute: RouteRecordRaw = {
	path: "order-parts/:id/edit",
	name: "ORDER_PART_EDIT",
	props: true,
	component: () => import("@module/OrderPart/pages/Edit.vue"),
	meta: {
		title: "order_part_edit_title",
		activeLinkGroup: "ORDER_PART_GROUP",
	},
};

const orderPartViewRoute: RouteRecordRaw = {
	path: "order-parts/:id",
	name: "ORDER_PART_VIEW",
	props: true,
	component: () => import("@module/OrderPart/pages/View.vue"),
	meta: {
		title: "order_part_view_title",
		activeLinkGroup: "ORDER_PART_GROUP",
	},
};

export function OrderPartRoutes(_sort: number): RouteRecordRaw[] {
	return [orderPartPageRoute, orderPartCreateRoute, orderPartEditRoute, orderPartViewRoute];
}
