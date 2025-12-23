<script setup lang="ts">
import { useRouter } from "vue-router";
import { RoleService, type RoleType, type RoleTypePartialType } from "@/service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import { ref } from "vue";
import Title from "@/components/Title.vue";

export interface Props {
	id: number;
}
const router = useRouter();

const { id } = defineProps<Props>();

const role = ref<RoleTypePartialType>({});

async function save(model: RoleType) {
	const role = {
		id: +id,
		name: model.name,
		description: model.description,
	};

	const response = await RoleService.createOrUpdate(role);

	if (!response) {
		return false;
	}
	router.push({
		name: "PAGE_ROLE",
	});
	return true;
}

async function getByID() {
	const response = await RoleService.getByID(+id);

	if (!response) return;

	role.value = response;
}

if (+id > 0) {
	getByID();
}
</script>

<template>
	<div class="flex! gap-x-4 items-center mb-3">
		<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
		<q-breadcrumbs>
			<q-breadcrumbs-el
				:label="$tl('roles_page')"
				icon="article"
				:to="{ name: 'PAGE_ROLE' }"
			/>
			<q-breadcrumbs-el :label="$tl('page_for_edit')" />
		</q-breadcrumbs>
	</div>
	<Form v-model="role" :save="save">
		<template #title>
			<Title class="mb-5">{{ $tl("update_role") }}</Title>
		</template>

		<template #name="{ model }">
			<Input v-model="model.name" label="name" class="col-lg-6 col-md-6 col-12" />
		</template>
		<template #description="{ model }">
			<Input
				v-model="model.description"
				label="description"
				class="col-lg-6 col-md-6 col-12"
			/>
		</template>

		<template #actions="{ loading }">
			<Button :loading="loading" type="submit" class="ml-auto">
				{{ $tl("save") }}
			</Button>
		</template>
	</Form>
</template>
