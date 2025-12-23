import { buildQuery } from "@/common";
import { DeviceService, SupplierService } from "@/service";

export async function searchDevices(params: string) {
	const searchParams = buildQuery({
		name: params,
	});

	const res = await DeviceService.search(searchParams);
	return res;
}
export async function searchSuppliers(params: string) {
	const searchParams = buildQuery({
		name: params,
	});

	const res = await SupplierService.search(searchParams);
	return res;
}
