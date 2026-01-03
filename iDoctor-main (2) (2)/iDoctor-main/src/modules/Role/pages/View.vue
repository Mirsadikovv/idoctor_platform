<script setup lang="ts">
import { useRouter } from "vue-router";
import { RoleService, type RoleType } from "@/service";
import { ref } from "vue";
import { useQuasar } from "quasar";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import EditRole from "./Edit.vue";

export interface Props {
	id: number;
}

const router = useRouter();
const $q = useQuasar();
const { id } = defineProps<Props>();

const model = ref<RoleType>({} as RoleType);

async function fetchRole() {
	const response = await RoleService.getByID(+id);
	if (!response) return;
	model.value = response;
}

async function deleteRole() {
	const response = await RoleService.delete(+id);
	if (!response) return false;
	
	router.push({ name: "PAGE_ROLE" });
	return true;
}

function confirmDelete() {
	$q.dialog({
		title: 'Подтверждение',
		message: 'Вы уверены, что хотите удалить эту роль?',
		cancel: true,
		persistent: true
	}).onOk(() => {
		deleteRole();
	});
}

if (+id > 0) {
	fetchRole();
}
</script>

<template>
	<div class="flex! gap-x-4 items-center mb-6">
		<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
		<q-breadcrumbs>
			<q-breadcrumbs-el
				:label="$tl('roles_page')"
				icon="article"
				:to="{ name: 'PAGE_ROLE' }"
			/>
			<q-breadcrumbs-el :label="model.name || $tl('view_role')" />
		</q-breadcrumbs>
		<q-space />
		
		<div class="flex gap-2" v-if="model.id">
			<ButtonDialog
				:label="$tl('edit_role')"
				icon="edit"
				color="primary"
				:fetch="fetchRole"
			>
				<EditRole :id="model.id" :fetch="fetchRole" />
			</ButtonDialog>
			
			<q-btn
				icon="delete"
				color="negative"
				flat
				round
				@click="confirmDelete"
			>
				<q-tooltip>{{ $tl('delete_role') }}</q-tooltip>
			</q-btn>
		</div>
	</div>

	<div v-if="model.id" class="q-pa-md">
		<q-card class="q-mb-lg">
			<q-card-section>
				<div class="text-h6 q-mb-md flex items-center">
					<q-icon name="security" class="q-mr-sm" />
					{{ $tl('role_details') }}
				</div>
				
				<q-list>
					<q-item>
						<q-item-section>
							<q-item-label overline>{{ $tl('name') }}</q-item-label>
							<q-item-label class="text-h6">{{ model.name }}</q-item-label>
						</q-item-section>
					</q-item>
					
					<q-item v-if="model.description">
						<q-item-section>
							<q-item-label overline>{{ $tl('description') }}</q-item-label>
							<q-item-label>{{ model.description }}</q-item-label>
						</q-item-section>
					</q-item>
					
					<q-item v-if="model.permissions">
						<q-item-section>
							<q-item-label overline>{{ $tl('permissions') }}</q-item-label>
							<q-item-label>{{ $tl('permissions_count') }}: {{ model.permissions?.length || 0 }}</q-item-label>
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