<script setup lang="ts">
import type { QBtnSlots, QBtnProps } from "quasar";
import { computed, useAttrs } from "vue";

const props = withDefaults(defineProps<QBtnProps>(), {
	flat: true,
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
});

const attrs = useAttrs();

const useSlots = defineSlots<QBtnSlots>();

const slots = computed(() => Object.keys(useSlots) as (keyof QBtnSlots)[]);
</script>

<template>
	<q-btn v-bind="{ ...props, ...attrs }" type="submit" class="bg-secondary text-white px-10!">
		<slot v-for="slot of slots" :name="slot" />
	</q-btn>
</template>
