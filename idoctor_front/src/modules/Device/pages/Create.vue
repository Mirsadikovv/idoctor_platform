<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref } from "vue";
import { DeviceService, type DeviceCreateType } from "../service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";

interface Props {
	fetch: Function;
}

const { fetch } = defineProps<Props>();

const deviceModel = ref<DeviceCreateType>({
	name: "",
	brand_name: ""
});
const onClose = inject<() => Promise<void>>("onClose");

async function save(model: DeviceCreateType) {
	const response = await DeviceService.create(model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}
</script>

<template>
	<Form v-model="deviceModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("create_device") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #name>
			<Input 
				v-model="deviceModel.name" 
				label="Device Name" 
				class="col-12" 
				required
			/>
		</template>
		<template #brand_name>
			<Input 
				v-model="deviceModel.brand_name" 
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