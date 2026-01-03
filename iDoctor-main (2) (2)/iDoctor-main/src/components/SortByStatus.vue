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
	options = ["ACTIVE", "INACTIVE"], // Исправлено: строковые значения для сортировки
	name = "active", // Исправлено: более подходящее имя для сортировки
	displayValue = "sort_by_status", // Исправлено: более подходящее значение для отображения
} = defineProps<Props>();

const emits = defineEmits<{
	(e: "sortChange", sort: string): void; // Исправлено: более подходящее название события
}>();

const router = useRouter();
const route = useRoute();

// Исправлено: убран let, изменен тип на string
const sortValue = ref<string>();

// Исправлено: название функции и параметров
async function sortChange(newSort: string) {
	const { query } = route;
	const updatedQuery = { ...query };

	if (!newSort) {
		delete updatedQuery[name]; // Удаляем параметр, если пустое значение
	} else {
		updatedQuery[name] = newSort;
	}

	await router.push({
		force: true,
		query: updatedQuery,
	});

	sortValue.value = newSort;
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
		@clear="sortChange('')"
		:clearable="true"
		@update:model-value="sortChange"
		color="accent"
	/>
</template>
