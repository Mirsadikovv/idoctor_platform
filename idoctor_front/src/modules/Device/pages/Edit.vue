<script setup lang="ts">
import { useRouter } from "vue-router";
import { DeviceService, type DeviceUpdateType, type DeviceType } from "../service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import { formRequired } from "@/common/validator";
import { ref } from "vue";

export interface Props {
	id: number;
}

const { id } = defineProps<Props>();

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const deviceModel = ref<DeviceType>({
	id: 0,
	name: "",
	brand_name: "",
	created_at: "",
});

async function getDeviceByID() {
	const data = await DeviceService.getByID(+id);
	if (data) {
		deviceModel.value = data;
	}
}

async function save(model: DeviceUpdateType) {
	const response = await DeviceService.update(+id, model);

	if (!response) return false;

	router.push({
		name: "DEVICE_PAGE",
	});
	return true;
}
</script>

<template>
	<PageLoading :find="getDeviceByID" #="{ loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
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
								:label="$tl('device_list')"
								icon="devices"
								:to="{ name: 'DEVICE_PAGE' }"
							/>
						</q-breadcrumbs>
					</div>
					<Form v-model="deviceModel" :save="save">
						<template #title>
							<Title class="mb-5">{{ $tl("edit_device") }}</Title>
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
				:add-button-route="{ name: 'DEVICE_CREATE' }"
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
</style>
