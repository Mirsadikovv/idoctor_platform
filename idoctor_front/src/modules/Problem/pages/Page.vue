<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref } from "vue";
import { ProblemService, type ProblemPageData } from "../service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import CreateProblem from "./Create.vue";
import EditProblem from "./Edit.vue";
import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import ConfirmDialog from "../components/ConfirmDialog.vue";

const problemPage = ref<ProblemPageData>({
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
	price: true,
	created_at: true,
};

const pikers = ref({});

async function page(query: string = "") {
	const params = new URLSearchParams(query);
	const searchParams = {
		page: params.get("page") ? +params.get("page")! : undefined,
		perpage: params.get("perpage") ? +params.get("perpage")! : undefined,
		name: params.get("name") || undefined,
		min_price: params.get("min_price") ? +params.get("min_price")! : undefined,
		max_price: params.get("max_price") ? +params.get("max_price")! : undefined,
	};

	const response = await ProblemService.page(searchParams);

	if (!response) return;

	problemPage.value = response;
}
</script>

<template>
	<PageLoading :find="page" #="{ fetch }">
		<div class="flex! gap-x-4 items-center mb-3">
			<q-breadcrumbs>
				<q-breadcrumbs-el :label="$tl('problem_list')" icon="medical_services" />
			</q-breadcrumbs>
			<q-space></q-space>

			<ButtonDialog label="create" :style="'width: auto;'" :fetch="fetch">
				<CreateProblem :fetch="fetch" />
			</ButtonDialog>
		</div>

		<ResponsiveTable :models="problemPage" hasOrder>
			<template #name:thead>
				{{ $tl("problem_name") }}
			</template>
			<template #name="{ model }">
				{{ model.name }}
			</template>

			<template #price:thead>
				{{ $tl("price") }}
			</template>
			<template #price="{ model }"> {{ model.price.toLocaleString() }} сум </template>

			<template #created_at:thead>
				{{ $tl("created_at") }}
			</template>
			<template #created_at="{ model }">
				{{ new Date(model.created_at).toLocaleDateString() }}
			</template>

			<template #edit:thead>
				<div class="text-center">
					{{ $tl("action") }}
				</div>
			</template>
			<template #edit="{ model }">
				<div class="text-center">
					<IconDialog
						icon="edit"
						:style="'width: 40%;'"
						:fetch="fetch"
						tooltipText="edit_problem"
						withTooltip
					>
						<EditProblem :id="model.id" :fetch="fetch" />
					</IconDialog>
					<IconDialog
						icon="delete"
						iconColor="negative"
						tooltipText="remove_problem"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" />
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
			<!-- Кастомный мобильный вид для медицинских проблем -->
			<template #card="{ model, orderNumber }">
				<q-item class="q-mb-md problem-item-bordered">
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="red" text-color="white" size="sm">
							<q-icon name="medical_services" />
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label>{{ $tl("problem_name") }}: {{ model.name }}</q-item-label>
						<q-item-label
							>{{ $tl("price") }}:
							{{ model.price.toLocaleString() }} сум</q-item-label
						>
						<q-item-label
							>{{ $tl("created_at") }}:
							{{ new Date(model.created_at).toLocaleDateString() }}</q-item-label
						>
						<q-item-label>№ {{ orderNumber }}</q-item-label>
					</q-item-section>
				</q-item>

				<!-- Кнопки действий -->
				<div class="card-actions flex justify-end gap-4">
					<ButtonDialog
						icon="edit"
						:style="'width: auto;'"
						:fetch="fetch"
						tooltipText="edit_problem"
						withTooltip
						flat
						round
						color="primary"
					>
						<EditProblem :id="model.id" :fetch="fetch" />
					</ButtonDialog>
					<ButtonDialog
						icon="delete"
						iconColor="negative"
						tooltipText="remove_problem"
						withTooltip
						flat
						round
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" />
					</ButtonDialog>
				</div>
			</template>
		</ResponsiveTable>
	</PageLoading>
</template>

<style scoped lang="scss">
.problem-item-bordered {
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
