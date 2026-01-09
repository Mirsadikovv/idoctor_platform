<script setup lang="ts">
import { ref } from "vue";

import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";

import {
	LanguageContentService,
	type LanguageContentPageData,
	type LanguageContentPartialType,
} from "@/service";
import { useLanguageStore } from "@/store/language-store";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";
import IconBtn from "@/components/quasar/btn/IconBtn.vue";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();
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
</script>

<template>
	<PageLoading :find="find" #="{ loading, fetch }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
					<ResponsiveTable :loading="loading" :models="models" :pick="pikers" hasOrder>
						<template #language:thead> </template>
						<template #language="{ model }">
							<q-chip color="secondary" text-color="white">
								{{ model?.language?.name }}
							</q-chip>
						</template>

						<template #key:thead> </template>
						<template #key="{ model }">
							<router-link
								:to="{ name: 'TRANSLATE_UPDATE', params: { id: model.key } }"
								class="text-primary text-decoration-none"
							>
								{{ model.key }}
							</router-link>
						</template>

						<template #value:thead> </template>
						<template #value></template>

						<template #edit="{ model }">
							<div class="text-center">
								<IconBtn
									v-if="$canPage('TRANSLATE_UPDATE')"
									:to="{
										name: 'TRANSLATE_UPDATE',
										params: { id: model.key },
									}"
									icon="edit"
								/>
								<IconBtn
									v-if="$canPage('TRANSLATE_VIEW')"
									:to="{
										name: 'TRANSLATE_VIEW',
										params: { id: model.key },
									}"
									icon="visibility"
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
						<!-- Кастомный мобильный вид для переводов -->
						<template #card="{ model, orderNumber }">
							<q-item
								class="translatedcontent-item-telegram"
								clickable
								:to="{
									name: $canPage('TRANSLATE_UPDATE') ? 'TRANSLATE_UPDATE' : 'TRANSLATE_VIEW',
									params: { id: model.key },
								}"
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
				</q-page>
			</q-page-container>

			<AppFooter
				:username="authStore.user?.username"
				:languages="$lang.languages"
				:current-language-id="$lang._currentLang?.id"
				:show-add-button="true"
				:add-button-route="{ name: 'TRANSLATE_CREATE' }"
				@toggle-drawer="toggleLeftDrawer"
				@go-to-profile="toggleLeftDrawer"
				@set-lang="setLang"
				@logout="logout"
			/>
		</q-layout>
	</PageLoading>
</template>

<style scoped>
@import "@/styles/telegram-app.scss";
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
