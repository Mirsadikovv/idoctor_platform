<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref } from "vue";
import { ProblemService, type ProblemUpdateType, type ProblemType } from "../service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";

export interface Props {
	id: string | number;
	fetch: Function;
}

const { id, fetch } = defineProps<Props>();

const problemModel = ref<ProblemType>({
	id: 0,
	name: "",
	price: 0,
	created_at: ""
});

const onClose = inject<() => Promise<void>>("onClose");

async function save(model: ProblemUpdateType) {
	const response = await ProblemService.update(+id, model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}

ProblemService.getByID(+id).then((data) => {
	if (data) {
		problemModel.value = data;
	}
});
</script>

<template>
	<Form v-model="problemModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("edit_problem") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #name="{ model }">
			<Input 
				v-model="model.name" 
				label="Problem Name" 
				class="col-12" 
				required
			/>
		</template>
		<template #price="{ model }">
			<Input 
				v-model="model.price" 
				label="Price" 
				type="number"
				:min="0"
				class="col-12" 
				required
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