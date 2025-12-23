<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref } from "vue";
import { ProblemService, type ProblemCreateType } from "../service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import { formNumber, formRequired } from "@/common/validator";

interface Props {
	fetch: Function;
}

const { fetch } = defineProps<Props>();

const problemModel = ref<Partial<ProblemCreateType>>({});
const onClose = inject<() => Promise<void>>("onClose");

async function save(model: ProblemCreateType) {
	const response = await ProblemService.create(model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}
</script>

<template>
	<Form v-model="problemModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("create_problem") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #name>
			<Input
				v-model="problemModel.name"
				label="Problem Name"
				class="col-12"
				:rules="[formRequired()]"
			/>
		</template>

		<template #price>
			<Input
				v-model.number="problemModel.price"
				:rules="[formRequired(), formNumber()]"
				label="Price"
				:min="0"
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
