<script setup lang="ts"></script>

<template>
	<div class="empty-page">
		<div class="content">
			<q-icon name="dashboard" class="main-icon" />
			<h1>{{ $tl("welcome_to_dashboard") }}</h1>
			<p>{{ $tl("empty_page_description") }}</p>
		</div>
	</div>
</template>

<style scoped>
.empty-page {
	height: calc(100vh - 150px);
	display: flex;
	align-items: center;
	justify-content: center;
	padding: 2rem;
}

.content {
	text-align: center;
	max-width: 500px;
}

.main-icon {
	font-size: 80px;
	color: #3b82f6;
	margin-bottom: 1rem;
}

h1 {
	font-size: 2rem;
	font-weight: 600;
	color: #1f2937;
	margin-bottom: 1rem;
}

p {
	color: #6b7280;
	font-size: 1.1rem;
	line-height: 1.5;
}
</style>
