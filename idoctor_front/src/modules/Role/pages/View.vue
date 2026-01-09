<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { RoleService, type RoleType } from "@/service";

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

const roleModel = ref<RoleType>({} as RoleType);

const loadRole = async () => {
	const data = await RoleService.getByID(+id);
	roleModel.value = data;
};
</script>

<template>
	<PageLoading :find="loadRole" #="{ loading }">
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
								:label="$tl('roles_page')"
								icon="article"
								:to="{ name: 'PAGE_ROLE' }"
							/>
						</q-breadcrumbs>
						<q-space />
					</div>

					<q-markup-table separator="cell" flat bordered>
						<tbody>
							<tr>
								<td class="font-bold text-left">{{ $tl("name") }}</td>
								<td>{{ roleModel?.name }}</td>
							</tr>
							<tr v-if="roleModel?.description">
								<td class="font-bold text-left">{{ $tl("description") }}</td>
								<td>{{ roleModel?.description }}</td>
							</tr>
							<tr v-if="roleModel?.permissions">
								<td class="font-bold text-left">{{ $tl("permissions") }}</td>
								<td>
									{{ $tl("permissions_count") }}:
									{{ roleModel?.permissions?.length || 0 }}
								</td>
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
				:add-button-route="{ name: 'ROLE_CREATE' }"
				add-button-icon="add"
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
