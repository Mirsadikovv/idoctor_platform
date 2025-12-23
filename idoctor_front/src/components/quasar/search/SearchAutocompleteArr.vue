<script setup lang="ts" generic="T, E extends T extends object ? T | T[keyof T] : T">
import { type QSelectProps } from "quasar";
import { ref, nextTick, watch, useAttrs } from "vue";
import { useRoute, useRouter } from "vue-router";
import { searchComplaintType } from "@/modules/Appeal/shared/utils";

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
	queryName: string;
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
}

const props = withDefaults(defineProps<Props<T, E>>(), {
	find: async (text: string = ""): Promise<any[]> => {
		const langId = 2;
		return await searchComplaintType(text, langId);
	},
	inputDebounce: 300,
	outlined: true,
	clearable: true,
});

const emits = defineEmits<{
	(e: "update", item: T | E | (T | E)[] | undefined): void;
	(e: "search", search: Record<string, any>): void;
	(e: "update:items", items: T[]): void;
}>();

const loading = ref(false);
const attrs = useAttrs();
const router = useRouter();
const route = useRoute();
const modelValue = defineModel<E>("modelValue");
const items = defineModel<T[]>("items");

async function find(text: string = "") {
	loading.value = true;
	try {
		const result = await props.find(text);
		items.value = result;
		emits("update:items", result);
	} catch (error) {
		console.error("Find error", error);
	} finally {
		loading.value = false;
	}
}

async function filter(
	inputValue: string = "",
	update: (callbackFn: () => void, afterFn?: (ref: any) => void) => void,
	_abortFn: () => void,
) {
	return update(async () => {
		await find(inputValue);
	});
}

watch(
	modelValue,
	async (val) => {
		await nextTick();
		await searchChange(val);
		emits("update", val);
	},
	{ deep: true },
);

async function searchChange(item: any) {
	const newQuery: Record<string, any> = { ...route.query };

	if (!item || (Array.isArray(item) && item.length === 0)) {
		delete newQuery[props.queryName];
	} else if (Array.isArray(item)) {
		newQuery[props.queryName] = item
			.map((i: any) => {
				if (typeof i === "object") {
					const valKey = props.optionValue as keyof T;
					return i?.[valKey] ?? i?.value ?? i?.id ?? "";
				}
				return String(i);
			})
			.filter(Boolean)
			.join(",");
	} else {
		newQuery[props.queryName] = String(item);
	}

	await router.replace({ query: newQuery });

	requestAnimationFrame(() => {
		emits("search", { ...newQuery });
	});
}
</script>

<template>
	<q-select
		:="{ ...props, ...attrs }"
		:label="$tl(props.label || props.queryName || 'search')"
		:loading="loading"
		use-input
		map-options
		emit-value
		v-model="modelValue"
		:options="items"
		@filter="filter"
		@focus="find()"
		color="accent"
		outlined
		use-chips
		multiple
	/>
</template>
