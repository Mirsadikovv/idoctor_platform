<script setup lang="ts">
import { ref } from "vue";
import { UserService, type UserPage } from "@/service";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const userPage = ref<UserPage>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const pick = {
	id: false,
	fullName: true,
	username: true,
	phoneNumber: true,
	dateOfBirth: true,
	gender: true,
};

const pikers = ref({});

async function page(query: string = "") {
	const response = await UserService.page(query);

	if (!response) return;

	userPage.value = response;
}
</script>

<template>
	<PageLoading :find="page" #="{ fetch, loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page
					:style="{
						height: 'calc(var(--app-height, 100vh) - 150px)',
					}"
					class="bg-white text-gray-900 overflow-auto p-4 pt-24"
				>
					<ResponsiveTable :models="userPage" hasOrder :pick="pikers" :loading="loading">
						<template #fullName:thead>{{ $tl("full_name") }}</template>
						<template #fullName="{ model }">
							<router-link
								:to="{ name: 'USER_VIEW', params: { id: model.id } }"
								class="text-primary text-decoration-none"
							>
								{{ model?.lastName }} {{ model?.firstName }} {{ model?.middleName }}
							</router-link>
						</template>

						<template #username:thead>{{ $tl("username") }}</template>
						<template #username="{ model }">
							{{ model?.username }}
						</template>

						<template #phoneNumber:thead>{{ $tl("phone_number") }}</template>
						<template #phoneNumber="{ model }">
							{{ model?.phoneNumber || "-" }}
						</template>

						<template #dateOfBirth:thead>{{ $tl("date_of_birth") }}</template>
						<template #dateOfBirth="{ model }">
							{{
								model?.dateOfBirth
									? new Date(model.dateOfBirth).toLocaleDateString()
									: "-"
							}}
						</template>

						<template #gender:thead>{{ $tl("gender") }}</template>
						<template #gender="{ model }">
							<q-chip v-if="model?.gender" color="secondary" outline>
								{{ $tl(model.gender) }}
							</q-chip>
							<span v-else>-</span>
						</template>

						<template #tfoot="{ totalPages }">
							<TablePaginate
								v-model:pikers="pikers"
								:total="totalPages"
								:pick="pick"
								@page="fetch"
							/>
						</template>
						<!-- Кастомный мобильный вид -->
						<template #card="{ model, orderNumber }">
							<q-item
								class="user-item-telegram"
								clickable
								:to="{ name: 'USER_EDIT', params: { id: model.id } }"
							>
								<q-item-section avatar v-if="orderNumber">
									<q-avatar color="primary" text-color="white" size="md">
										{{ orderNumber }}
									</q-avatar>
								</q-item-section>

								<q-item-section>
									<q-item-label class="text-weight-bold text-h6">
										{{ model?.lastName }} {{ model?.firstName }}
										{{ model?.middleName }}
									</q-item-label>
									<q-item-label caption class="text-body2">
										@{{ model?.username || $tl("no_username") }}
									</q-item-label>
									<q-item-label
										caption
										class="text-body2"
										v-if="model?.phoneNumber"
									>
										{{ model?.phoneNumber }}
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
