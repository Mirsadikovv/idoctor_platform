<script setup lang="ts">
import { QBtn } from "quasar";
import { ref } from "vue";
import { useRouter } from "vue-router";

import { AuthService, type LanguageType, type User } from "@/service";
import { useLanguageStore } from "@/store/language-store";
import { LanguageContentService } from "@/modules/TranslatedContent/service";
import { useAuthStore } from "@/store/auth-store";
import { AuthLayoutRoute } from "@/router";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import { getClientURL } from "@/common/urls";
import PageLoading from "@/components/PageLoading.vue";
const router = useRouter();
const languageStore = useLanguageStore();
const authStore = useAuthStore();

const userInfo = ref<Partial<User>>({});

async function find(_query: string) {
	const userProfile = await AuthService.me();

	if (userProfile) {
		userInfo.value = {
			...userProfile,
			photo: userProfile?.photo ? `data:image/jpeg;base64,${userProfile?.photo}` : null,
		};
	}
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

/**
 * Redirect to the empty page.
 */
function onDashboard(): void {
	router.push({ name: "PAGE_EMPTY" });
}
</script>

<template>
	<PageLoading :find="find">
		<q-layout view="hHh Lpr lff" class="modern-layout">
			<q-header class="clean-header">
				<q-toolbar class="h-18 gap-x-3 clean-toolbar">
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

			<!-- Main Content -->
			<q-page-container>
				<q-page class="modern-page">
					<div class="page-container">
						<!-- Profile Section -->
						<div class="profile-section">
							<div class="profile-card">
								<!-- User Avatar & Info -->
								<div class="user-header">
									<div class="user-info">
										<h1 class="user-name">
											{{ authStore.user?.lastName }}
											{{ authStore.user?.firstName }}
											{{ authStore.user?.middleName }}
										</h1>
									</div>
								</div>

								<!-- User Details Grid -->
								<div class="details-grid">
									<div class="detail-item">
										<q-icon name="wc" />
										<div class="detail-content">
											<span class="detail-label">{{ $tl("gender") }}</span>
											<span class="detail-value">
												{{ $tl(authStore.user?.telegramUsername || "N/A") }}
											</span>
										</div>
									</div>
									<div class="detail-item">
										<q-icon name="cake" />
										<div class="detail-content">
											<span class="detail-label">
												{{ $tl("dateOfBirth") }}
											</span>
											<span class="detail-value">
												{{ authStore.user?.dateOfBirth }}
											</span>
										</div>
									</div>
								</div>

								<!-- Quick Actions -->
								<div class="quick-actions">
									<q-btn
										@click="onDashboard"
										icon="dashboard"
										:label="$tl('go_to_dashboard')"
										color="primary"
										class="action-card"
										no-caps
									/>
								</div>
							</div>
						</div>
					</div>
				</q-page>
			</q-page-container>
		</q-layout>
	</PageLoading>
</template>

<style lang="scss" scoped>
@media screen and (max-width: 400px) {
	.adapt-padding {
		padding: 8px !important;
	}
}

.display-none {
	display: none;
}

/* Header aligned with BaseLayout */
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

/* Modern Layout */
.modern-layout {
	min-height: 100vh;
	background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
}

/* Main Page */
.modern-page {
	background: transparent;
	padding: 2rem;
}

.page-container {
	max-width: 1200px;
	margin: 0 auto;
	display: grid;
	gap: 2rem;
	grid-template-columns: 400px 1fr;

	@media (max-width: 1024px) {
		grid-template-columns: 1fr;
		max-width: 600px;
	}
}

/* Profile Section */
.profile-section {
	display: flex;
	flex-direction: column;
}

.profile-card {
	background: rgba(255, 255, 255, 0.95);
	backdrop-filter: blur(10px);
	border-radius: 24px;
	border: 1px solid rgba(0, 0, 0, 0.05);
	box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
	padding: 2rem;
	transition: all 0.3s ease;

	&:hover {
		transform: translateY(-4px);
		box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
	}
}

.user-header {
	display: flex;
	align-items: center;
	gap: 1.5rem;
	margin-bottom: 2rem;
}

.user-info {
	flex: 1;
}

.user-name {
	font-size: 1.5rem;
	font-weight: 700;
	color: #1e293b;
	margin: 0 0 0.5rem 0;
	line-height: 1.2;
}

.details-grid {
	display: grid;
	gap: 1rem;
	margin-bottom: 2rem;
}

.detail-item {
	display: flex;
	align-items: center;
	gap: 1rem;
	padding: 1rem;
	background: rgba(248, 250, 252, 0.8);
	border-radius: 16px;
	border: 1px solid rgba(0, 0, 0, 0.03);
	transition: all 0.3s ease;

	&:hover {
		background: rgba(248, 250, 252, 1);
		transform: translateX(4px);
	}

	.q-icon {
		color: $primary;
		font-size: 1.25rem;
	}
}

.detail-content {
	display: flex;
	flex-direction: column;
	gap: 0.25rem;
}

.detail-label {
	font-size: 0.875rem;
	color: #64748b;
	font-weight: 500;
}

.detail-value {
	font-weight: 600;
	color: #1e293b;
}

.quick-actions {
	display: flex;
	flex-direction: column;
	gap: 0.75rem;
}

.action-card {
	border-radius: 16px !important;
	padding: 1rem 1.5rem !important;
	font-weight: 600 !important;
	transition: all 0.3s ease !important;
	justify-content: start !important;
	text-transform: none !important;
	text-align: start !important;

	&:hover {
		transform: translateY(-2px);
		box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
	}

	&:active {
		transform: translateY(0) !important;
	}
}
</style>
