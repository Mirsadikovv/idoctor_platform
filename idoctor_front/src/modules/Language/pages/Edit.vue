<script setup lang="ts">
import { useRouter } from "vue-router";
import { LanguageService, type LanguagePartialType, type LanguageType } from "@/service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import { ref } from "vue";
import Title from "@/components/Title.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";

export interface Props {
	id: number;
}

const { id } = defineProps<Props>();

const LanguageModel = ref<LanguagePartialType>({});
const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();

async function save(model: LanguageType) {
	const response = await LanguageService.update(+id, model);

	if (!response) return false;

	router.push({
		name: "LANGUAGE_PAGE",
	});
	return true;
}

async function getByID() {
	const response = await LanguageService.getByID(+id);

	if (!response) return;

	LanguageModel.value = response;
}
</script>

<template>
	<PageLoading :find="getByID" #="{ loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
					<div class="flex! gap-x-4 items-center mb-3">
						<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
						<q-breadcrumbs>
							<q-breadcrumbs-el
								:label="$tl('language_list')"
								icon="article"
								:to="{ name: 'LANGUAGE_PAGE' }"
							/>
						</q-breadcrumbs>
					</div>
					<Form v-model="LanguageModel" :save="save">
						<template #title>
							<Title class="mb-5">{{ $tl("edit_lang") }}</Title>
						</template>

						<template #name="{ model }">
							<Input
								v-model="model.name"
								label="name"
								class="col-lg-6 col-md-6 col-12"
							/>
						</template>
						<template #description="{ model }">
							<Input
								v-model="model.description"
								label="description"
								class="col-lg-6 col-md-6 col-12"
							/>
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
				:add-button-route="{ name: 'LANGUAGE_CREATE' }"
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
