<script setup lang="ts">
import { useRouter } from "vue-router";
import { OrderPartService, type OrderPartCreateType } from "@/service";
import Form from "@/components/quasar/form/Form.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import { ref } from "vue";
import Title from "@/components/Title.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";
import { formRequired, formNumber } from "@/common/validator";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { searchParts, searchSuppliers } from "../utils";
interface Props {
	orderId: number;
}

const { orderId } = defineProps<Props>();

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();

const orderModel = ref<Partial<OrderPartCreateType>>({});

async function save(model: OrderPartCreateType) {
	const response = await OrderPartService.create({ ...model, order_id: orderId });

	if (!response) return false;

	router.push({
		name: "ORDER_PART_PAGE",
	});

	return true;
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
							:label="$tl('order_part_list')"
							:to="{ name: 'ORDER_PART_PAGE' }"
						/>
					</q-breadcrumbs>
				</div>
				<Form v-model="orderModel" :save="save">
					<template #title>
						<Title class="mb-5">{{ $tl("create_order") }}</Title>
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
