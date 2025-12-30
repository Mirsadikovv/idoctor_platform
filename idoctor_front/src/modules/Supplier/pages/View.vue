<script setup lang="ts">
import { useRouter } from "vue-router";
import { SupplierService, type SupplierType } from "../service";
import { ref } from "vue";
import { useQuasar } from "quasar";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import EditSupplier from "./Edit.vue";

export interface Props {
	id: number;
}

const router = useRouter();
const $q = useQuasar();
const { id } = defineProps<Props>();

const model = ref<SupplierType>({} as SupplierType);

async function fetchSupplier() {
	const response = await SupplierService.getByID(+id);
	if (!response) return;
	model.value = response;
}

async function deleteSupplier() {
	const response = await SupplierService.delete(+id);
	if (!response) return false;
	
	router.push({ name: "SUPPLIER_PAGE" });
	return true;
}

async function restoreSupplier() {
	const response = await SupplierService.restore(+id);
	if (!response) return false;
	
	await fetchSupplier();
	return true;
}

function confirmDelete() {
	$q.dialog({
		title: 'Подтверждение',
		message: 'Вы уверены, что хотите удалить этого поставщика?',
		cancel: true,
		persistent: true
	}).onOk(() => {
		deleteSupplier();
	});
}

function confirmRestore() {
	$q.dialog({
		title: 'Подтверждение',
		message: 'Вы уверены, что хотите восстановить этого поставщика?',
		cancel: true,
		persistent: true
	}).onOk(() => {
		restoreSupplier();
	});
}

if (+id > 0) {
	fetchSupplier();
}
</script>

<template>
	<div class="flex! gap-x-4 items-center mb-6">
		<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
		<q-breadcrumbs>
			<q-breadcrumbs-el
				:label="$tl('supplier_list')"
				icon="business"
				:to="{ name: 'SUPPLIER_PAGE' }"
			/>
			<q-breadcrumbs-el :label="model.name || $tl('supplier_view_title')" />
		</q-breadcrumbs>
		<q-space />
		
		<div class="flex gap-2" v-if="model.id">
			<ButtonDialog
				v-if="!model.deleted_at"
				:label="$tl('edit_supplier')"
				icon="edit"
				color="primary"
				:fetch="fetchSupplier"
			>
				<EditSupplier :id="model.id" :fetch="fetchSupplier" />
			</ButtonDialog>
			
			<q-btn
				v-if="model.deleted_at"
				icon="sync"
				color="positive"
				flat
				round
				@click="confirmRestore"
			>
				<q-tooltip>{{ $tl('restore_supplier') }}</q-tooltip>
			</q-btn>
			
			<q-btn
				v-if="!model.deleted_at"
				icon="delete"
				color="negative"
				flat
				round
				@click="confirmDelete"
			>
				<q-tooltip>{{ $tl('remove_supplier') }}</q-tooltip>
			</q-btn>
		</div>
	</div>

	<div v-if="model.id" class="q-pa-md">
		<q-card class="q-mb-lg">
			<q-card-section>
				<div class="text-h6 q-mb-md flex items-center">
					<q-icon name="business" class="q-mr-sm" color="teal" />
					{{ $tl('supplier_details') }}
				</div>
				
				<q-list>
					<q-item>
						<q-item-section>
							<q-item-label overline>{{ $tl('supplier_name') }}</q-item-label>
							<q-item-label class="text-h6">{{ model.name }}</q-item-label>
						</q-item-section>
					</q-item>
					
					<q-item v-if="model.created_at">
						<q-item-section>
							<q-item-label overline>{{ $tl('created_at') }}</q-item-label>
							<q-item-label>{{ new Date(model.created_at).toLocaleDateString() }}</q-item-label>
						</q-item-section>
					</q-item>
					
					<q-item>
						<q-item-section>
							<q-item-label overline>{{ $tl('status') }}</q-item-label>
							<q-chip :color="!model.deleted_at ? 'positive' : 'negative'" text-color="white">
								{{ !model.deleted_at ? $tl("ACTIVE") : $tl("DELETED") }}
							</q-chip>
						</q-item-section>
					</q-item>
				</q-list>
			</q-card-section>
		</q-card>
	</div>
	
	<div v-else class="flex justify-center q-pa-xl">
		<q-spinner color="primary" size="3em" />
	</div>
</template>