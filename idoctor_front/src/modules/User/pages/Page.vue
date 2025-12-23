<script setup lang="ts">
import { ref } from "vue";
import { UserService, type UserPage } from "@/service";

import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import IconBtn from "@/components/quasar/btn/IconBtn.vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";

import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import Button from "@/components/quasar/btn/Button.vue";
import CreateUser from "./Create.vue";
import EditUser from "./Edit.vue";
import ConfirmDialog from "../components/ConfirmDialog.vue";

const userPage = ref<UserPage>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const pick = {
	id: false,
	edit: false,
	fullName: true,
	username: true,
	phoneNumber: true,
	dateOfBirth: true,
	gender: true,
};

const pikers = ref({});

async function page(query: string = "") {
	const response = await UserService.page(query);

	if (!response) return;

	userPage.value = response;
}
</script>

<template>
	<PageLoading :find="page" #="{ fetch, loading }">
		<div class="flex! gap-x-4 items-center mb-3">
			<q-breadcrumbs>
				<q-breadcrumbs-el :label="$tl('user_list')" icon="article" />
			</q-breadcrumbs>
			<q-space></q-space>

			<ButtonDialog label="create" :style="'width: auto'" :fetch="fetch">
				<CreateUser :fetch="fetch" />
			</ButtonDialog>
		</div>

		<ResponsiveTable :models="userPage" hasOrder :pick="pikers" :loading="loading">
			<template #fullName:thead>{{ $tl("full_name") }}</template>
			<template #fullName="{ model }">
				{{ model?.lastName }} {{ model?.firstName }} {{ model?.middleName }}
			</template>

			<template #username:thead>{{ $tl("username") }}</template>
			<template #username="{ model }">
				{{ model?.username }}
			</template>

			<template #phoneNumber:thead>{{ $tl("phone_number") }}</template>
			<template #phoneNumber="{ model }">
				{{ model?.phoneNumber || "-" }}
			</template>

			<template #dateOfBirth:thead>{{ $tl("date_of_birth") }}</template>
			<template #dateOfBirth="{ model }">
				{{ model?.dateOfBirth ? new Date(model.dateOfBirth).toLocaleDateString() : "-" }}
			</template>

			<template #gender:thead>{{ $tl("gender") }}</template>
			<template #gender="{ model }">
				<q-chip v-if="model?.gender" color="secondary" outline>
					{{ $tl(model.gender) }}
				</q-chip>
				<span v-else>-</span>
			</template>

			<template #edit:thead>
				<div class="text-center">
					{{ $tl("action") }}
				</div>
			</template>
			<template #edit="{ model }">
				<div class="text-center">
					<IconBtn
						:to="{ name: 'USER_VIEW', params: { id: model.id } }"
						icon="visibility"
						tooltipText="view_user"
						withTooltip
					/>
					<IconDialog
						icon="edit"
						:style="'width: 40%;'"
						:fetch="fetch"
						tooltipText="edit_user"
						withTooltip
					>
						<EditUser :id="model.id" :fetch="fetch" />
					</IconDialog>
					<IconDialog
						iconColor="negative"
						icon="delete"
						tooltipText="delete_user"
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
			<!-- Кастомный мобильный вид -->
			<template #card="{ model, orderNumber }">
				<q-item class="q-mb-md user-item-bordered">
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="primary" text-color="white" size="sm">
							{{ orderNumber }}
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label>
							{{ $tl("full_name") }}: {{ model?.lastName }} {{ model?.firstName }}
							{{ model?.middleName }}
						</q-item-label>
						<q-item-label>
							{{ $tl("username") }}: {{ model?.username || "-" }}
						</q-item-label>
						<q-item-label>
							{{ $tl("phone_number") }}: {{ model?.phoneNumber || "-" }}
						</q-item-label>
						<q-item-label v-if="model?.gender">
							{{ $tl("gender") }}:
							<q-chip color="secondary" outline size="sm">
								{{ $tl(model.gender) }}
							</q-chip>
						</q-item-label>
					</q-item-section>
				</q-item>

				<!-- Кнопки действий -->
				<div class="card-actions flex justify-end gap-4">
					<Button
						:to="{ name: 'USER_VIEW', params: { id: model.id } }"
						icon="visibility"
						tooltipText="view_user"
						withTooltip
					/>
					<ButtonDialog
						icon="edit"
						:style="'width: auto;'"
						:fetch="fetch"
						tooltipText="edit_user"
						withTooltip
					>
						<EditUser :id="model.id" :fetch="fetch" />
					</ButtonDialog>
					<ButtonDialog
						iconColor="negative"
						icon="delete"
						tooltipText="delete_user"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="true" />
					</ButtonDialog>
				</div>
			</template>
		</ResponsiveTable>
	</PageLoading>
</template>

<style scoped lang="scss">
.user-item-bordered {
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
