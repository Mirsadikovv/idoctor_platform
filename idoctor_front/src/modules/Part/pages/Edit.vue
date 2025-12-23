<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref, onMounted } from "vue";
import { PartService, type PartUpdateType, type PartType } from "../service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import { formRequired } from "@/common/validator";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { searchDevices, searchSuppliers } from "../utils";

export interface Props {
	id: string | number;
	fetch: Function;
}

const { id, fetch } = defineProps<Props>();

const partModel = ref<Partial<PartType>>({});

const onClose = inject<() => Promise<void>>("onClose");

onMounted(async () => {
	const [partData] = await Promise.all([PartService.getByID(+id)]);

	if (partData) {
		partModel.value = partData;
	}
});

async function save(model: PartUpdateType) {
	const response = await PartService.update(+id, model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}
</script>

<template>
	<Form v-model="partModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("edit_part") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #name="{ model }">
			<Input
				v-model="model.name"
				label="Part Name"
				class="col-12"
				:rules="[formRequired()]"
			/>
		</template>

		<template #device_id="{ model }">
			<Autocomplete
				v-model="model.device_id"
				label="Device"
				class="col-12"
				:rules="[formRequired()]"
				:find="searchDevices"
				:option-label="(val) => val.name + ' | ' + val.brand_name"
				option-value="id"
			/>
		</template>

		<template #supplier_id="{ model }">
			<Autocomplete
				v-model="model.supplier_id"
				label="Supplier"
				class="col-12"
				:rules="[formRequired()]"
				:find="searchSuppliers"
				option-label="name"
				option-value="id"
			/>
		</template>
		<template #actions="{ loading }">
			<div class="col-12">
				<div class="row my-row justify-center">
					<div class="col max-w-240px!">
						<Button :loading="loading" v-close-popup ui-type="danger" class="w-full!">
							{{ $tl("cancel") }}
						</Button>
					</div>
					<div class="col max-w-240px!">
						<Button :loading="loading" type="submit" class="w-full!">
							{{ $tl("save") }}
						</Button>
					</div>
				</div>
			</div>
		</template>
	</Form>
</template>
