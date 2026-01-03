<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from "vue";
import { useTemplateRef } from "vue";
import { useQuasar } from "quasar";

const fsRef = useTemplateRef<HTMLElement>("fsRef");
const fsIcon = ref<"fullscreen" | "fullscreen_exit">("fullscreen");
const isFs = ref<boolean>(false);
const quasar = useQuasar();

const toggleFs = () => {
	try {
		const element = fsRef.value;
		if (!element) return;

		element.classList.toggle("p-4");
		quasar.fullscreen.toggle(element);

		isFs.value = !isFs.value;
		fsIcon.value = isFs.value ? "fullscreen_exit" : "fullscreen";
	} catch (error) {
		console.error("Fullscreen xatosi:", error);
	}
};

const handleEsc = (event: KeyboardEvent) => {
	if (event.key === "Escape" && isFs.value) {
		toggleFs();
	}
};

onMounted(() => {
	document.addEventListener("keydown", handleEsc);
});

onBeforeUnmount(() => {
	document.removeEventListener("keydown", handleEsc);
});
</script>

<template>
	<div ref="fsRef" class="">
		<slot v-bind="{ isFs, toggleFs, fsIcon }"></slot>
	</div>
</template>
