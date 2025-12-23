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
		:class="{ 'rounded-2xl my-1 mr-1': !!activeGroupList }"
		exact
		:active="active(route.meta?.activeLinkGroup)"
		:to="to"
		replace
		clickable
	>
		<q-item-section avatar>
			<q-icon :name="icon" />
		</q-item-section>
		<q-item-section>{{ $tg(title) }} </q-item-section>
	</q-item>
</template>

<style scoped lang="scss">
.q-item {
	position: relative;
	transition: all 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94);

	&:hover:not(.active__link) {
		background: rgba($active-link, 0.08);

		&::after {
			transform: translateY(-50%) scale(1);
		}
	}
}

.active__link {
	color: white;
	background: $active-link;
	position: relative;
	overflow: hidden;

	&::before {
		content: "";
		position: absolute;
		top: 0;
		left: -100%;
		width: 100%;
		height: 100%;
		background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
		animation: slideShine 5s infinite;
	}
}

@keyframes slideShine {
	0% {
		left: -100%;
	}
	50% {
		left: 100%;
	}
	100% {
		left: 100%;
	}
}

.q-item__section--avatar {
	min-width: 0 !important;
}
</style>
