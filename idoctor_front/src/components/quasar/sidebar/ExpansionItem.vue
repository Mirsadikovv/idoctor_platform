<script setup lang="ts">
import { computed, provide, reactive } from "vue";
import type { ExtendedProps } from "./component";
import { useRoute } from "vue-router";

const props = defineProps<ExtendedProps>();

const route = useRoute();

const activeList = reactive(new Set<string>());

provide("activeList", activeList);

const active = computed(() => activeList.has(route.meta?.activeLinkGroup as string));
</script>

<template>
	<q-expansion-item
		v-bind="props"
		:model-value="active"
		:label="//@ts-ignore
		$tg(props.label!)"
		expand-icon-class="min-w-[30px]"
	>
		<slot></slot>
	</q-expansion-item>
</template>

<style scoped lang="scss">
.active__link {
	color: white;
	background: $active-link;
}
.q-item__section--avatar {
	min-width: 0 !important;
	padding: 0 !important;
}
</style>
