<script setup lang="ts">
import type { ProcessedIcons } from "@/common";
import IconBtn from "@/components/quasar/btn/IconBtn.vue";
import { provide, ref } from "vue";

interface Props {
	icon: ProcessedIcons;
	iconColor?: string;
	withTooltip?: boolean;
	tooltipText?: string;
	isDisable?: boolean;
	style?: string;
}
const { icon, iconColor, withTooltip, tooltipText, isDisable, style } = defineProps<Props>();
const model = ref(false);

// функция закрытия
async function onClose() {
	await new Promise((resolve) => setTimeout(resolve, 500));
	model.value = false;
}

// экспортируем, чтобы ref был доступен через `ref="..."` (опционально)
defineExpose({ model });

// предоставляем onClose для вложенных компонентов
provide("onClose", onClose);
</script>

<template>
	<IconBtn :icon="icon" @click="model = true" :disable="isDisable" :color="iconColor">
		<q-tooltip
			v-if="withTooltip"
			anchor="top middle"
			self="center right"
			class="bg-accent text-white"
		>
			{{ $tl(tooltipText) }}
		</q-tooltip>
	</IconBtn>

	<q-dialog v-model="model" persistent>
		<q-card style="max-width: 99%; max-height: 95%" :style="style">
			<slot></slot>
		</q-card>
	</q-dialog>
</template>
