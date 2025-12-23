<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref } from "vue";
import { OrderService, type OrderPageData } from "../service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import CreateOrder from "./Create.vue";
import EditOrder from "./Edit.vue";
import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import IconBtn from "@/components/quasar/btn/IconBtn.vue";
import ConfirmDialog from "../components/ConfirmDialog.vue";
import { orderStatusOptions, paymentStatusOptions } from "../utils";
import Button from "@/components/quasar/btn/Button.vue";

const orderPage = ref<OrderPageData>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const pick = {
	id: false,
	edit: false,
	client_name: true,
	master_name: true,
	problem_names: true,
	part_names: true,
	price: true,
	status: true,
	payment_status: true,
	created_at: true,
	deleted_at: true,
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

			<template #edit:thead>
				<div class="text-center">
					{{ $tl("action") }}
				</div>
			</template>
			<template #edit="{ model }">
				<div class="text-center">
					<IconBtn
						:to="{ name: 'ORDER_VIEW', params: { id: model.id } }"
						icon="visibility"
						tooltipText="view_order"
						withTooltip
					/>
					<IconDialog
						v-if="!model.deleted_at"
						icon="edit"
						:style="'width: 50%;'"
						:fetch="fetch"
						tooltipText="edit_order"
						withTooltip
					>
						<EditOrder :id="model.id" :fetch="fetch" />
					</IconDialog>
					<IconDialog
						v-if="model.deleted_at"
						icon="sync"
						iconColor="positive"
						tooltipText="restore_order"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="false" />
					</IconDialog>
					<IconDialog
						v-if="!model.deleted_at"
						icon="delete"
						iconColor="negative"
						tooltipText="remove_order"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="true" />
					</IconDialog>
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
			<!-- Кастомный мобильный вид для заказов -->
			<template #card="{ model, orderNumber }">
				<q-item class="q-mb-md order-item-bordered">
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="orange" text-color="white" size="sm">
							<q-icon name="assignment" />
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label>
							{{ $tl("client") }}: {{ model?.client_id || "-" }}
						</q-item-label>
						<q-item-label>
							{{ $tl("master") }}: {{ model.master_id || "-" }}
						</q-item-label>
						<q-item-label v-if="model.price">
							{{ $tl("price") }}: {{ model.price?.toLocaleString() }} сум
						</q-item-label>
						<q-item-label>
							{{ $tl("status") }}:
							<q-chip color="primary" outline size="sm">
								{{
									orderStatusOptions.find((item) => item.value === model.status)
										?.label
								}}
							</q-chip>
						</q-item-label>
						<q-item-label>
							{{ $tl("payment_status") }}:
							<q-chip
								:color="model.payment_status === 'paid' ? 'positive' : 'warning'"
								outline
								size="sm"
							>
								{{
									paymentStatusOptions.find(
										(item) => item.value === model.payment_status,
									)?.label
								}}
							</q-chip>
						</q-item-label>
						<q-item-label>№ {{ orderNumber }}</q-item-label>
					</q-item-section>
				</q-item>

				<!-- Кнопки действий -->
				<div class="card-actions flex justify-end gap-4">
					<Button
						:to="{ name: 'ORDER_VIEW', params: { id: model.id } }"
						icon="visibility"
						tooltipText="view_order"
						withTooltip
					/>
					<ButtonDialog
						v-if="!model.deleted_at"
						icon="edit"
						:style="'width: auto;'"
						:fetch="fetch"
						tooltipText="edit_order"
						withTooltip
						flat
						round
						color="primary"
					>
						<EditOrder :id="model.id" :fetch="fetch" />
					</ButtonDialog>
					<ButtonDialog
						v-if="model.deleted_at"
						icon="sync"
						iconColor="positive"
						tooltipText="restore_order"
						withTooltip
						flat
						round
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="false" />
					</ButtonDialog>
					<ButtonDialog
						v-if="!model.deleted_at"
						icon="delete"
						iconColor="negative"
						tooltipText="remove_order"
						withTooltip
						flat
						round
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="true" />
					</ButtonDialog>
				</div>
			</template>
		</ResponsiveTable>
	</PageLoading>
</template>

<style scoped lang="scss">
.order-item-bordered {
	border: 1px solid rgba(0, 0, 0, 0.12);
	border-radius: 8px;
	padding: 12px;
}

.card-actions {
	border-top: 1px solid rgba(0, 0, 0, 0.1);
	padding-top: 12px;

	@media (max-width: 480px) {
		flex-direction: column;
		gap: 8px !important;

		:deep(.q-btn) {
			width: 100% !important;
		}
	}
}
</style>
