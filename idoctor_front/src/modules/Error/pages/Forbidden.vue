<script setup lang="ts">
import NotFound from "@/assets/Na_Nov_26.png";
</script>

<template>
	<div class="flex flex-col h-full justify-center items-center">
		<div
			class="flex! w-full max-w-[900px] mx-auto bg-[#d4dfff] text-[#1a3d9e] p-4 text-base rounded-xl mt-10"
			style="border: 2px solid #6780e4"
		>
			<q-icon name="info" size="24px" color="#d4dfff" class="cursor-pointer mr-3"></q-icon>
			{{ $tl("info_about_forbidden") }}
		</div>
		<img class="max-w-[1000px]" :src="NotFound" alt="404" />

		<q-btn :to="{ name: 'LOGIN_AUTH' }" class="bg-accent text-white mt-4" flat>
			{{ $tl("back_to_login_page") }}
		</q-btn>
	</div>
</template>
