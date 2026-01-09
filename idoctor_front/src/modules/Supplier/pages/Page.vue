<script setup lang="ts">
import { ref } from "vue";
import { SupplierService, type SupplierPageData } from "../service";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import IconBtn from "@/components/quasar/btn/IconBtn.vue";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const supplierPage = ref<SupplierPageData>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const pick = {
	id: false,
	name: true,
	created_at: true,
	status: true,
};

const pikers = ref({});

async function page(query: string = "") {
	const params = new URLSearchParams(query);
	const searchParams = {
		page: params.get("page") ? +params.get("page")! : undefined,
		perpage: params.get("perpage") ? +params.get("perpage")! : undefined,
		name: params.get("name") || undefined,
		include_deleted: true,
	};

	const response = await SupplierService.pageWithDelete(searchParams);

	if (!response) return;

	supplierPage.value = response;
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
					<ResponsiveTable :models="supplierPage" hasOrder :loading="loading">
						<template #name:thead>
							{{ $tl("supplier_name") }}
						</template>
						<template #name="{ model }">
							<router-link
								:to="{ name: 'SUPPLIER_VIEW', params: { id: model.id } }"
								class="text-primary text-decoration-none"
							>
								{{ model.name }}
							</router-link>
						</template>

						<template #created_at:thead>
							{{ $tl("created_at") }}
						</template>
						<template #created_at="{ model }">
							{{ new Date(model.created_at).toLocaleDateString() }}
						</template>

						<template #status:thead>
							{{ $tl("status") }}
						</template>
						<template #status="{ model }">
							<q-chip
								:color="!model.deleted_at ? 'positive' : 'negative'"
								text-color="white"
							>
								{{ !model.deleted_at ? $tl("ACTIVE") : $tl("DELETED") }}
							</q-chip>
						</template>

						<template #edit="{ model }">
							<div class="text-center">
								<IconBtn
									v-if="$canPage('SUPPLIER_EDIT')"
									:to="{
										name: 'SUPPLIER_EDIT',
										params: { id: model.id },
									}"
									icon="edit"
								/>
								<IconBtn
									v-if="$canPage('SUPPLIER_VIEW')"
									:to="{
										name: 'SUPPLIER_VIEW',
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
						<!-- Кастомный мобильный вид для поставщиков -->
						<template #card="{ model, orderNumber }">
							<q-item
								class="supplier-item-telegram"
								clickable
								:to="{
									name: $canPage('SUPPLIER_EDIT') ? 'SUPPLIER_EDIT' : 'SUPPLIER_VIEW',
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
										{{ new Date(model.created_at).toLocaleDateString() }}
									</q-item-label>
									<!-- Чип статуса -->
									<div class="q-mt-xs">
										<q-chip
											:color="!model.deleted_at ? 'positive' : 'negative'"
											outline
											size="sm"
											dense
										>
											{{ !model.deleted_at ? $tl("ACTIVE") : $tl("DELETED") }}
										</q-chip>
									</div>
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
				:add-button-route="{ name: 'SUPPLIER_CREATE' }"
				add-button-icon="business"
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
