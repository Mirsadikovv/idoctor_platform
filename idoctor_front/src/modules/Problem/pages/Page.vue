<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref } from "vue";
import { ProblemService, type ProblemPageData } from "../service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import CreateProblem from "./Create.vue";

const problemPage = ref<ProblemPageData>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const pick = {
	id: false,
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
				<router-link
					:to="{ name: 'PROBLEM_VIEW', params: { id: model.id } }"
					class="text-primary text-decoration-none"
				>
					{{ model.name }}
				</router-link>
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
				<q-item
					class="problem-item-telegram"
					clickable
					:to="{ name: 'PROBLEM_VIEW', params: { id: model.id } }"
				>
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="primary" text-color="white" size="md">
							{{ orderNumber }}
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label class="text-weight-bold text-h6">
							{{ model.name }}
						</q-item-label>
						<q-item-label caption class="text-body2">
							{{ $tl("price") }}: {{ model.price.toLocaleString() }} сум
						</q-item-label>
						<q-item-label caption class="text-body2">
							{{ new Date(model.created_at).toLocaleDateString() }}
						</q-item-label>
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
.problem-item-telegram {
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
