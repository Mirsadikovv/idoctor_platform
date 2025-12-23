<script setup lang="ts">
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";

export interface Props {
	total: number;
	maxPages?: number;
	name?: string;
}

const { total, name = "page", maxPages = 6 } = defineProps<Props>();

const router = useRouter();
const route = useRoute();

const page = ref(currentPage());

const emits = defineEmits<{
	(e: "page", page: number): void;
}>();

function currentPage() {
	const queryPage = +route.query[name]!;
	const page = isNaN(queryPage) ? 1 : queryPage;
	return page < total ? page : total;
}

async function pageChange(page: number) {
	const { query } = route;

	query[name] = page as any;

	await router.push({
		force: true,
		query: {
			...query,
		},
	});

	emits("page", page);
}
</script>

<template>
	<q-pagination
		class="flex flex-center"
		@update:model-value="pageChange"
		v-model="page"
		:max="total"
		:max-pages="maxPages"
		boundary-numbers
		boundary-links
		icon-first="keyboard_double_arrow_left"
		icon-last="keyboard_double_arrow_right"
		color="secondary"
	/>
</template>
