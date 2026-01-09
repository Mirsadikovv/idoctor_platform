<script setup lang="ts">
import { dateFormat, dateValid, isValidDate, parseDate } from "@/common";
import { ref } from "vue";
import { useRoute, useRouter } from "vue-router";

export interface Props {
	name?: string;
	label?: string;
	modelValue?: string;
	maskInput?: string;
	maskDate?: string;
	dense?: boolean;
	withoutTime?: boolean;
}

const {
	name = "search",
	label = "search",
	maskDate,
	maskInput,
	dense,
	withoutTime,
} = defineProps<Props>();

const emits = defineEmits<{
	(e: "search", search: string): void;
}>();

const router = useRouter();

const route = useRoute();

let search = ref(currentSearch());

function currentSearch(): string {
	return dateFormat(route.query[name] as string, "DD-MM-YYYY") || "";
}

async function searchChange(item: any) {
	const { query } = route;
	// TODO Fix me
	if (isValidDate(item)) {
		// Получаем объект Date из строки в формате DD-MM-YYYY
		const parsed = parseDate(item, "DD-MM-YYYY"); // допустим, возвращает обычный Date

		// Получаем компоненты даты
		const day = parsed!.getDate();
		const month = parsed!.getMonth();
		const year = parsed!.getFullYear();

		// Создаём UTC-дейту без сдвига временной зоны
		const utcDate = new Date(Date.UTC(year, month, day, 0, 0, 0));

		if (withoutTime) {
			query[name] = utcDate.toISOString().split("T")[0];
		} else {
			query[name] = utcDate.toISOString();
		}

		await router.push({
			force: true,
			query: {
				...query,
			},
		});

		emits("search", item);
	}
}

async function onClear() {
	const { query } = route;

	// Создаем копию query и удаляем нужный параметр
	const { [name]: removed, ...filteredQuery } = query;

	// Дополнительно фильтруем null/undefined значения
	const cleanedQuery = Object.fromEntries(
		Object.entries(filteredQuery).filter(([_, value]) => value != null),
	);

	// Обновляем поиск
	search.value = "";

	await router.replace({ query: cleanedQuery });
	emits("search", "");
}
</script>

<template>
	<q-input
		v-model="search"
		@update:model-value="searchChange"
		@clear="onClear"
		:placeholder="maskDate || 'DD-MM-YYYY'"
		:mask="maskInput || '##-##-####'"
		:label="$tl(label)"
		:rules="[dateValid($tl('date_will_be_dd_mm_yyyy'))]"
		:validate="isValidDate"
		color="accent"
		debounce="700"
		clearable
		outlined
		:dense="dense"
	>
		<template #prepend>
			<q-icon name="event" class="cursor-pointer mr-2">
				<q-popup-proxy
					ref="datePicker"
					transition-show="scale"
					transition-hide="scale"
					cover
				>
					<q-date
						@update:model-value="searchChange"
						v-model="search"
						:mask="maskDate || 'DD-MM-YYYY'"
						color="accent"
						landscape
					>
						<div class="row items-center justify-end">
							<q-btn v-close-popup :label="$tl('close')" color="accent" flat />
						</div>
					</q-date>
				</q-popup-proxy>
			</q-icon>
		</template>
	</q-input>
</template>
