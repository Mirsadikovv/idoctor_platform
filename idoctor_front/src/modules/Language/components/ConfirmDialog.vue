<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { LanguageService, type LanguagePartialType } from "../../Language/service";
import { inject, ref } from "vue";
import Button from "@/components/quasar/btn/Button.vue";

interface Props {
	id: number;
	fetch: Function;
	isRemove?: boolean;
}

const { fetch, id, isRemove } = defineProps<Props>();

const onClose = inject<() => Promise<void>>("onClose");

const LanguageModel = ref<LanguagePartialType>({});

async function confirm() {
	if (isRemove) {
		await LanguageService.delete(id);
	} else {
		await LanguageService.restore(id);
	}
	await onClose?.();
	fetch();
	return true;
}
</script>

<template>
	<Form v-model="LanguageModel" :save="confirm">
		<template #title>
			<div class="mb-2 text-center! w-full! text-2xl text-primary font-700">
				{{ $tl(isRemove ? "do_you_want_remove_lang" : "do_you_want_restore_lang") }}?
			</div>
		</template>
		<template #actions="{ loading }">
			<div class="flex! w-100%! mx-auto gap-x-2 mt-3">
				<Button ui-type="danger" class="w-full! bg-negative" color="negative" v-close-popup>
					{{ $tl("cancel") }}
				</Button>

				<Button :loading="loading" type="submit" class="w-full!">
					{{ $tl("confirm") }}
				</Button>
			</div>
		</template>
	</Form>
</template>
