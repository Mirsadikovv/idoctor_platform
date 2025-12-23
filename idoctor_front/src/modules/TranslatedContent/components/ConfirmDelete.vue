<script setup lang="ts">
import { inject } from "vue";
import { LanguageContentService } from "@/service";

interface Props {
	id: number;
	keyValue: string;
	fetch: Function;
}

const { fetch, keyValue, id } = defineProps<Props>();

const onClose = inject<() => Promise<void>>("onClose");

async function confirm() {
	await LanguageContentService.delete(id, keyValue);
	await onClose?.();
	fetch();
	return true;
}
</script>

<template>
	<q-card class="min-w-[460px] rounded-3xl shadow-2xl bg-white border border-gray-100">
		<q-card-section class="text-center py-8 px-8">
			<div class="mb-6">
				<div
					class="w-20 h-20 mx-auto bg-red-100 rounded-full flex items-center justify-center shadow-inner"
				>
					<q-icon name="delete" size="2.5rem" class="text-red-500" />
				</div>
			</div>

			<div class="text-xl font-bold text-gray-800 mb-3">
				{{ $tl("delete_confirm") }}
			</div>
		</q-card-section>

		<q-separator class="bg-gray-100" />

		<q-card-actions class="px-8 py-6">
			<div class="flex w-full gap-4">
				<Button
					outline
					no-caps
					class="flex-1 py-3 rounded-xl font-medium border-2 border-gray-300 text-gray-700 hover:bg-gray-50 hover:border-gray-400 transition-all duration-200"
					:label="$tl('cancel') || $tl('no')"
					v-close-popup
				/>
				<Button
					ui-type="danger"
					class="flex-1 py-3 rounded-xl font-medium bg-red-500 text-white hover:bg-red-600 shadow-lg hover:shadow-xl transition-all duration-200"
					:label="$tl('delete') || $tl('yes')"
					@click="confirm"
				/>
			</div>
		</q-card-actions>
	</q-card>
</template>
