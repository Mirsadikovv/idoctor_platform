<script setup lang="ts" generic="T, E extends T extends object ? T | T[keyof T] : T">
import { QSelect, type QSelectProps } from "quasar";
import { ref, useAttrs } from "vue";
// T extends Record<string, any> ? T | T[keyof T] : T

export type QSelectOverrides = Omit<
	QSelectProps,
	| "modelValue"
	| "onUpdate:modelValue"
	| "onClear"
	| "onRemove"
	| "onAdd"
	| "onNewValue"
	| "useInput"
	| "loading"
	| "optionValue"
	| "optionLabel"
	| "optionDisable"
	| "mapOptions"
	| "emitValue"
>;

export interface Props<T, E> extends QSelectOverrides {
	items?: T[];
	modelValue?: E[] | E | any;
	find?: (text: string | "") => Promise<T[]>;
	optionValue?:
		| ((option: T) => E)
		| (T extends Record<string, any> ? Extract<keyof T, string> : string | undefined);
	optionLabel?:
		| ((option: T) => any)
		| (T extends Record<string, any> ? Extract<keyof T, string> : string | undefined);

	optionDisable?:
		| ((option: T) => boolean)
		| (T extends Record<string, any> ? Extract<keyof T, string> : string | undefined);

	"onUpdate:modelValue"?: (value: T) => void;
	onAdd?: (details: { index: number; value: T }) => void;
	onClear?: (value: T) => void;
	onRemove?: (details: { index: number; value: T }) => void;
	onNewValue?: (
		inputValue: string,
		doneFn: (item?: T, mode?: "add" | "add-unique" | "toggle") => void,
	) => void;
}

const props = withDefaults(defineProps<Props<T, E>>(), {
	find: async () => [],
	inputDebounce: 300,
	outlined: true,
	clearable: true,
	virtualScrollHorizontal: undefined,
	error: undefined,
	noErrorIcon: undefined,
	reactiveRules: undefined,
	lazyRules: undefined,
	stackLabel: undefined,
	hideHint: undefined,
	dark: undefined,
	loading: undefined,
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
	displayValueHtml: undefined,
	hideSelected: undefined,
	hideDropdownIcon: undefined,
	optionsDense: undefined,
	optionsDark: undefined,
	optionsHtml: undefined,
	optionsCover: undefined,
	menuShrink: undefined,
	popupNoRouteDismiss: undefined,
	useChips: undefined,
	fillInput: undefined,
	mapOptions: undefined,
	disableTabSelection: undefined,
	emitValue: undefined,
	color: "secondary",
	isStatic: false,
});

const emits = defineEmits<{
	(e: "update", item: T | number): void;
	(e: "update:items", items: T[]): void;
}>();

const loading = ref(false);

const attrs = useAttrs();

const modelValue = defineModel<E>("modelValue");

const items = defineModel<T[]>("items");

async function find(text: string = "") {
	loading.value = true;
	try {
		const result = await props.find(text);
		items.value = result;
		updateItems(result);
	} catch (error) {
	} finally {
		loading.value = false;
	}
}

async function filter(
	inputValue: string = "",
	update: (callbackFn: () => void, afterFn?: (ref: QSelect) => void) => void,
	_abortFn: () => void,
) {
	return update(async () => {
		await find(inputValue);
	});
}

function update(item: T) {
	emits("update", item);
}

function updateItems(item: T[]) {
	emits("update:items", item);
}
</script>

<template>
	<q-select
		v-bind="{ ...props, ...attrs }"
		:loading="loading"
		use-input
		map-options
		emit-value
		v-model="modelValue"
		:options="items"
		@update:model-value="update"
		@filter="filter"
		:label="$tl(label)"
	>
		<slote></slote>
	</q-select>
</template>
