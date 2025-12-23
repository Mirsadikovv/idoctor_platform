<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref } from "vue";
import { DeviceService, type DeviceUpdateType, type DeviceType } from "../service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";

export interface Props {
	id: string | number;
	fetch: Function;
}

const { id, fetch } = defineProps<Props>();

const deviceModel = ref<DeviceType>({
	id: 0,
	name: "",
	brand_name: "",
	created_at: ""
});

const onClose = inject<() => Promise<void>>("onClose");

async function save(model: DeviceUpdateType) {
	const response = await DeviceService.update(+id, model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}

DeviceService.getByID(+id).then((data) => {
	if (data) {
		deviceModel.value = data;
	}
});
</script>

<template>
	<Form v-model="deviceModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("edit_device") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #name="{ model }">
			<Input 
				v-model="model.name" 
				label="Device Name" 
				class="col-12" 
				required
			/>
		</template>
		<template #brand_name="{ model }">
			<Input 
				v-model="model.brand_name" 
				label="Brand Name" 
				class="col-12" 
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