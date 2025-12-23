<script setup lang="ts">
import { useRouter } from "vue-router";
import { RoleService, type RoleType, type RoleTypePartialType } from "@/service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import { ref } from "vue";
import Title from "@/components/Title.vue";
import { formRequired } from "@/common/validator";

const router = useRouter();

const role = ref<RoleTypePartialType>({});

async function save(model: RoleType) {
	const role = {
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
			<q-breadcrumbs-el :label="$tl('page_for_create')" />
		</q-breadcrumbs>
	</div>
	<Form v-model="role" :save="save">
		<template #title>
			<Title class="mb-5">{{ $tl("create_role") }}</Title>
		</template>

		<template #name="{ model }">
			<Input
				:rules="[formRequired()]"
				v-model="model.name"
				label="name"
				class="col-lg-6 col-md-6 col-12"
			/>
		</template>
		<template #description="{ model }">
			<Input
				v-model="model.description"
				label="description"
				class="col-lg-6 col-md-6 col-12"
			/>
		</template>

		<template #actions="{ loading }">
			<div class="row my-row justify-center">
				<q-space />
				<div class="col-auto max-w-200px! w-100%!">
					<Button :loading="loading" type="submit" class="w-100%!">
						{{ $tl("save") }}
					</Button>
				</div>
			</div>
		</template>
	</Form>
</template>
