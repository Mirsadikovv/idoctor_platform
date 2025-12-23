<script setup lang="ts">
import { ref } from "vue";
import Select from "@/components/quasar/form/Select.vue";
import { useLanguageStore } from "@/store/language-store";
const { tl } = useLanguageStore();
const status = ref(false);
const statuses = [
	{ value: true, label: tl("active") },
	{ value: false, label: tl("inactive") },
];
</script>
<template>
	<Select :options="statuses" v-model="status" map-options emit-value></Select>
</template>
