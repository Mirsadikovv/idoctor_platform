<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref } from "vue";
import { SupplierService, type SupplierCreateType } from "../service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import { formRequired } from "@/common/validator";

interface Props {
	fetch: Function;
}

const { fetch } = defineProps<Props>();

const supplierModel = ref<Partial<SupplierCreateType>>({});
const onClose = inject<() => Promise<void>>("onClose");

async function save(model: SupplierCreateType) {
	const response = await SupplierService.create(model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}
</script>

<template>
	<Form v-model="supplierModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("create_supplier") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #name>
			<Input
				v-model="supplierModel.name"
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