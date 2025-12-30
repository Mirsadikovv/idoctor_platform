<script setup lang="ts">
import { useRouter } from "vue-router";
import { DeviceService, type DeviceCreateType } from "../service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import { formRequired } from "@/common/validator";
import { ref } from "vue";

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const deviceModel = ref<DeviceCreateType>({
	name: "",
	brand_name: ""
});

async function save(model: DeviceCreateType) {
	const response = await DeviceService.create(model);

	if (!response) return false;

	router.push({
		name: "DEVICE_PAGE",
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
				class="bg-white text-gray-900 overflow-auto p-4 pt-16"
			>
				<div class="flex! gap-x-4 items-center mb-3">
					<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
					<q-breadcrumbs>
						<q-breadcrumbs-el
							:label="$tl('device_list')"
							icon="devices"
							:to="{ name: 'DEVICE_PAGE' }"
						/>
						<q-breadcrumbs-el :label="$tl('page_for_create')" />
					</q-breadcrumbs>
				</div>
				<Form v-model="deviceModel" :save="save">
					<template #title>
						<Title class="mb-5">{{ $tl("create_device") }}</Title>
					</template>

					<template #name="{ model }">
						<Input 
							v-model="model.name" 
							label="Device Name" 
							class="col-lg-6 col-md-6 col-12"
							:rules="[formRequired()]"
						/>
					</template>
					<template #brand_name="{ model }">
						<Input 
							v-model="model.brand_name" 
							label="Brand Name" 
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