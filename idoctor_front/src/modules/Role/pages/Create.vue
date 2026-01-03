<script setup lang="ts">
import { useRouter } from "vue-router";
import {
	RoleService,
	type PermissionType,
	type RoleType,
	type RoleTypePartialType,
} from "@/service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import { ref } from "vue";
import Title from "@/components/Title.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import { formRequired } from "@/common/validator";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { searchPermission } from "../utils";

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const role = ref<RoleTypePartialType>({});
const permissions = ref<Partial<PermissionType>[]>([]);

async function save(model: RoleType) {
	const _permissions: Record<string, string[]> = {};

	permissions.value.forEach((p) => {
		_permissions[String(p.path)] = [];
	});

	permissions.value.forEach((p) => {
		_permissions[p.path!].push(p.method!);
	});

	const role = {
		name: model.name,
		description: model.description,
	};

	const response = await RoleService.create(role);

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
	<q-layout view="hHh Lpr lff">
		<q-page-container>
			<q-page
				:style="{
					height: 'calc(var(--app-height, 100vh) - 150px)',
				}"
				class="bg-white text-gray-900 overflow-auto p-4 pt-24"
			>
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

					<template #permissions>
						<Autocomplete
							use-chips
							class="col-12 max-h-400px overflow-auto"
							v-model="permissions"
							:option-label="(p) => `${p.method} => ${p.path}`"
							:find="searchPermission"
							label="permissions"
							multiple
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
			</q-page>
		</q-page-container>

		<AppFooter
			:username="authStore.user?.username"
			:languages="$lang.languages"
			:current-language-id="$lang._currentLang?.id"
			:show-add-button="false"
			@toggle-drawer="toggleLeftDrawer"
			@go-to-profile="toggleLeftDrawer"
			@set-lang="setLang"
			@logout="logout"
		/>
	</q-layout>
</template>

<style scoped>
@import "@/styles/telegram-app.scss";
</style>
