<script setup lang="ts">
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import { useAuthStore } from "@/store/auth-store";
import AppFooter from "@/components/AppFooter.vue";
import { useRouter } from "vue-router";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
const { containerStyle } = useTelegramViewport();
const router = useRouter();

const cards = [
	{
		title: "Статистика мастеров",
		description:
			"Просмотр статистики по каждому мастеру: количество заказов, доход, средний чек",
		route: { name: "STATISTICS_MASTERS" },
		icon: "people",
	},
	{
		title: "Статистика заказов",
		description: "Общая статистика заказов: количество, статусы, доходы, расходы, прибыль",
		route: { name: "STATISTICS_ORDERS" },
		icon: "assessment",
	},
	{
		title: "Статистика платежей",
		description: "Статистика по типам оплаты: наличные, карта, переводы, онлайн",
		route: { name: "STATISTICS_PAYMENTS" },
		icon: "payments",
	},
	{
		title: "Доход по периодам",
		description: "Доходы и расходы с группировкой по дням, неделям, месяцам или годам",
		route: { name: "STATISTICS_REVENUE" },
		icon: "trending_up",
	},
	{
		title: "Популярные детали",
		description: "Топ используемых деталей по количеству использований и прибыли",
		route: { name: "STATISTICS_TOP_PARTS" },
		icon: "stars",
	},
];
</script>

<template>
	<q-layout view="hHh Lpr lff">
		<q-page-container>
			<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
				<div class="text-h5 text-weight-medium q-mb-md">{{ $tl("statistics") }}</div>

				<div class="row q-col-gutter-md">
					<div
						v-for="card in cards"
						:key="card.route.name"
						class="col-12 col-sm-6 col-md-4"
					>
						<q-card
							flat
							bordered
							class="cursor-pointer shadow-2 transition-all hover:shadow-8"
							@click="router.push(card.route)"
						>
							<q-card-section class="flex items-center">
								<q-icon
									:name="card.icon"
									size="32px"
									color="primary"
									class="q-mr-md"
								/>
								<div class="flex-1">
									<div class="text-subtitle1 text-weight-medium">
										{{ card.title }}
									</div>
									<div class="text-caption text-grey-7">
										{{ card.description }}
									</div>
								</div>
								<q-icon name="chevron_right" color="grey-6" />
							</q-card-section>
						</q-card>
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
</template>

<style scoped>
@import "@/styles/telegram-app.scss";

.transition-all {
	transition: all 0.2s ease;
}

.hover\:shadow-8:hover {
	box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
}
</style>
