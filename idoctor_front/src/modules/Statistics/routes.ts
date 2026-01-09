import type { RouteRecordRaw } from "vue-router";

const statisticsPageRoute: RouteRecordRaw = {
	path: "statistics",
	name: "STATISTICS_PAGE",
	component: () => import("@module/Statistics/pages/Page.vue"),
	meta: {
		title: "statistics",
		activeLinkGroup: "STATISTICS_GROUP",
		permissions: ["statistics"],
		sidebar: {
			label: "statistics",
			icon: "analytics",
			isExpandedGroup: false,
		},
	},
};

const statisticsMastersRoute: RouteRecordRaw = {
	path: "statistics/masters",
	name: "STATISTICS_MASTERS",
	component: () => import("@module/Statistics/pages/Masters/Page.vue"),
	meta: {
		title: "statistics_masters",
		activeLinkGroup: "STATISTICS_GROUP",
		permissions: ["statistics"],
	},
};

const statisticsOrdersRoute: RouteRecordRaw = {
	path: "statistics/orders",
	name: "STATISTICS_ORDERS",
	component: () => import("@module/Statistics/pages/Orders/Page.vue"),
	meta: {
		title: "statistics_orders",
		activeLinkGroup: "STATISTICS_GROUP",
		permissions: ["statistics"],
	},
};

const statisticsPaymentsRoute: RouteRecordRaw = {
	path: "statistics/payments",
	name: "STATISTICS_PAYMENTS",
	component: () => import("@module/Statistics/pages/Payments/Page.vue"),
	meta: {
		title: "statistics_payments",
		activeLinkGroup: "STATISTICS_GROUP",
		permissions: ["statistics"],
	},
};

const statisticsRevenueRoute: RouteRecordRaw = {
	path: "statistics/revenue",
	name: "STATISTICS_REVENUE",
	component: () => import("@module/Statistics/pages/Revenue/Page.vue"),
	meta: {
		title: "statistics_revenue",
		activeLinkGroup: "STATISTICS_GROUP",
		permissions: ["statistics"],
	},
};

const statisticsTopPartsRoute: RouteRecordRaw = {
	path: "statistics/top-parts",
	name: "STATISTICS_TOP_PARTS",
	component: () => import("@module/Statistics/pages/TopParts/Page.vue"),
	meta: {
		title: "statistics_top_parts",
		activeLinkGroup: "STATISTICS_GROUP",
		permissions: ["statistics"],
	},
};

export function StatisticsRoutes(sort: number): RouteRecordRaw[] {
	return [
		statisticsPageRoute,
		statisticsMastersRoute,
		statisticsOrdersRoute,
		statisticsPaymentsRoute,
		statisticsRevenueRoute,
		statisticsTopPartsRoute,
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
