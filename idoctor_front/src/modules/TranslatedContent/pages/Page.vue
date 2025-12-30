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
			<template #key="{ model }">
				<router-link
					:to="{ name: 'TRANSLATED_CONTENT_VIEW', params: { id: model.key } }"
					class="text-primary text-decoration-none"
				>
					{{ model.key }}
				</router-link>
			</template>

			<template #value:thead> </template>
			<template #value></template>

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
				<q-item
					class="translatedcontent-item-telegram"
					clickable
					:to="{ name: 'TRANSLATED_CONTENT_VIEW', params: { id: model.key } }"
				>
					<q-item-section avatar v-if="orderNumber">
						<q-avatar color="primary" text-color="white" size="md">
							{{ orderNumber }}
						</q-avatar>
					</q-item-section>

					<q-item-section>
						<q-item-label class="text-weight-bold text-h6">
							{{ model.key }}
						</q-item-label>
						<q-item-label caption class="text-body2">
							{{ $tl("language") }}: {{ model?.language?.name }}
						</q-item-label>
						<q-item-label caption class="text-body2">
							{{ model.value || "-" }}
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
.translatedcontent-item-telegram {
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
