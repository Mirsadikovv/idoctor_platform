<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { ProblemService, type ProblemType } from "../service";

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

const problemModel = ref<ProblemType>({} as ProblemType);

const loadProblem = async () => {
	const data = await ProblemService.getByID(+id);
	problemModel.value = data;
};
</script>

<template>
	<PageLoading :find="loadProblem" #="{ loading }">
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
								:label="$tl('problem_list')"
								icon="medical_services"
								:to="{ name: 'PROBLEM_PAGE' }"
							/>
						</q-breadcrumbs>
						<q-space />
					</div>

					<q-markup-table separator="cell" flat bordered>
						<tbody>
							<tr>
								<td class="font-bold text-left">{{ $tl("problem_name") }}</td>
								<td>{{ problemModel?.name }}</td>
							</tr>
							<tr v-if="problemModel?.price">
								<td class="font-bold text-left">{{ $tl("price") }}</td>
								<td class="text-positive font-bold">
									{{ problemModel?.price?.toLocaleString() }} сум
								</td>
							</tr>
							<tr v-if="problemModel?.created_at">
								<td class="font-bold text-left">{{ $tl("created_at") }}</td>
								<td>
									{{ new Date(problemModel?.created_at).toLocaleDateString() }}
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
				:add-button-route="{ name: 'PROBLEM_CREATE' }"
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
