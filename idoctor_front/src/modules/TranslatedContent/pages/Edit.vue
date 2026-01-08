<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref } from "vue";
import { formRequired } from "@/common";
import { initalBatch } from "../utils";
import { LanguageContentService, type LangContentCreateOrUpdate } from "@/service";

import Title from "@/components/Title.vue";
import Form from "@/components/quasar/form/Form.vue";
import Input from "@/components/quasar/form/Input.vue";
import Button from "@/components/quasar/btn/Button.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";
import { useLanguageStore } from "@/store/language-store";

export interface Props {
	id: string;
}

const { id } = defineProps<Props>();

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();
const langStore = useLanguageStore();

const model = ref<Partial<LangContentCreateOrUpdate>>({});

async function save(newModel: LangContentCreateOrUpdate) {
	const response = await LanguageContentService.createOrUpdate({
		...newModel,
		contents: newModel.contents.map((item) => ({ ...item, key: newModel.key })),
	});

	if (!response) return false;

	router.push({
		name: "TRANSLATED_CONTENT",
	});
	return true;
}

async function getByKey() {
	// Сначала инициализируем базовую структуру
	const initialData = await initalBatch();
	model.value = initialData;

	// Затем загружаем данные по ключу
	const response = await LanguageContentService.getByKey(id);
	if (!response) return;

	model.value = {
		category: response.category,
		key: response.key,
		contents: langStore._languages
			.sort((a, b) => a.id - b.id)
			.map((item) => ({
				category: response.category,
				langName: item.name,
				LangDescp: item.description,
				languageId:
					response.contents.find((i) => i.languageId === item.id)?.languageId || item.id,
				value: response.contents.find((i) => i.languageId === item.id)?.value || "",
				key: response.key,
			})),
	};
}
</script>

<template>
	<PageLoading :find="getByKey" #="{ loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
					<div class="flex! gap-x-4 items-center mb-3">
						<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
						<q-breadcrumbs>
							<q-breadcrumbs-el
								:label="$tl('translated_content_page')"
								icon="article"
								:to="{ name: 'TRANSLATED_CONTENT' }"
							/>
						</q-breadcrumbs>
					</div>
					<Form :model-value="model" :save="save">
						<template #title>
							<Title class="mb-5">{{ $tl("update_translate") }}</Title>
						</template>

						<template #key="{ model }">
							<Input
								v-model="model.key"
								label="key"
								class="col-lg-6 col-md-6 col-12"
								:rules="[formRequired($tl('this_field_is_required'))]"
								:readonly="!!model.key"
							/>
						</template>

						<template #separator>
							<div class="col-12">
								<q-separator class="mx--4! my-4px!" />
							</div>
						</template>

						<template #values>
							<div
								v-if="model?.contents"
								class="col-lg-4 col-md-6 col-12"
								v-for="item in model.contents"
								:key="item.languageId"
							>
								<div class="mb-2 text-base font-medium">
									{{ item.langName }} | {{ item.LangDescp }}
								</div>
								<Input
									v-model="item.value"
									label="value"
									:rules="[formRequired($tl('this_field_is_required'))]"
								/>
							</div>
						</template>

						<template #actions="{ loading }">
							<Button :loading="loading" type="submit" class="ml-auto">
								{{ $tl("save") }}
							</Button>
						</template>
					</Form>
				</q-page>
			</q-page-container>

			<AppFooter
				:username="authStore.user?.username"
				:languages="$lang.languages"
				:current-language-id="$lang._currentLang?.id"
				:show-add-button="true"
				:add-button-route="{ name: 'TRANSLATED_CONTENT_CREATE' }"
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
</style>
