import type { RouteRecordRaw } from "vue-router";

const pageRoutes: RouteRecordRaw[] = [];

export function routers(): RouteRecordRaw[] {
	return pageRoutes;
}

export function registerRoutes(...routers: RouteRecordRaw[]) {
	pageRoutes.push(...routers);
}
