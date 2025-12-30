<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref } from "vue";
import { LanguageService, type LanguagePageData } from "@/service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import CreateLang from "@module/Language/pages/Create.vue";

const languagePage = ref<LanguagePageData>({
	data: [],
	totalRows: 0,
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
});

const pick = {
	id: false,
	name: true,
	description: true,
	status: true,
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
			<template #name="{ model }">
				<router-link
					:to="{ name: 'LANGUAGE_VIEW', params: { id: model.id } }"
					class="text-primary text-decoration-none"
				>
					{{ model.name }}
				</router-link>
			</template>

			<template #description:thead> </template>
			<template #description="{ model }">
				{{ model.description }}
			</template>

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
				<q-item
					class="language-item-telegram"
					clickable
					:to="{ name: 'LANGUAGE_VIEW', params: { id: model.id } }"
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
							{{ model.description }}
						</q-item-label>
						<!-- Чип статуса -->
						<div class="q-mt-xs">
							<q-chip
								:color="!model.deletedAt ? 'positive' : 'negative'"
								outline
								size="sm"
								dense
							>
								{{ !model.deletedAt ? $tl("active") : $tl("deleted") }}
							</q-chip>
						</div>
					</q-item-section>

					<q-item-section side>
						<q-icon name="chevron_right" color="grey-6" />
					</q-item-section>
				</q-item>
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
.language-item-telegram {
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
