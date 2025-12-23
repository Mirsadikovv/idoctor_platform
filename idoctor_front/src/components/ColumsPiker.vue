<script setup lang="ts">
import { useRoute } from "vue-router";
import { computed, ref } from "vue";

export interface Props {
	picker: Record<string, boolean>;
	name?: string;
}

const { picker, name } = defineProps<Props>();

const model = defineModel({ default: {} });

const route = useRoute();

const routeName = computed(() => route.name as string);

const { columns, update } = buildColumsPiker(picker);

const pick = ref(parseLocalStorage("columns", columns));

function getColums(key: string): Record<string, string[]> | undefined {
	const raw = localStorage.getItem(key);

	if (!raw) {
		return undefined;
	}

	try {
		const cols = JSON.parse(raw);
		return cols;
	} catch (_) {}

	return undefined;
}

function parseLocalStorage(key: string, _default: string[]): string[] {
	try {
		const cols = getColums(key);

		if (!cols) {
			return [..._default];
		}

		const col = cols[routeName.value];

		if (!col?.length) {
			return [..._default];
		}

		const omits = Object.keys(picker).filter((k) => picker[k] === false);

		const result = [...new Set([...col, ...omits])];

		omits.forEach((key) => {
			//@ts-ignore
			model.value[key] = false;
		});

		col.forEach((key) => {
			//@ts-ignore
			model.value[key] = true;
		});

		return result;
	} catch (_) {}

	return [..._default];
}

function setLocalStorage(val: any) {
	if (!val) {
		return;
	}

	try {
		const key = "columns";

		const cols = getColums(key) || {};

		cols[routeName.value] = val;

		localStorage.setItem(key, JSON.stringify(cols));
	} catch (_) {}
}

function buildColumsPiker(args: Record<string, boolean>) {
	let columns: string[] = [];

	const omits: Record<string, boolean> = {};
	{
		const keys = Object.keys(args);

		keys.filter((key) => args[key] === false).forEach((key) => {
			omits[key] = false;
		});

		columns = keys.filter((key) => args[key] === true);
	}

	const update = (arg: string[]) => {
		const tmp = {
			...omits,
		};

		if (!arg) {
			model.value = tmp;
			return;
		}

		arg.forEach((key) => {
			tmp[key] = true;
		});

		model.value = tmp;

		setLocalStorage(arg);
	};

	return {
		update: update,
		columns: columns,
	};
}
</script>
<template>
	<!-- @vue-ignore -->
	<q-select
		v-model="pick"
		:options="columns"
		:option-label="(key) => $tl(key)"
		@update:model-value="update"
		:display-value="name || 'columns'"
		borderless
		dense
		multiple
		color="secondary"
	/>
</template>
