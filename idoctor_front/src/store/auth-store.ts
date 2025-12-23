import { RemoveToken, SetToken, Token } from "@/common";
import type { User } from "@/modules/User/service";
import { defineStore } from "pinia";

export type AuthStoreType = ReturnType<typeof useAuthStore>;

type State = {
	token: string;
	user: User | undefined;
};

export const useAuthStore = defineStore("auth", {
	state(): State {
		return {
			token: Token(),
			user: undefined,
		};
	},
	actions: {
		setUser(user: User) {
			this.user = user;
		},
		setToken(token: string) {
			this.token = token;
			SetToken(token);
		},
		removeUser() {
			this.user = undefined;
		},
		removeToken() {
			this.token = "";
			RemoveToken();
		},
		removeSession() {
			this.removeToken();
			this.removeUser();
		},
		page(pagename: string) {
			const { pages } = this.user || {};
			return pages?.[pagename] || {};
		},
		canPage(pagename: string) {
			const { pages } = this.user || {};
			return pages ? pagename in pages : false;
		},
		can(pagename: string, ...key: string[]) {
			// TODO: FIX ME

			return true;

			const keys = this.page(pagename) || [];

			if (!keys || !key.length) return false;

			if (!key.length) return true;

			return keys?.some((k: string) => key.includes(k));
		},
	},
	getters: {
		hasToken(state): Readonly<boolean> {
			return !!state.token;
		},
		getToken(state): Readonly<string> {
			return state.token;
		},
		getUser(state): Readonly<User> {
			return state.user as Readonly<User>;
		},
		hasUser(state): Readonly<boolean> {
			return !!state.user;
		},
		pages(state) {
			return state.user?.pages || {};
		},
		routes(state) {
			const routes = state.user?.pages || [];

			if (routes.length) {
				return Object.keys(routes);
			}

			return [];
		},
		permissions(state) {
			return state.user?.role?.permissions || [];
		},

		len(state) {
			return state.user?.role?.permissions?.length || -1;
		},
	},
});
