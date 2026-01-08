<script setup lang="ts">
import { ref } from "vue";
import { AuthService, LanguageContentService, type LanguageType } from "@/service";
import PageLoading from "@/components/PageLoading.vue";

import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import { useRouter } from "vue-router";
import { useLanguageStore } from "@/store/language-store";
import { useAuthStore } from "@/store/auth-store";
import { AuthLayoutRoute } from "@/router";
import { getAccessibleRoutes, type AccessibleRoute } from "@/common";
import { useTelegramViewport } from "@/composables/useTelegramViewport";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
// import { buildSidebar } from "@/common";

const router = useRouter();
const authStore = useAuthStore();
const languageStore = useLanguageStore();
const { containerStyle } = useTelegramViewport();

const pages = ref<AccessibleRoute[]>([]);
const confirm = ref(false);

async function getRoles(_query: string = "") {
	pages.value = getAccessibleRoutes();
}

async function setLang(language: LanguageType) {
	const globalContent = LanguageContentService.getLanguageContents(language.id);

	const [global] = await Promise.all([globalContent]);

	languageStore.setGlobalLang(global);

	router.push({
		name: router.currentRoute.value.name as string,
		params: {
			...router.currentRoute.value.params,
			lang: language.name,
		},
		query: router.currentRoute.value.query,
	});
}

async function logout() {
	await AuthService.signOut().catch(() => {});
	authStore.removeSession();
	router.clearRoutes();
	router.addRoute(AuthLayoutRoute);
	await router.push({
		name: "LOGIN_AUTH",
	});
}
</script>

<template>
	<PageLoading :find="getRoles" #="{ loading }">
		<!-- Скелетон для загрузки -->
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page :style="containerStyle" class="bg-gray-100 text-gray-900 overflow-auto p-4">
					<q-scroll-area
						class="h-full bg-transparent p-0"
						visible
						:thumb-style="{
							right: '5px',
							borderRadius: '5px',
							backgroundColor: '#6b7280',
							width: '6px',
						}"
						:bar-style="{ width: '0px' }"
						ref="firstRef"
					>
						<div class="flex flex-col gap-3 h-full">
							<q-item
								v-for="pageName in pages"
								:key="pageName.name"
								clickable
								@click="router.push({ name: pageName.name })"
								class="telegram-page-item rounded-xl transition-all duration-200 min-h-56px bg-white border border-gray-200 hover:bg-blue-50! hover:border-blue-300! active:bg-blue-100! shadow-sm hover:shadow-md!"
							>
								<q-item-section>
									<q-item-label class="text-base font-medium text-gray-800!">
										{{ $tl(pageName.label) }}
									</q-item-label>
								</q-item-section>
								<q-item-section side>
									<q-icon name="chevron_right" color="gray-400" size="20px" />
								</q-item-section>
							</q-item>
						</div>
					</q-scroll-area>
				</q-page>
			</q-page-container>

			<q-footer class="clean-header shadow-lg border-b border-white/20">
				<q-toolbar
					class="h-18 gap-x-3 clean-toolbar bg-transparent text-white px-6 justify-end"
				>
					<!-- Mobile Menu (Telegram optimized) -->
					<q-btn-dropdown
						flat
						aria-label="Menu"
						size="lg"
						dropdown-icon="settings"
						class="flex clean-menu-dropdown bg-transparent! text-white! rounded-lg! transition-all duration-200 min-h-44px! min-w-44px!"
					>
						<q-list
							class="telegram-menu-list min-w-200px max-w-320px bg-white/95 backdrop-blur-20 rounded-xl border border-black/10 shadow-2xl p-2"
						>
							<!-- Profile Item -->
							<q-item
								clickable
								@click="router.push({ name: 'PAGE_PROFILE' })"
								class="telegram-menu-item rounded-lg my-1 transition-all duration-200 min-h-48px bg-green-50 border border-green-100 hover:bg-green-100! active:bg-green-150!"
								v-close-popup
							>
								<q-item-section avatar>
									<q-icon name="account_box" color="positive" size="20px" />
								</q-item-section>
								<q-item-section>
									<q-item-label class="text-sm font-medium text-gray-800">
										{{ authStore.user?.username }}
									</q-item-label>
								</q-item-section>
							</q-item>

							<q-separator class="my-1 bg-gray-200" />

							<!-- Language Selector -->
							<q-expansion-item
								icon="language"
								:label="$tl('language')"
								header-class="text-gray-800! font-medium px-3 py-2 rounded-lg hover:bg-gray-50!"
								class="telegram-menu-item my-1"
							>
								<q-item
									v-for="language of $lang.languages"
									:key="language.id"
									clickable
									:class="[
										'mx-2 my-1 rounded-lg min-h-40px transition-colors',
										language.id === $lang._currentLang?.id
											? 'bg-blue-100! text-blue-800! font-semibold'
											: 'hover:bg-blue-50!',
									]"
									@click="setLang(language)"
									v-close-popup
								>
									<q-item-section>
										<q-item-label class="text-sm">{{
											language.name
										}}</q-item-label>
									</q-item-section>
									<q-item-section side>
										<q-icon
											v-if="language.id === $lang._currentLang?.id"
											name="check_circle"
											size="16px"
											color="positive"
										/>
									</q-item-section>
								</q-item>
							</q-expansion-item>

							<q-separator class="my-1 bg-gray-200" />

							<!-- Logout Mobile -->
							<q-item
								clickable
								@click="confirm = true"
								class="telegram-menu-item rounded-lg my-1 transition-all duration-200 min-h-48px bg-red-50 border border-red-100 hover:bg-red-100! active:bg-red-150!"
								v-close-popup
							>
								<q-item-section avatar>
									<q-icon name="logout" color="negative" size="20px" />
								</q-item-section>
								<q-item-section>
									<q-item-label class="text-sm font-medium text-red-600">
										{{ $tl("logout") }}
									</q-item-label>
								</q-item-section>
							</q-item>
						</q-list>
					</q-btn-dropdown>

					<!-- Logout desktop -->
					<ButtonDialog
						:classBtn="'hidden lg:flex! md:hidden! bg-red-600 text-white rounded-lg px-3 py-2 transition-all duration-200 hover:bg-red-700! hover:-translate-y-0.5 hover:shadow-lg active:translate-y-0'"
						icon="logout"
						label="logout"
					>
						<q-card class="min-w-300px w-400px rounded-xl overflow-hidden">
							<q-card-section
								class="bg-gradient-to-r from-blue-600 to-blue-700 text-white p-6"
							>
								<div class="text-2xl font-bold">{{ $tl("logout_confirm") }}</div>
								<div class="text-lg opacity-90 mt-1">
									{{ $tl("are_you_sure") }}?
								</div>
							</q-card-section>

							<q-card-actions align="center" class="flex gap-3 p-4">
								<q-btn
									no-caps
									outline
									class="flex-1 py-2 px-4 border-2 border-blue-500 text-blue-600 rounded-lg font-medium hover:bg-blue-50 transition-colors"
									:label="$tl('no')"
									v-close-popup
								/>
								<q-btn
									no-caps
									class="flex-1 py-2 px-4 bg-red-600 text-white rounded-lg font-medium hover:bg-red-700 transition-colors"
									:label="$tl('yes')"
									v-close-popup
									@click="logout()"
								/>
							</q-card-actions>
						</q-card>
					</ButtonDialog>
				</q-toolbar>
			</q-footer>

			<!-- Mobile Logout Confirmation Dialog -->
			<q-dialog v-model="confirm" persistent>
				<q-card class="min-w-300px w-400px rounded-xl overflow-hidden">
					<q-card-section
						class="bg-gradient-to-r from-blue-600 to-blue-700 text-white p-6"
					>
						<div class="text-2xl font-bold">{{ $tl("logout_confirm") }}</div>
						<div class="text-lg opacity-90 mt-1">{{ $tl("are_you_sure") }}?</div>
					</q-card-section>

					<q-card-actions align="center" class="flex gap-3 p-4">
						<q-btn
							no-caps
							outline
							color="secondary"
							class="flex-1 py-2 px-4 border-2 border-blue-500 text-blue-600 rounded-lg font-medium hover:bg-blue-50 transition-colors"
							:label="$tl('no')"
							v-close-popup
						/>
						<q-btn
							no-caps
							outline
							color="negative"
							class="flex-1 py-2 px-4"
							:label="$tl('yes')"
							v-close-popup
							@click="logout()"
						/>
					</q-card-actions>
				</q-card>
			</q-dialog>
		</q-layout>
	</PageLoading>
</template>

<style scoped lang="scss">
/* ========== Telegram Web App Variables ========== */
:root {
	--tg-color-scheme: var(--tg-theme-color-scheme, light);
	--tg-bg-color: var(--tg-theme-bg-color, #ffffff);
	--tg-text-color: var(--tg-theme-text-color, #000000);
	--tg-hint-color: var(--tg-theme-hint-color, #999999);
	--tg-button-color: var(--tg-theme-button-color, #3390ec);
	--tg-button-text-color: var(--tg-theme-button-text-color, #ffffff);
	--app-height: var(--tg-viewport-height, 100vh);
	--tg-safe-area-inset-top: env(safe-area-inset-top, 0);
	--tg-safe-area-inset-bottom: env(safe-area-inset-bottom, 0);
	--tg-safe-area-inset-left: env(safe-area-inset-left, 0);
	--tg-safe-area-inset-right: env(safe-area-inset-right, 0);
}

@media (prefers-color-scheme: dark) {
	:root {
		--tg-bg-color: var(--tg-theme-bg-color, #212121);
		--tg-text-color: var(--tg-theme-text-color, #ffffff);
		--tg-hint-color: var(--tg-theme-hint-color, #cccccc);
	}
}

/* ========== Telegram UX Optimizations ========== */
* {
	-webkit-user-select: none;
	-moz-user-select: none;
	-ms-user-select: none;
	user-select: none;
	-webkit-tap-highlight-color: transparent;
}

input,
textarea,
[contenteditable] {
	-webkit-user-select: text;
	-moz-user-select: text;
	-ms-user-select: text;
	user-select: text;
}

/* ========== Custom Styles for Quasar Components ========== */
.q-item__section--avatar {
	min-width: 0 !important;
}

@media screen and (max-width: 480px) {
	.adapt-padding {
		padding: 6px !important;
	}
}

/* ========== Page Navigation Styles ========== */
.telegram-page-item {
	background: rgba(255, 255, 255, 0.95);
	backdrop-filter: blur(10px);
	border: 1px solid rgba(0, 0, 0, 0.05) !important;
	box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
	padding: 12px 16px;
	margin: 0 4px;

	&:hover {
		background: rgba(59, 130, 246, 0.05) !important;
		border-color: rgba(59, 130, 246, 0.2) !important;
		transform: translateY(-1px);
		box-shadow: 0 4px 16px rgba(59, 130, 246, 0.1) !important;
	}

	&:active {
		background: rgba(59, 130, 246, 0.1) !important;
		transform: translateY(0);
		box-shadow: 0 2px 8px rgba(59, 130, 246, 0.15) !important;
	}

	.q-item__section--avatar {
		color: var(--tg-button-color, #3390ec);
		transition: all 0.2s ease;
	}

	&:hover .q-item__section--avatar {
		transform: translateX(2px);
	}
}

/* Mobile-first responsive adjustments */
@media screen and (max-width: 480px) {
	.telegram-page-item {
		min-height: 64px;
		margin: 0 2px;
		padding: 16px;
		border-radius: 16px;

		.q-item__label {
			font-size: 16px !important;
			line-height: 1.4;
		}

		.q-item__section--avatar {
			min-width: 16px;
		}
	}
}

/* Touch optimizations for mobile */
@media (hover: none) and (pointer: coarse) {
	.telegram-page-item {
		&:hover {
			transform: none;
		}

		&:active {
			background: rgba(59, 130, 246, 0.15) !important;
			border-color: rgba(59, 130, 246, 0.3) !important;
		}
	}
}

/* Dark mode support */
@media (prefers-color-scheme: dark) {
	.telegram-page-item {
		background: rgba(33, 33, 33, 0.95);
		border-color: rgba(255, 255, 255, 0.1);
		color: var(--tg-text-color, #ffffff);

		&:hover {
			background: rgba(59, 130, 246, 0.15) !important;
			border-color: rgba(59, 130, 246, 0.3) !important;
		}

		.q-item__label {
			color: var(--tg-text-color, #ffffff);
		}
	}
}

/* Отключение анимаций на слабых устройствах */
@media (prefers-reduced-motion: reduce) {
	.app-logo-text,
	.skeleton-animate,
	.telegram-page-item {
		animation: none;
	}
	* {
		transition: none !important;
	}
}
</style>
