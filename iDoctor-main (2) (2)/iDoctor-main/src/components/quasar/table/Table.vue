<script setup lang="ts" generic="T">
import type { PageDataType } from "@/service";
import SkeletonTable from "./SkeletonTable.vue";

export interface Props<T> {
	models: PageDataType<T>;
	loading?: boolean;
	pick?: Record<string, boolean>;
	traction?: (model: T, index: number) => any;
	customRowStyle?: (model: T, index: number) => any;
	whithoutMaxHeight?: boolean;
	hasOrder?: boolean;
}

type TheadSlot = {
	[K in keyof T as `${K & string}:thead`]: (_: { key: keyof T }) => any;
};

type TafterSlot = {
	[K in keyof T as `${K & string}:tafter`]: (_: { key: keyof T }) => any;
};

type BodySlot = {
	[key in keyof T]: (_: { model: T; index: number }) => any;
};

type DynamicBodySlot = {
	[key: string]: (_: { model: T; index: number }) => any;
};

type DynamicHeadSlot = {
	[key: `${string}:thead`]: (_: { key: string }) => any;
};

type ExistsType = (key: keyof T) => boolean;

type SlotTheadFoot = {
	before: (_: { models: T[]; exists: ExistsType; totalPages: number }) => any;
	thead: (_: { models: T[]; exists: ExistsType; totalPages: number }) => any;
	tfoot: (_: { models: T[]; exists: ExistsType; totalPages: number }) => any;
};

const slots = defineSlots<
	BodySlot & TheadSlot & SlotTheadFoot & DynamicBodySlot & DynamicHeadSlot & TafterSlot
>();

const {
	models = {
		data: [],
		totalPages: 0,
		currentPage: 0,
		pageSize: 0,
	},
	loading = false,
	pick,
	traction = (_: T, _i: number) => undefined,
	customRowStyle = (_: T, _i: number) => undefined,
	whithoutMaxHeight = false,
} = defineProps<Props<T>>();

function fields(slots: any) {
	const keys = Object.keys(slots).filter(
		(k) =>
			!k.endsWith(":thead") &&
			!k.endsWith("thead") &&
			!k.endsWith("tfoot") &&
			!k.endsWith("before") &&
			!k.endsWith("tafter"),
	);

	if (!Object.keys(pick || {}).length) {
		return keys;
	}

	return keys.filter((k) => {
		//@ts-ignore
		const val = pick[k];
		if (val === undefined) {
			return false;
		}

		if (val === false) {
			return true;
		}

		return true;
	});
}

function exists(_key: string): boolean {
	if (!pick) {
		return true;
	}

	return true;
}
</script>

<template>
	<q-markup-table
		v-if="!loading"
		class="qqy"
		:style="whithoutMaxHeight ? '' : 'max-height: calc(100vh - 280px)'"
		wrap-cells
		separator="cell"
		bordered
		dense
		flat
	>
		<thead>
			<slot
				v-if="!!slots['thead']"
				:name="//@ts-ignore
				`thead`"
				:="{ models: models.data, exists, totalPages: models.totalPages }"
			></slot>

			<tr>
				<th
					class="text-center text-base! font-bold! w-1% px-4! sticky top-0 bg-white! z-100 select-auto!"
					style="min-height: 40px"
					v-if="hasOrder"
				>
					№
				</th>
				<th
					class="text-left text-base! font-bold! sticky top-0 bg-white! z-100 py-1! select-auto!"
					style="min-height: 40px"
					v-for="(key, i) of fields(slots)"
					:key="i"
				>
					<slot
						:name="//@ts-ignore
						`${key}:thead`"
						:="{ key }"
					>
						{{ $tl(key) }}
					</slot>
				</th>
			</tr>

			<tr v-if="Object.keys(slots).some((k) => k.endsWith('tafter'))">
				<th
					class="text-left text-base! font-bold! sticky top-0 bg-white! z-100"
					style="min-height: 40px"
					v-for="(key, i) of fields(slots)"
					:key="i"
				>
					<slot
						:name="//@ts-ignore
						`${key}:tafter`"
						:="{ key }"
					>
					</slot>
				</th>
			</tr>
		</thead>

		<tbody>
			<slot
				v-if="!!slots['before']"
				:name="//@ts-ignore
				`before`"
				:="{ models: models.data, exists, totalPages: models.totalPages }"
			></slot>
			<tr v-if="!models.data.length">
				<td
					:colspan="fields(slots).length + (hasOrder ? 1 : 0)"
					class="text-center py-16 text-grey-6"
					style="height: 200px"
				>
					{{ $tl("no_data") }}
				</td>
			</tr>
			<tr
				v-else
				v-for="(data, i) of models.data"
				:key="i"
				:class="traction(data, i)"
				:style="customRowStyle(data, i)"
			>
				<td
					class="text-center text-sm! w-1% px-4!"
					style="min-height: 40px"
					v-if="hasOrder"
				>
					{{ (models.currentPage - 1) * models.pageSize + i + 1 }}
				</td>
				<td
					class="text-left text-sm!"
					style="min-height: 40px"
					v-for="(field, j) in fields(slots)"
					:key="j"
				>
					<slot
						:="{ model: data, index: i }"
						:name="//@ts-ignore
						`${field as keyof T}`"
					>
						{{
							//@ts-ignore
							data?.[field]
						}}
					</slot>
				</td>
			</tr>
		</tbody>
		<tfoot v-if="!!slots['tfoot']" class="sticky bottom-0 z-100 bg-white">
			<slot
				:name="// @ts-ignore
				`tfoot`"
				:="{ models: models.data, exists, totalPages: models.totalPages }"
			>
			</slot>
		</tfoot>
	</q-markup-table>
	<SkeletonTable v-else></SkeletonTable>
</template>

<style scoped lang="scss">
.qqy > table {
	position: relative;
	border-collapse: separate;
	border-spacing: 0;
}

// Анимация строк
.qqy tbody tr {
	transition: all 0.3s ease;

	&:hover {
		background: rgba($primary, 0.08);
		box-shadow: 0 2px 8px rgba($primary, 0.15);
		border-left: 4px solid $primary;
	}

	&:nth-child(even):not(:hover) {
		background: rgba(0, 0, 0, 0.02);
	}
}

// Анимация заголовков
.qqy th {
	background: #f8f9fa;
	transition: all 0.3s ease;

	&:hover {
		background: $primary;
		font-weight: 600;
		box-shadow: 0 2px 4px rgba($primary, 0.3);
	}
}
</style>
