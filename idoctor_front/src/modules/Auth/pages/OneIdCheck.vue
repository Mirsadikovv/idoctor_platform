<script setup lang="ts">
import WelcomeBg from "@/assets/Bg_Main_Mini_v2.jpg";
import Logo from "@/assets/LogoIcon.svg";
import { AuthService } from "@/service";
import { useRoute } from "vue-router";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/store/auth-store";
import { normalaizeRoute } from "@/plugins/router.plugin";
import { normalaizeLanguage } from "@/plugins/router.plugin";
import { useLanguageStore } from "@/store/language-store";

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const laguageStore = useLanguageStore();

async function signIn(code: string) {
	await normalaizeRoute(router);

	await normalaizeLanguage();

	await router.push({
		name: "PAGE_PROFILE",
		params: {
			lang: laguageStore.currentLang?.name,
		},
	});

	return true;
}

if (route.query.code) {
	signIn(route.query.code as string);
}
</script>

<template>
	<div class="relative">
		<!-- Тёмная прозрачная подложка -->
		<div class="fixed inset-0 z--1 bg-black/40 backdrop-blur"></div>

		<!-- Фоновое изображение -->
		<img
			class="fixed inset-0 w-screen h-screen object-cover z--2"
			:src="WelcomeBg"
			alt="WelcomeBg"
		/>

		<div class="flex items-center justify-center h-screen px-4 max-w-500px mx-auto">
			<!-- Контейнер с содержимым -->
			<div
				class="relative flex flex-col items-center justify-center gap-4 p-10 backdrop-blur backdrop-filter rounded-2xl"
			>
				<!-- Логотип -->
				<img class="mx-auto h-[200px]" :src="Logo" alt="logo" />

				<!-- Контент -->
				<div
					class="absolute z--1 top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2"
				>
					<q-spinner color="secondary" size="440px" :thickness="0.5" />
				</div>
			</div>
		</div>
	</div>
</template>
