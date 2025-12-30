<script setup lang="ts">
import { useRouter } from "vue-router";
import { UserService, type UserCreateType } from "../service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import { formRequired } from "@/common/validator";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { searchRole } from "../utils";
import DatePicker from "@/components/quasar/form/DatePicker.vue";
import { ref } from "vue";

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const userModel = ref<Partial<UserCreateType>>({});

const genderOptions = [
	{ label: "Male", value: "male" },
	{ label: "Female", value: "female" },
];

async function save(model: UserCreateType) {
	const response = await UserService.create(model);

	if (!response) return false;

	router.push({
		name: "PAGE_USER",
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
				class="bg-white text-gray-900 overflow-auto p-4 pt-16"
			>
				<div class="flex! gap-x-4 items-center mb-3">
					<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
					<q-breadcrumbs>
						<q-breadcrumbs-el
							:label="$tl('user_list')"
							icon="article"
							:to="{ name: 'PAGE_USER' }"
						/>
						<q-breadcrumbs-el :label="$tl('page_for_create')" />
					</q-breadcrumbs>
				</div>
				<Form v-model="userModel" :save="save">
					<template #title>
						<Title class="mb-5">{{ $tl("create_user") }}</Title>
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
						<DatePicker
							v-model="model.dateOfBirth"
							label="Date of Birth"
							class="col-12"
						/>
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
