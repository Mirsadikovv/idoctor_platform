<script setup lang="ts" generic="T extends string">
import { QTab, QTabPanel, QTabPanels, QTabs, type QTabsProps } from "quasar";
import { onMounted } from "vue";
export interface Props<T> extends QTabsProps {
	active?: T;
	items?: T[];
	isFullWidth?: boolean;
}
type OptionsType = "before" | "after";
type ArrayType = NonNullable<typeof props.items>[number];
type TablType = `tab${Capitalize<ArrayType>}`;
type PanelType = `panel${Capitalize<ArrayType>}`;

type SlotsValue = {
	index: number;
	name: string;
};

type SlotType = {
	[key in TablType | PanelType | OptionsType]: (_: SlotsValue) => any;
};

type AnySlotType = {
	[key: string]: (_: SlotsValue) => any;
};

type ALLType = SlotType & Record<string, AnySlotType | never>;

const props = withDefaults(defineProps<Props<T>>(), {
	narrowIndicator: false,
	outsideArrows: true,
	mobileArrows: true,
	inlineLabel: true,
	isFullWidth: false,
	align: "justify",
	vertical: undefined,
	stretch: undefined,
	shrink: undefined,
	switchIndicator: undefined,
	noCaps: undefined,
	dense: undefined,
});

const modelValue = defineModel("tab");

defineSlots<ALLType>();

function tabs(slots: SlotsValue) {
	return Object.keys(slots || {}).filter((s) => s.startsWith("tab"));
}

function panels($slots: SlotsValue) {
	return Object.keys($slots || {}).filter((s) => s.startsWith("panel"));
}

function rename(slot: string) {
	return slot?.replace("tab", "")?.replace("panel", "")?.toLowerCase() as PanelType;
}

onMounted(() => {
	if (!modelValue.value) {
		modelValue.value = 0;
	}
});
</script>

<template>
	<q-tabs
		v-bind="props"
		v-model="modelValue"
		v-bind:="$attrs"
		:style="isFullWidth ? { width: '100%' } : {}"
		class="text-primary bg-[#dee3ef] inline-block rounded"
		active-bg-color="info"
		active-color="info"
		indicator-color="info"
		active-class="text-white overflow-visible"
	>
		<slot :name="`before`" :="{ index: 0, name: '' }"></slot>
		<q-tab
			v-for="(slot, index) of tabs($slots as any)"
			:name="index"
			:label="$tl(rename(slot))"
			class=""
		>
			<slot :name="(slot as TablType)" :="{ index, name: slot }"></slot>
		</q-tab>
		<slot :name="`after`" :="{ index: 0, name: '' }"></slot>
	</q-tabs>
	<q-tab-panels
		v-model="modelValue"
		animated
		transition-duration="500"
		transition-next="fade"
		transition-prev="fade"
	>
		<q-tab-panel
			v-for="(slot, index) of panels($slots as any)"
			:key="index"
			:name="index"
			class="p-0! bg-light!"
		>
			<slot :name="(slot as PanelType)" :="{ index, name: slot }">
				{{ slot }}
			</slot>
		</q-tab-panel>
	</q-tab-panels>
</template>
