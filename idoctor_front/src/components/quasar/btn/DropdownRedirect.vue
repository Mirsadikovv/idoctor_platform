<script setup lang="ts">
import type { TLastWorkflowPartial } from "@/service";
import { useRouter } from "vue-router";

interface ActionItem {
	key: string;
	route: string;
}

interface Props {
	label?: string;
	actions: ActionItem[];
	withTooltip?: boolean;
	tooltipText?: string;
	isDisable?: boolean;
	style?: string;
	classBtn?: string;
	colorBtn?: string;
	fetch?: Function;
	id?: number;
	version?: number;
	latestWorkflow?: TLastWorkflowPartial;
	redirect?: string;
	[key: string]: any;
}

const props = defineProps<Props>();
const router = useRouter();
const {
	label,
	actions,
	withTooltip,
	tooltipText,
	isDisable,
	classBtn,
	colorBtn,
	id,
	redirect,
	version,
} = props;

// функция открытия модалки с выбранным компонентом
async function openAction(action: ActionItem) {
	router.push({
		name: action.route,
		params: {
			id: id,
		},
		query: {
			version: version,
			redirect: redirect || "",
		},
	});
}
</script>

<template>
	<q-btn-dropdown
		no-caps
		:color="colorBtn || 'white'"
		flat
		:class="`bg-info ${classBtn}`"
		:label="$tl(label || 'choose_action')"
		:disable="isDisable"
	>
		<q-tooltip
			v-if="withTooltip"
			anchor="top middle"
			self="center right"
			class="bg-accent text-white"
		>
			{{ $tl(tooltipText) }}
		</q-tooltip>

		<q-list>
			<q-item
				v-for="action in actions"
				:key="action.key"
				clickable
				v-close-popup
				@click="openAction(action)"
			>
				<q-item-section>
					<q-item-label>{{ $tl(`${action.key}_BTN`) }}</q-item-label>
				</q-item-section>
			</q-item>
		</q-list>
	</q-btn-dropdown>
</template>

<style lang="scss">
.clean-logout-btn {
	border-radius: 8px !important;
	transition: $transition-fast;

	&:hover {
		transform: translateY(-1px);
		box-shadow: 0 4px 12px rgba(50, 38, 220, 0.2);
	}

	&:active {
		transform: translateY(0);
	}
}
</style>
