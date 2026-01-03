<script setup lang="ts">
import type { TLastWorkflowPartial } from "@/service";
import { ref, provide } from "vue";

interface ActionItem {
	key: string;
	component: any;
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

const {
	label,
	actions,
	withTooltip,
	tooltipText,
	isDisable,
	style,
	classBtn,
	colorBtn,
	latestWorkflow,
	redirect,
	...other
} = props;

const model = ref(false);
const selectedAction = ref<ActionItem | null>(null);

// функция открытия модалки с выбранным компонентом
async function openAction(action: ActionItem) {
	selectedAction.value = action;
	model.value = true;
}

// функция закрытия
async function onClose() {
	await new Promise((resolve) => setTimeout(resolve, 500));
	model.value = false;
	selectedAction.value = null;
}

// функция открытия
async function onOpen() {
	model.value = true;
}

// экспортируем для доступа через ref
defineExpose({ model });

// предоставляем onClose и onOpen для вложенных компонентов
provide("onClose", onClose);
provide("onOpen", onOpen);
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
					<q-item-label>{{ $tl(action.key) }}</q-item-label>
				</q-item-section>
			</q-item>
		</q-list>
	</q-btn-dropdown>

	<q-dialog v-model="model" persistent>
		<q-card style="max-width: 99%; max-height: 95%; width: 1000px" :style="style">
			<component
				v-if="selectedAction"
				:is="selectedAction.component"
				v-bind="other"
				:latestWorkflow="latestWorkflow"
				:redirect="redirect"
			/>
		</q-card>
	</q-dialog>
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
