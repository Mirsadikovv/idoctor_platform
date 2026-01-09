export interface RouteMeta extends Record<string | number | symbol, any> {}
import type { ProcessedIcons } from "@/common";
import type { LanguageStoreType } from "./store/language-store";
import type { AuthStoreType } from "./store/auth-store";

declare module "vue-router" {
	interface RouteMeta extends Record<string | number | symbol, any> {
		title: string;
		activeLinkGroup?: string;
		sort?: number;
		expanded?: {
			icon: ProcessedIcons;
			label: string;
		};
		sidebar?: {
			isExpandedGroup: boolean;
			icon: ProcessedIcons;
			label: string;
			activeTab?: string;
		};
	}

	interface RouteLocationNormalizedLoadedGeneric extends Record<string | number | symbol, any> {
		query: {
			[key: string]: string | undefined;
			page: number;
			perpage: number;
		};
		params: {
			[key: string]: string | undefined;
			id: number;
		};
	}
}

declare module "@vue/runtime-core" {
	interface ComponentCustomProperties {
		$lang: LanguageStoreType;
		$tl: (key?: string) => string | undefined;
		$tg: (key?: string) => string | undefined;
		$changeLang: (language: LanguageType) => Promise<void>;
		$auth: AuthStoreType;
		$canPage: (pagename: string) => boolean;
		$can: (...key: string[]) => boolean;
	}

	export interface ComponentCustomProperties {
		vPermissions: Directive<HTMLElement, string[], string, "page">;
	}
}
