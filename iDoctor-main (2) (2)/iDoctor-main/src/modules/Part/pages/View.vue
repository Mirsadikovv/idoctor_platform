<script setup lang="ts">
import { useRouter } from "vue-router";
import { PartService, type PartType } from "../service";
import { ref } from "vue";
import { useQuasar } from "quasar";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import EditPart from "./Edit.vue";

export interface Props {
	id: number;
}

const router = useRouter();
const $q = useQuasar();
const { id } = defineProps<Props>();

const model = ref<PartType>({} as PartType);

async function fetchPart() {
	const response = await PartService.getByID(+id);
	if (!response) return;
	model.value = response;
}

async function deletePart() {
	const response = await PartService.delete(+id);
	if (!response) return false;
	
	router.push({ name: "PART_PAGE" });
	return true;
}

async function restorePart() {
	const response = await PartService.restore(+id);
	if (!response) return false;
	
	await fetchPart();
	return true;
}

function confirmDelete() {
	$q.dialog({
		title: 'Подтверждение',
		message: 'Вы уверены, что хотите удалить эту запчасть?',
		cancel: true,
		persistent: true
	}).onOk(() => {
		deletePart();
	});
}

function confirmRestore() {
	$q.dialog({
		title: 'Подтверждение',
		message: 'Вы уверены, что хотите восстановить эту запчасть?',
		cancel: true,
		persistent: true
	}).onOk(() => {
		restorePart();
	});
}

if (+id > 0) {
	fetchPart();
}
</script>

<template>
	<div class="flex! gap-x-4 items-center mb-6">
		<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
		<q-breadcrumbs>
			<q-breadcrumbs-el
				:label="$tl('part_list')"
				icon="build"
				:to="{ name: 'PART_PAGE' }"
			/>
			<q-breadcrumbs-el :label="model.name || $tl('part_view_title')" />
		</q-breadcrumbs>
		<q-space />
		
		<div class="flex gap-2" v-if="model.id">
			<ButtonDialog
				v-if="!model.deleted_at"
				:label="$tl('edit_part')"
				icon="edit"
				color="primary"
				:fetch="fetchPart"
			>
				<EditPart :id="model.id" :fetch="fetchPart" />
			</ButtonDialog>
			
			<q-btn
				v-if="model.deleted_at"
				icon="sync"
				color="positive"
				flat
				round
				@click="confirmRestore"
			>
				<q-tooltip>{{ $tl('restore_part') }}</q-tooltip>
			</q-btn>
			
			<q-btn
				v-if="!model.deleted_at"
				icon="delete"
				color="negative"
				flat
				round
				@click="confirmDelete"
			>
				<q-tooltip>{{ $tl('remove_part') }}</q-tooltip>
			</q-btn>
		</div>
	</div>

	<div v-if="model.id" class="q-pa-md">
		<q-card class="q-mb-lg">
			<q-card-section>
				<div class="text-h6 q-mb-md flex items-center">
					<q-icon name="build" class="q-mr-sm" color="orange" />
					{{ $tl('part_details') }}
				</div>
				
				<q-list>
					<q-item>
						<q-item-section>
							<q-item-label overline>{{ $tl('part_name') }}</q-item-label>
							<q-item-label class="text-h6">{{ model.name }}</q-item-label>
						</q-item-section>
					</q-item>
					
					<q-item v-if="model.device_name">
						<q-item-section>
							<q-item-label overline>{{ $tl('device') }}</q-item-label>
							<q-item-label>{{ model.device_name }}</q-item-label>
						</q-item-section>
					</q-item>
					
					<q-item v-if="model.supplier_name">
						<q-item-section>
							<q-item-label overline>{{ $tl('supplier') }}</q-item-label>
							<q-item-label>{{ model.supplier_name }}</q-item-label>
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