<script setup lang="ts">
// Определяем модель как объединение типов string и Date
import { ref } from "vue";

export interface Props {
	label?: string;
}
const { label } = defineProps<Props>();
const date = defineModel<string>({});

const datePicker = ref();

function handleDateChange() {
	datePicker.value.hide();
}

const showDatePicker = () => {
	datePicker.value.show();
};
</script>

<template>
	<q-input
		:label="$tl(label)"
		@click="showDatePicker"
		outlined
		v-model="date"
		color="accent"
		placeholder="YYYY-MM-DD"
		mask="####-##-##"
	>
		<template #append>
			<q-icon name="event" class="cursor-pointer">
				<q-popup-proxy
					ref="datePicker"
					cover
					transition-show="scale"
					transition-hide="scale"
				>
					<q-date @update:model-value="handleDateChange" v-model="date" mask="YYYY-MM-DD">
						<div class="row items-center justify-end">
							<q-btn v-close-popup :label="$tl('close')" color="accent" flat />
						</div>
					</q-date>
				</q-popup-proxy>
			</q-icon>
		</template>
	</q-input>
</template>
