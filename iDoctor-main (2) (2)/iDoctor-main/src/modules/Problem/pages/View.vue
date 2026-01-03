<script setup lang="ts">
import { useRouter } from "vue-router";
import { ProblemService, type ProblemType } from "../service";
import { ref } from "vue";
import { useQuasar } from "quasar";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import EditProblem from "./Edit.vue";

export interface Props {
	id: number;
}

const router = useRouter();
const $q = useQuasar();
const { id } = defineProps<Props>();

const model = ref<ProblemType>({} as ProblemType);

async function fetchProblem() {
	const response = await ProblemService.getByID(+id);
	if (!response) return;
	model.value = response;
}

async function deleteProblem() {
	const response = await ProblemService.delete(+id);
	if (!response) return false;
	
	router.push({ name: "PROBLEM_PAGE" });
	return true;
}

function confirmDelete() {
	$q.dialog({
		title: 'Подтверждение',
		message: 'Вы уверены, что хотите удалить эту проблему?',
		cancel: true,
		persistent: true
	}).onOk(() => {
		deleteProblem();
	});
}

if (+id > 0) {
	fetchProblem();
}
</script>

<template>
	<div class="flex! gap-x-4 items-center mb-6">
		<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
		<q-breadcrumbs>
			<q-breadcrumbs-el
				:label="$tl('problem_list')"
				icon="medical_services"
				:to="{ name: 'PROBLEM_PAGE' }"
			/>
			<q-breadcrumbs-el :label="model.name || $tl('problem_view_title')" />
		</q-breadcrumbs>
		<q-space />
		
		<div class="flex gap-2" v-if="model.id">
			<ButtonDialog
				:label="$tl('edit_problem')"
				icon="edit"
				color="primary"
				:fetch="fetchProblem"
			>
				<EditProblem :id="model.id" :fetch="fetchProblem" />
			</ButtonDialog>
			
			<q-btn
				icon="delete"
				color="negative"
				flat
				round
				@click="confirmDelete"
			>
				<q-tooltip>{{ $tl('remove_problem') }}</q-tooltip>
			</q-btn>
		</div>
	</div>

	<div v-if="model.id" class="q-pa-md">
		<q-card class="q-mb-lg">
			<q-card-section>
				<div class="text-h6 q-mb-md flex items-center">
					<q-icon name="medical_services" class="q-mr-sm" color="red" />
					{{ $tl('problem_details') }}
				</div>
				
				<q-list>
					<q-item>
						<q-item-section>
							<q-item-label overline>{{ $tl('problem_name') }}</q-item-label>
							<q-item-label class="text-h6">{{ model.name }}</q-item-label>
						</q-item-section>
					</q-item>
					
					<q-item v-if="model.price">
						<q-item-section>
							<q-item-label overline>{{ $tl('price') }}</q-item-label>
							<q-item-label class="text-h6 text-weight-bold text-positive">
								{{ model.price.toLocaleString() }} сум
							</q-item-label>
						</q-item-section>
					</q-item>
					
					<q-item v-if="model.created_at">
						<q-item-section>
							<q-item-label overline>{{ $tl('created_at') }}</q-item-label>
							<q-item-label>{{ new Date(model.created_at).toLocaleDateString() }}</q-item-label>
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