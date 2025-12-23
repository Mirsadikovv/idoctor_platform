<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref } from "vue";
import { LanguageService, type LanguagePartialType, type LanguageType } from "@/service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";

export interface Props {
	id: string | number;
	fetch: Function;
}

const { id, fetch } = defineProps<Props>();

const LanguageModel = ref<LanguagePartialType>({});

const onClose = inject<() => Promise<void>>("onClose");

async function save(model: LanguageType) {
	const response = await LanguageService.update(+id, model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}

LanguageService.getByID(+id).then((data) => (LanguageModel.value = data));
</script>

<template>
	<Form v-model="LanguageModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("edit_lang") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #name="{ model }">
			<Input v-model="model.name" label="name" class="col-12" />
		</template>
		<template #description="{ model }">
			<Input v-model="model.description" label="description" class="col-12" />
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
