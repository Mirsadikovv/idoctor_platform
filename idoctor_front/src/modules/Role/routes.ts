import type { RouteRecordRaw } from "vue-router";

const rolePageRoute: RouteRecordRaw = {
	path: "roles",
	name: "PAGE_ROLE",
	component: () => import("@module/Role/pages/Page.vue"),
	meta: {
		title: "roles_page",
		activeLinkGroup: "ROLE_GROUP",
		contents: [],
		sidebar: {
			label: "roles_page",
			icon: "person",
			isExpandedGroup: false,
		},
	},
};
const roleCreateRoute: RouteRecordRaw = {
	path: "role-create",
	name: "CREATE_ROLE",
	component: () => import("@module/Role/pages/Create.vue"),
	meta: {
		title: "create_role",
		activeLinkGroup: "ROLE_GROUP",
		contents: [],
	},
};

const roleUpdateRoute: RouteRecordRaw = {
	path: "role-update/:id",
	name: "EDIT_ROLE",
	props: true,
	component: () => import("@module/Role/pages/Edit.vue"),
	meta: {
		title: "update_role",
		activeLinkGroup: "ROLE_GROUP",
		contents: [],
	},
};

export function RoleRoutes(sort: number): RouteRecordRaw[] {
	return [rolePageRoute, roleCreateRoute, roleUpdateRoute].map((route) => {
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
