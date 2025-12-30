<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref } from "vue";
import { OrderService, type OrderPageData } from "../service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import { orderStatusOptions, paymentStatusOptions } from "../utils";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const orderPage = ref<OrderPageData>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const pick = {
	id: true,
	client_name: true,
	master_name: true,
	problem_names: true,
	part_names: true,
	price: true,
	status: true,
	payment_status: true,
	created_at: true,
};

const pikers = ref({});

async function page(query: string = "") {
	const response = await OrderService.pageWithDelete(query);

	if (!response) return;

	orderPage.value = response;
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
					class="bg-white text-gray-900 overflow-auto p-4 pt-20"
				>
					<ResponsiveTable :models="orderPage" hasOrder>
						<template #id:thead>{{ $tl("id") }}</template>
						<template #id="{ model }">
							<router-link
								:to="{ name: 'ORDER_VIEW', params: { id: model.id } }"
								class="text-primary text-decoration-none"
							>
								#{{ model.id }}
							</router-link>
						</template>

						<template #client_name:thead>
							{{ $tl("client") }}
						</template>
						<template #client_name="{ model }">
							{{ model?.client_id }}
						</template>

						<template #master_name:thead>
							{{ $tl("master") }}
						</template>
						<template #master_name="{ model }">
							{{ model.master_id }}
						</template>

						<template #price:thead>
							{{ $tl("price") }}
						</template>
						<template #price="{ model }">
							{{ model.price?.toLocaleString() }} сум
						</template>

						<template #status:thead>
							{{ $tl("status") }}
						</template>
						<template #status="{ model }">
							<q-chip color="primary" outline>
								{{
									orderStatusOptions.find((item) => item.value === model.status)
										?.label
								}}
							</q-chip>
						</template>

						<template #payment_status:thead>
							{{ $tl("payment_status") }}
						</template>
						<template #payment_status="{ model }">
							<q-chip
								:color="model.payment_status === 'paid' ? 'positive' : 'warning'"
								outline
							>
								{{
									paymentStatusOptions.find(
										(item) => item.value === model.payment_status,
									)?.label
								}}
							</q-chip>
						</template>

						<template #tfoot="{ totalPages }">
							<TablePaginate
								v-model:pikers="pikers"
								:total="totalPages"
								:pick="pick"
								@page="fetch"
							/>
						</template>
						<!-- Кастомный мобильный вид для заказов -->
						<template #card="{ model, orderNumber }">
							<q-item
								class="order-item-telegram"
								clickable
								:to="{ name: 'ORDER_EDIT', params: { id: model.id } }"
							>
								<q-item-section avatar v-if="orderNumber">
									<q-avatar color="primary" text-color="white" size="md">
										{{ orderNumber }}
									</q-avatar>
								</q-item-section>

								<q-item-section>
									<q-item-label class="text-weight-bold text-h6">
										{{ $tl("order") }} #{{ model.id }}
									</q-item-label>
									<q-item-label caption class="text-body2">
										{{ $tl("client") }}: {{ model?.client_id || "-" }}
									</q-item-label>
									<q-item-label caption class="text-body2" v-if="model.price">
										{{ model.price?.toLocaleString() }} сум
									</q-item-label>
									<!-- Чипы статусов -->
									<div class="q-mt-xs flex gap-1">
										<q-chip
											:color="
												model.payment_status === 'paid'
													? 'positive'
													: 'warning'
											"
											outline
											size="sm"
											dense
										>
											{{
												paymentStatusOptions.find(
													(item) => item.value === model.payment_status,
												)?.label
											}}
										</q-chip>
										<q-chip color="primary" outline size="sm" dense>
											{{
												orderStatusOptions.find(
													(item) => item.value === model.status,
												)?.label
											}}
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
				:add-button-route="{ name: 'ORDER_CREATE' }"
				add-button-icon="add_circle"
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
.order-item-telegram {
	max-height: 120px;
	min-height: 90px;
	background: white;
	padding: 12px;
	transition: all 0.2s ease;
	cursor: pointer;

	&:hover {
		background: rgba(0, 0, 0, 0.02);
		border-color: rgba(0, 0, 0, 0.12);
		transform: translateY(-1px);
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
	}

	&:active {
		transform: translateY(0);
		box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
	}

	.q-item__section--avatar {
		padding-right: 16px;
	}

	.q-item__section--side {
		padding-left: 8px;
	}
}
</style>
