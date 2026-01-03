import "@/modules";
import { LoginRoute } from "@/modules/Auth/router";
import { forbiddenPage, notfoundPage } from "@/modules/Error/router/error";
import type { RouteRecordRaw } from "vue-router";

export const AuthLayoutRoute: RouteRecordRaw = {
	path: "/auth",
	component: () => import("@layout/EmptyLayout.vue"),
	children: [LoginRoute, notfoundPage, forbiddenPage],
};
