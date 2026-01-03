import { LanguageContentService } from "@/modules/TranslatedContent/service";
import { AuthLayoutRoute } from "@/router";
import { routers } from "@/router/layout";
import { AuthService, LanguageService } from "@/service";
import { useAuthStore } from "@/store/auth-store";
import { useLanguageStore } from "@/store/language-store";
import type { App } from "vue";
import { createRouter, createWebHistory, type Router, type RouteRecordRaw } from "vue-router";
import { OnError, OnRequest } from "./axios.plugin";
import { ProgileRoute } from "@/modules/Auth/router";
import { flattenRoutes } from "@/common";
import { TelegramWebApp } from "@/common/telegram";
import { Notify } from "quasar";

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

	await normalaizeRoute(router);
	await normalaizeLanguage();

	const languageStore = useLanguageStore();

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

	router.clearRoutes();

	if (!authStore.hasToken) {
		Notify.create({
			type: "warning",
			message: "DEBUG: No auth token found, checking Telegram",
			timeout: 20000,
			actions: [{ icon: "close", color: "white" }],
		});

		if (TelegramWebApp.isAvailable()) {
			Notify.create({
				type: "positive",
				message: "DEBUG: Telegram Web App is available",
				timeout: 20000,
				actions: [{ icon: "close", color: "white" }],
			});

			TelegramWebApp.initialize();
			console.log(TelegramWebApp.initialize());

			// Получаем Telegram ID пользователя
			const telegramId = TelegramWebApp.getTelegramId();
			Notify.create({
				type: "info",
				message: `DEBUG: Telegram ID: ${telegramId || "not found"}`,
				timeout: 20000,
				actions: [{ icon: "close", color: "white" }],
			});

			if (!telegramId) {
				Notify.create({
					type: "negative",
					message: "DEBUG: No Telegram ID found, redirecting to auth",
					timeout: 20000,
					actions: [{ icon: "close", color: "white" }],
				});
				authStore.removeSession();
				return router.addRoute(AuthLayoutRoute);
			}

			const res = await AuthService.signInTelegram(telegramId.toString());
			Notify.create({
				type: "positive",
				message: "DEBUG: Telegram sign-in successful",
				timeout: 20000,
				actions: [{ icon: "close", color: "white" }],
			});
			authStore.setToken(res.token);
		} else {
			Notify.create({
				type: "warning",
				message: "DEBUG: Telegram Web App not available",
				timeout: 20000,
				actions: [{ icon: "close", color: "white" }],
			});
		}

		if (!authStore.hasToken) {
			Notify.create({
				type: "negative",
				message: "DEBUG: Still no token, redirecting to auth layout",
				timeout: 20000,
				actions: [{ icon: "close", color: "white" }],
			});
			authStore.removeSession();
			return router.addRoute(AuthLayoutRoute);
		}
	}

	const user = await AuthService.me();

	if (!user) {
		authStore.removeSession();
		return router.addRoute(AuthLayoutRoute);
	}

	authStore.setUser(user);

	router.addRoute(ProgileRoute);

	const flatRoutes = flattenRoutes(routers()).filter((p) => p.name);

	// FIX_ME: || import.meta.env.PROD
	// if (import.meta.env.DEV) {
	const layout: RouteRecordRaw = {
		path: "/:lang?/admin",
		props: true,
		component: () => import("@layout/EmptyLayout.vue"),
		children: flatRoutes,
	};

	router.addRoute(layout);

	// } else {
	// 	const children = Object.keys(user.pages)
	// 		.map((routeName: string) => flatRoutes.find((route) => route.name === routeName)!)
	// 		.filter((route) => !!route);

	// 	const layout: RouteRecordRaw = {
	// 		path: "/:lang?/admin",
	// 		props: true,
	// 		component: () => import("@layout/BaseLayout.vue"),
	// 		children: [...children, emptyRoute],
	// 	};

	// 	router.addRoute(layout);
	// }
}

export async function normalaizeLanguage() {
	const languageStore = useLanguageStore();

	const langs = await LanguageService.search();
	languageStore.setLanguages(langs);

	const content = await LanguageContentService.getLanguageContents(languageStore.langID);
	languageStore.setGlobalLang(content);
}
