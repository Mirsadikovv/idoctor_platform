<script setup lang="ts">
export interface Props {
	size?: string | number;
	color?: string;
	thickness?: number;
	text?: string;
	overlay?: boolean;
	fullscreen?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
	size: "60px",
	color: "primary",
	thickness: 4,
	text: "",
	overlay: false,
	fullscreen: false,
});
</script>

<template>
	<div
		:class="[
			'loader-container',
			{
				'loader-overlay': props.overlay,
				'loader-fullscreen': props.fullscreen,
			},
		]"
	>
		<div class="loader-content">
			<q-spinner-dots :size="props.size" :color="props.color" />
			<div v-if="props.text" class="loader-text">
				{{ $tl(props.text) }}
			</div>
		</div>
	</div>
</template>

<style lang="scss" scoped>
.loader-container {
	position: fixed;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
	z-index: 9999;

	&.loader-overlay {
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		transform: none;
		display: flex;
		align-items: center;
		justify-content: center;
		background: rgba(255, 255, 255, 0.9);
		backdrop-filter: blur(4px);
	}

	&.loader-fullscreen {
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		transform: none;
		display: flex;
		align-items: center;
		justify-content: center;
		background: rgba(255, 255, 255, 0.95);
		backdrop-filter: blur(8px);
		min-height: 100vh;
		z-index: 99999;
	}
}

.loader-content {
	display: flex;
	flex-direction: column;
	align-items: center;
	gap: 16px;
}

.loader-text {
	font-size: 14px;
	color: $primary;
	font-weight: 500;
	animation: pulse 2s ease-in-out infinite;
}

@keyframes pulse {
	0%,
	100% {
		opacity: 1;
	}
	50% {
		opacity: 0.6;
	}
}
</style>
