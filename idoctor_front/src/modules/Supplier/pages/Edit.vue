<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref } from "vue";
import { SupplierService, type SupplierUpdateType, type SupplierType } from "../service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import { formRequired } from "@/common/validator";

export interface Props {
	id: string | number;
	fetch: Function;
}

const { id, fetch } = defineProps<Props>();

const supplierModel = ref<SupplierType>({
	id: 0,
	name: "",
	created_at: ""
});

const onClose = inject<() => Promise<void>>("onClose");

async function save(model: SupplierUpdateType) {
	const response = await SupplierService.update(+id, model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}

SupplierService.getByID(+id).then((data) => {
	if (data) {
		supplierModel.value = data;
	}
});
</script>

<template>
	<Form v-model="supplierModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("edit_supplier") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #name="{ model }">
			<Input
				v-model="model.name"
				label="Supplier Name"
				class="col-12"
				:rules="[formRequired()]"
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