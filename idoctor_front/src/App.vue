<template>
	<RouterView />
</template>
<script setup lang="ts">
import "@unocss/reset/tailwind.css";
</script>

<style>
html,
body {
	overscroll-behavior: none;
	touch-action: manipulation;
}
</style>
