<script setup lang="ts">
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";

export interface Props {
	options?: number[];
	name?: string;
	displayValue?: string;
}

const { options = [10, 15, 25], name = "perpage", displayValue = "perpage" } = defineProps<Props>();

const emits = defineEmits<{
	(e: "perPage", page: number): void;
}>();

const router = useRouter();
const route = useRoute();
let perPage = ref<number>(currentPerPage());

function currentPerPage() {
	const queryPerPage = +route.query[name]!;
	const perPage = isNaN(queryPerPage) ? options[0] : queryPerPage;
	return perPage;
}

async function perPageChnage(perPage: number) {
	const { query } = route;

	query[name] = perPage as any;

	await router.push({
		force: true,
		query: {
			...query,
		},
	});

	emits("perPage", perPage);
}
</script>

<template>
	<q-select
		v-model="perPage"
		:label="$tl(displayValue)"
		:options="options"
		borderless
		dense
		@update:model-value="perPageChnage"
		color="secondary"
		class="min-w-[80px]"
	/>
</template>
