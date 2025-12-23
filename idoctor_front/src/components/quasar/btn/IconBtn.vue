<script setup lang="ts">
import type { ProcessedIcons } from "@/common";
import type { QBtnProps } from "quasar";

// Объединяем все пропсы в один интерфейс
export interface Props extends QBtnProps {
	icon: ProcessedIcons;
	withTooltip?: boolean;
	tooltipText?: string;
}

const props = withDefaults(defineProps<Props>(), {
	color: "secondary",
	flat: false,
	round: true,
	replace: undefined,
	outline: true,
	unelevated: undefined,
	rounded: undefined,
	push: undefined,
	square: undefined,
	glossy: undefined,
	fab: undefined,
	fabMini: undefined,
	noCaps: undefined,
	noWrap: undefined,
	dense: undefined,
	ripple: undefined,
	stack: undefined,
	stretch: undefined,
	loading: undefined,
	disable: undefined,
	darkPercentage: undefined,
	size: "12px",
	padding: "xs",
	withTooltip: false,
	tooltipText: "",
});
</script>

<template>
	<q-btn class="mr-2" v-bind="{ ...props, ...$attrs }">
		<q-tooltip
			v-if="props.withTooltip"
			anchor="top middle"
			self="center left"
			class="bg-accent text-white"
		>
			{{ $tl(props.tooltipText) }}
		</q-tooltip>
		<slot />
	</q-btn>
</template>
