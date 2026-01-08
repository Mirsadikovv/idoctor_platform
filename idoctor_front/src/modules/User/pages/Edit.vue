<script setup lang="ts">
import { useRouter } from "vue-router";
import { UserService, type UserUpdateType } from "../service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import { formRequired } from "@/common/validator";
import { ref } from "vue";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { searchRole } from "../utils";
import { getId } from "@/common";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import DatePicker from "@/components/quasar/form/DatePicker.vue";

export interface Props {
	id: number;
}

const { id } = defineProps<Props>();

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();

const userModel = ref<Partial<UserUpdateType>>({});
const genderOptions = [
	{ label: "Male", value: "male" },
	{ label: "Female", value: "female" },
];

async function getUserByID() {
	const userData = await UserService.getByID(+id);
	if (userData) {
		userModel.value = {
			username: userData.username,
			roleId: { id: userData.roleId, name: userData.role },
		};
	}
}

async function save(model: UserUpdateType) {
	const response = await UserService.update(+id, { ...model, roleId: getId(model.roleId) });

	if (!response) return false;

	router.push({
		name: "USER_PAGE",
	});
	return true;
}
</script>

<template>
	<PageLoading :find="getUserByID" #="{ loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
					<div class="flex! gap-x-4 items-center mb-3">
						<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
						<q-breadcrumbs>
							<q-breadcrumbs-el
								:label="$tl('user_list')"
								icon="article"
								:to="{ name: 'USER_PAGE' }"
							/>
						</q-breadcrumbs>
					</div>
					<Form v-model="userModel" :save="save">
						<template #title>
							<Title class="mb-5">{{ $tl("edit_user_credentials") }}</Title>
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
							<Input
								v-model="model.phoneNumber"
								label="Phone Number"
								class="col-12"
								:rules="[formRequired()]"
								mask="+### (##) ###-##-##"
							/>
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
							<Button :loading="loading" type="submit" class="ml-auto">
								{{ $tl("save") }}
							</Button>
						</template>
					</Form>
				</q-page>
			</q-page-container>

			<AppFooter
				:username="authStore.user?.username"
				:languages="$lang.languages"
				:current-language-id="$lang._currentLang?.id"
				:show-add-button="true"
				:add-button-route="{ name: 'USER_CREATE' }"
				add-button-icon="person_add"
				@toggle-drawer="toggleLeftDrawer"
				@go-to-profile="toggleLeftDrawer"
				@set-lang="setLang"
				@logout="logout"
			/>
		</q-layout>
	</PageLoading>
</template>

<style scoped>
@import "@/styles/telegram-app.scss";
</style>
