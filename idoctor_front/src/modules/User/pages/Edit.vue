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

export interface Props {
	id: number;
}

const { id } = defineProps<Props>();

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const userModel = ref<Partial<UserUpdateType>>({});

async function getUserByID() {
	const userData = await UserService.getByID(+id);
	if (userData) {
		userModel.value = {
			username: userData.username,
		};
	}
}

async function save(model: UserUpdateType) {
	const response = await UserService.update(+id, model);

	if (!response) return false;

	router.push({
		name: "PAGE_USER",
	});
	return true;
}
</script>

<template>
	<PageLoading :find="getUserByID" #="{ loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page
					:style="{
						height: 'calc(var(--app-height, 100vh) - 150px)',
					}"
					class="bg-white text-gray-900 overflow-auto p-4 pt-20"
				>
					<div class="flex! gap-x-4 items-center mb-3">
						<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
						<q-breadcrumbs>
							<q-breadcrumbs-el
								:label="$tl('user_list')"
								icon="article"
								:to="{ name: 'PAGE_USER' }"
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
								label="New Password"
								type="password"
								class="col-12"
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
