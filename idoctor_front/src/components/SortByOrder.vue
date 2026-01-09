<script setup lang="ts">
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";

export interface Props {
	options?: string[];
	name?: string;
	displayValue?: string;
	label?: string;
}

const {
	options = ["ASC", "DESC"], // Исправлено: строковые значения для сортировки
	name = "order", // Исправлено: более подходящее имя для сортировки
	displayValue = "sort_by_order", // Исправлено: более подходящее значение для отображения
} = defineProps<Props>();

const emits = defineEmits<{
	(e: "sortChange", sort: string): void; // Исправлено: более подходящее название события
}>();

const router = useRouter();
const route = useRoute();

// Исправлено: убран let, изменен тип на string
const sortValue = ref<string>(currentSortValue());

function currentSortValue(): string {
	// Исправлено: тип возвращаемого значения
	const querySort = route.query[name] as string; // Исправлено: приведение к строке
	const currentSort = querySort && options.includes(querySort) ? querySort : options[0];
	return currentSort;
}

// Исправлено: название функции и параметров
async function sortChange(newSort: string) {
	const { query } = route;
	query[name] = newSort; // Исправлено: уже строка, не нужно преобразование

	await router.push({
		force: true,
		query: {
			...query,
		},
	});

	sortValue.value = newSort; // Исправлено: обновление реактивного значения
	emits("sortChange", newSort);
}
</script>

<template>
	<q-select
		v-model="sortValue"
		:display-value="$tl(displayValue)"
		:options="options"
		:option-label="(option) => $tl(option) || ''"
		outlined
		dense
		:clearable="false"
		@update:model-value="sortChange"
		color="accent"
	/>
</template>
