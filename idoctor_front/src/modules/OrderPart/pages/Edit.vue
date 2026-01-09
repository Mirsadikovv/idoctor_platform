<script setup lang="ts">
import { useRouter } from "vue-router";
import { OrderPartService, type OrderPartUpdateType, type OrderPartType } from "@/service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { ref } from "vue";
import Title from "@/components/Title.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import { formRequired, formNumber } from "@/common/validator";
import { searchParts, searchSuppliers } from "../utils";

export interface Props {
	id: number;
}

const { id } = defineProps<Props>();

const orderModel = ref<Partial<OrderPartType>>({});
const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

async function save(model: OrderPartUpdateType) {
	const response = await OrderPartService.update(+id, model);

	if (!response) return false;

	router.push({
		name: "ORDER_PART_PAGE",
	});
	return true;
}

async function getByID() {
	const response = await OrderPartService.getByID(+id);

	if (!response) return;

	orderModel.value = response;
}
</script>

<template>
	<PageLoading :find="getByID" #="{ loading }">
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
								:label="$tl('order_part_list')"
								icon="assignment"
								:to="{ name: 'ORDER_PART_PAGE' }"
							/>
						</q-breadcrumbs>
					</div>
					<Form v-model="orderModel" :save="save">
						<template #title>
							<Title class="mb-5">{{ $tl("edit_order_part") }}</Title>
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

						<template #part_id="{ model }">
							<Autocomplete
								v-model="model.part_id"
								label="part_id"
								class="col-lg-4 col-md-6 col-12"
								:rules="[formRequired()]"
								:find="searchParts"
								option-label="name"
								option-value="id"
							/>
						</template>

						<template #price="{ model }">
							<Input
								v-model.number="model.price"
								label="price"
								class="col-12"
								:rules="[formRequired(), formNumber()]"
							/>
						</template>

						<template #income_price="{ model }">
							<Input
								v-model.number="model.income_price"
								label="income_price"
								class="col-12"
								:rules="[formRequired(), formNumber()]"
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
				:add-button-route="{ name: 'ORDER_PART_CREATE' }"
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
