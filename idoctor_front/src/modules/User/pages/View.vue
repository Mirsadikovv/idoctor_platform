<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { UserService, type UserPartial } from "@/service";

import AppFooter from "@/components/AppFooter.vue";
import { useAuthStore } from "@/store/auth-store";
import { useAppNavigation } from "@/composables/useAppNavigation";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";

export interface Props {
	id: number | string;
}
const { id } = defineProps<Props>();

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const userModel = ref<UserPartial>({});

const loadUser = async () => {
	const data = await UserService.findByID(+id);
	userModel.value = data;
};
</script>
<template>
	<PageLoading :find="loadUser" #="{ loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
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
								:label="$tl('user_list')"
								icon="article"
								:to="{ name: 'USER_PAGE' }"
							/>
						</q-breadcrumbs>
					</div>
					<q-markup-table separator="cell" flat bordered>
						<tbody>
							<tr>
								<td class="font-bold text-left">{{ $tl("fullName") }}</td>
								<td>
									{{ userModel?.lastName }}
									{{ userModel?.firstName }}
									{{ userModel?.middleName }}
								</td>
							</tr>
							<tr>
								<td class="font-bold text-left">{{ $tl("gender") }}</td>
								<td>{{ $tl(userModel?.gender) }}</td>
							</tr>

							<tr>
								<td class="font-bold text-left">{{ $tl("username") }}</td>
								<td>{{ userModel?.username }}</td>
							</tr>

							<tr>
								<td class="font-bold text-left">{{ $tl("dateOfBirth") }}</td>
								<td>{{ userModel?.dateOfBirth }}</td>
							</tr>
						</tbody>
					</q-markup-table>
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
