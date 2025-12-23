<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";

import { inject, ref, onMounted } from "vue";
import { UserService, type UserUpdateType } from "../service";
import { formRequired } from "@/common/validator";

export interface Props {
	id: string | number;
	fetch: Function;
}

const { id, fetch } = defineProps<Props>();

const userModel = ref<Partial<UserUpdateType>>({});
const onClose = inject<() => Promise<void>>("onClose");

onMounted(async () => {
	const userData = await UserService.getByID(+id);
	if (userData) {
		userModel.value = {
			username: userData.username,
		};
	}
});

async function save(model: UserUpdateType) {
	const response = await UserService.update(+id, model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}
</script>

<template>
	<Form v-model="userModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("edit_user_credentials") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #username="{ model }">
			<Input
				v-model="model.username"
				label="Username"
				class="col-12"
				:rules="[formRequired()]"
			/>
		</template>

		<template #password="{ model }">
			<Input v-model="model.password" label="New Password" type="password" class="col-12" />
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
