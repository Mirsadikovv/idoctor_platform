<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref, onMounted } from "vue";
import { PartService, type PartPageData } from "../service";
import { DeviceService, SupplierService } from "@/service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import CreatePart from "./Create.vue";
import EditPart from "./Edit.vue";
import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import ConfirmDialog from "../components/ConfirmDialog.vue";
import { buildQuery } from "@/common";

const partPage = ref<PartPageData>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const devices = ref<Array<{ name: string; id: number }>>([]);
const suppliers = ref<Array<{ name: string; id: number }>>([]);

const pick = {
	id: false,
	edit: false,
	name: true,
	device_name: true,
	supplier_name: true,
	created_at: true,
	deleted_at: true,
};

const pikers = ref({});

const queryLimit = buildQuery({
	limit: 100,
});

onMounted(async () => {
	const [deviceData, supplierData] = await Promise.all([
		DeviceService.search(queryLimit),
		SupplierService.search(queryLimit),
	]);

	if (deviceData) {
		devices.value = deviceData;
	}

	if (supplierData) {
		suppliers.value = supplierData;
	}
});

async function page(query: string = "") {
	const params = new URLSearchParams(query);
	const searchParams = {
		page: params.get("page") ? +params.get("page")! : undefined,
		perpage: params.get("perpage") ? +params.get("perpage")! : undefined,
		name: params.get("name") || undefined,
		device_id: params.get("device_id") ? +params.get("device_id")! : undefined,
		supplier_id: params.get("supplier_id") ? +params.get("supplier_id")! : undefined,
		include_deleted: true,
	};

	const response = await PartService.pageWithDelete(searchParams);

	if (!response) return;

	partPage.value = response;
}

function getDeviceName(deviceId: number): string {
	return devices.value.find((device) => device.id === deviceId)?.name || "-";
}

function getSupplierName(supplierId: number): string {
	return suppliers.value.find((supplier) => supplier.id === supplierId)?.name || "-";
}
</script>

<template>
	<PageLoading :find="page" #="{ fetch }">
		<div class="flex! gap-x-4 items-center mb-3">
			<q-breadcrumbs>
				<q-breadcrumbs-el :label="$tl('part_list')" icon="build" />
			</q-breadcrumbs>
			<q-space></q-space>

			<ButtonDialog label="create" :style="'width: auto;'" :fetch="fetch">
				<CreatePart :fetch="fetch" />
			</ButtonDialog>
		</div>

		<ResponsiveTable :models="partPage" hasOrder>
			<template #name:thead>
				{{ $tl("part_name") }}
			</template>
			<template #name="{ model }">
				{{ model.name }}
			</template>

			<template #device_name:thead>
				{{ $tl("device") }}
			</template>
			<template #device_name="{ model }">
				{{ getDeviceName(model.device_id) }}
			</template>

			<template #supplier_name:thead>
				{{ $tl("supplier") }}
			</template>
			<template #supplier_name="{ model }">
				{{ getSupplierName(model.supplier_id) }}
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
						tooltipText="edit_part"
						withTooltip
					>
						<EditPart :id="model.id" :fetch="fetch" />
					</IconDialog>
					<IconDialog
						v-if="model.deleted_at"
						icon="sync"
						iconColor="positive"
						tooltipText="restore_part"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="false" />
					</IconDialog>
					<IconDialog
						v-if="!model.deleted_at"
						icon="delete"
						iconColor="negative"
						tooltipText="remove_part"
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
			<!-- Кастомный мобильный вид для деталей -->
			<template #card="{ model, orderNumber }">
				<q-item class="q-mb-md part-item-bordered">
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="indigo" text-color="white" size="sm">
							<q-icon name="build" />
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label>
							<span class="text-bold">{{ $tl("part_name") }}:</span>
							{{ model.name }}
						</q-item-label>
						<q-item-label>
							<span class="text-bold">{{ $tl("device") }}:</span>
							{{ getDeviceName(model.device_id) }}
						</q-item-label>
						<q-item-label>
							<span class="text-bold">{{ $tl("supplier") }}:</span>
							{{ getSupplierName(model.supplier_id) }}
						</q-item-label>
						<q-item-label>
							<span class="text-bold">{{ $tl("created_at") }}:</span>
							{{ new Date(model.created_at).toLocaleDateString() }}
						</q-item-label>
						<q-item-label>
							<span class="text-bold">{{ $tl("status") }}:</span>
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
						tooltipText="edit_part"
						withTooltip
						flat
						round
						color="primary"
					>
						<EditPart :id="model.id" :fetch="fetch" />
					</ButtonDialog>
					<ButtonDialog
						v-if="model.deleted_at"
						icon="sync"
						iconColor="positive"
						tooltipText="restore_part"
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
						tooltipText="remove_part"
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
.part-item-bordered {
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
