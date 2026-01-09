<script setup lang="ts">
import { ref } from "vue";
import { StatisticsService } from "../../service/StatisticsService";
import { OrderStatus, PaymentStatus, PaymentType } from "../../service/types";
import type { OrderStatistics } from "../../service/types";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useBreakpoints } from "@/composables/useBreakpoints";
import { useAuthStore } from "@/store/auth-store";
import AppFooter from "@/components/AppFooter.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();
const { isMobile } = useBreakpoints();

const data = ref<OrderStatistics | null>(null);

const startDate = ref();
const endDate = ref();

const masterId = ref<number | undefined>(undefined);
const status = ref<OrderStatus | undefined>(undefined);
const paymentStatus = ref<PaymentStatus | undefined>(undefined);
const paymentType = ref<PaymentType | undefined>(undefined);

const statusOptions = [
	{ label: "В ожидании", value: OrderStatus.PENDING },
	{ label: "В работе", value: OrderStatus.IN_PROGRESS },
	{ label: "Завершен", value: OrderStatus.COMPLETED },
	{ label: "Отменен", value: OrderStatus.CANCELLED },
];

const paymentStatusOptions = [
	{ label: "Оплачен", value: PaymentStatus.PAID },
	{ label: "Не оплачен", value: PaymentStatus.UNPAID },
	{ label: "Частично оплачен", value: PaymentStatus.PARTIAL },
];

const paymentTypeOptions = [
	{ label: "Наличные", value: PaymentType.CASH },
	{ label: "Карта", value: PaymentType.CARD },
	{ label: "Перевод", value: PaymentType.TRANSFER },
	{ label: "Онлайн", value: PaymentType.ONLINE },
];

async function page(_query: string = "") {
	if (!startDate.value || !endDate.value) {
		return;
	}

	if (startDate.value > endDate.value) {
		(await import("@/common/Notify")).ErrorNotify(
			"Дата начала не может быть больше даты окончания",
		);
		return;
	}

	const result = await StatisticsService.getOrderStatistics({
		start_date: startDate.value,
		end_date: endDate.value,
		master_id: masterId.value,
		status: status.value,
		payment_status: paymentStatus.value,
		payment_type: paymentType.value,
	});

	if (result) {
		data.value = result;
	}
}
</script>

<template>
	<PageLoading :find="page" #="{ fetch, loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
					<div class="text-h5 text-weight-medium q-mb-md">Статистика заказов</div>

					<q-card flat bordered class="q-mb-md bg-white shadow-2">
						<q-card-section>
							<div class="text-subtitle2 q-mb-sm">Фильтры</div>
							<div class="row q-col-gutter-md">
								<div class="col-12 col-sm-6 col-md-4">
									<q-input
										v-model="startDate"
										type="date"
										label="Дата начала"
										outlined
										dense
									/>
								</div>
								<div class="col-12 col-sm-6 col-md-4">
									<q-input
										v-model="endDate"
										type="date"
										label="Дата окончания"
										outlined
										dense
									/>
								</div>
								<div class="col-12 col-sm-6 col-md-4">
									<q-input
										v-model.number="masterId"
										type="number"
										label="ID мастера"
										outlined
										dense
										clearable
									/>
								</div>
								<div class="col-12 col-sm-6 col-md-4">
									<q-select
										v-model="status"
										:options="statusOptions"
										label="Статус заказа"
										outlined
										dense
										clearable
										emit-value
										map-options
									/>
								</div>
								<div class="col-12 col-sm-6 col-md-4">
									<q-select
										v-model="paymentStatus"
										:options="paymentStatusOptions"
										label="Статус оплаты"
										outlined
										dense
										clearable
										emit-value
										map-options
									/>
								</div>
								<div class="col-12 col-sm-6 col-md-4">
									<q-select
										v-model="paymentType"
										:options="paymentTypeOptions"
										label="Тип оплаты"
										outlined
										dense
										clearable
										emit-value
										map-options
									/>
								</div>
								<div class="col-12">
									<q-btn
										color="primary"
										label="Применить"
										icon="search"
										@click="fetch"
										:loading="loading"
										class="full-width"
									/>
								</div>
							</div>
						</q-card-section>
					</q-card>

					<q-card flat bordered class="bg-white shadow-2" v-if="!data">
						<q-card-section class="text-center text-grey-7">
							Нет данных для отображения
						</q-card-section>
					</q-card>

					<div v-else>
						<!-- Статусы заказов -->
						<div class="text-subtitle1 text-weight-medium q-mb-sm text-grey-8">
							Статусы заказов
						</div>
						<div class="row q-col-gutter-md q-mb-md">
							<div class="col-12 col-sm-6 col-md-3">
								<q-card flat bordered class="bg-blue-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="shopping_cart"
											:size="isMobile ? '32px' : '40px'"
											color="blue-7"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-blue-9"
											>
												{{ data.total_orders }}
											</div>
											<div class="text-caption text-grey-8">
												Всего заказов
											</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
							<div class="col-12 col-sm-6 col-md-3">
								<q-card flat bordered class="bg-green-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="check_circle"
											:size="isMobile ? '32px' : '40px'"
											color="green-7"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-green-9"
											>
												{{ data.completed_orders }}
											</div>
											<div class="text-caption text-grey-8">Завершено</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
							<div class="col-12 col-sm-6 col-md-3">
								<q-card flat bordered class="bg-orange-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="pending"
											:size="isMobile ? '32px' : '40px'"
											color="orange-7"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-orange-9"
											>
												{{ data.in_progress_orders }}
											</div>
											<div class="text-caption text-grey-8">В работе</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
							<div class="col-12 col-sm-6 col-md-3">
								<q-card flat bordered class="bg-red-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="cancel"
											:size="isMobile ? '32px' : '40px'"
											color="red-7"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-red-9"
											>
												{{ data.cancelled_orders }}
											</div>
											<div class="text-caption text-grey-8">Отменено</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
						</div>

						<!-- Финансы -->
						<div class="text-subtitle1 text-weight-medium q-mb-sm text-grey-8">
							Финансовые показатели
						</div>
						<div class="row q-col-gutter-md q-mb-md">
							<div class="col-12 col-sm-6 col-md-4">
								<q-card flat bordered class="bg-teal-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="trending_up"
											:size="isMobile ? '32px' : '40px'"
											color="teal-7"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-teal-9"
											>
												{{ data.total_revenue.toLocaleString() }}
											</div>
											<div class="text-caption text-grey-8">Доход</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
							<div class="col-12 col-sm-6 col-md-4">
								<q-card flat bordered class="bg-purple-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="trending_down"
											:size="isMobile ? '32px' : '40px'"
											color="purple-7"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-purple-9"
											>
												{{ data.total_cost.toLocaleString() }}
											</div>
											<div class="text-caption text-grey-8">Расход</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
							<div class="col-12 col-sm-6 col-md-4">
								<q-card flat bordered class="bg-indigo-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="savings"
											:size="isMobile ? '32px' : '40px'"
											color="indigo-7"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-indigo-9"
											>
												{{ data.total_profit.toLocaleString() }}
											</div>
											<div class="text-caption text-grey-8">Прибыль</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
						</div>

						<!-- Оплата -->
						<div class="text-subtitle1 text-weight-medium q-mb-sm text-grey-8">
							Статус оплаты
						</div>
						<div class="row q-col-gutter-md">
							<div class="col-12 col-sm-6 col-md-4">
								<q-card flat bordered class="bg-cyan-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="paid"
											:size="isMobile ? '32px' : '40px'"
											color="cyan-7"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-cyan-9"
											>
												{{ data.paid_orders }}
											</div>
											<div class="text-caption text-grey-8">Оплачено</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
							<div class="col-12 col-sm-6 col-md-4">
								<q-card flat bordered class="bg-pink-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="credit_card_off"
											:size="isMobile ? '32px' : '40px'"
											color="pink-7"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-pink-9"
											>
												{{ data.unpaid_orders }}
											</div>
											<div class="text-caption text-grey-8">Не оплачено</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
							<div class="col-12 col-sm-6 col-md-4">
								<q-card flat bordered class="bg-amber-1">
									<q-card-section
										:class="isMobile ? 'q-pa-sm' : 'q-pa-md'"
										class="flex items-center"
										:style="isMobile ? 'gap: 8px' : 'gap: 12px'"
									>
										<q-icon
											name="calculate"
											:size="isMobile ? '32px' : '40px'"
											color="amber-8"
										/>
										<div class="flex-1">
											<div
												:class="isMobile ? 'text-h6' : 'text-h5'"
												class="text-weight-bold text-amber-9"
											>
												{{ data.average_order_price.toLocaleString() }}
											</div>
											<div class="text-caption text-grey-8">Средний чек</div>
										</div>
									</q-card-section>
								</q-card>
							</div>
						</div>
					</div>
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
	</PageLoading>
</template>

<style scoped>
@import "@/styles/telegram-app.scss";
</style>
