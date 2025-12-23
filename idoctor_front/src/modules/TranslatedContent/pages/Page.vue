<script setup lang="ts">
import { ref } from "vue";

import PageLoading from "@/components/PageLoading.vue";
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import Form from "@/components/quasar/form/Form.vue";
import Expasion from "@/components/quasar/Expasion.vue";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import Search from "@/components/quasar/search/Search.vue";

import { useRouter } from "vue-router";
import {
	LanguageContentService,
	type LanguageContentPageData,
	type LanguageContentPartialType,
	type LanguagePartialType,
} from "@/service";
import { useLanguageStore } from "@/store/language-store";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import CreateTranslate from "../components/Create.vue";
import UpdateTranslate from "../components/Update.vue";
import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import ConfirmDelete from "../components/ConfirmDelete.vue";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { searchLangs } from "../utils";
import Button from "@/components/quasar/btn/Button.vue";
const router = useRouter();
const langStore = useLanguageStore();

let searchModel = ref<LanguageContentPartialType>({
	languageId: langStore.langID,
	language: langStore.lang!,
});

const models = ref<LanguageContentPageData>({
	data: [],
	totalRows: 0,
	totalPages: 0,
	pageSize: 0,
	currentPage: 0,
});

const pick = {
	id: false,
	edit: false,
	language: true,
	key: true,
	value: true,
};

const pikers = ref({});

async function find(query: string) {
	const response = await LanguageContentService.page(
		query,
		searchModel.value.language?.id || langStore.langID,
	);
	if (!response) return;
	models.value = response;
}

async function save() {
	searchModel.value = {};
	await find("");
	router.replace({ query: {} });
	return true;
}

async function updateLang(item: LanguagePartialType) {
	searchModel.value.language = item;
	await find("");
}
</script>

<template>
	<PageLoading :find="find" #="{ loading, fetch }">
		<div class="flex! gap-x-4 items-center mb-3">
			<q-breadcrumbs>
				<q-breadcrumbs-el :label="$tl('translated_content_page')" icon="article" />
			</q-breadcrumbs>
			<q-space></q-space>

			<ButtonDialog :label="$tl('create')" :style="'width: auto;'">
				<CreateTranslate />
			</ButtonDialog>
		</div>

		<Expasion class="mb-4">
			<Form is-cleaned-style :model-value="searchModel" :save="save" paddingMini>
				<template #language="{ model }">
					<Autocomplete
						label="language"
						v-model="model.language"
						:find="searchLangs"
						@update="
							(q) => {
								updateLang(q as LanguagePartialType);
							}
						"
						class="col-lg-4 col-md-6 col-12"
						option-label="description"
					/>
				</template>
				<template #key="{ model }">
					<Search
						query-name="key"
						label="key"
						v-model="model.key"
						@search="fetch"
						input-debounce="500"
						class="col-lg-4 col-md-6 col-12"
					/>
				</template>

				<template #value="{ model }">
					<Search
						query-name="value"
						label="value"
						v-model="model.value"
						@search="fetch"
						input-debounce="500"
						class="col-lg-4 col-md-6 col-12"
					/>
				</template>

				<template #actions="{ loading }">
					<div class="col-12">
						<Button @click="save()" :loading="loading" class="ml-auto">
							{{ $tl("save") }}
						</Button>
					</div>
				</template>
			</Form>
		</Expasion>

		<ResponsiveTable :loading="loading" :models="models" :pick="pikers" hasOrder>
			<template #language:thead> </template>
			<template #language="{ model }">
				<q-chip color="secondary" text-color="white">{{ model?.language?.name }} </q-chip>
			</template>

			<template #key:thead> </template>
			<template #key></template>

			<template #value:thead> </template>
			<template #value></template>

			<template #edit:thead>
				<div class="text-center">
					{{ $tl("action") }}
				</div>
			</template>
			<template #edit="{ model }">
				<div class="text-center w-100px mx-auto">
					<IconDialog
						:key="model.id"
						icon="edit"
						:style="'width: auto;'"
						tooltipText="update_translate"
						withTooltip
					>
						<UpdateTranslate :keyWord="model.key" :fetch="fetch" />
					</IconDialog>

					<IconDialog
						icon="delete"
						iconColor="negative"
						tooltipText="remove_lang"
						withTooltip
					>
						<ConfirmDelete
							:fetch="fetch"
							:id="model.id"
							:key-value="model.key"
							:isRemove="true"
						/>
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
			<!-- Кастомный мобильный вид для переводов -->
			<template #card="{ model, orderNumber }">
				<q-item class="q-mb-md translatedcontent-item-bordered">
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="purple" text-color="white" size="sm">
							<q-icon name="translate" />
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label>
							<strong>№ {{ orderNumber }}</strong>
						</q-item-label>

						<q-item-label>
							<strong>Язык:</strong>
							{{ model?.language?.name }}
						</q-item-label>
						<q-item-label><strong>Ключ:</strong> {{ model.key }}</q-item-label>
						<q-item-label class="text-caption">
							<strong>Значение:</strong> {{ model.value || "-" }}
						</q-item-label>
					</q-item-section>
				</q-item>

				<!-- Кнопки действий -->
				<div class="card-actions flex justify-end gap-4">
					<ButtonDialog
						:key="model.id"
						icon="edit"
						:style="'width: auto;'"
						tooltipText="update_translate"
						withTooltip
						flat
						round
						color="primary"
					>
						<UpdateTranslate :keyWord="model.key" :fetch="fetch" />
					</ButtonDialog>
					<ButtonDialog
						icon="delete"
						iconColor="negative"
						tooltipText="remove_lang"
						withTooltip
						flat
						round
					>
						<ConfirmDelete
							:fetch="fetch"
							:id="model.id"
							:key-value="model.key"
							:isRemove="true"
						/>
					</ButtonDialog>
				</div>
			</template>
		</ResponsiveTable>
	</PageLoading>
</template>

<style scoped lang="scss">
.translatedcontent-item-bordered {
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
