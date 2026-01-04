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

/**
 * // Order Status Constants
const (
  StatusPending    = "PENDING"
  StatusReceived   = "RECEIVED"
  StatusDiagnostic = "DIAGNOSTIC"
  StatusCompleted  = "COMPLETED"
  StatusCancelled  = "CANCELLED"
)

// Payment Type Constants
const (
  PaymentTypeCash   = "CASH"
  PaymentTypeCard   = "CARD"
  PaymentTypeOnline = "ONLINE"
)

// Payment Status Constants
const (
  PaymentStatusPending   = "PENDING"
  PaymentStatusPaid      = "PAID"
  PaymentStatusRefunded  = "REFUNDED"
  PaymentStatusCancelled = "CANCELLED"
)
 * 
 */

export const orderStatusOptions = [
	{ label: "PENDING", value: "PENDING" },
	{ label: "RECEIVED", value: "RECEIVED" },
	{ label: "DIAGNOSTIC", value: "DIAGNOSTIC" },
	{ label: "COMPLETED", value: "COMPLETED" },
	{ label: "CANCELLED", value: "CANCELLED" },
];

export const paymentTypeOptions = [
	{ label: "CASH", value: "CASH" },
	{ label: "CARD", value: "CARD" },
	{ label: "ONLINE", value: "ONLINE" },
];

export const paymentStatusOptions = [
	{ label: "PENDING", value: "PENDING" },
	{ label: "PAID", value: "PAID" },
	{ label: "REFUNDED", value: "REFUNDED" },
	{ label: "CANCELLED", value: "CANCELLED" },
];
