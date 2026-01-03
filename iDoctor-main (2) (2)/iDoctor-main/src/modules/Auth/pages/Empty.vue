<template>
	<div class="telegram-empty-page">
		<div class="telegram-content">
			<!-- Animated Dashboard Icon -->
			<div class="telegram-icon-wrapper">
				<q-icon name="dashboard" class="telegram-main-icon text-blue-500 animate-pulse" />
				<div class="telegram-icon-background"></div>
			</div>

			<!-- Welcome Text -->
			<div class="telegram-text-content">
				<h1
					class="text-2xl sm:text-3xl lg:text-4xl font-bold text-gray-800 mb-4 leading-tight text-center"
				>
					{{ $tl("welcome_to_dashboard") }}
				</h1>
				<p
					class="text-base sm:text-lg text-gray-600 leading-relaxed text-center max-w-md mx-auto"
				>
					{{ $tl("empty_page_description") }}
				</p>
			</div>
		</div>
	</div>
</template>

<style scoped>
.telegram-empty-page {
	height: calc(var(--tg-viewport-height, 100vh) - 74px);
	min-height: calc(100dvh - 74px);
	display: flex;
	align-items: center;
	justify-content: center;
	padding: max(env(safe-area-inset-top, 0), 1rem) max(env(safe-area-inset-right, 0), 1rem)
		max(env(safe-area-inset-bottom, 0), 1rem) max(env(safe-area-inset-left, 0), 1rem);
}

.telegram-content {
	text-align: center;
	max-width: 500px;
	width: 100%;
	padding: 1.5rem;
}

.telegram-icon-wrapper {
	position: relative;
	display: inline-block;
	margin-bottom: 2rem;
}

.telegram-main-icon {
	font-size: 5rem;
	position: relative;
	z-index: 2;
}

.telegram-icon-background {
	position: absolute;
	top: 50%;
	left: 50%;
	width: 120px;
	height: 120px;
	background: radial-gradient(circle, rgba(59, 130, 246, 0.1) 0%, transparent 70%);
	border-radius: 50%;
	transform: translate(-50%, -50%);
	animation: iconGlow 3s ease-in-out infinite alternate;
}

@keyframes iconGlow {
	0% {
		transform: translate(-50%, -50%) scale(1);
		opacity: 0.5;
	}
	100% {
		transform: translate(-50%, -50%) scale(1.1);
		opacity: 0.8;
	}
}

@media (max-width: 768px) {
	.telegram-empty-page {
		height: calc(var(--tg-viewport-height, 100vh) - 74px);
		padding: 1rem;
	}
	.telegram-content {
		padding: 1rem;
	}
	.telegram-main-icon {
		font-size: 4rem;
	}
	.telegram-icon-background {
		width: 100px;
		height: 100px;
	}
}

@media (max-width: 480px) {
	.telegram-empty-page {
		height: calc(var(--tg-viewport-height, 100vh) - 74px);
		padding: 0.75rem;
	}
	.telegram-content {
		padding: 0.75rem;
	}
	.telegram-main-icon {
		font-size: 3.5rem;
	}
}

@media (max-width: 360px) {
	.telegram-main-icon {
		font-size: 3rem;
	}
	.telegram-icon-background {
		width: 80px;
		height: 80px;
	}
}

@media (prefers-reduced-motion: reduce) {
	.telegram-icon-background {
		animation: none;
		opacity: 0.6;
	}
}
</style>
