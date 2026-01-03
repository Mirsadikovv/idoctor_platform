<script
	setup
	lang="ts"
	generic="T extends string | number | object | FileList | null | undefined | boolean"
>
import type { QInputProps, QInputSlots } from "quasar";
import { computed } from "vue";

export interface Props<T> extends Omit<QInputProps, "modelValue"> {
	modelValue?: T;
}

type QInputSlotsType = keyof QInputSlots;

const props = withDefaults(defineProps<Props<T>>(), {
	outlined: true,
	fillMask: undefined,
	reverseFillMask: undefined,
	unmaskedValue: undefined,
	error: undefined,
	noErrorIcon: undefined,
	reactiveRules: undefined,
	lazyRules: undefined,
	stackLabel: undefined,
	hideHint: undefined,
	dark: undefined,
	loading: undefined,
	clearable: undefined,
	filled: undefined,
	borderless: undefined,
	standout: undefined,
	labelSlot: undefined,
	bottomSlots: undefined,
	hideBottomSpace: undefined,
	counter: undefined,
	rounded: undefined,
	square: undefined,
	dense: undefined,
	itemAligned: undefined,
	disable: undefined,
	readonly: undefined,
	autofocus: undefined,
	autogrow: undefined,
});

const useSlots = defineSlots<QInputSlots>();

const slots = computed(() => Object.keys(useSlots) as QInputSlotsType[]);

const modelValue = defineModel<T>();
</script>

<template>
	<q-select :="props" :label="$tl(label?.toLowerCase())" v-model="modelValue">
		<slot v-for="slot of slots" :name="slot"></slot>
	</q-select>
</template>
