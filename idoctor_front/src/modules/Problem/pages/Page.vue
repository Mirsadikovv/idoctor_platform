<script setup lang="ts">
import { ref } from "vue";
import { ProblemService, type ProblemPageData } from "../service";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";
import IconBtn from "@/components/quasar/btn/IconBtn.vue";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();

const problemPage = ref<ProblemPageData>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const pick = {
	id: false,
	name: true,
	price: true,
	created_at: true,
};

const pikers = ref({});

async function page(query: string = "") {
	const params = new URLSearchParams(query);
	const searchParams = {
		page: params.get("page") ? +params.get("page")! : undefined,
		perpage: params.get("perpage") ? +params.get("perpage")! : undefined,
		name: params.get("name") || undefined,
		min_price: params.get("min_price") ? +params.get("min_price")! : undefined,
		max_price: params.get("max_price") ? +params.get("max_price")! : undefined,
	};

	const response = await ProblemService.page(searchParams);

	if (!response) return;

	problemPage.value = response;
}
</script>

<template>
	<PageLoading :find="page" #="{ fetch, loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
					<ResponsiveTable :models="problemPage" hasOrder :loading="loading">
						<template #name:thead>
							{{ $tl("problem_name") }}
						</template>
						<template #name="{ model }">
							<router-link
								:to="{ name: 'PROBLEM_VIEW', params: { id: model.id } }"
								class="text-primary text-decoration-none"
							>
								{{ model.name }}
							</router-link>
						</template>

						<template #price:thead>
							{{ $tl("price") }}
						</template>
						<template #price="{ model }">
							{{ model.price.toLocaleString() }} сум
						</template>

						<template #created_at:thead>
							{{ $tl("created_at") }}
						</template>
						<template #created_at="{ model }">
							{{ new Date(model.created_at).toLocaleDateString() }}
						</template>

						<template #edit="{ model }">
							<div class="text-center">
								<IconBtn
									v-if="$canPage('PROBLEM_EDIT')"
									:to="{
										name: 'PROBLEM_EDIT',
										params: { id: model.id },
									}"
									icon="edit"
								/>
								<IconBtn
									v-if="$canPage('PROBLEM_VIEW')"
									:to="{
										name: 'PROBLEM_VIEW',
										params: { id: model.id },
									}"
									icon="visibility"
								/>
							</div>
						</template>

						<template #tfoot="{ totalPages }">
							<TablePaginate
								v-model:pikers="pikers"
								:total="totalPages"
								:pick="pick"
								@page="fetch"
							/>
						</template>
						<!-- Кастомный мобильный вид для медицинских проблем -->
						<template #card="{ model, orderNumber }">
							<q-item
								class="problem-item-telegram"
								clickable
								:to="{
									name: $canPage('PROBLEM_EDIT') ? 'PROBLEM_EDIT' : 'PROBLEM_VIEW',
									params: { id: model.id },
								}"
							>
								<q-item-section avatar v-if="orderNumber">
									<q-avatar color="primary" text-color="white" size="md">
										{{ orderNumber }}
									</q-avatar>
								</q-item-section>

								<q-item-section>
									<q-item-label class="text-weight-bold text-h6">
										{{ model.name }}
									</q-item-label>
									<q-item-label caption class="text-body2">
										{{ $tl("price") }}: {{ model.price.toLocaleString() }} сум
									</q-item-label>
									<q-item-label caption class="text-body2">
										{{ new Date(model.created_at).toLocaleDateString() }}
									</q-item-label>
								</q-item-section>

								<q-item-section side>
									<q-icon name="chevron_right" color="grey-6" />
								</q-item-section>
							</q-item>
						</template>
					</ResponsiveTable>
				</q-page>
			</q-page-container>

			<AppFooter
				:username="authStore.user?.username"
				:languages="$lang.languages"
				:current-language-id="$lang._currentLang?.id"
				:show-add-button="true"
				:add-button-route="{ name: 'PROBLEM_CREATE' }"
				add-button-icon="medical_services"
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
