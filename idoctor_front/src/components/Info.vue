<script setup lang="ts">
export interface Props {
	label: string;
	value: string | number;
	error?: boolean;
}
defineProps<Props>();
</script>
<template>
	<div class="info" :class="{ 'info-error': error }">
		<span class="info-label">
			{{ $tl(label) }}
		</span>
		<span class="info-value">
			{{ value }}
		</span>
	</div>
</template>

<style lang="scss">
.info {
	height: 56px;
	border-radius: 4px;
	padding: 0 12px;
	border: 1px dashed rgba(0, 0, 0, 0.24);
	display: flex;
	flex-direction: column;
	align-items: flex-start;
	justify-content: flex-start;
	align-content: flex-start;
	position: relative;
	padding-top: 9px;
	cursor: text;
	transition: all 0.2s ease-in-out;
	overflow: hidden;
	&:not(.info-error):hover {
		border-color: #000;
	}
	&-label {
		color: rgba(0, 0, 0, 0.6);
		font-size: 12px;
		line-height: 1.25;
		font-weight: 400;
		letter-spacing: 0.00937em;
	}
	&-value {
		color: rgba(0, 0, 0, 0.87);
		white-space: nowrap;
		overflow: hidden;
		display: -webkit-box;
		-webkit-line-clamp: 1; /* number of lines to show */
		line-clamp: 1;
		-webkit-box-orient: vertical;
		width: 100%;
		text-overflow: ellipsis;
		white-space: nowrap;
		overflow: hidden;
	}
	&-error {
		border-color: $negative;
	}
}
</style>
