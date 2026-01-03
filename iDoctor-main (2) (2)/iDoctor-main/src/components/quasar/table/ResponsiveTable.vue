<script setup lang="ts" generic="T">
import type { PageDataType } from "@/service";
import Table from './Table.vue';
import MobileList from './MobileList.vue';
import { useBreakpoints } from '@/composables/useBreakpoints';
import { computed } from 'vue';

export interface Props<T> {
	models: PageDataType<T>;
	loading?: boolean;
	pick?: Record<string, boolean>;
	traction?: (model: T, index: number) => any;
	customRowStyle?: (model: T, index: number) => any;
	whithoutMaxHeight?: boolean;
	hasOrder?: boolean;
	mobileView?: 'cards' | 'accordion';
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

type CardSlot = {
	card: (_: { model: T; index: number; fields: string[]; orderNumber?: number }) => any;
};

type SlotTheadFoot = {
	before: (_: { models: T[]; exists?: ExistsType; totalPages: number }) => any;
	thead: (_: { models: T[]; exists?: ExistsType; totalPages: number }) => any;
	tfoot: (_: { models: T[]; exists?: ExistsType; totalPages: number }) => any;
};

const slots = defineSlots<
	BodySlot & 
	TheadSlot & 
	SlotTheadFoot & 
	DynamicBodySlot & 
	DynamicHeadSlot & 
	TafterSlot &
	CardSlot
>();

const {
	mobileView = 'cards',
	...tableProps
} = defineProps<Props<T>>();

const { isMobile } = useBreakpoints();

// Вычисляемые свойства для передачи в компоненты
const tableSlots = computed(() => {
	const result: Record<string, any> = {};
	Object.keys(slots).forEach(key => {
		if (key !== 'card') {
			result[key] = slots[key as keyof typeof slots];
		}
	});
	return result;
});

const mobileSlots = computed(() => {
	const result: Record<string, any> = {};
	Object.keys(slots).forEach(key => {
		// Исключаем слоты которые не нужны мобильному компоненту
		if (!['thead', 'customRowStyle', 'whithoutMaxHeight'].includes(key)) {
			result[key] = slots[key as keyof typeof slots];
		}
	});
	return result;
});
</script>

<template>
	<Table 
		v-if="!isMobile" 
		v-bind="tableProps"
	>
		<template 
			v-for="(_, name) in tableSlots" 
			:key="name" 
			#[name]="slotData"
		>
			<slot :name="name" v-bind="slotData" />
		</template>
	</Table>
	
	<MobileList 
		v-else 
		v-bind="{
			models: tableProps.models,
			loading: tableProps.loading,
			pick: tableProps.pick,
			traction: tableProps.traction,
			hasOrder: tableProps.hasOrder
		}"
	>
		<template 
			v-for="(_, name) in mobileSlots" 
			:key="name" 
			#[name]="slotData"
		>
			<slot :name="name" v-bind="slotData" />
		</template>
	</MobileList>
</template>

<style scoped lang="scss">
// Дополнительные стили если нужны для переходов
.responsive-table-container {
	transition: all 0.3s ease;
}
</style>