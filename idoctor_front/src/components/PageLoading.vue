<script setup lang="ts">
import { ref } from "vue";
import { useRoute } from "vue-router";

export interface Props {
	find: (query: string) => Promise<any>;
	fetch?: boolean;
}

const { find, fetch = true } = defineProps<Props>();
const route = useRoute();

let loading = ref(true);

async function _find() {
	loading.value = true;

	try {
		const { query } = route;

		const queries = Object.keys(query).map((key) => `${key}=${query[key] ?? ""}`);

		await find(queries.join("&"));
	} catch (error) {
	} finally {
		setTimeout(() => {
			loading.value = false;
		}, 250);
	}
}

if (fetch) {
	_find();
}
</script>

<template>
	<slot :="{ loading, fetch: _find }"></slot>
</template>
