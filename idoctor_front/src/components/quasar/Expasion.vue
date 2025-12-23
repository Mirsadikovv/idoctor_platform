<script setup lang="ts">
export interface Props {
	label?: string;
	icon?: string;
	defaultOpened?: boolean;
	clazz?: string;
	paddingOff?: boolean;
	dense?: boolean;
}

const { label, icon, defaultOpened, clazz, paddingOff, dense } = defineProps<Props>();
</script>

<template>
	<q-list class="rounded bg-white" bordered separator :class="clazz">
		<q-expansion-item
			:dense="dense"
			expand-separator
			:default-opened="defaultOpened"
			:icon="icon ? icon : 'manage_search'"
			expand-icon-class="text-white"
			header-class="font-bold text-white text-base bg-gradient-header"
			:label="!!label ? $tl(label) : $tl('search_and_filters')"
		>
			<q-separator />
			<div :class="paddingOff ? 'p-0' : 'p-4'">
				<slot />
			</div>
		</q-expansion-item>
	</q-list>
</template>
<style lang="scss">
.q-item__section--avatar {
	min-width: 0 !important;
}
.q-expansion-item__container {
	min-height: 40px !important;
}
</style>
