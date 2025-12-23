<script setup lang="ts">
import { ref } from "vue";
import { RoleService, type RolePageData, type RoleTypePartialType } from "@/service";
import PageLoading from "@/components/PageLoading.vue";
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";

import { useRouter } from "vue-router";
import Expasion from "@/components/quasar/Expasion.vue";
import Form from "@/components/quasar/form/Form.vue";
import Search from "@/components/quasar/search/Search.vue";

import Button from "@/components/quasar/btn/Button.vue";
import IconBtn from "@/components/quasar/btn/IconBtn.vue";
const router = useRouter();

const roles = ref<RolePageData>({
	currentPage: 1,
	data: [],
	pageSize: 0,
	totalPages: 0,
	totalRows: 0,
});

const pick = {
	id: false,
	edit: false,
	name: true,
	description: true,
};

const pikers = ref({});

async function getRoles(query: string = "") {
	const response = await RoleService.page(query);

	if (!response) return;

	roles.value = response;
}

let searchModel = ref<RoleTypePartialType>({});

async function save() {
	searchModel.value = {};
	await getRoles("");
	router.replace({ query: {} });
	return true;
}
</script>

<template>
	<PageLoading :find="getRoles" #="{ fetch, loading }">
		<div class="flex! gap-x-4 items-center mb-3">
			<q-breadcrumbs>
				<q-breadcrumbs-el :label="$tl('roles_page')" icon="article" />
			</q-breadcrumbs>
			<q-space></q-space>

			<Button
				:to="{
					name: 'CREATE_ROLE',
				}"
			>
				{{ $tl("create") }}
			</Button>
		</div>
		<Expasion class="mb-4">
			<Form :model-value="searchModel" :save="save" is-cleaned-style>
				<template #fio="{ model }">
					<Search
						query-name="name"
						label="name"
						v-model="model.name"
						@search="fetch"
						input-debounce="500"
						class="col-lg-3 col-md-6 col-12"
					></Search>
				</template>

				<template #actions="{ loading }">
					<div class="col-12">
						<Button :loading="loading" type="submit" class="ml-auto">
							{{ $tl("save") }}
						</Button>
					</div>
				</template>
			</Form>
		</Expasion>

		<ResponsiveTable :models="roles" :pick="pikers" :loading="loading" has-order>
			<template #name:thead> </template>
			<template #name="{ model }">
				{{ model.name }}
			</template>

			<template #description:thead> </template>
			<template #description="{ model }">
				{{ model.description }}
			</template>

			<template #edit:thead>
				<div class="text-center">
					{{ $tl("action") }}
				</div>
			</template>
			<template #edit="{ model }">
				<div class="text-center">
					<IconBtn
						:to="{
							name: 'EDIT_ROLE',
							params: { id: model.id },
						}"
						icon="edit"
					/>
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
			<!-- Кастомный мобильный вид для ролей -->
			<template #card="{ model, orderNumber }">
				<q-item class="q-mb-md role-item-bordered">
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="secondary" text-color="white" size="sm">
							<q-icon name="security" />
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label>{{ $tl("name") }}: {{ model.name }}</q-item-label>
						<q-item-label
							>{{ $tl("description") }}: {{ model.description || "-" }}</q-item-label
						>
						<q-item-label>№ {{ orderNumber }}</q-item-label>
					</q-item-section>
				</q-item>

				<!-- Кнопки действий -->
				<div class="card-actions flex justify-end gap-4">
					<Button
						:to="{
							name: 'EDIT_ROLE',
							params: { id: model.id },
						}"
						icon="edit"
						color="primary"
					>
						{{ $tl("edit") }}
					</Button>
				</div>
			</template>
		</ResponsiveTable>
	</PageLoading>
</template>

<style scoped lang="scss">
.role-item-bordered {
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
