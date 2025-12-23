import type { RouteRecordRaw } from "vue-router";

const userPageRoute: RouteRecordRaw = {
	path: "users",
	name: "USER_PAGE",
	component: () => import("@module/User/pages/Page.vue"),
	meta: {
		title: "users_page",
		activeLinkGroup: "USER_GROUP",
		sidebar: {
			label: "users_page",
			icon: "account_circle",
			isExpandedGroup: false,
		},
	},
};

const userViewRoute: RouteRecordRaw = {
	path: "users-view/:id",
	name: "USER_VIEW",
	props: true,
	component: () => import("@module/User/pages/View.vue"),
	meta: {
		title: "users_page",
		activeLinkGroup: "USER_GROUP",
	},
};

const userCreateRoute: RouteRecordRaw = {
	path: "users/create",
	name: "USER_CREATE",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "user_create_title",
		activeLinkGroup: "USER_GROUP",
	},
};

const userEditRoute: RouteRecordRaw = {
	path: "users/:id/edit",
	name: "USER_EDIT", 
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "user_edit_title",
		activeLinkGroup: "USER_GROUP",
	},
};

const userDeleteRoute: RouteRecordRaw = {
	path: "",
	name: "USER_DELETE",
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "users_page",
		activeLinkGroup: "USER_GROUP",
	},
};

export function UserRoutes(sort: number): RouteRecordRaw[] {
	return [
		userPageRoute,
		userCreateRoute,
		userEditRoute,
		userViewRoute,
		userDeleteRoute,
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
