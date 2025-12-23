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

const scrollAreaStyle = {
	thumbStyle: {
		right: "5px",
		borderRadius: "5px",
		backgroundColor: "hsl(210, 4%, 42%)",
		width: "6px",
	},
	barStyle: {
		width: "0px",
	},
};

const router = useRouter();
const languageStore = useLanguageStore();
const authStore = useAuthStore();

const drawerOpen = ref(true);
const showMobileMenu = ref(false);
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
			<q-header class="clean-header">
				<q-toolbar class="h-18 gap-x-3 clean-toolbar">
					<q-btn
						size="lg"
						color="primary"
						:icon="iconToggle"
						@click="toggleLeftDrawer"
						class="clean-btn"
						dense
						round
						flat
					/>

					<q-toolbar-title class="flex items-center gap-4">
						<div class="min-w-100px app-logo-text">iDoctor</div>
					</q-toolbar-title>

					<div class="display-none xl:flex! md:flex! items-center gap-x-2">
						<q-btn-dropdown
							class="display-none md:flex! sm:hidden! clean-profile-lang"
							:label="$lang._currentLang?.name.toUpperCase()"
							no-caps
							flat
						>
							<q-list class="clean-lang-list">
								<q-item
									v-for="value in $lang.languages"
									:key="value.id"
									clickable
									v-close-popup
									@click="setLang(value)"
									:class="['clean-lang-item']"
								>
									<q-item-section>
										<q-item-label class="clean-lang-label">
											{{ value.name.toUpperCase() }}
										</q-item-label>
									</q-item-section>
									<q-item-section>
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
										/>
									</q-item-section>
								</q-item>
							</q-list>
						</q-btn-dropdown>
					</div>

					<q-btn
						flat
						icon="format_indent_decrease"
						aria-label="Menu"
						size="lg"
						class="lg:hidden! md:block! clean-btn"
						@click="showMobileMenu = !showMobileMenu"
					/>

					<!-- <div class="display-none xl:flex! md:flex! items-center">
						<div
							@click="() => router.push({ name: 'PAGE_PROFILE' })"
							class="clean-profile-chip"
						>
							<q-icon name="person" class="mr-2" />
							{{ authStore.user?.lastName }}
							{{ authStore.user?.firstName?.charAt(0) }}.
							{{ authStore.user?.middleName?.charAt(0) }}.
						</div>
					</div> -->

					<ButtonDialog
						:classBtn="'clean-logout-btn bg-negative adapt-padding p-2 display-none lg:flex! md:hidden! h-40px!'"
						icon="logout"
						label="logout"
					>
						<q-card class="min-w-[300px] w-400px">
							<q-card-section class="bg-blue text-white">
								<div class="text-3xl">{{ $tl("logout_confirm") }}</div>
								<div class="text-xl">{{ $tl("are_you_sure") }} ?</div>
							</q-card-section>

							<q-card-actions align="center" class="flex no-wrap gap-x-2">
								<q-btn
									no-caps
									outline
									class="full-width"
									:label="$tl('no')"
									color="blue"
									v-close-popup
								/>
								<q-btn
									unelevated
									no-caps
									class="full-width"
									:label="$tl('yes')"
									color="negative"
									v-close-popup
									@click="logout()"
								/>
							</q-card-actions>
						</q-card>
					</ButtonDialog>
				</q-toolbar>
			</q-header>

			<q-drawer
				show-if-above
				v-model="drawerOpen"
				side="left"
				class="simple-sidebar font-700 select-text! flex flex-col justify-between no-wrap"
			>
				<q-scroll-area
					class="fit simple-scroll"
					visible
					:thumb-style="scrollAreaStyle.thumbStyle"
					:bar-style="scrollAreaStyle.barStyle"
					ref="firstRef"
				>
					<SideList />
				</q-scroll-area>
			</q-drawer>

			<q-drawer v-model="showMobileMenu" side="right" overlay class="modern-mobile-menu">
				<q-list class="modern-mobile-list">
					<!-- Название приложения -->
					<q-item class="modern-mobile-item app-name">
						<q-item-section avatar>
							<q-icon name="apps" />
						</q-item-section>
						<q-item-section>
							<q-item-label>{{ $tl("app_name") }}</q-item-label>
						</q-item-section>
					</q-item>

					<q-separator class="modern-separator" />

					<!-- Профиль пользователя -->
					<q-item
						clickable
						@click="router.push({ name: 'PAGE_PROFILE' })"
						class="modern-mobile-item profile"
					>
						<q-item-section avatar>
							<q-icon name="account_box" />
						</q-item-section>
						<q-item-section>
							<q-item-label>
								{{ authStore.user?.lastName }}
								{{ authStore.user?.firstName?.charAt(0) }}.
								{{ authStore.user?.middleName?.charAt(0) }}.
							</q-item-label>
						</q-item-section>
					</q-item>

					<q-separator class="modern-separator" />

					<!-- Языковой селектор -->
					<q-expansion-item
						icon="language"
						:label="$tl('language')"
						header-class="modern-mobile-expansion"
						class="modern-mobile-item"
					>
						<q-item
							v-for="language of $lang.languages"
							:key="language.id"
							clickable
							:class="
								language.id === $lang._currentLang?.id
									? `modern-lang-active`
									: `modern-lang-item`
							"
							@click="
								setLang(language);
								showMobileMenu = false;
							"
						>
							<q-item-section>
								<q-item-label>{{ language.name }}</q-item-label>
							</q-item-section>
						</q-item>
					</q-expansion-item>

					<q-separator class="modern-separator" />

					<!-- Logout -->
					<q-item
						clickable
						@click="
							confirm = true;
							showMobileMenu = false;
						"
						class="modern-mobile-item logout"
					>
						<q-item-section avatar>
							<q-icon name="logout" />
						</q-item-section>
						<q-item-section>
							<q-item-label>{{ $tl("logout") }}</q-item-label>
						</q-item-section>
					</q-item>
				</q-list>
			</q-drawer>

			<q-page-container>
				<q-page style="max-height: calc(100vh - 150px)" class="overflow-auto p-4 bg-light">
					<router-view />
				</q-page>
			</q-page-container>
		</q-layout>
	</PageLoading>

	<q-dialog v-model="confirm" persistent>
		<q-card class="min-w-[300px] w-400px">
			<q-card-section class="bg-blue text-white">
				<div class="lg:text-3xl text-xl">{{ $tl("logout_confirm") }}</div>
			</q-card-section>

			<q-card-actions align="center" class="flex no-wrap gap-x-2">
				<q-btn
					no-caps
					outline
					class="full-width"
					:label="$tl('no')"
					color="blue"
					v-close-popup
				/>
				<q-btn
					unelevated
					no-caps
					class="full-width"
					:label="$tl('yes')"
					color="negative"
					v-close-popup
					@click="logout()"
				/>
			</q-card-actions>
		</q-card>
	</q-dialog>
</template>

<style lang="scss">
@media screen and (max-width: 400px) {
	.adapt-padding {
		padding: 8px !important;
	}
}

.q-item__section--avatar {
	min-width: 0 !important;
}

.display-none {
	display: none;
}

/* Clean Header Styles */
.clean-header {
	background: linear-gradient(135deg, #1a237e 0%, #3949ab 50%, #5e35b1 100%);
	box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
	border-bottom: 1px solid rgba(255, 255, 255, 0.2);
}

.clean-toolbar {
	background: transparent;
	color: white !important;
	padding: 0 24px;
}

.clean-btn {
	background: transparent !important;
	color: white !important;
	border-radius: 8px !important;
	transition: $transition-fast;

	&:hover {
		background: rgba(255, 255, 255, 0.15) !important;
		transform: scale(1.05);
	}

	&:active {
		transform: scale(0.95);
	}
}

.clean-profile-chip {
	height: 40px !important;

	display: flex;
	align-items: center;
	padding: 8px 12px;
	background: rgba(255, 255, 255, 0.15);
	border-radius: 8px;
	color: white;
	font-weight: 500;
	font-size: 14px;
	cursor: pointer;
	transition: $transition-fast;
	border: 1px solid rgba(255, 255, 255, 0.2);
	backdrop-filter: blur(10px);

	&:hover {
		background: rgba(255, 255, 255, 0.2);
		transform: translateY(-1px);
		box-shadow: 0 4px 12px rgba(255, 255, 255, 0.15);
	}

	&:active {
		transform: translateY(0);
	}

	.q-icon {
		color: white;
		font-size: 18px;
	}
}

.clean-profile-lang {
	height: 40px !important;

	display: flex;
	align-items: center;
	padding: 7px 12px;
	background: rgba(255, 255, 255, 0.15);
	border-radius: 8px;
	color: white;
	font-weight: 500;
	font-size: 14px;
	cursor: pointer;
	transition: $transition-fast;
	border: 1px solid rgba(255, 255, 255, 0.2);
	backdrop-filter: blur(10px);

	&:hover {
		background: rgba(255, 255, 255, 0.2);
		transform: translateY(-1px);
		box-shadow: 0 4px 12px rgba(255, 255, 255, 0.15);
	}

	&:active {
		transform: translateY(0);
	}

	.q-icon {
		color: white;
		font-size: 18px;
	}
}

.clean-logout-btn {
	border-radius: 8px !important;
	transition: $transition-fast;

	&:hover {
		transform: translateY(-1px);
		box-shadow: 0 4px 12px rgba(50, 38, 220, 0.2);
	}

	&:active {
		transform: translateY(0);
	}
}

/* Logo Text Styles */
.app-logo-text {
	font-size: 1.75rem;
	font-weight: 700;
	color: white;
	text-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
	letter-spacing: 0.5px;
	background: linear-gradient(135deg, rgba(255, 255, 255, 1), rgba(255, 255, 255, 0.8));
	-webkit-background-clip: text;
	-webkit-text-fill-color: transparent;
	background-clip: text;
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

/* Language List Styles */
.clean-lang-list {
	min-width: 70px;
}

.clean-lang-item {
	min-height: 30px;
	border-radius: 8px !important;
	margin: 2px 0;
	transition: all 0.3s ease;
	color: $primary;
	position: relative;
	text-align: center !important;
	font-size: 12px;

	&:hover {
		background: rgba(43, 65, 129, 0.1) !important;
		transform: translateY(-1px);
		box-shadow: 0 2px 8px rgba(43, 65, 129, 0.15);
	}

	&:active {
		transform: translateY(0);
	}
}

.clean-lang-label {
	font-size: 13px;
	letter-spacing: 0.5px;
}

.q-item__section {
	color: inherit;
}

.q-item__label {
	font-size: 14px;
}

.q-item__label--caption {
	color: rgba(43, 65, 129, 0.7);
	font-size: 12px;
}

/* Simple Sidebar Styles */
.simple-sidebar {
	background: #f9fafb;
	border-right: 1px solid #e5e7eb;
	color: #374151 !important;
}

.simple-scroll {
	background: transparent;
	padding: 0;
}

/* Modern Mobile Menu Styles */
.modern-mobile-menu {
	background: linear-gradient(135deg, $gradient-start 0%, $gradient-end 100%);
	backdrop-filter: $backdrop-blur;
	-webkit-backdrop-filter: $backdrop-blur;
	border-left: 1px solid $glass-border;
	box-shadow: $modern-shadow-lg;
}

.modern-mobile-list {
	background: transparent;
	padding: 20px 0;
}

.modern-mobile-item {
	margin: 4px 12px;
	border-radius: $modern-radius;
	transition: $transition-base;
	color: rgba(255, 255, 255, 0.9);
	border: 1px solid transparent;

	&:hover {
		background: rgba(255, 255, 255, 0.15) !important;
		border-color: rgba(255, 255, 255, 0.2);
		transform: translateX(-4px);
		box-shadow: $modern-shadow;

		.q-icon {
			transform: scale(1.1);
		}
	}

	&:active {
		transform: translateX(-2px);
	}

	.q-icon {
		transition: $transition-fast;
		color: rgba(255, 255, 255, 0.9);
	}

	&.app-name {
		background: rgba(255, 255, 255, 0.1);
		border: 1px solid rgba(255, 255, 255, 0.2);

		.q-icon {
			color: #5a9bd5;
		}
	}

	&.profile {
		.q-icon {
			color: #4caf50;
		}
	}

	&.logout {
		.q-icon {
			color: #d32f2f;
		}

		.q-item-label {
			color: #d32f2f;
		}
	}
}

.modern-separator {
	background: rgba(255, 255, 255, 0.2);
	margin: 8px 20px;
}

.modern-mobile-expansion {
	color: rgba(255, 255, 255, 0.9) !important;

	.q-icon {
		color: #f9a825 !important;
	}
}

.modern-lang-item {
	margin: 2px 20px;
	border-radius: 8px;
	transition: $transition-base;
	color: rgba(255, 255, 255, 0.8);

	&:hover {
		background: rgba(255, 255, 255, 0.1) !important;
		color: white;
		transform: translateX(-2px);
	}
}

.modern-lang-active {
	margin: 2px 20px;
	border-radius: 8px;
	background: rgba(91, 155, 213, 0.3) !important;
	color: white !important;
	border: 1px solid rgba(91, 155, 213, 0.5);
	box-shadow: 0 2px 8px rgba(91, 155, 213, 0.2);
}

/* Loading Skeleton Styles */
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

.skeleton-btn {
	width: 40px;
	height: 40px;
	background: #e5e7eb;
	border-radius: 8px;
}

.skeleton-logo {
	width: 120px;
	height: 40px;
	background: #e5e7eb;
	border-radius: 8px;
}

.skeleton-spacer {
	flex: 1;
}

.skeleton-profile {
	width: 180px;
	height: 40px;
	background: #e5e7eb;
	border-radius: 8px;
}

.skeleton-lang {
	width: 80px;
	height: 40px;
	background: #e5e7eb;
	border-radius: 8px;
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

@media screen and (max-width: 768px) {
	.skeleton-sidebar {
		display: none;
	}

	.skeleton-profile,
	.skeleton-lang {
		display: none;
	}

	.skeleton-toolbar {
		gap: 8px;
	}
}
</style>
