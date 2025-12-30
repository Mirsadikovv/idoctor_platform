<script setup lang="ts">
import { useRouter } from "vue-router";
import { OrderService, type OrderCreateType } from "../service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { ref } from "vue";
import Title from "@/components/Title.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";
import { formRequired, formNumber } from "@/common/validator";
import {
	searchClients,
	searchMasters,
	searchProblems,
	searchParts,
	orderStatusOptions,
	paymentStatusOptions,
	paymentTypeOptions,
} from "../utils";

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const orderModel = ref<Partial<OrderCreateType>>({
	part_ids: [],
	problem_ids: [],
});

async function save(model: OrderCreateType) {
	const response = await OrderService.create(model);

	if (!response) return false;

	router.push({
		name: "ORDER_PAGE",
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
				class="bg-white text-gray-900 overflow-auto p-4 pt-20"
			>
				<div class="flex! gap-x-4 items-center mb-3">
					<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
					<q-breadcrumbs>
						<q-breadcrumbs-el
							:label="$tl('order_list')"
							icon="assignment"
							:to="{ name: 'ORDER_PAGE' }"
						/>
						<q-breadcrumbs-el :label="$tl('page_for_create')" />
					</q-breadcrumbs>
				</div>
				<Form v-model="orderModel" :save="save">
					<template #title>
						<Title class="mb-5">{{ $tl("create_order") }}</Title>
					</template>

					<template #client_id="{ model }">
						<Autocomplete
							v-model="model.client_id"
							label="Client"
							class="col-12"
							:find="searchClients"
							option-label="username"
							option-value="id"
						/>
					</template>

					<template #master_id="{ model }">
						<Autocomplete
							v-model="model.master_id"
							label="Master"
							class="col-12"
							:find="searchMasters"
							option-label="username"
							option-value="id"
						/>
					</template>

					<template #problem_ids="{ model }">
						<Autocomplete
							v-model="model.problem_ids"
							label="Problems"
							class="col-12"
							:find="searchProblems"
							option-label="name"
							option-value="id"
							multiple
							use-chips
						/>
					</template>

					<template #part_ids="{ model }">
						<Autocomplete
							v-model="model.part_ids"
							label="Parts"
							class="col-12"
							:find="searchParts"
							option-label="name"
							option-value="id"
							multiple
							use-chips
						/>
					</template>

					<template #price="{ model }">
						<Input
							v-model.number="model.price"
							label="Price"
							class="col-12"
							:rules="[formRequired(), formNumber()]"
						/>
					</template>

					<template #status="{ model }">
						<Autocomplete
							v-model="model.status"
							:find="
								async () => {
									return orderStatusOptions;
								}
							"
							label="Status"
							class="col-12"
							option-label="label"
							option-value="value"
							:rules="[formRequired()]"
						/>
					</template>

					<template #payment_status="{ model }">
						<Autocomplete
							v-model="model.payment_status"
							:find="
								async () => {
									return paymentStatusOptions;
								}
							"
							label="Payment Status"
							option-label="label"
							option-value="value"
							class="col-12"
							:rules="[formRequired()]"
						/>
					</template>

					<template #payment_type="{ model }">
						<Autocomplete
							v-model="model.payment_type"
							:find="
								async () => {
									return paymentTypeOptions;
								}
							"
							label="Payment Type"
							class="col-12"
							option-label="label"
							option-value="value"
							:rules="[formRequired()]"
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
