import { buildQuery } from "@/common";
import { RoleService } from "@/service";

export async function searchPermission(
	params: string = "",
	type: string = "",
	page_name: string = "",
) {
	const query = buildQuery({
		perpage: 20,
		type: type,
		page_name: page_name,
		search: params,
	});
	const data = await RoleService.permissions(query);
	return data;
}

export async function getAllPermissions(params: string = "") {
	const query = buildQuery({
		search: params,
		without_paginate: true,
	});
	const data = await RoleService.permissions(query);
	return data;
}
