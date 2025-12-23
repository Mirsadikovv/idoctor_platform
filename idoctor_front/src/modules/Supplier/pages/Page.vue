<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref } from "vue";
import { SupplierService, type SupplierPageData } from "../service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import CreateSupplier from "./Create.vue";
import EditSupplier from "./Edit.vue";
import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import ConfirmDialog from "../components/ConfirmDialog.vue";

const supplierPage = ref<SupplierPageData>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const pick = {
	id: false,
	edit: false,
	name: true,
	created_at: true,
	deleted_at: true,
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
	<PageLoading :find="page" #="{ fetch }">
		<div class="flex! gap-x-4 items-center mb-3">
			<q-breadcrumbs>
				<q-breadcrumbs-el :label="$tl('supplier_list')" icon="business" />
			</q-breadcrumbs>
			<q-space></q-space>

			<ButtonDialog label="create" :style="'width: auto;'" :fetch="fetch">
				<CreateSupplier :fetch="fetch" />
			</ButtonDialog>
		</div>

		<ResponsiveTable :models="supplierPage" hasOrder>
			<template #name:thead>
				{{ $tl("supplier_name") }}
			</template>
			<template #name="{ model }">
				{{ model.name }}
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
				<q-chip :color="!model.deleted_at ? 'positive' : 'negative'" text-color="white">
					{{ !model.deleted_at ? $tl("ACTIVE") : $tl("DELETED") }}
				</q-chip>
			</template>

			<template #edit:thead>
				<div class="text-center">
					{{ $tl("action") }}
				</div>
			</template>
			<template #edit="{ model }">
				<div class="text-center">
					<IconDialog
						v-if="!model.deleted_at"
						icon="edit"
						:style="'width: 40%;'"
						:fetch="fetch"
						tooltipText="edit_supplier"
						withTooltip
					>
						<EditSupplier :id="model.id" :fetch="fetch" />
					</IconDialog>
					<IconDialog
						v-if="model.deleted_at"
						icon="sync"
						iconColor="positive"
						tooltipText="restore_supplier"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="false" />
					</IconDialog>
					<IconDialog
						v-if="!model.deleted_at"
						icon="delete"
						iconColor="negative"
						tooltipText="remove_supplier"
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
			<!-- Кастомный мобильный вид для поставщиков -->
			<template #card="{ model, orderNumber }">
				<q-item class="q-mb-md supplier-item-bordered">
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="teal" text-color="white" size="sm">
							<q-icon name="business" />
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label>{{ $tl("supplier_name") }}: {{ model.name }}</q-item-label>
						<q-item-label
							>{{ $tl("created_at") }}:
							{{ new Date(model.created_at).toLocaleDateString() }}</q-item-label
						>
						<q-item-label>
							{{ $tl("status") }}:
							<q-chip
								:color="!model.deleted_at ? 'positive' : 'negative'"
								text-color="white"
								size="sm"
							>
								{{ !model.deleted_at ? $tl("ACTIVE") : $tl("DELETED") }}
							</q-chip>
						</q-item-label>
						<q-item-label>№ {{ orderNumber }}</q-item-label>
					</q-item-section>
				</q-item>

				<!-- Кнопки действий -->
				<div class="card-actions flex justify-end gap-4">
					<ButtonDialog
						v-if="!model.deleted_at"
						icon="edit"
						:style="'width: auto;'"
						:fetch="fetch"
						tooltipText="edit_supplier"
						withTooltip
						flat
						round
						color="primary"
					>
						<EditSupplier :id="model.id" :fetch="fetch" />
					</ButtonDialog>
					<ButtonDialog
						v-if="model.deleted_at"
						icon="sync"
						iconColor="positive"
						tooltipText="restore_supplier"
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
						tooltipText="remove_supplier"
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
.supplier-item-bordered {
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
