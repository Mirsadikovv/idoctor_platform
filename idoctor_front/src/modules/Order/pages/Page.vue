<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref } from "vue";
import { OrderService, type OrderPageData } from "../service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import CreateOrder from "./Create.vue";
import { orderStatusOptions, paymentStatusOptions } from "../utils";

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
	<PageLoading :find="page" #="{ fetch }">
		<div class="flex! gap-x-4 items-center mb-3">
			<q-breadcrumbs>
				<q-breadcrumbs-el :label="$tl('order_list')" icon="assignment" />
			</q-breadcrumbs>
			<q-space></q-space>

			<ButtonDialog label="create" :style="'width: auto;'" :fetch="fetch">
				<CreateOrder :fetch="fetch" />
			</ButtonDialog>
		</div>

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
			<template #price="{ model }"> {{ model.price?.toLocaleString() }} сум </template>

			<template #status:thead>
				{{ $tl("status") }}
			</template>
			<template #status="{ model }">
				<q-chip color="primary" outline>
					{{ orderStatusOptions.find((item) => item.value === model.status)?.label }}
				</q-chip>
			</template>

			<template #payment_status:thead>
				{{ $tl("payment_status") }}
			</template>
			<template #payment_status="{ model }">
				<q-chip :color="model.payment_status === 'paid' ? 'positive' : 'warning'" outline>
					{{
						paymentStatusOptions.find((item) => item.value === model.payment_status)
							?.label
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
					:to="{ name: 'ORDER_VIEW', params: { id: model.id } }"
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
								:color="model.payment_status === 'paid' ? 'positive' : 'warning'"
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
									orderStatusOptions.find((item) => item.value === model.status)
										?.label
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
	</PageLoading>
</template>

<style scoped lang="scss">
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
