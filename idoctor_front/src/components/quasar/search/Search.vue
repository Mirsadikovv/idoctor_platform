<script setup lang="ts" generic="T extends string | number | null | undefined">
import type { QInputProps, QInputSlots } from "quasar";
import { computed, ref } from "vue";
import { useRoute, useRouter } from "vue-router";

export interface Props<T> extends Omit<QInputProps, "modelValue"> {
	modelValue?: T;
	modelModifiers?: Record<string, boolean>;
	queryName?: string;
}

type QInputSlotsType = keyof QInputSlots;

const props = withDefaults(defineProps<Props<T>>(), {
	outlined: true,
	clearable: true,
	debounce: 400,
	color: "accent",
	queryName: "search",
	fillMask: undefined,
	reverseFillMask: undefined,
	unmaskedValue: undefined,
	error: undefined,
	noErrorIcon: undefined,
	reactiveRules: undefined,
	lazyRules: true,
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
	autogrow: undefined,
});

const emits = defineEmits<{
	(e: "search", search: string): void;
}>();

const useSlots = defineSlots<QInputSlots>();

const slots = computed(() => Object.keys(useSlots) as QInputSlotsType[]);

const router = useRouter();
const route = useRoute();

const modelValue = defineModel<T>();

const search = ref(currentSearch());

function currentSearch(): string {
	return (route.query[props.queryName] as string) || "";
}

async function searchChange(search: string | null | number) {
	const _search = `${search ?? ""}`;

	await router.push({
		force: true,
		query: {
			[props.queryName]: _search,
		},
	});

	emits("search", _search);
}

function clear() {
	search.value = "";
	searchChange("");
}

defineExpose({
	clear,
});
</script>

<template>
	<q-input :="props" :label="$tl(label)" v-model="search" @update:model-value="searchChange">
		<slot v-for="slot of slots" :name="slot"></slot>
	</q-input>
</template>
