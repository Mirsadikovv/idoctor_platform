import type { RouteRecordRaw } from "vue-router";

export const emptyRoute: RouteRecordRaw = {
	path: "empty",
	name: "PAGE_EMPTY",
	component: () => import("@/modules/Auth/pages/Empty.vue"),
	meta: {
		title: "empty_route",
	},
};

const LoginRoute: RouteRecordRaw = {
	path: "login",
	name: "LOGIN_AUTH",
	component: () => import("@/modules/Auth/pages/OneId.vue"),
	meta: {
		title: "login_title",
	},
};

const ProgileRoute: RouteRecordRaw = {
	path: "/:lang/profile",
	name: "PAGE_PROFILE",
	component: () => import("@/modules/Auth/pages/Profile.vue"),
	meta: {
		title: "profile_title",
	},
};

export function AuthRoutes(sort: number): RouteRecordRaw[] {
	return [LoginRoute, ProgileRoute, emptyRoute].map((route) => {
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

// Экспортирую отдельно для обратной совместимости, так как эти роуты используются без sidebar
export { LoginRoute, ProgileRoute };
