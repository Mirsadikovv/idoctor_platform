<script setup lang="ts">
import { useRoute } from "vue-router";
import { activeList, type LinkProps } from "./component";

const { title, to, icon, active, activeLinkGroup } = defineProps<LinkProps>();
const route = useRoute();

const activeGroupList = activeList();

activeGroupList?.add(activeLinkGroup);
</script>

<template>
	<q-item
		active-class="active__link"
		:class="{ 'rounded-2xl my-1 ': !!activeGroupList }"
		:active="active(route.meta?.activeLinkGroup)"
		:to="to"
		exact
		replace
		clickable
	>
		<q-item-section>
			<q-icon :name="icon" />
		</q-item-section>
		<q-item-section>{{ $tg(title) }} </q-item-section>
	</q-item>
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
