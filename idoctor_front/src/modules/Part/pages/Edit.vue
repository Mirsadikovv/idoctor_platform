<script setup lang="ts">
import { useRouter } from "vue-router";
import { PartService, type PartUpdateType, type PartType } from "../service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Title from "@/components/Title.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";
import { formRequired } from "@/common/validator";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { searchDevices, searchSuppliers } from "../utils";
import { ref } from "vue";

export interface Props {
	id: number;
}

const { id } = defineProps<Props>();

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();

const partModel = ref<Partial<PartType>>({});

async function getPartByID() {
	const partData = await PartService.getByID(+id);
	if (partData) {
		partModel.value = partData;
	}
}

async function save(model: PartUpdateType) {
	const response = await PartService.update(+id, model);

	if (!response) return false;

	router.push({
		name: "PART_PAGE",
	});
	return true;
}
</script>

<template>
	<PageLoading :find="getPartByID" #="{ loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
					<div class="flex! gap-x-4 items-center mb-3">
						<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
						<q-breadcrumbs>
							<q-breadcrumbs-el
								:label="$tl('part_list')"
								icon="build"
								:to="{ name: 'PART_PAGE' }"
							/>
						</q-breadcrumbs>
					</div>
					<Form v-model="partModel" :save="save">
						<template #title>
							<Title class="mb-5">{{ $tl("edit_part") }}</Title>
						</template>

						<template #name="{ model }">
							<Input
								v-model="model.name"
								label="Part Name"
								class="col-lg-4 col-md-6 col-12"
								:rules="[formRequired()]"
							/>
						</template>

						<template #device_id="{ model }">
							<Autocomplete
								v-model="model.device_id"
								label="Device"
								class="col-lg-4 col-md-6 col-12"
								:rules="[formRequired()]"
								:find="searchDevices"
								:option-label="(val) => val.name + ' | ' + val.brand_name"
								option-value="id"
							/>
						</template>

						<template #supplier_id="{ model }">
							<Autocomplete
								v-model="model.supplier_id"
								label="Supplier"
								class="col-lg-4 col-md-6 col-12"
								:rules="[formRequired()]"
								:find="searchSuppliers"
								option-label="name"
								option-value="id"
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
				:add-button-route="{ name: 'PART_CREATE' }"
				add-button-icon="build"
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
