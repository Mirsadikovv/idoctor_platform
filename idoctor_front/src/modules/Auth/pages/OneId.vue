<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/store/auth-store";
import { AuthService, type AuthUser } from "@/modules/Auth/service";
import Form from "@/components/quasar/form/Form.vue";
import Input from "@/components/quasar/form/Input.vue";
import Submit from "@/components/quasar/form/Submit.vue";
import { normalaizeLanguage, normalaizeRoute } from "@/plugins/router.plugin";
import { useLanguageStore } from "@/store/language-store";

const router = useRouter();
const authStore = useAuthStore();
const laguageStore = useLanguageStore();

const form = ref<AuthUser>({
	username: "",
	password: "",
});

const handleLogin = async (formData: AuthUser): Promise<boolean> => {
	try {
		const authData = await AuthService.signIn(formData);
		authStore.setToken(authData.token);

		const user = await AuthService.me();
		authStore.setUser(user);

		await normalaizeRoute(router);

		await normalaizeLanguage();

		await router.push({
			name: "PAGE_PROFILE",
			params: {
				lang: laguageStore.currentLang?.name,
			},
		});
		return true;
	} catch (error) {
		console.error("Login error:", error);
		return false;
	}
};
</script>

<template>
	<div class="auth-container">
		<!-- Анимированный фон -->
		<div class="animated-background">
			<div class="floating-orbs">
				<div class="orb orb-1"></div>
				<div class="orb orb-2"></div>
				<div class="orb orb-3"></div>
				<div class="orb orb-4"></div>
				<div class="orb orb-5"></div>
			</div>
			<div class="gradient-mesh"></div>
		</div>

		<div class="content-wrapper">
			<!-- Контейнер с содержимым -->
			<div class="login-card">
				<!-- Иконка приложения -->
				<div class="app-icon">
					<div class="icon-inner">
						<svg viewBox="0 0 24 24" fill="currentColor" class="medical-icon">
							<path
								d="M19 8h-2v3h-3v2h3v3h2v-3h3v-2h-3V8zM4 6H2v14c0 1.1.9 2 2 2h14v-2H4V6zm16-4H8c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm0 14H8V4h12v12z"
							/>
						</svg>
					</div>
				</div>

				<!-- Заголовок -->
				<div class="welcome-text">
					<h1>{{ $tl("iDoctor") }}</h1>
					<p>{{ $tl("welcome_to_idoctor") }}</p>
				</div>

				<!-- Форма авторизации -->
				<div class="form-container">
					<Form
						v-model="form"
						:save="handleLogin"
						:is-cleaned-style="true"
						:clazz="'bg-transparent border-none!'"
					>
						<template #login="{ model }">
							<Input
								v-model="model.username"
								label="login"
								class="col-12 login-input"
								:rules="[(val: string) => !!val || 'Поле обязательно для заполнения']"
							/>
						</template>

						<template #password="{ model }">
							<Input
								v-model="model.password"
								label="password"
								type="password"
								class="col-12 login-input"
								:rules="[(val: string) => !!val || 'Поле обязательно для заполнения']"
							/>
						</template>

						<template #actions="{ loading }">
							<Submit
								:loading="loading"
								no-caps
								:label="$tl('login')"
								class="login-button"
								unelevated
								rounded
							/>
						</template>
					</Form>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.auth-container {
	position: relative;
	width: 100vw;
	height: 100vh;
	overflow: hidden;
	display: flex;
	align-items: center;
	justify-content: center;
}

/* Анимированный фон */
.animated-background {
	position: fixed;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	background: linear-gradient(135deg, #1a237e 0%, #3949ab 50%, #5e35b1 100%);
	z-index: -2;
}

.gradient-mesh {
	position: absolute;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	background: radial-gradient(circle at 20% 30%, rgba(26, 35, 126, 0.4) 0%, transparent 50%),
		radial-gradient(circle at 80% 70%, rgba(57, 73, 171, 0.4) 0%, transparent 50%),
		radial-gradient(circle at 40% 80%, rgba(94, 53, 177, 0.4) 0%, transparent 50%);
	animation: meshShift 20s ease-in-out infinite;
}

@keyframes meshShift {
	0%,
	100% {
		transform: translateX(0) translateY(0) rotate(0deg);
	}
	33% {
		transform: translateX(-20px) translateY(-20px) rotate(120deg);
	}
	66% {
		transform: translateX(20px) translateY(-10px) rotate(240deg);
	}
}

.floating-orbs {
	position: absolute;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
}

.orb {
	position: absolute;
	border-radius: 50%;
	backdrop-filter: blur(10px);
	animation: float 15s infinite;
}

.orb-1 {
	width: 120px;
	height: 120px;
	top: 10%;
	left: 10%;
	background: linear-gradient(45deg, rgba(26, 35, 126, 0.2), rgba(26, 35, 126, 0.4));
	animation-delay: 0s;
	animation-duration: 20s;
}

.orb-2 {
	width: 80px;
	height: 80px;
	top: 20%;
	right: 20%;
	background: linear-gradient(45deg, rgba(57, 73, 171, 0.2), rgba(57, 73, 171, 0.4));
	animation-delay: 5s;
	animation-duration: 25s;
}

.orb-3 {
	width: 60px;
	height: 60px;
	bottom: 30%;
	left: 20%;
	background: linear-gradient(45deg, rgba(94, 53, 177, 0.2), rgba(94, 53, 177, 0.4));
	animation-delay: 10s;
	animation-duration: 18s;
}

.orb-4 {
	width: 100px;
	height: 100px;
	bottom: 20%;
	right: 15%;
	background: linear-gradient(45deg, rgba(63, 81, 181, 0.2), rgba(63, 81, 181, 0.4));
	animation-delay: 3s;
	animation-duration: 22s;
}

.orb-5 {
	width: 40px;
	height: 40px;
	top: 50%;
	left: 50%;
	background: linear-gradient(45deg, rgba(121, 134, 203, 0.2), rgba(121, 134, 203, 0.4));
	animation-delay: 8s;
	animation-duration: 16s;
}

@keyframes float {
	0%,
	100% {
		transform: translateX(0) translateY(0) scale(1);
		opacity: 0.3;
	}
	25% {
		transform: translateX(30px) translateY(-30px) scale(1.1);
		opacity: 0.6;
	}
	50% {
		transform: translateX(-20px) translateY(-60px) scale(0.9);
		opacity: 0.4;
	}
	75% {
		transform: translateX(-30px) translateY(20px) scale(1.05);
		opacity: 0.7;
	}
}

/* Контент */
.content-wrapper {
	position: relative;
	z-index: 10;
	width: 100%;
	max-width: 400px;
	padding: 1rem;
}

.login-card {
	background: rgba(255, 255, 255, 0.1);
	backdrop-filter: blur(20px);
	border: 1px solid rgba(255, 255, 255, 0.2);
	border-radius: 24px;
	padding: 3rem 2rem;
	box-shadow: 0 25px 45px rgba(0, 0, 0, 0.1);
	text-align: center;
	animation: slideUp 0.8s ease-out;
}

@keyframes slideUp {
	from {
		opacity: 0;
		transform: translateY(30px);
	}
	to {
		opacity: 1;
		transform: translateY(0);
	}
}

.app-icon {
	margin-bottom: 2rem;
}

.icon-inner {
	width: 80px;
	height: 80px;
	margin: 0 auto;
	background: linear-gradient(135deg, #1a237e, #3949ab);
	border-radius: 20px;
	display: flex;
	align-items: center;
	justify-content: center;
	box-shadow: 0 15px 35px rgba(26, 35, 126, 0.4);
	animation: iconPulse 2s ease-in-out infinite;
}

@keyframes iconPulse {
	0%,
	100% {
		transform: scale(1);
		box-shadow: 0 15px 35px rgba(26, 35, 126, 0.4);
	}
	50% {
		transform: scale(1.05);
		box-shadow: 0 20px 40px rgba(26, 35, 126, 0.6);
	}
}

.medical-icon {
	width: 40px;
	height: 40px;
	color: white;
}

.welcome-text {
	margin-bottom: 2.5rem;
}

.welcome-text h1 {
	font-size: 2rem;
	font-weight: 700;
	color: white;
	margin: 0 0 0.5rem 0;
	text-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
}

.welcome-text p {
	font-size: 1rem;
	color: rgba(255, 255, 255, 0.8);
	margin: 0;
	font-weight: 400;
}

.form-container {
	text-align: left;
}

:deep(.login-input .q-field__control) {
	background: rgba(255, 255, 255, 0.1) !important;
	border-radius: 12px !important;
	border: 1px solid rgba(255, 255, 255, 0.2) !important;
	backdrop-filter: blur(10px);
}

:deep(.login-input .q-field__control:hover) {
	background: rgba(255, 255, 255, 0.15) !important;
	border-color: rgba(255, 255, 255, 0.3) !important;
}

:deep(.login-input .q-field__control:focus-within) {
	background: rgba(255, 255, 255, 0.2) !important;
	border-color: rgba(255, 255, 255, 0.5) !important;
	box-shadow: 0 0 0 2px rgba(255, 255, 255, 0.1) !important;
}

:deep(.login-input .q-field__native) {
	color: white !important;
	font-weight: 500;
}

:deep(.login-input .q-field__label) {
	color: rgba(255, 255, 255, 0.7) !important;
	font-weight: 500;
}

:deep(.login-input .q-field__marginal) {
	color: rgba(255, 255, 255, 0.6) !important;
}

:deep(.login-button) {
	background: linear-gradient(135deg, #1a237e, #3949ab) !important;
	color: white !important;
	width: 100%;
	height: 48px;
	font-weight: 600 !important;
	font-size: 1rem !important;
	margin-top: 1rem !important;
	box-shadow: 0 15px 35px rgba(26, 35, 126, 0.4) !important;
	transition: all 0.3s ease !important;
}

:deep(.login-button:hover) {
	transform: translateY(-2px) !important;
	box-shadow: 0 20px 40px rgba(26, 35, 126, 0.6) !important;
}

:deep(.login-button:active) {
	transform: translateY(0) !important;
}

/* Мобильная адаптация */
@media (max-width: 768px) {
	.content-wrapper {
		max-width: 350px;
		padding: 0.5rem;
	}

	.login-card {
		padding: 2rem 1.5rem;
	}

	.welcome-text h1 {
		font-size: 1.75rem;
	}

	.icon-inner {
		width: 70px;
		height: 70px;
	}

	.medical-icon {
		width: 35px;
		height: 35px;
	}
}

@media (max-width: 480px) {
	.login-card {
		padding: 1.5rem 1rem;
		border-radius: 20px;
	}
}
</style>
