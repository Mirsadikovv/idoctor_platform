<script setup lang="ts">
import { inject, ref } from "vue";
import { formRequired } from "@/common";
import { initalBatch } from "../utils";
import { LanguageContentService, type LangContentCreateOrUpdate } from "@/service";

import Title from "@/components/Title.vue";
import Form from "@/components/quasar/form/Form.vue";
import Input from "@/components/quasar/form/Input.vue";
import Button from "@/components/quasar/btn/Button.vue";

const model = ref<LangContentCreateOrUpdate>();

initalBatch().then((res) => (model.value = res));

const onClose = inject<() => Promise<void>>("onClose");

async function save(newModel: LangContentCreateOrUpdate) {
	const response = await LanguageContentService.createOrUpdate({
		...newModel,
		contents: newModel.contents.map((item) => ({ ...item, key: newModel.key })),
	});

	if (!response) return false;
	await onClose?.();
	return true;
}
</script>

<template>
	<Form :model-value="model" :save="save">
		<template #title>
			<Title class="mb-5">{{ $tl("add_translate") }} </Title>
		</template>

		<template #key="{ model }">
			<Input
				v-model="model.key"
				label="key"
				class="col-12"
				:rules="[formRequired($tl('this_field_is_required'))]"
			/>
		</template>

		<template #separator>
			<div class="col-12">
				<q-separator class="mx--4! my-4px!" />
			</div>
		</template>

		<template #values>
			<div
				v-if="model?.contents"
				class="col-lg-4 col-md-6 col-12"
				v-for="item in model.contents"
				:key="item.languageId"
			>
				<div class="mb-2 text-base font-medium">
					{{ item.langName }} | {{ item.LangDescp }}
				</div>
				<Input
					v-model="item.value"
					label="value"
					:rules="[formRequired($tl('this_field_is_required'))]"
				/>
			</div>
		</template>

		<template #actions="{ loading }">
			<div class="col-12">
				<div class="row my-row justify-center">
					<div class="col max-w-220px!">
						<Button
							:loading="loading"
							class="w-full!"
							ui-type="danger"
							v-close-popup
						>
							{{ $tl("cancel") }}
						</Button>
					</div>
					<div class="col max-w-220px!">
						<Button :loading="loading" type="submit" class="w-full!">
							{{ $tl("save") }}
						</Button>
					</div>
				</div>
			</div>
		</template>
	</Form>
</template>
