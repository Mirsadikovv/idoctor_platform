<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import { ref } from "vue";
import { LanguageService, type LanguagePageData } from "@/service";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import EditLang from "@module/Language/pages/Edit.vue";
import ConfirmDialog from "@module/Language/components/ConfirmDialog.vue";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

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
	<PageLoading :find="page" #="{ fetch, loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page
					:style="{
						height: 'calc(var(--app-height, 100vh) - 150px)',
					}"
					class="bg-white text-gray-900 overflow-auto p-4 pt-20"
				>
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
							<q-chip
								:color="!model.deletedAt ? 'positive' : 'negative'"
								text-color="white"
							>
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
									<ConfirmDialog
										:fetch="fetch"
										:id="model.id"
										:isRemove="false"
									/>
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
								:to="{ name: 'LANGUAGE_EDIT', params: { id: model.id } }"
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
				</q-page>
			</q-page-container>

			<AppFooter
				:username="authStore.user?.username"
				:languages="$lang.languages"
				:current-language-id="$lang._currentLang?.id"
				:show-add-button="true"
				:add-button-route="{ name: 'LANGUAGE_CREATE' }"
				add-button-icon="add_circle"
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
