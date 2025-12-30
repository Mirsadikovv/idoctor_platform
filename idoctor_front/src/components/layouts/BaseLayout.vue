<script setup lang="ts">
import { computed, ref } from "vue";
import { useRouter } from "vue-router";

import { AuthService, type LanguageType } from "@/service";
import { useLanguageStore } from "@/store/language-store";
import { LanguageContentService } from "@/modules/TranslatedContent/service";
import { useAuthStore } from "@/store/auth-store";
import { AuthLayoutRoute } from "@/router";
import { buildSidebar } from "@/common";
import ButtonDialog from "../quasar/dialog/ButtonDialog.vue";
import PageLoading from "../PageLoading.vue";

const router = useRouter();
const languageStore = useLanguageStore();
const authStore = useAuthStore();

const drawerOpen = ref(true);
const confirm = ref(false);

function toggleLeftDrawer() {
	drawerOpen.value = !drawerOpen.value;
}

const iconToggle = computed(() => {
	return drawerOpen.value ? "menu_open" : "menu";
});

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
async function find(_params: string) {}
const SideList = buildSidebar();
</script>

<template>
	<PageLoading :find="find" #="{ loading }">
		<!-- Скелетон для загрузки -->
		<div v-if="loading" class="loading-skeleton">
			<div class="skeleton-header">
				<div class="skeleton-toolbar">
					<div class="skeleton-btn skeleton-animate"></div>
					<div class="skeleton-logo skeleton-animate"></div>
					<div class="skeleton-spacer"></div>
					<div class="skeleton-profile skeleton-animate"></div>
					<div class="skeleton-lang skeleton-animate"></div>
					<div class="skeleton-btn skeleton-animate"></div>
				</div>
			</div>
			<div class="skeleton-content">
				<div class="skeleton-sidebar">
					<div class="skeleton-menu-item skeleton-animate" v-for="n in 8" :key="n"></div>
				</div>
				<div class="skeleton-page">
					<div class="skeleton-page-content">
						<div
							class="skeleton-text-line skeleton-animate"
							v-for="n in 12"
							:key="n"
						></div>
					</div>
				</div>
			</div>
		</div>

		<q-layout view="hHh Lpr lff" v-else>
			<q-footer
				class="clean-header bg-gradient-to-r from-blue-900 via-blue-700 to-purple-700 shadow-lg border-b border-white/20"
			>
				<q-toolbar class="h-18 gap-x-3 clean-toolbar bg-transparent text-white px-6">
					<q-btn
						size="lg"
						:icon="iconToggle"
						@click="toggleLeftDrawer"
						class="clean-btn bg-transparent! text-white! rounded-lg! transition-all duration-200 min-h-44px! min-w-44px! hover:bg-white/15! active:bg-white/20!"
						dense
						round
						flat
					/>

					<q-toolbar-title class="flex items-center gap-4"> </q-toolbar-title>

					<div class="hidden xl:flex! md:flex! items-center gap-x-3">
						<!-- Desktop Profile Button -->
						<q-btn-dropdown
							class="hidden lg:flex! clean-desktop-profile"
							no-caps
							flat
							icon="account_circle"
						>
							<template v-slot:label>
								<div class="flex items-center gap-2 text-white">
									<q-avatar size="32px" class="bg-white/20">
										<q-icon name="account_circle" size="20px" color="white" />
									</q-avatar>
									<div class="flex flex-col items-start">
										<span class="text-sm font-medium">
											{{ authStore.user?.username || "User" }}
										</span>
										<span class="text-xs text-white/80">
											{{ $tl("profile") }}
										</span>
									</div>
								</div>
							</template>

							<q-list class="desktop-profile-menu min-w-280px">
								<!-- User Info -->
								<q-item class="desktop-profile-info">
									<q-item-section avatar>
										<q-avatar
											size="48px"
											class="bg-gradient-to-br from-blue-500 to-purple-600"
										>
											<q-icon
												name="account_circle"
												size="28px"
												color="white"
											/>
										</q-avatar>
									</q-item-section>
									<q-item-section>
										<q-item-label class="text-base font-semibold text-gray-800">
											{{ authStore.user?.firstName }}
											{{ authStore.user?.lastName }}
										</q-item-label>
										<q-item-label caption class="text-sm text-gray-600">
											{{ authStore.user?.username }}
										</q-item-label>
									</q-item-section>
								</q-item>

								<q-separator class="my-2" />

								<!-- Profile Action -->
								<q-item
									clickable
									@click="router.push({ name: 'PAGE_PROFILE' })"
									class="desktop-menu-item hover:bg-blue-50"
									v-close-popup
								>
									<q-item-section avatar class="min-w-0">
										<q-icon name="person" color="blue" size="20px" />
									</q-item-section>
									<q-item-section>
										<q-item-label class="text-gray-700 font-medium">
											{{ $tl("profile") }}
										</q-item-label>
									</q-item-section>
								</q-item>

								<!-- Settings (if needed) -->
								<q-item
									clickable
									class="desktop-menu-item hover:bg-gray-50"
									v-close-popup
								>
									<q-item-section avatar class="min-w-0">
										<q-icon name="settings" color="gray" size="20px" />
									</q-item-section>
									<q-item-section>
										<q-item-label class="text-gray-700 font-medium">
											{{ $tl("settings") || "Settings" }}
										</q-item-label>
									</q-item-section>
								</q-item>

								<q-separator class="my-2" />

								<!-- Logout -->
								<q-item
									clickable
									@click="confirm = true"
									class="desktop-menu-item hover:bg-red-50"
									v-close-popup
								>
									<q-item-section avatar class="min-w-0">
										<q-icon name="logout" color="negative" size="20px" />
									</q-item-section>
									<q-item-section>
										<q-item-label class="text-red-600 font-medium">
											{{ $tl("logout") }}
										</q-item-label>
									</q-item-section>
								</q-item>
							</q-list>
						</q-btn-dropdown>

						<!-- Desktop Language Selector -->
						<q-btn-dropdown
							class="hidden md:flex! sm:hidden! clean-profile-lang"
							:label="$lang._currentLang?.name.toUpperCase()"
							no-caps
							flat
						>
							<q-list class="clean-lang-list min-w-70px">
								<q-item
									v-for="value in $lang.languages"
									:key="value.id"
									clickable
									v-close-popup
									@click="setLang(value)"
									class="clean-lang-item rounded-lg mx-1 my-0.5 transition-colors hover:bg-blue-50"
								>
									<q-item-section>
										<q-item-label
											class="text-center text-sm font-medium text-blue-700"
										>
											{{ value.name.toUpperCase() }}
										</q-item-label>
									</q-item-section>
									<q-item-section side class="ml-2">
										<q-icon
											v-if="value.id === $lang._currentLang?.id"
											name="check_circle"
											size="16px"
											color="positive"
										/>
										<q-icon
											v-else
											size="16px"
											name="arrow_forward"
											color="blue"
											class="opacity-60"
										/>
									</q-item-section>
								</q-item>
							</q-list>
						</q-btn-dropdown>
					</div>

					<!-- Mobile Menu (Telegram optimized) -->
					<q-btn-dropdown
						flat
						aria-label="Menu"
						size="lg"
						class="lg:hidden! md:flex! clean-menu-dropdown bg-transparent! text-white! rounded-lg! transition-all duration-200 min-h-44px! min-w-44px!"
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
									<q-item-label caption class="text-xs text-gray-600">{{
										$tl("profile")
									}}</q-item-label>
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

							<!-- Logout Item -->
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

			<q-drawer
				show-if-above
				v-model="drawerOpen"
				side="left"
				class="bg-gray-50 border-r border-gray-200 text-gray-700! font-bold flex flex-col justify-between"
			>
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
					<SideList />
				</q-scroll-area>
			</q-drawer>

			<q-page-container>
				<q-page
					:style="{
						height: 'calc(var(--app-height, 100vh) - 150px)',
						'padding-top': 'var(--tg-safe-area-inset-top), 1rem)',
						'padding-bottom': 'var(--tg-safe-area-inset-bottom), 1rem)',
						'padding-left': 'max(var(--tg-safe-area-inset-left), 1rem)',
						'padding-right': 'max(var(--tg-safe-area-inset-right), 1rem)',
					}"
					class="bg-white text-gray-900 overflow-auto p-4 pt-20"
				>
					<router-view />
				</q-page>
			</q-page-container>
		</q-layout>
	</PageLoading>

	<q-dialog v-model="confirm" persistent>
		<q-card class="min-w-300px w-400px rounded-xl overflow-hidden shadow-2xl">
			<q-card-section class="bg-gradient-to-r from-blue-600 to-blue-700 text-white p-6">
				<div class="text-xl lg:text-2xl font-bold">{{ $tl("logout_confirm") }}</div>
				<div class="text-base opacity-90 mt-1">{{ $tl("are_you_sure") }}?</div>
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
	</q-dialog>
</template>

<style lang="scss">
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

/* Desktop Profile Dropdown */
.clean-desktop-profile {
	background: rgba(255, 255, 255, 0.15) !important;
	border: 1px solid rgba(255, 255, 255, 0.2) !important;
	border-radius: 12px !important;
	backdrop-filter: blur(10px);
	transition: all 0.2s ease;

	&:hover {
		background: rgba(255, 255, 255, 0.25) !important;
		transform: translateY(-1px);
		box-shadow: 0 4px 12px rgba(255, 255, 255, 0.15);
	}
}

.desktop-profile-menu {
	background: rgba(255, 255, 255, 0.98);
	backdrop-filter: blur(20px);
	border-radius: 16px;
	border: 1px solid rgba(0, 0, 0, 0.08);
	box-shadow: 0 12px 40px rgba(0, 0, 0, 0.12);
	padding: 12px;
}

.desktop-profile-info {
	background: rgba(59, 130, 246, 0.05);
	border: 1px solid rgba(59, 130, 246, 0.1);
	border-radius: 12px;
	margin-bottom: 8px;
	padding: 16px 12px;
}

.desktop-menu-item {
	border-radius: 8px;
	margin: 2px 0;
	min-height: 44px;
	transition: all 0.2s ease;

	&:hover {
		transform: translateX(2px);
	}
}

/* Language Selector Styles */
.clean-profile-lang {
	background: rgba(255, 255, 255, 0.15) !important;
	border: 1px solid rgba(255, 255, 255, 0.2) !important;
	border-radius: 8px !important;
	backdrop-filter: blur(10px);
	transition: all 0.2s ease;

	&:hover {
		background: rgba(255, 255, 255, 0.25) !important;
		transform: translateY(-1px);
		box-shadow: 0 4px 12px rgba(255, 255, 255, 0.15);
	}
}

.clean-lang-list {
	background: rgba(255, 255, 255, 0.98);
	backdrop-filter: blur(20px);
	border-radius: 12px;
	border: 1px solid rgba(0, 0, 0, 0.08);
	box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
	padding: 8px;
}

/* App Logo Animation */
.app-logo-text {
	background: linear-gradient(135deg, rgba(255, 255, 255, 1), rgba(255, 255, 255, 0.8));
	-webkit-background-clip: text;
	-webkit-text-fill-color: transparent;
	background-clip: text;
	letter-spacing: 0.5px;
	animation: logoGlow 3s ease-in-out infinite alternate;
}

@keyframes logoGlow {
	0% {
		text-shadow: 0 0 10px rgba(255, 255, 255, 0.5), 0 0 20px rgba(255, 255, 255, 0.3);
	}
	100% {
		text-shadow: 0 0 20px rgba(255, 255, 255, 0.8), 0 0 30px rgba(255, 255, 255, 0.5);
	}
}

/* ========== Loading Skeleton ========== */
.loading-skeleton {
	height: 100vh;
	background: #f9fafb;
}

.skeleton-header {
	height: 72px;
	background: white;
	border-bottom: 1px solid #e5e7eb;
	box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.skeleton-toolbar {
	height: 100%;
	display: flex;
	align-items: center;
	padding: 0 24px;
	gap: 16px;
}

.skeleton-btn,
.skeleton-logo,
.skeleton-profile,
.skeleton-lang {
	background: #e5e7eb;
	border-radius: 8px;
}

.skeleton-btn {
	width: 40px;
	height: 40px;
}
.skeleton-logo {
	width: 120px;
	height: 40px;
}
.skeleton-profile {
	width: 180px;
	height: 40px;
}
.skeleton-lang {
	width: 80px;
	height: 40px;
}
.skeleton-spacer {
	flex: 1;
}

.skeleton-content {
	display: flex;
	height: calc(100vh - 72px);
}

.skeleton-sidebar {
	width: 300px;
	background: #f9fafb;
	border-right: 1px solid #e5e7eb;
	padding: 16px;
	display: flex;
	flex-direction: column;
	gap: 8px;
}

.skeleton-menu-item {
	height: 48px;
	background: #e5e7eb;
	border-radius: 8px;
}

.skeleton-page {
	flex: 1;
	padding: 24px;
	background: #ffffff;
}

.skeleton-page-content {
	display: flex;
	flex-direction: column;
	gap: 16px;
}

.skeleton-text-line {
	height: 16px;
	background: #e5e7eb;
	border-radius: 4px;

	&:nth-child(odd) {
		width: 100%;
	}
	&:nth-child(even) {
		width: 85%;
	}
	&:nth-child(3n) {
		width: 70%;
	}
}

.skeleton-animate {
	background: linear-gradient(90deg, #e5e7eb 25%, #f3f4f6 50%, #e5e7eb 75%);
	background-size: 200% 100%;
	animation: skeleton-loading 1.5s infinite;
}

@keyframes skeleton-loading {
	0% {
		background-position: 200% 0;
	}
	100% {
		background-position: -200% 0;
	}
}

/* ========== Responsive Design ========== */
@media screen and (max-width: 768px) {
	.skeleton-sidebar,
	.skeleton-profile,
	.skeleton-lang {
		display: none;
	}
	.skeleton-toolbar {
		gap: 8px;
	}
}

@media screen and (max-width: 480px) {
	.adapt-padding {
		padding: 6px !important;
	}
}

/* Отключение анимаций на слабых устройствах */
@media (prefers-reduced-motion: reduce) {
	.app-logo-text,
	.skeleton-animate {
		animation: none;
	}
	* {
		transition: none !important;
	}
}
</style>
