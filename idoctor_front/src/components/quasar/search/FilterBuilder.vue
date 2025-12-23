<script lang="ts" setup>
import { computed } from "vue";
import Expasion from "../Expasion.vue";

export interface FilterConfig {
	componentName: string;
	modelKey: string;
	label?: string;
	name?: string;
	class?: string;
	placeholder?: string;
	clearable?: boolean;
	inputDebounce?: string | number;
	optionLabel?: string | ((option: any) => string);
	queryName?: string;

	// Функции-колбэки
	find?: () => Promise<any[]> | any[];
	update?: (value: any) => void | Promise<void>;
	search?: (value: any) => void | Promise<void>;
	clear?: () => void | Promise<void>;

	// Дополнительные пропсы для компонента
	props?: Record<string, any>;
}

interface Props {
	filters: FilterConfig[];
	modelValue: Record<string, any>;
	loading?: boolean;
	paddingOff?: boolean;
	rowClass?: string;
	showActions?: boolean;
	actionsClass?: string;
	clearButtonText?: string;

	// Колбэки для основных действий
	onFind?: (query: string) => Promise<void>;
	onClearAll?: () => Promise<boolean>;
}

const props = withDefaults(defineProps<Props>(), {
	loading: false,
	paddingOff: true,
	rowClass: "",
	showActions: true,
	actionsClass: "col-12 flex",
	clearButtonText: "Clear",
});

const emit = defineEmits<{
	"update:modelValue": [value: Record<string, any>];
	search: [filter: FilterConfig, value: any];
	update: [filter: FilterConfig, value: any];
	clear: [filter: FilterConfig];
}>();

const searchModel = computed({
	get: () => props.modelValue,
	set: (value) => emit("update:modelValue", value),
});

// Получение пропсов для конкретного фильтра
function getFilterProps(filter: FilterConfig) {
	const baseProps = {
		label: filter.label,
		name: filter.name,
		class: filter.class,
		placeholder: filter.placeholder,
		clearable: filter.clearable ?? true,
		inputDebounce: filter.inputDebounce,
		optionLabel: filter.optionLabel,
		queryName: filter.queryName,
		find: filter.find,
	};

	// Добавляем дополнительные пропсы если есть
	return { ...baseProps, ...filter.props };
}

// Обработчик обновления значения
async function handleUpdate(filter: FilterConfig, value: any) {
	// Обновляем модель
	searchModel.value = { ...searchModel.value, [filter.modelKey]: value };

	// Вызываем кастомный колбэк если есть
	if (filter.update) {
		await filter.update(value);
	}

	// Эмитим событие
	emit("update", filter, value);

	// Вызываем поиск если указано
	if (props.onFind) {
		await props.onFind("");
	}
}

// Обработчик поиска
async function handleSearch(filter: FilterConfig, value: any) {
	if (filter.search) {
		await filter.search(value);
	}

	emit("search", filter, value);

	if (props.onFind) {
		await props.onFind("");
	}
}

// Обработчик очистки конкретного фильтра
async function handleClear(filter: FilterConfig, _event?: any) {
	if (filter.clear) {
		await filter.clear();
	}

	emit("clear", filter);
}

// Общая очистка всех фильтров
async function onClear() {
	if (props.onClearAll) {
		return await props.onClearAll();
	}

	// Дефолтная логика очистки
	searchModel.value = {};
	if (props.onFind) {
		await props.onFind("");
	}
	return true;
}
</script>

<template>
	<Expasion :padding-off="paddingOff">
		<div class="row [&>*]:p2 items-start!" :class="rowClass">
			<!-- Динамическое создание компонентов фильтров -->
			<component
				v-for="(filter, index) in filters"
				:key="`filter-${index}`"
				:is="filter.componentName"
				v-bind="getFilterProps(filter)"
				v-model="searchModel[filter.modelKey]"
				@update="handleUpdate(filter, $event)"
				@search="handleSearch(filter, $event)"
				@clear="handleClear(filter, $event)"
			/>

			<!-- Блок с кнопками действий -->
			<div v-if="showActions" :class="actionsClass">
				<q-space></q-space>
				<slot name="actions" :loading="loading" :onClear="onClear">
					<q-btn
						@click="onClear()"
						:loading="loading"
						class="bg-white text-primary pt1! px-8!"
						outline
					>
						{{ clearButtonText }}
					</q-btn>
				</slot>
			</div>
		</div>
	</Expasion>
</template>
