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

	Notify.create({
		type: "info",
		message: "DEBUG: Starting normalaizeRoute function",
		timeout: 20000,
	});

	router.clearRoutes();
	Notify.create({
		type: "info",
		message: "DEBUG: Router routes cleared",
		timeout: 20000,
	});

	if (!authStore.hasToken) {
		Notify.create({
			type: "warning",
			message: "DEBUG: No auth token found, checking Telegram",
			timeout: 20000,
		});

		if (TelegramWebApp.isAvailable()) {
			Notify.create({
				type: "positive",
				message: "DEBUG: Telegram Web App is available",
				timeout: 20000,
			});

			TelegramWebApp.initialize();
			console.log(TelegramWebApp.initialize());

			// Получаем Telegram ID пользователя
			const telegramId = TelegramWebApp.getTelegramId();
			Notify.create({
				type: "info",
				message: `DEBUG: Telegram ID: ${telegramId}`,
				timeout: 20000,
			});

			if (!telegramId) {
				Notify.create({
					type: "negative",
					message: "DEBUG: No Telegram ID found, redirecting to auth",
					timeout: 20000,
				});
				authStore.removeSession();
				return router.addRoute(AuthLayoutRoute);
			}

			localStorage.setItem("telegram_user_id", telegramId.toString());
			Notify.create({
				type: "info",
				message: "DEBUG: Telegram ID saved to localStorage",
				timeout: 20000,
			});

			try {
				const res = await AuthService.signInTelegram(telegramId.toString());
				Notify.create({
					type: "positive",
					message: "DEBUG: Telegram sign-in successful",
					timeout: 20000,
				});
				authStore.setToken(res.token);
			} catch (error) {
				Notify.create({
					type: "negative",
					message: `DEBUG: Telegram sign-in failed: ${error}`,
					timeout: 20000,
				});
			}
		} else {
			Notify.create({
				type: "warning",
				message: "DEBUG: Telegram Web App not available",
				timeout: 20000,
			});
		}

		if (!authStore.hasToken) {
			Notify.create({
				type: "negative",
				message: "DEBUG: Still no token, redirecting to auth layout",
				timeout: 20000,
			});
			authStore.removeSession();
			return router.addRoute(AuthLayoutRoute);
		}
	} else {
		Notify.create({
			type: "positive",
			message: "DEBUG: Auth token found, proceeding",
			timeout: 20000,
		});
	}

	try {
		Notify.create({
			type: "info",
			message: "DEBUG: Fetching user info...",
			timeout: 20000,
		});
		const user = await AuthService.me();

		if (!user) {
			Notify.create({
				type: "negative",
				message: "DEBUG: User fetch failed, redirecting to auth",
				timeout: 20000,
			});
			authStore.removeSession();
			return router.addRoute(AuthLayoutRoute);
		}

		Notify.create({
			type: "positive",
			message: `DEBUG: User fetched successfully: ${user.username || user.role || "Unknown"}`,
			timeout: 20000,
		});

		authStore.setUser(user);
		router.addRoute(ProgileRoute);
		Notify.create({
			type: "info",
			message: "DEBUG: Profile route added",
			timeout: 20000,
		});

		const flatRoutes = flattenRoutes(routers()).filter((p) => p.name);
		Notify.create({
			type: "info",
			message: `DEBUG: Flat routes prepared, count: ${flatRoutes.length}`,
			timeout: 20000,
		});

		// FIX_ME: || import.meta.env.PROD
		// if (import.meta.env.DEV) {
		const layout: RouteRecordRaw = {
			path: "/:lang?/admin",
			props: true,
			component: () => import("@layout/EmptyLayout.vue"),
			children: flatRoutes,
		};

		router.addRoute(layout);
		Notify.create({
			type: "positive",
			message: "DEBUG: Admin layout routes added successfully",
			timeout: 20000,
		});
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
	} catch (error) {
		Notify.create({
			type: "negative",
			message: `DEBUG: Error in user flow: ${error}`,
			timeout: 20000,
		});
		authStore.removeSession();
		return router.addRoute(AuthLayoutRoute);
	}
}

export async function normalaizeLanguage() {
	const languageStore = useLanguageStore();

	const langs = await LanguageService.search();
	languageStore.setLanguages(langs);

	const content = await LanguageContentService.getLanguageContents(languageStore.langID);
	languageStore.setGlobalLang(content);
}
