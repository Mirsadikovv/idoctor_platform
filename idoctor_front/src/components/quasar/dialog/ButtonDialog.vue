<script setup lang="ts">
import type { ProcessedIcons } from "@/common";
import { ref, provide } from "vue";
import Button from "@/components/quasar/btn/Button.vue";

interface Props {
	icon?: ProcessedIcons;
	label?: string;
	withTooltip?: boolean;
	tooltipText?: string;
	isDisable?: boolean;
	style?: string;
	classBtn?: string;
	colorBtn?: string;
}
const { icon, label, withTooltip, tooltipText, isDisable, style, classBtn, colorBtn } =
	defineProps<Props>();

const model = ref(false);

// функция открытия
async function onOpen() {
	model.value = true;
}

// функция закрытия
async function onClose() {
	await new Promise((resolve) => setTimeout(resolve, 500));
	model.value = false;
}

// экспортируем, чтобы ref был доступен через `ref="..."` (опционально)
defineExpose({ model });

// предоставляем onClose и onOpen для вложенных компонентов
provide("onClose", onClose);
provide("onOpen", onOpen);
</script>

<template>
	<Button
		:icon="icon"
		@click="model = true"
		:disable="isDisable"
		:color="colorBtn"
		:class="classBtn"
	>
		{{ $tl(label) }}
		<q-tooltip
			v-if="withTooltip"
			anchor="top middle"
			self="center right"
			class="bg-accent text-white"
		>
			{{ $tl(tooltipText) }}
		</q-tooltip>
	</Button>

	<q-dialog v-model="model" persistent>
		<q-card style="max-width: 99%; max-height: 95%" :style="style">
			<slot></slot>
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
