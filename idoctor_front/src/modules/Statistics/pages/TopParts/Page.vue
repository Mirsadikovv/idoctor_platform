<script setup lang="ts">
import { ref, computed } from "vue";
import { StatisticsService } from "../../service/StatisticsService";
import type { TopPartsStatistics } from "../../service/types";
import type { PageDataType } from "@/service";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";
import AppFooter from "@/components/AppFooter.vue";
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();

const data = ref<TopPartsStatistics[]>([]);

const startDate = ref();
const endDate = ref();

const masterId = ref<number | undefined>(undefined);
const limit = ref(10);

const pick = {
	part_name: true,
	usage_count: true,
	total_revenue: true,
	total_cost: true,
	total_profit: true,
};

const tableData = computed<PageDataType<TopPartsStatistics>>(() => ({
	totalRows: data.value.length,
	totalPages: 1,
	currentPage: 1,
	pageSize: data.value.length,
	data: data.value,
}));

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

	const result = await StatisticsService.getTopParts({
		start_date: startDate.value,
		end_date: endDate.value,
		master_id: masterId.value,
		limit: limit.value,
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
					<div class="text-h5 text-weight-medium q-mb-md">Популярные детали</div>

					<q-card flat bordered class="q-mb-md bg-white shadow-2">
						<q-card-section>
							<div class="text-subtitle2 q-mb-sm">Фильтры</div>
							<div class="row q-col-gutter-md">
								<div class="col-12 col-sm-6 col-md-3">
									<q-input
										v-model="startDate"
										type="date"
										label="Дата начала"
										outlined
										dense
									/>
								</div>
								<div class="col-12 col-sm-6 col-md-3">
									<q-input
										v-model="endDate"
										type="date"
										label="Дата окончания"
										outlined
										dense
									/>
								</div>
								<div class="col-12 col-sm-6 col-md-3">
									<q-input
										v-model.number="masterId"
										type="number"
										label="ID мастера"
										outlined
										dense
										clearable
									/>
								</div>
								<div class="col-12 col-sm-6 col-md-3">
									<q-input
										v-model.number="limit"
										type="number"
										label="Лимит"
										outlined
										dense
										min="1"
										max="100"
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

					<q-card
						flat
						bordered
						class="bg-white shadow-2"
						v-if="!loading && data.length === 0"
					>
						<q-card-section class="text-center text-grey-7">
							Нет данных для отображения
						</q-card-section>
					</q-card>

					<ResponsiveTable
						v-else
						:models="tableData"
						:pick="pick"
						:loading="loading"
						has-order
					>
						<template #part_name:thead></template>
						<template #part_name="{ model }">
							{{ model.part_name }}
						</template>

						<template #usage_count:thead></template>
						<template #usage_count="{ model }">
							{{ model.usage_count }}
						</template>

						<template #total_revenue:thead></template>
						<template #total_revenue="{ model }">
							{{ model.total_revenue.toLocaleString() }}
						</template>

						<template #total_cost:thead></template>
						<template #total_cost="{ model }">
							{{ model.total_cost.toLocaleString() }}
						</template>

						<template #total_profit:thead></template>
						<template #total_profit="{ model }">
							{{ model.total_profit.toLocaleString() }}
						</template>

						<template #card="{ model, orderNumber }">
							<q-item clickable>
								<q-item-section avatar v-if="orderNumber">
									<q-avatar color="primary" text-color="white" size="md">
										{{ orderNumber }}
									</q-avatar>
								</q-item-section>

								<q-item-section>
									<q-item-label class="text-xl">
										<strong>{{ model.part_name }}</strong>
									</q-item-label>
									<q-item-label caption class="text-body2">
										Использований: {{ model.usage_count }}
									</q-item-label>
									<q-item-label caption class="text-body2">
										Доход: {{ model.total_revenue.toLocaleString() }} | Расход:
										{{ model.total_cost.toLocaleString() }} | Прибыль:
										{{ model.total_profit.toLocaleString() }}
									</q-item-label>
								</q-item-section>

								<q-item-section side>
									<q-icon name="chevron_right" color="grey-6" />
								</q-item-section>
							</q-item>
						</template>
					</ResponsiveTable>
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
