import type { RouteRecordRaw } from "vue-router";

const problemPageRoute: RouteRecordRaw = {
	path: "problems",
	name: "PROBLEM_PAGE",
	component: () => import("@module/Problem/pages/Page.vue"),
	meta: {
		title: "problem_page_title",
		activeLinkGroup: "PROBLEM_GROUP",
		sidebar: {
			label: "problem_page_title",
			icon: "medical_services",
			isExpandedGroup: false,
		},
	},
};

const problemCreateRoute: RouteRecordRaw = {
	path: "problems/create",
	name: "PROBLEM_CREATE",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "problem_create_title",
		activeLinkGroup: "PROBLEM_GROUP",
	},
};

const problemEditRoute: RouteRecordRaw = {
	path: "problems/:id/edit",
	name: "PROBLEM_EDIT",
	props: true,
	component: () => import("@layout/EmptyLayout.vue"),
	meta: {
		title: "problem_edit_title",
		activeLinkGroup: "PROBLEM_GROUP",
	},
};

export function ProblemRoutes(sort: number): RouteRecordRaw[] {
	return [problemPageRoute, problemCreateRoute, problemEditRoute].map((route) => {
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