import type { PageDataType } from "@/service";

export function useTableData<T>() {
	function extractFields(slots: any, pick?: Record<string, boolean>) {
		const keys = Object.keys(slots).filter(
			(k) =>
				!k.endsWith(":thead") &&
				!k.endsWith("thead") &&
				!k.endsWith("tfoot") &&
				!k.endsWith("before") &&
				!k.endsWith("tafter") &&
				!k.endsWith("card")
		);

		if (!Object.keys(pick || {}).length) {
			return keys;
		}

		return keys.filter((k) => {
			const val = pick![k];
			if (val === undefined) {
				return false;
			}

			return true;
		});
	}

	function hasSlot(slotName: string, slots: any): boolean {
		return !!slots[slotName];
	}

	function getSlotNames(slots: any) {
		return {
			thead: Object.keys(slots).filter(k => k.endsWith(":thead")),
			tafter: Object.keys(slots).filter(k => k.endsWith(":tafter")),
			body: Object.keys(slots).filter(k => !k.includes(":")),
			before: hasSlot("before", slots),
			tfoot: hasSlot("tfoot", slots),
			card: hasSlot("card", slots)
		};
	}

	function calculateOrderNumber(
		models: PageDataType<T>, 
		index: number
	): number {
		return (models.currentPage - 1) * models.pageSize + index + 1;
	}

	return {
		extractFields,
		hasSlot,
		getSlotNames,
		calculateOrderNumber
	};
}