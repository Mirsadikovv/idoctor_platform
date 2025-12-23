import { buildQuery } from "@/common";
import { UserService, ProblemService, PartService } from "@/service";
import type { User, ProblemType, PartType } from "@/service";

export async function searchClients(query: string): Promise<User[]> {
	const params = buildQuery({
		fullname: query,
	});

	const users = await UserService.search(params);

	return users;
}

export async function searchMasters(query: string): Promise<User[]> {
	const params = buildQuery({
		fullname: query,
	});

	const users = await UserService.search(params);

	return users;
}

export async function searchProblems(query: string): Promise<ProblemType[]> {
	const params = buildQuery({
		name: query,
	});

	const problems = await ProblemService.search(params);

	return problems;
}

export async function searchParts(query: string): Promise<PartType[]> {
	const params = buildQuery({
		name: query,
	});

	const parts = await PartService.search(params);

	return parts;
}

export const orderStatusOptions = [
	{ label: "Новый", value: "new" },
	{ label: "В работе", value: "in_progress" },
	{ label: "Выполнен", value: "completed" },
	{ label: "Отменен", value: "cancelled" },
];

export const paymentStatusOptions = [
	{ label: "Не оплачен", value: "unpaid" },
	{ label: "Частично оплачен", value: "partially_paid" },
	{ label: "Оплачен", value: "paid" },
];

export const paymentTypeOptions = [
	{ label: "Наличные", value: "cash" },
	{ label: "Карта", value: "card" },
	{ label: "Безнал", value: "transfer" },
	{ label: "Онлайн", value: "online" },
];
