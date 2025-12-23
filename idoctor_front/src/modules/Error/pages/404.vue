<script setup lang="ts">
import NotFound from "@/assets/Na_Nov_26.png";
import { computed } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
// Анимированный текст 404
const animatedNumbers = computed(() => ["4", "0", "4"]);
</script>

<template>
	<div class="error-page-container">
		<!-- Фоновые элементы -->
		<div class="background-pattern"></div>
		<div class="floating-shapes">
			<div class="shape shape-1"></div>
			<div class="shape shape-2"></div>
			<div class="shape shape-3"></div>
			<div class="shape shape-4"></div>
		</div>

		<!-- Основной контент -->
		<div class="content-wrapper">
			<!-- Анимированная большая надпись 404 -->
			<div class="error-number-container">
				<div
					v-for="(num, index) in animatedNumbers"
					:key="index"
					class="error-number"
					:style="{ animationDelay: `${index * 0.2}s` }"
				>
					{{ num }}
				</div>
			</div>

			<!-- Заголовок и описание -->
			<div class="text-content">
				<h1 class="error-title">
					{{ $tl("page_not_found") }}
				</h1>
				<p class="error-description">
					{{ $tl("info_about_not_found") }}
				</p>
			</div>

			<!-- Иллюстрация -->
			<div class="image-container">
				<div class="image-background"></div>
				<img class="error-image" :src="NotFound" alt="404 Illustration" />
			</div>

			<!-- Кнопки действий -->
			<div class="actions-container">
				<q-btn @click="router.back()" class="secondary-btn" outline rounded size="lg">
					<q-icon name="arrow_back" class="mr-2" />
					{{ $tl("go_back") }}
				</q-btn>

				<q-btn
					:to="{ name: 'LOGIN_AUTH' }"
					class="primary-btn"
					unelevated
					rounded
					size="lg"
				>
					<q-icon name="home" class="mr-2" />
					{{ $tl("back_to_login_page") }}
				</q-btn>
			</div>
		</div>
	</div>
</template>

<style scoped>
.error-page-container {
	height: 100vh;
	display: flex;
	align-items: center;
	justify-content: center;
	position: relative;
	overflow: hidden;
	background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	padding: 1rem;
}

/* Фоновый паттерн */
.background-pattern {
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	background-image: radial-gradient(
			circle at 25% 25%,
			rgba(255, 255, 255, 0.1) 2px,
			transparent 2px
		),
		radial-gradient(circle at 75% 75%, rgba(255, 255, 255, 0.1) 2px, transparent 2px);
	background-size: 60px 60px;
	animation: patternMove 20s linear infinite;
}

@keyframes patternMove {
	0% {
		transform: translate(0, 0);
	}
	100% {
		transform: translate(60px, 60px);
	}
}

/* Плавающие фигуры */
.floating-shapes {
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	pointer-events: none;
}

.shape {
	position: absolute;
	border-radius: 50%;
	background: rgba(255, 255, 255, 0.1);
	animation: float 6s ease-in-out infinite;
}

.shape-1 {
	width: 100px;
	height: 100px;
	top: 10%;
	left: 10%;
	animation-delay: 0s;
}

.shape-2 {
	width: 60px;
	height: 60px;
	top: 20%;
	right: 20%;
	animation-delay: 2s;
}

.shape-3 {
	width: 80px;
	height: 80px;
	bottom: 30%;
	left: 15%;
	animation-delay: 4s;
}

.shape-4 {
	width: 120px;
	height: 120px;
	bottom: 10%;
	right: 10%;
	animation-delay: 1s;
}

@keyframes float {
	0%,
	100% {
		transform: translateY(0) rotate(0deg);
	}
	50% {
		transform: translateY(-20px) rotate(180deg);
	}
}

/* Основной контент */
.content-wrapper {
	position: relative;
	z-index: 10;
	text-align: center;
	max-width: 800px;
	width: 100%;
}

/* Анимированные числа 404 */
.error-number-container {
	display: flex;
	justify-content: center;
	gap: 1rem;
	margin-bottom: 1.5rem;
}

.error-number {
	font-size: clamp(4rem, 12vw, 8rem);
	font-weight: 900;
	color: white;
	text-shadow: 0 0 30px rgba(255, 255, 255, 0.5);
	animation: bounce 2s ease-in-out infinite;
	background: linear-gradient(45deg, #fff, #f0f0f0);
	-webkit-background-clip: text;
	-webkit-text-fill-color: transparent;
	background-clip: text;
}

@keyframes bounce {
	0%,
	20%,
	50%,
	80%,
	100% {
		transform: translateY(0);
	}
	40% {
		transform: translateY(-20px);
	}
	60% {
		transform: translateY(-10px);
	}
}

/* Текстовый контент */
.text-content {
	margin-bottom: 2rem;
}

.error-title {
	font-size: clamp(2rem, 5vw, 3rem);
	font-weight: 700;
	color: white;
	margin-bottom: 1rem;
	text-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
}

.error-description {
	font-size: clamp(1rem, 2.5vw, 1.25rem);
	color: rgba(255, 255, 255, 0.9);
	line-height: 1.6;
	max-width: 600px;
	margin: 0 auto;
}

/* Контейнер изображения */
.image-container {
	position: relative;
	margin-bottom: 2rem;
	display: flex;
	justify-content: center;
}

.image-background {
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
	width: 120%;
	height: 120%;
	background: radial-gradient(circle, rgba(255, 255, 255, 0.1) 0%, transparent 70%);
	border-radius: 50%;
	animation: pulse 3s ease-in-out infinite;
}

@keyframes pulse {
	0%,
	100% {
		transform: translate(-50%, -50%) scale(1);
	}
	50% {
		transform: translate(-50%, -50%) scale(1.1);
	}
}

.error-image {
	max-width: 100%;
	max-height: 300px;
	width: auto;
	height: auto;
	position: relative;
	z-index: 2;
	filter: drop-shadow(0 10px 30px rgba(0, 0, 0, 0.3));
}

/* Кнопки действий */
.actions-container {
	display: flex;

	gap: 1rem;
	align-items: center;
	justify-content: center;
}

.primary-btn {
	background: linear-gradient(135deg, #4f46e5 0%, #7c3aed 100%) !important;
	color: white !important;
	min-width: 200px;
	height: 48px;
	font-weight: 600;
	box-shadow: 0 8px 25px rgba(79, 70, 229, 0.4) !important;
	transition: all 0.3s ease !important;
}

.primary-btn:hover {
	transform: translateY(-2px) !important;
	box-shadow: 0 12px 35px rgba(79, 70, 229, 0.6) !important;
}

.secondary-btn {
	border: 2px solid rgba(255, 255, 255, 0.8) !important;
	color: white !important;
	min-width: 200px;
	height: 48px;
	font-weight: 600;
	backdrop-filter: blur(10px);
	background: rgba(255, 255, 255, 0.1) !important;
	transition: all 0.3s ease !important;
}

.secondary-btn:hover {
	background: rgba(255, 255, 255, 0.2) !important;
	border-color: white !important;
	transform: translateY(-2px) !important;
}

/* Мобильная адаптация */
@media (max-width: 768px) {
	.error-page-container {
		padding: 1rem;
	}

	.error-number-container {
		gap: 0.5rem;
		margin-bottom: 1.5rem;
	}

	.text-content {
		margin-bottom: 2rem;
	}

	.image-container {
		margin-bottom: 2rem;
	}

	.error-image {
		max-height: 250px;
	}

	.actions-container {
		gap: 0.75rem;
	}

	.primary-btn,
	.secondary-btn {
		min-width: 280px;
		width: 100%;
		max-width: 300px;
	}
}

@media (max-width: 480px) {
	.error-number-container {
		margin-bottom: 1rem;
	}

	.error-image {
		max-height: 200px;
	}

	.floating-shapes {
		display: none;
	}
}
</style>
