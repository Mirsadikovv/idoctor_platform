<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { OrderService, type OrderType } from "../service";

import AppFooter from "@/components/AppFooter.vue";
import { useAuthStore } from "@/store/auth-store";
import { useAppNavigation } from "@/composables/useAppNavigation";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";

export interface Props {
	id: number | string;
}
const { id } = defineProps<Props>();

const router = useRouter();
const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const orderModel = ref<OrderType>({} as OrderType);

const loadOrder = async () => {
	const data = await OrderService.getByID(+id);
	orderModel.value = data;
};
</script>

<template>
	<PageLoading :find="loadOrder" #="{ loading }">
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
								:label="$tl('order_list')"
								icon="receipt"
								:to="{ name: 'ORDER_PAGE' }"
							/>
						</q-breadcrumbs>
						<q-btn
							flat
							color="white"
							class="bg-secondary"
							:label="$tl('order_parts')"
							:to="{ name: 'ORDER_PART_PAGE', params: { orderId: id } }"
						/>
					</div>

					<q-markup-table separator="cell" flat bordered>
						<tbody>
							<tr>
								<td class="font-bold text-left">{{ $tl("order_id") }}</td>
								<td>#{{ orderModel?.id }}</td>
							</tr>
							<tr v-if="orderModel?.price">
								<td class="font-bold text-left">{{ $tl("price") }}</td>
								<td class="text-positive font-bold">
									{{ orderModel?.price?.toLocaleString() }} сум
								</td>
							</tr>
							<tr v-if="orderModel?.status">
								<td class="font-bold text-left">{{ $tl("status") }}</td>
								<td>
									<q-chip color="primary" text-color="white">
										{{ orderModel?.status }}
									</q-chip>
								</td>
							</tr>
							<tr v-if="orderModel?.payment_status">
								<td class="font-bold text-left">{{ $tl("payment_status") }}</td>
								<td>
									<q-chip
										:color="
											orderModel?.payment_status === 'paid'
												? 'positive'
												: 'warning'
										"
										text-color="white"
									>
										{{ orderModel?.payment_status }}
									</q-chip>
								</td>
							</tr>
							<tr v-if="orderModel?.payment_type">
								<td class="font-bold text-left">{{ $tl("payment_type") }}</td>
								<td>{{ orderModel?.payment_type }}</td>
							</tr>
							<tr v-if="orderModel?.created_at">
								<td class="font-bold text-left">{{ $tl("created_at") }}</td>
								<td>{{ new Date(orderModel?.created_at).toLocaleDateString() }}</td>
							</tr>
							<tr v-if="orderModel?.deleted_at">
								<td class="font-bold text-left">{{ $tl("deleted_at") }}</td>
								<td>{{ new Date(orderModel?.deleted_at).toLocaleDateString() }}</td>
							</tr>
						</tbody>
					</q-markup-table>
				</q-page>
			</q-page-container>

			<AppFooter
				:username="authStore.user?.username"
				:languages="$lang.languages"
				:current-language-id="$lang._currentLang?.id"
				:show-add-button="true"
				:add-button-route="{ name: 'ORDER_CREATE' }"
				add-button-icon="add"
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
