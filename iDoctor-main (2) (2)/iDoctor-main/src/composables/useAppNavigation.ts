import { useRouter } from "vue-router";
import { useAuthStore } from "@/store/auth-store";
import { useLanguageStore } from "@/store/language-store";
import { AuthLayoutRoute } from "@/router";
import { AuthService, LanguageContentService, type LanguageType } from "@/service";

export function useAppNavigation() {
	const router = useRouter();
	const authStore = useAuthStore();
	const languageStore = useLanguageStore();

	function toggleLeftDrawer() {
		router.push({ 
			name: "PAGE_PROFILE", 
			params: { lang: languageStore.currentLang?.name } 
		});
	}

	async function setLang(language: LanguageType) {
		const globalContent = LanguageContentService.getLanguageContents(language.id);
		const [global] = await Promise.all([globalContent]);
		
		languageStore.setGlobalLang(global);

		router.push({
			name: router.currentRoute.value.name as string,
			params: {
				...router.currentRoute.value.params,
				lang: language.name,
			},
			query: router.currentRoute.value.query,
		});
	}

	async function logout() {
		await AuthService.signOut().catch(() => {});
		authStore.removeSession();
		router.clearRoutes();
		router.addRoute(AuthLayoutRoute);
		await router.push({ name: "LOGIN_AUTH" });
	}

	return {
		toggleLeftDrawer,
		setLang,
		logout
	};
}