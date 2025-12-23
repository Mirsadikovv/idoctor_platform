import { buildQuery, Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { User } from "@/service";

export interface AuthUser {
	username: string;
	password: string;
}

export interface RefreshBody {
	organization_id: number;
	role_id: number;
	soato_id: number;
}

export interface AuthToken {
	token: string;
}

class AuthService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async signIn(authUser: AuthUser) {
		const { data } = await api.post<AuthToken>("/auth/sign-in", authUser);
		return data;
	}
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async signUp(user: User) {
		await api.post("/auth/sign-up", user);
		return true;
	}

	@Try({
		async onSuccess(result) {
			(await import("@/common/Notify")).SuccesNotify(result.statusText);
		},
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async signOut() {
		await api.post("/auth/sign-out");
		return true;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async me() {
		const { data } = await api.post<User>(`/auth/me`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async refreshToken(refresh: Partial<RefreshBody>) {
		const query = buildQuery(refresh);
		const { data } = await api.post<AuthToken>(`/auth/refresh?${query}`);
		return data;
	}
}

const authService = new AuthService();

export { authService as AuthService };
