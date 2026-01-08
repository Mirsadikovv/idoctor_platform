<script setup lang="ts">
import { ref } from "vue";
import { RoleService, type RolePageData } from "@/service";
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

const roles = ref<RolePageData>({
	currentPage: 1,
	data: [],
	pageSize: 0,
	totalPages: 0,
	totalRows: 0,
});

const pick = {
	id: false,
	name: true,
	description: true,
};

const pikers = ref({});

async function getRoles(query: string = "") {
	const response = await RoleService.page(query);

	if (!response) return;

	roles.value = response;
}
</script>

<template>
	<PageLoading :find="getRoles" #="{ fetch, loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
					<ResponsiveTable :models="roles" :pick="pikers" :loading="loading" has-order>
						<template #name:thead> </template>
						<template #name="{ model }">
							<router-link
								:to="{ name: 'ROLE_VIEW', params: { id: model.id } }"
								class="text-primary text-decoration-none"
							>
								{{ model.name }}
							</router-link>
						</template>

						<template #description:thead> </template>
						<template #description="{ model }">
							{{ model.description }}
						</template>

						<template #edit="{ model }">
							<div class="text-center">
								<IconBtn
									v-if="$canPage('ROLE_EDIT')"
									:to="{
										name: 'ROLE_EDIT',
										params: { id: model.id },
									}"
									icon="edit"
								/>
								<IconBtn
									v-if="$canPage('ROLE_VIEW')"
									:to="{
										name: 'ROLE_VIEW',
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
						<!-- Кастомный мобильный вид для ролей -->
						<template #card="{ model, orderNumber }">
							<q-item
								class="role-item-telegram"
								clickable
								:to="{
									name: $canPage('ROLE_EDIT') ? 'ROLE_EDIT' : 'ROLE_VIEW',
									params: { id: model.id },
								}"
							>
								<q-item-section avatar v-if="orderNumber">
									<q-avatar color="primary" text-color="white" size="md">
										{{ orderNumber }}
									</q-avatar>
								</q-item-section>

								<q-item-section>
									<q-item-label class="text-xl">
										<strong>{{ $tl("name") }}: </strong> {{ model.name }}
									</q-item-label>
									<q-item-label
										caption
										class="text-body2"
										v-if="model.description"
									>
										{{ model.description }}
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
				:add-button-route="{ name: 'CREATE_ROLE' }"
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
