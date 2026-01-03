<template>
	<div class="loading-skeleton">
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
</template>

<style scoped>
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

@media (prefers-reduced-motion: reduce) {
	.skeleton-animate {
		animation: none;
	}
}
</style>