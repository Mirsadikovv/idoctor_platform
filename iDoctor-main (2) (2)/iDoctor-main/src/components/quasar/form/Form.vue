<script setup lang="ts" generic="T extends Record<string, any>">
import { QForm } from "quasar";
import { computed, ref } from "vue";

export interface Props<T> {
	save: (modelValue: Required<T>) => Promise<boolean>;
	modelValue?: Partial<T>;
	actionClass?: string;
	paddingMini?: boolean;
	isCleanedStyle?: boolean;
	clazz?: string;
}

type KEY = keyof Required<T>;

type RecordType = Required<T>;

type ModelKeySlots = {
	[K in KEY]: (_: { model: RecordType; loading: boolean }) => any;
};

type RecordSlot = {
	[key: string]: (_: { model: RecordType; loading: boolean }) => any;
};

type ActionSlot = {
	actions: (_: { model: RecordType; loading: boolean }) => any;
};

const { save, clazz } = defineProps<Props<RecordType>>();

const modelValue = defineModel<RecordType | Partial<T>>();
const refModel = ref({} as RecordType);
const reactiveModel = computed(() => modelValue.value ?? refModel.value);

const slots = defineSlots<RecordSlot & ActionSlot & ModelKeySlots>();

const form = ref<QForm>()!;

const loading = ref(false);

function fields(slots: any) {
	const keys = Object.keys(slots).filter((k) => !k.endsWith("actions") && !k.endsWith("before"));
	return keys;
}

async function success(modelValue: any) {
	try {
		loading.value = true;

		const isSave = await save(modelValue);

		if (!isSave) {
			return;
		}

		form.value && form.value!.reset();
	} catch (error) {
		console.log(error);
	} finally {
		loading.value = false;
	}
}

function validate() {
	return form.value?.validate();
}

function resetValidation() {
	form.value?.resetValidation();
}
defineExpose({ validate, resetValidation });
</script>

<template>
	<q-card
		flat
		:bordered="!isCleanedStyle"
		square
		style="border-radius: 16px !important"
		:class="clazz"
	>
		<q-card-section :class="isCleanedStyle ? 'p-0!' : 'py-20px!'">
			<slot
				:name="//@ts-ignore
				`title`"
				:="{ model: (reactiveModel as RecordType), loading }"
			></slot>
			<q-form class="row my-row" ref="form" @submit.prevent="success(reactiveModel)">
				<slot
					:name="//@ts-ignore
					`before`"
					:="{ model: (reactiveModel as RecordType), loading }"
				></slot>
				<template v-for="field of fields(slots).filter((item) => item !== 'title')">
					<slot
						:="{ model: (reactiveModel as RecordType), loading }"
						:name="//@ts-ignore
						`${field}`"
					>
					</slot>
				</template>
				<div class="col-12" :class="actionClass">
					<slot
						:name="//@ts-ignore
						`actions`"
						:="{ model: (reactiveModel as RecordType), loading }"
					></slot>
				</div>
			</q-form>
		</q-card-section>
	</q-card>
</template>
