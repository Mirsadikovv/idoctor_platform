<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref } from "vue";
import { UserService, type UserCreateType } from "../service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import { formRequired } from "@/common/validator";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { searchRole } from "../utils";
import DatePicker from "@/components/quasar/form/DatePicker.vue";

interface Props {
	fetch: Function;
}

const { fetch } = defineProps<Props>();

const userModel = ref<Partial<UserCreateType>>({});
const onClose = inject<() => Promise<void>>("onClose");

const genderOptions = [
	{ label: "Male", value: "male" },
	{ label: "Female", value: "female" },
];

async function save(model: UserCreateType) {
	const response = await UserService.create(model);

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
				{{ $tl("create_user") }}
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
			<Input
				v-model="model.password"
				label="Password"
				type="password"
				class="col-12"
				:rules="[formRequired()]"
			/>
		</template>

		<template #roleId="{ model }">
			<Autocomplete
				v-model="model.roleId"
				:find="searchRole"
				label="role"
				class="col-12"
				:rules="[formRequired()]"
				option-label="name"
				option-value="id"
			/>
		</template>

		<template #firstName="{ model }">
			<Input v-model="model.firstName" label="First Name" class="col-12" />
		</template>

		<template #lastName="{ model }">
			<Input v-model="model.lastName" label="Last Name" class="col-12" />
		</template>

		<template #middleName="{ model }">
			<Input v-model="model.middleName" label="Middle Name" class="col-12" />
		</template>

		<template #phoneNumber="{ model }">
			<Input v-model="model.phoneNumber" label="Phone Number" class="col-12" />
		</template>

		<template #dateOfBirth="{ model }">
			<DatePicker v-model="model.dateOfBirth" label="Date of Birth" class="col-12" />
		</template>

		<template #gender="{ model }">
			<Autocomplete
				v-model="model.gender"
				label="Gender"
				class="col-12"
				:find="async () => genderOptions"
				option-label="label"
				option-value="value"
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
