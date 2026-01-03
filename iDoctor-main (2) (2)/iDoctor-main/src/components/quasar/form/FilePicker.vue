<script setup lang="ts" generic="T extends File | FileList | File[] | null | undefined">
import type { QFileProps } from "quasar";

export interface Props<T> extends Omit<QFileProps, "modelValue"> {
	modelValue?: T;
	modelModifiers?: Record<string, boolean>;
}

const props = withDefaults(defineProps<Props<T>>(), {
	outlined: true,
	error: undefined,
	noErrorIcon: undefined,
	reactiveRules: undefined,
	lazyRules: true,
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
	multiple: undefined,
	appendToFiles: undefined,
	useChips: undefined,
	displayValue: undefined,
	counterLabel: undefined,
	tabindex: undefined,
	inputClass: undefined,
	inputStyle: undefined,
});

const modelValue = defineModel<T>();
</script>

<template>
	<q-file :="props" :label="$tl(label)" v-model="modelValue" clearable>
		<template v-for="(_, slot) of $slots" #[slot]="scope">
			<slot :name="slot" v-bind="scope || {}"></slot>
		</template>
	</q-file>
</template>
