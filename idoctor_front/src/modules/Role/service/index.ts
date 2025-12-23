import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { IdType, PageDataType } from "@/service";

export type RoleType = {
	id: number;
	name: string;
	description: string;
	pages: Record<string, string[]>;
	permissions: Record<string, string[]>;

	type: string;
};

export type RoleCreateType = {
	id: number;
	name: string;
	description: string;
	pages: Record<string, string[]>;
	permissions: Record<string, string[]>;
};

export type PermissionType = {
	uuid: string;
	path: string;
	method: string;
};

export type PermissionPartialType = {
	items: Partial<PermissionType>[];
};

export type RoleTypePartialType = Partial<RoleType>;
export type RolePageData = PageDataType<RoleType>;

class RoleService {
	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async page(query: string = "") {
		const { data } = await api.get<RolePageData>(`/role/page?${query}`);
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
	async permissions(q: string = "") {
		const { data } = await api.get<PermissionPartialType>(`/role/permissions?${q}`);
		return data.items;
	}

	@Try({
		async onError(err) {
			(await import("@/common/Notify")).ErrorNotify(
				// @ts-ignore
				err?.response?.data.message || err.message,
			);
		},
	})
	async getByID(id: number) {
		const { data } = await api.get<RoleType>(`/role/${id}`);
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
	async search(query: string = "") {
		const { data } = await api.get<RoleType[]>(`/role/search?${query}`);

		return data || [];
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
	async createOrUpdate(role: Partial<RoleCreateType>) {
		const { data } = await api.put<IdType>(`/role`, role);

		return data;
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
	async delete(languageId: number, categoery: string, key: string) {
		await api.delete(`/language-content/${languageId}/${categoery}/${key}`);
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
	async restore(id: number) {
		await api.patch(`/language/${id}/restore`);
		return true;
	}
}

const roleService = new RoleService();

export { roleService as RoleService };
