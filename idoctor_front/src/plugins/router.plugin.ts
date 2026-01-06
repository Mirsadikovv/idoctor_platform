import { LanguageContentService } from "@/modules/TranslatedContent/service";
import { AuthLayoutRoute } from "@/router";
import { routers } from "@/router/layout";
import { AuthService, LanguageService } from "@/service";
import { useAuthStore } from "@/store/auth-store";
import { useLanguageStore } from "@/store/language-store";
import type { App } from "vue";
import { createRouter, createWebHistory, type Router, type RouteRecordRaw } from "vue-router";
import { OnError, OnRequest } from "./axios.plugin";
import { emptyRoute, ProgileRoute } from "@/modules/Auth/router";
import { admin, client, flattenRoutes, master } from "@/common";
import { TelegramWebApp } from "@/common/telegram";

export async function routerPlugin(app: App<unknown>) {
	// Инициализируем Telegram Web App если доступен

	const router = createRouter({
		history: createWebHistory(),
		routes: [],
	});

	const authStore = useAuthStore();

	OnRequest((response) => {
		if (!authStore.hasToken) return;
		response.headers["Authorization"] = `Bearer ${authStore.getToken}`;
	});

	OnError(async (error) => {
		if (error.response?.status === 401) {
			authStore.removeSession();
			router.clearRoutes();
			router.addRoute(AuthLayoutRoute);
			router.push({ name: "LOGIN_AUTH" });
		}
	});
	const languageStore = useLanguageStore();

	await normalaizeRoute(router);

	await normalaizeLanguage();

	router.afterEach((to, _from, failuer) => {
		if (failuer) {
			return;
		}
		document.title = languageStore.tl(to.meta.title) || APP_TITLE;
	});

	router.beforeEach(async (to, _from, next) => {
		const routes = router.getRoutes();

		let _next = next;

		const route: RouteRecordRaw | undefined = routes.find(
			(route) => !!route.name && route.name === to.name,
		);

		if (!route) {
			const layout = routes.find((route) => !route.name);

			const child =
				layout!.children.find((route) => route && route.meta?.sidebar) ||
				layout!.children[0];

			_next = () => {
				next({
					name: child.name,
					params: {
						lang: to.params.lang || languageStore.lang?.name,
					},
				});
			};
		}
		const activeLang =
			languageStore.languages.find((lang) => lang.name === to.params.lang) ||
			languageStore.lang ||
			languageStore.languages[0];

		languageStore.setCurrentLang(activeLang);

		return _next();
	});

	app.use(router);

	await router.isReady();
}

export async function normalaizeRoute(router: Router) {
	const authStore = useAuthStore();
	const languageStore = useLanguageStore();

	router.clearRoutes();

	if (!authStore.hasToken) {
		if (TelegramWebApp.isAvailable()) {
			TelegramWebApp.initialize();

			// Получаем Telegram ID пользователя
			const telegramId = TelegramWebApp.getTelegramId();

			if (!telegramId) {
				authStore.removeSession();
				return router.addRoute(AuthLayoutRoute);
			}

			localStorage.setItem("telegram_user_id", telegramId.toString());

			const res = await AuthService.signInTelegram(telegramId.toString());

			authStore.setToken(res.token);
		}

		if (!authStore.hasToken) {
			authStore.removeSession();
			return router.addRoute(AuthLayoutRoute);
		}
	}

	const user = await AuthService.me();

	if (!user) {
		authStore.removeSession();
		return router.addRoute(AuthLayoutRoute);
	}

	let userWithPages = { ...user };

	if (user.role === "admin") {
		userWithPages = { ...user, pages: [...admin] };
		authStore.setUser({ ...user, pages: [...admin] });
	}
	if (user.role === "master") {
		userWithPages = { ...user, pages: [...master] };
		authStore.setUser({ ...user, pages: [...master] });
	}
	if (user.role === "user") {
		userWithPages = { ...user, pages: [...client] };
		authStore.setUser({ ...user, pages: [...client] });
	}

	router.addRoute(ProgileRoute);

	const flatRoutes = flattenRoutes(routers()).filter((p) => p.name);

	// FIX_ME: || import.meta.env.PROD
	// if (import.meta.env.DEV) {
	// 	const layout: RouteRecordRaw = {
	// 		path: "/:lang?/admin",
	// 		props: true,
	// 		component: () => import("@layout/EmptyLayout.vue"),
	// 		children: flatRoutes,
	// 	};

	// 	router.addRoute(layout);

	// 	await router.push({
	// 		name: "PAGE_PROFILE",
	// 		params: {
	// 			lang: languageStore.currentLang?.name,
	// 		},
	// 	});
	// } else {

	const children = userWithPages.pages
		.map((routeName: string) => flatRoutes.find((route) => route.name === routeName)!)
		.filter((route) => !!route);

	const layout: RouteRecordRaw = {
		path: "/:lang?/admin",
		props: true,
		component: () => import("@layout/EmptyLayout.vue"),
		children: [...children, emptyRoute],
	};

	router.addRoute(layout);

	await router.push({
		name: "PAGE_PROFILE",
		params: {
			lang: languageStore.currentLang?.name,
		},
	});
	// }
}

export async function normalaizeLanguage() {
	const languageStore = useLanguageStore();

	const langs = await LanguageService.search();
	languageStore.setLanguages(langs);

	const content = await LanguageContentService.getLanguageContents(languageStore.langID);
	languageStore.setGlobalLang(content);
}
