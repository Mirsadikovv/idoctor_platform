import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { PageDataType } from "@/service";

export type User = {
	id: number;
	username: string;
	firstName: string;
	lastName: string;
	middleName: string;
	dateOfBirth: string;
	gender: "male" | "female";
	phoneNumber: string;
	roleId: number;
	telegramId: number;
	telegramUsername: string;
	languageCode: string;
	lastVisit: string;

	fullName: string;
};

export type UserCreateType = {
	username: string;
	password: string;
	roleId: number;
	firstName?: string;
	lastName?: string;
	middleName?: string;
	dateOfBirth?: string;
	gender?: "male" | "female";
	phoneNumber?: string;
	telegramId?: number;
	telegramUsername?: string;
	languageCode?: string;
};

export type UserUpdateType = {
	username?: string;
	password?: string;
};

export type UserPartial = Partial<User>;
export type UserPage = PageDataType<User>;

export type UserSearchParams = {
	page?: number;
	perpage?: number;
	fullname?: string;
	gender?: "male" | "female";
	order?: string;
	sort_by?: string;
	limit?: number;
};

class UserService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(params: UserSearchParams | string = {}) {
		let queryString = "";

		if (typeof params === "string") {
			queryString = params;
		} else {
			const searchParams = new URLSearchParams();
			Object.entries(params).forEach(([key, value]) => {
				if (value !== undefined) {
					searchParams.append(key, String(value));
				}
			});
			queryString = searchParams.toString();
		}

		const { data } = await api.get<UserPage>(`/user/page?${queryString}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async search(params: UserSearchParams | string = {}) {
		let queryString = "";

		if (typeof params === "string") {
			queryString = params;
		} else {
			const searchParams = new URLSearchParams();
			Object.entries(params).forEach(([key, value]) => {
				if (value !== undefined) {
					searchParams.append(key, String(value));
				}
			});
			queryString = searchParams.toString();
		}

		const { data } = await api.get<User[]>(`/user/search?${queryString}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async getByID(id: number) {
		const { data } = await api.get<User>(`/user/${id}`);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async create(user: UserCreateType) {
		const { data } = await api.post(`/user`, user);
		return data;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async update(id: number, user: UserUpdateType) {
		await api.patch(`/user/${id}`, user);
		return true;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async delete(id: number) {
		await api.delete(`/user/${id}`);
		return true;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				err?.response?.data.message || err.message,
			);
		},
	})
	async restore(id: number) {
		await api.patch(`/user/${id}/restore`);
		return true;
	}

	// Legacy compatibility method
	async findByID(id: number) {
		return this.getByID(id);
	}
}

const userService = new UserService();

export { userService as UserService };
