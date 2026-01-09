<script setup lang="ts">
import { useRouter } from "vue-router";
import { LanguageService, type LanguageType } from "@/service";
import { ref } from "vue";
import { useQuasar } from "quasar";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import AppFooter from "@/components/AppFooter.vue";
import EditLang from "./Edit.vue";

export interface Props {
	id: number;
}

const router = useRouter();
const $q = useQuasar();
const { id } = defineProps<Props>();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();

const model = ref<LanguageType>({} as LanguageType);

async function fetchLanguage() {
	const response = await LanguageService.getByID(+id);
	if (!response) return;
	model.value = response;
}

async function deleteLanguage() {
	const response = await LanguageService.delete(+id);
	if (!response) return false;
	
	router.push({ name: "LANGUAGE_PAGE" });
	return true;
}

async function restoreLanguage() {
	const response = await LanguageService.restore(+id);
	if (!response) return false;
	
	await fetchLanguage();
	return true;
}

function confirmDelete() {
	$q.dialog({
		title: 'Подтверждение',
		message: 'Вы уверены, что хотите удалить этот язык?',
		cancel: true,
		persistent: true
	}).onOk(() => {
		deleteLanguage();
	});
}

function confirmRestore() {
	$q.dialog({
		title: 'Подтверждение',
		message: 'Вы уверены, что хотите восстановить этот язык?',
		cancel: true,
		persistent: true
	}).onOk(() => {
		restoreLanguage();
	});
}

if (+id > 0) {
	fetchLanguage();
}
</script>

<template>
	<q-layout view="hHh Lpr lff">
		<q-page-container>
			<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
				<div class="flex! gap-x-4 items-center mb-3">
					<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
					<q-breadcrumbs>
						<q-breadcrumbs-el
							:label="$tl('language_list')"
							icon="g_translate"
							:to="{ name: 'LANGUAGE_PAGE' }"
						/>
						<q-breadcrumbs-el :label="model.name || $tl('language_view_title')" />
					</q-breadcrumbs>
					<q-space />
					
					<div class="flex gap-2" v-if="model.id">
						<ButtonDialog
							v-if="!model.deletedAt"
							:label="$tl('edit_language')"
							icon="edit"
							color="primary"
							:fetch="fetchLanguage"
						>
							<EditLang :id="model.id" :fetch="fetchLanguage" />
						</ButtonDialog>
						
						<q-btn
							v-if="model.deletedAt"
							icon="sync"
							color="positive"
							flat
							round
							@click="confirmRestore"
						>
							<q-tooltip>{{ $tl('restore_language') }}</q-tooltip>
						</q-btn>
						
						<q-btn
							v-if="!model.deletedAt"
							icon="delete"
							color="negative"
							flat
							round
							@click="confirmDelete"
						>
							<q-tooltip>{{ $tl('remove_language') }}</q-tooltip>
						</q-btn>
					</div>
				</div>

				<div v-if="model.id" class="q-pa-md">
					<q-card class="q-mb-lg">
						<q-card-section>
							<div class="text-h6 q-mb-md flex items-center">
								<q-icon name="g_translate" class="q-mr-sm" color="blue" />
								{{ $tl('language_details') }}
							</div>
							
							<q-list>
								<q-item>
									<q-item-section>
										<q-item-label overline>{{ $tl('language_name') }}</q-item-label>
										<q-item-label class="text-h6">{{ model.name }}</q-item-label>
									</q-item-section>
								</q-item>
								
								<q-item v-if="model.description">
									<q-item-section>
										<q-item-label overline>{{ $tl('description') }}</q-item-label>
										<q-item-label>{{ model.description }}</q-item-label>
									</q-item-section>
								</q-item>
								
								<q-item>
									<q-item-section>
										<q-item-label overline>{{ $tl('status') }}</q-item-label>
										<q-chip :color="!model.deletedAt ? 'positive' : 'negative'" text-color="white">
											{{ !model.deletedAt ? $tl("active") : $tl("deleted") }}
										</q-chip>
									</q-item-section>
								</q-item>
							</q-list>
						</q-card-section>
					</q-card>
				</div>
				
				<div v-else class="flex justify-center q-pa-xl">
					<q-spinner color="primary" size="3em" />
				</div>
			</q-page>
		</q-page-container>

		<AppFooter
			:username="authStore.user?.username"
			:languages="$lang.languages"
			:current-language-id="$lang._currentLang?.id"
			:show-add-button="true"
			:add-button-route="{ name: 'LANGUAGE_CREATE' }"
			add-button-icon="add"
			@toggle-drawer="toggleLeftDrawer"
			@go-to-profile="toggleLeftDrawer"
			@set-lang="setLang"
			@logout="logout"
		/>
	</q-layout>
</template>