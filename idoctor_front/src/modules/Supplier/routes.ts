import type { RouteRecordRaw } from "vue-router";

const supplierPageRoute: RouteRecordRaw = {
	path: "suppliers",
	name: "SUPPLIER_PAGE",
	component: () => import("@module/Supplier/pages/Page.vue"),
	meta: {
		title: "supplier_page_title",
		activeLinkGroup: "SUPPLIER_GROUP",
		sidebar: {
			label: "supplier_page_title",
			icon: "business",
			isExpandedGroup: false,
		},
	},
};

const supplierCreateRoute: RouteRecordRaw = {
	path: "suppliers/create",
	name: "SUPPLIER_CREATE",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "supplier_create_title",
		activeLinkGroup: "SUPPLIER_GROUP",
	},
};

const supplierEditRoute: RouteRecordRaw = {
	path: "suppliers/:id/edit",
	name: "SUPPLIER_EDIT",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "supplier_edit_title",
		activeLinkGroup: "SUPPLIER_GROUP",
	},
};

export function SupplierRoutes(sort: number): RouteRecordRaw[] {
	return [supplierPageRoute, supplierCreateRoute, supplierEditRoute].map((route) => {
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