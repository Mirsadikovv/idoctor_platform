<script setup lang="ts" generic="T extends string">
import { QRouteTab, QTabPanel, QTabPanels, QTabs, type QTabsProps } from "quasar";
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
	align: "justify",
	vertical: undefined,
	stretch: undefined,
	shrink: undefined,
	switchIndicator: undefined,
	noCaps: undefined,
	dense: undefined,
	isFullWidth: false,
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
	<div class="row items-center gap-x-4">
		<div class="col">
			<q-tabs
				v-bind="props"
				v-model="modelValue"
				v-bind:="$attrs"
				:style="isFullWidth ? { width: '100%' } : {}"
				class="text-primary bg-[#dee3ef] inline-block rounded-sm"
				active-bg-color="accent"
				active-color="accent"
				indicator-color="accent"
				active-class="text-white overflow-visible"
				align="justify"
			>
				<slot :name="`before`" :="{ index: 0, name: '' }"></slot>

				<q-route-tab
					v-for="(slot, index) of tabs($slots as any)"
					no-caps
					:name="index"
					:label="$tl(rename(slot))"
					:to="{ query: { tab: rename(slot) } }"
				>
					<slot :name="(slot as TablType)" :="{ index, name: slot }"></slot>
				</q-route-tab>
				<slot :name="`after`" :="{ index: 0, name: '' }"></slot>
			</q-tabs>
		</div>
		<div class="col-auto" v-if="$slots['after-tab']">
			<slot name="after-tab" v-bind="{ index: 0, name: 'afterTab' }"></slot>
		</div>
	</div>
	<div class="col-auto" v-if="$slots['after-tab']">
		<slot name="after-tab" v-bind="{ index: 0, name: 'afterTab' }"></slot>
	</div>
	<q-tab-panels
		v-model="modelValue"
		animated
		transition-duration="500"
		transition-next="fade"
		transition-prev="fade"
	>
		<q-tab-panel
			class="p-0!"
			v-for="(slot, index) of panels($slots as any)"
			:key="index"
			:name="index"
		>
			<slot :name="(slot as PanelType)" :="{ index, name: slot }">
				{{ slot }}
			</slot>
		</q-tab-panel>
	</q-tab-panels>
</template>
