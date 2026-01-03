<script setup lang="ts">
import { useRouter } from "vue-router";
import { DeviceService, type DeviceType } from "../service";
import { ref } from "vue";
import { useQuasar } from "quasar";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import EditDevice from "./Edit.vue";

export interface Props {
	id: number;
}

const router = useRouter();
const $q = useQuasar();
const { id } = defineProps<Props>();

const model = ref<DeviceType>({} as DeviceType);

async function fetchDevice() {
	const response = await DeviceService.getByID(+id);
	if (!response) return;
	model.value = response;
}

async function deleteDevice() {
	const response = await DeviceService.delete(+id);
	if (!response) return false;
	
	router.push({ name: "DEVICE_PAGE" });
	return true;
}

function confirmDelete() {
	$q.dialog({
		title: 'Подтверждение',
		message: 'Вы уверены, что хотите удалить это устройство?',
		cancel: true,
		persistent: true
	}).onOk(() => {
		deleteDevice();
	});
}

if (+id > 0) {
	fetchDevice();
}
</script>

<template>
	<div class="flex! gap-x-4 items-center mb-6">
		<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
		<q-breadcrumbs>
			<q-breadcrumbs-el
				:label="$tl('device_list')"
				icon="devices"
				:to="{ name: 'DEVICE_PAGE' }"
			/>
			<q-breadcrumbs-el :label="model.name || $tl('device_view_title')" />
		</q-breadcrumbs>
		<q-space />
		
		<div class="flex gap-2" v-if="model.id && !model.deleted_at">
			<ButtonDialog
				:label="$tl('edit_device')"
				icon="edit"
				color="primary"
				:fetch="fetchDevice"
			>
				<EditDevice :id="model.id" :fetch="fetchDevice" />
			</ButtonDialog>
			
			<q-btn
				icon="delete"
				color="negative"
				flat
				round
				@click="confirmDelete"
			>
				<q-tooltip>{{ $tl('remove_device') }}</q-tooltip>
			</q-btn>
		</div>
	</div>

	<div v-if="model.id" class="q-pa-md">
		<q-card class="q-mb-lg">
			<q-card-section>
				<div class="text-h6 q-mb-md flex items-center">
					<q-icon name="devices" class="q-mr-sm" />
					{{ $tl('device_details') }}
				</div>
				
				<q-list>
					<q-item>
						<q-item-section>
							<q-item-label overline>{{ $tl('device_name') }}</q-item-label>
							<q-item-label class="text-h6">{{ model.name }}</q-item-label>
						</q-item-section>
					</q-item>
					
					<q-item v-if="model.brand_name">
						<q-item-section>
							<q-item-label overline>{{ $tl('brand_name') }}</q-item-label>
							<q-item-label>{{ model.brand_name }}</q-item-label>
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