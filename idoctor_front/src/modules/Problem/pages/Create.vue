<script setup lang="ts">
import { useRouter } from "vue-router";
import { ProblemService, type ProblemCreateType } from "../service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import { formNumber, formRequired } from "@/common/validator";
import { ref } from "vue";

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const problemModel = ref<Partial<ProblemCreateType>>({});

async function save(model: ProblemCreateType) {
	const response = await ProblemService.create(model);

	if (!response) return false;

	router.push({
		name: "PROBLEM_PAGE",
	});

	return true;
}
</script>

<template>
	<q-layout view="hHh Lpr lff">
		<q-page-container>
			<q-page
				:style="{
					height: 'calc(var(--app-height, 100vh) - 150px)',
				}"
				class="bg-white text-gray-900 overflow-auto p-4 pt-24"
			>
				<div class="flex! gap-x-4 items-center mb-3">
					<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
					<q-breadcrumbs>
						<q-breadcrumbs-el
							:label="$tl('problem_list')"
							icon="medical_services"
							:to="{ name: 'PROBLEM_PAGE' }"
						/>
						<q-breadcrumbs-el :label="$tl('page_for_create')" />
					</q-breadcrumbs>
				</div>
				<Form v-model="problemModel" :save="save">
					<template #title>
						<Title class="mb-5">{{ $tl("create_problem") }}</Title>
					</template>

					<template #name="{ model }">
						<Input
							v-model="model.name"
							label="Problem Name"
							class="col-lg-6 col-md-6 col-12"
							:rules="[formRequired()]"
						/>
					</template>

					<template #price="{ model }">
						<Input
							v-model.number="model.price"
							:rules="[formRequired(), formNumber()]"
							label="Price"
							:min="0"
							class="col-lg-6 col-md-6 col-12"
						/>
					</template>

					<template #actions="{ loading }">
						<div class="row my-row justify-center">
							<q-space />
							<div class="col-auto max-w-200px! w-100%!">
								<Button :loading="loading" type="submit" class="w-100%!">
									{{ $tl("save") }}
								</Button>
							</div>
						</div>
					</template>
				</Form>
			</q-page>
		</q-page-container>

		<AppFooter
			:username="authStore.user?.username"
			:languages="$lang.languages"
			:current-language-id="$lang._currentLang?.id"
			:show-add-button="false"
			@toggle-drawer="toggleLeftDrawer"
			@go-to-profile="toggleLeftDrawer"
			@set-lang="setLang"
			@logout="logout"
		/>
	</q-layout>
</template>

<style scoped>
@import "@/styles/telegram-app.scss";
</style>
