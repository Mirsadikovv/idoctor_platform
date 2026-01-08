<script setup lang="ts">
import { useRouter } from "vue-router";
import { OrderService, type OrderUpdateType, type OrderType } from "../service";
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
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";
import { formRequired, formNumber } from "@/common/validator";
import {
	searchClients,
	orderStatusOptions,
	paymentStatusOptions,
	paymentTypeOptions,
} from "../utils";
import DatePicker from "@/components/quasar/form/DatePicker.vue";
import { parseDate } from "@/common";

export interface Props {
	id: number;
}

const { id } = defineProps<Props>();

const orderModel = ref<Partial<OrderType>>({});
const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();

async function save(model: Partial<OrderUpdateType>) {
	model.deadline = parseDate(model.deadline, "YYYY-MM-DD")?.toISOString();

	const response = await OrderService.update(+id, model);

	if (!response) return false;

	router.push({
		name: "ORDER_PAGE",
	});
	return true;
}

async function getByID() {
	const response = await OrderService.getByID(+id);

	if (!response) return;

	orderModel.value = response;
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
								:label="$tl('order_list')"
								icon="assignment"
								:to="{ name: 'ORDER_PAGE' }"
							/>
						</q-breadcrumbs>
					</div>
					<Form v-model="orderModel" :save="save">
						<template #title>
							<Title class="mb-5 justify-between">
								<div>
									{{ $tl("edit_order") }}
								</div>

								<q-btn
									flat
									color="white"
									class="bg-secondary"
									:label="$tl('order_parts')"
									:to="{ name: 'ORDER_PART_PAGE', params: { orderId: id } }"
								/>
							</Title>
						</template>

						<template #client_name="{ model }">
							<Input
								v-model="model.client_name"
								label="client_name"
								class="col-lg-6 col-md-6 col-12"
							/>
						</template>

						<template #client_phone="{ model }">
							<Input
								v-model="model.client_phone"
								label="client_phone"
								class="col-lg-6 col-md-6 col-12"
								:rules="[formRequired()]"
							/>
						</template>

						<template #deadline="{ model }">
							<DatePicker
								v-model="model.deadline"
								:rules="[formRequired($tl('this_field_is_required'))]"
								label="Deadline"
								class="col-lg-6 col-md-6 col-12"
							/>
						</template>

						<template #price="{ model }">
							<Input
								v-model.number="model.price"
								label="Price"
								class="col-lg-6 col-md-6 col-12"
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
								class="col-lg-6 col-md-6 col-12"
								:rules="[formRequired()]"
								:option-label="(v) => $tl(v.label)"
								option-value="value"
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
								label="PaymentStatus"
								class="col-lg-6 col-md-6 col-12"
								:rules="[formRequired()]"
								:option-label="(v) => $tl(v.label)"
								option-value="value"
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
								label="PaymentType"
								class="col-lg-6 col-md-6 col-12"
								:rules="[formRequired()]"
								:option-label="(v) => $tl(v.label)"
								option-value="value"
							/>
						</template>

						<template #client_id="{ model }">
							<Autocomplete
								v-model="model.client_id"
								label="Client"
								class="col-lg-6 col-md-6 col-12"
								:find="searchClients"
								option-label="username"
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
				:add-button-route="{ name: 'ORDER_CREATE' }"
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
