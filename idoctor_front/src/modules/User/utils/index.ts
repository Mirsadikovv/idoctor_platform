import { RoleService } from "@/service";
import { UserService, type User } from "../service";

export async function searchUsers(query: string): Promise<User[]> {
	if (!query || query.length < 2) return [];

	const users = await UserService.search({
		fullname: query,
		limit: 20,
	});

	return users || [];
}

export function getUserFullName(user: User | undefined): string {
	if (!user) return "";

	const parts = [user.lastName, user.firstName, user.middleName].filter(Boolean);
	return parts.join(" ");
}

export function formatUserForDisplay(user: User): string {
	const fullName = getUserFullName(user);
	return fullName ? `${fullName} (${user.username})` : user.username;
}

export async function searchRole(params: string) {
	const roles = await RoleService.search(params);
	return roles;
}
