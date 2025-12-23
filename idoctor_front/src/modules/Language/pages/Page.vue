<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref } from "vue";
import { LanguageService, type LanguagePageData } from "@/service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import CreateLang from "@module/Language/pages/Create.vue";
import EditLang from "@module/Language/pages/Edit.vue";

import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import ConfirmDialog from "../components/ConfirmDialog.vue";

const languagePage = ref<LanguagePageData>({
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
	description: true,
	deletedAt: true,
};

const pikers = ref({});

async function page(query: string = "") {
	const response = await LanguageService.pageWithDelete(query);

	if (!response) return;

	languagePage.value = response;
}
</script>

<template>
	<PageLoading :find="page" #="{ fetch }">
		<div class="flex! height-full gap-x-4 items-center mb-3">
			<q-breadcrumbs>
				<q-breadcrumbs-el :label="$tl('language_list')" icon="article" />
			</q-breadcrumbs>
			<q-space></q-space>

			<ButtonDialog label="create" :style="'width: 100%;'" :fetch="fetch">
				<CreateLang :fetch="fetch" />
			</ButtonDialog>
		</div>

		<ResponsiveTable :models="languagePage" hasOrder>
			<template #name:thead> </template>
			<template #name></template>

			<template #description:thead> </template>
			<template #description></template>

			<template #status:thead> </template>
			<template #status="{ model }">
				<q-chip :color="!model.deletedAt ? 'positive' : 'negative'" text-color="white">
					{{ !model.deletedAt ? $tl("ACTIVE") : $tl("DELETED") }}
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
						v-if="!model.deletedAt"
						icon="edit"
						:style="'width: 40%;'"
						:fetch="fetch"
						tooltipText="edit_lang"
						withTooltip
					>
						<EditLang :id="model.id" :fetch="fetch" />
					</IconDialog>

					<IconDialog
						v-if="model.deletedAt"
						icon="sync"
						iconColor="positive"
						tooltipText="restore_lang"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="false" />
					</IconDialog>
					<IconDialog
						v-if="!model.deletedAt"
						icon="delete"
						iconColor="negative"
						tooltipText="remove_lang"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="true" />
					</IconDialog>
				</div>
			</template>

			<!-- Кастомный мобильный вид для языков -->
			<template #card="{ model, orderNumber }">
				<q-item class="q-mb-md language-item-bordered">
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="primary" text-color="white" size="sm">
							{{ orderNumber }}
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label>{{ $tl("NAME") }}: {{ model.name }} </q-item-label>
						<q-item-label>
							{{ $tl("DESCRIPTION") }}: {{ model.description }}
						</q-item-label>
						<q-item-label>
							{{ $tl("STATUS") }}:
							<q-chip
								:color="!model.deletedAt ? 'positive' : 'negative'"
								text-color="white"
								size="sm"
							>
								{{ !model.deletedAt ? $tl("ACTIVE") : $tl("DELETED") }}
							</q-chip>
						</q-item-label>
					</q-item-section>
				</q-item>

				<div class="card-actions flex justify-end gap-4">
					<ButtonDialog
						v-if="!model.deletedAt"
						icon="edit"
						color="primary"
						size="sm"
						round
						flat
						:fetch="fetch"
						tooltipText="edit_lang"
						withTooltip
					>
						<EditLang :id="model.id" :fetch="fetch" />
					</ButtonDialog>

					<ButtonDialog
						v-if="model.deletedAt"
						icon="sync"
						iconColor="positive"
						color="positive"
						size="sm"
						round
						flat
						:fetch="fetch"
						tooltipText="restore_lang"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="false" />
					</ButtonDialog>

					<ButtonDialog
						v-if="!model.deletedAt"
						icon="delete"
						iconColor="negative"
						color="negative"
						size="sm"
						round
						flat
						:fetch="fetch"
						tooltipText="remove_lang"
						withTooltip
					>
						<ConfirmDialog :fetch="fetch" :id="model.id" :isRemove="true" />
					</ButtonDialog>
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
		</ResponsiveTable>
	</PageLoading>
</template>

<style scoped lang="scss">
.language-item-bordered {
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
