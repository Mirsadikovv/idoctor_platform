<script setup lang="ts">
import type { QBtnProps } from "quasar";
import { type VNode } from "vue";

export interface Props extends QBtnProps {
	withTooltip?: boolean;
	tooltipText?: string;
	uiType?: "default" | "danger" | "positive" | "error" | "secondary";
}

const props = withDefaults(defineProps<Props>(), {
	replace: undefined,
	outline: undefined,
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
	round: undefined,
	darkPercentage: undefined,
	padding: "8px 10px",
	withTooltip: false,
	tooltipText: "",
	align: "between",
	flat: true,
	color: "accent",
	uiType: "default",
});

const slots = defineSlots<{
	default: () => VNode[];
	label: () => VNode[];
	loading: () => VNode[];
}>();
</script>

<template>
	<q-btn
		:class="[
			'min-w-[100px]! clean-logout-btn bg-secondary flex',
			{
				'btn-danger': props.uiType === 'danger',
				'btn-positive': props.uiType === 'positive',
				'btn-error': props.uiType === 'error',
				'btn-default': props.uiType === 'default',
				'btn-secondary': props.uiType === 'secondary',
			},
		]"
		v-bind="{
			...{
				...props,
				color:
					props.uiType === 'danger'
						? 'red-6'
						: props.uiType === 'positive'
						? 'green-6'
						: props.uiType === 'error'
						? 'red-6'
						: props.uiType === 'secondary'
						? '#424242'
						: 'primary',
				outline: props.uiType === 'danger',
				flat: false,
				textColor:
					props.uiType === 'danger'
						? 'red-6'
						: props.uiType === 'error'
						? 'white'
						: props.uiType === 'secondary'
						? 'white'
						: undefined,
				uiType: undefined,
			},
			...$attrs,
		}"
	>
		<q-tooltip
			v-if="props.withTooltip"
			anchor="top middle"
			self="center left"
			class="bg-accent text-white"
		>
			{{ $tl(props.tooltipText) }}
		</q-tooltip>
		<template v-for="(slotFn, name) in slots" #[name]>
			<component :is="slotFn" />
		</template>
	</q-btn>
</template>

<style lang="scss">
.q-btn__content {
	gap: 5px;
	justify-content: center;
	width: 100% !important;
}

// Default button styles - filled blue button
.btn-default {
	@apply text-white!;

	&:hover {
		@apply bg-blue-700!;
	}
}

// Danger button styles - no border, light red background
.btn-danger {
	@apply bg-red-50! text-red-600!;

	&:hover {
		@apply bg-red-100!;
	}
}

// Positive button styles - light green background
.btn-positive {
	@apply bg-green-50! text-green-600!;

	&:hover {
		@apply bg-green-100!;
	}
}

// Error button styles - filled red button
.btn-error {
	@apply text-white!;

	&:hover {
		@apply bg-red-700!;
	}
}

// Secondary button styles - filled secondary button
.btn-secondary {
	background-color: #757575 !important;
	color: white !important;

	&:hover {
		background-color: #616161 !important;
	}
}
</style>
