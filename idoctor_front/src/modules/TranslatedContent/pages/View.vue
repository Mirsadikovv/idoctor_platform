<script setup lang="ts">
import { useRouter } from "vue-router";
import { LanguageContentService, type LangContentType } from "@/service";
import { ref } from "vue";
import { useQuasar } from "quasar";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import UpdateTranslate from "../components/Update.vue";

export interface Props {
	id: string;
}

const router = useRouter();
const $q = useQuasar();
const { id } = defineProps<Props>();

const model = ref<LangContentType>();

async function fetchTranslatedContent() {
	const response = await LanguageContentService.getByKey(id);
	if (!response) return;
	model.value = response;
}

async function deleteTranslatedContent() {
	const response = await LanguageContentService.delete(id);
	if (!response) return false;

	router.push({ name: "TRANSLATED_CONTENT" });
	return true;
}

function confirmDelete() {
	$q.dialog({
		title: "Подтверждение",
		message: "Вы уверены, что хотите удалить этот перевод?",
		cancel: true,
		persistent: true,
	}).onOk(() => {
		deleteTranslatedContent();
	});
}
fetchTranslatedContent();
</script>

<template>
	<div class="flex! gap-x-4 items-center mb-6">
		<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
		<q-breadcrumbs>
			<q-breadcrumbs-el
				:label="$tl('translated_content_page')"
				icon="g_translate"
				:to="{ name: 'TRANSLATED_CONTENT' }"
			/>
		</q-breadcrumbs>
		<q-space />

		<div class="flex gap-2" v-if="id">
			<ButtonDialog
				:label="$tl('update_translate')"
				icon="edit"
				color="primary"
				:fetch="fetchTranslatedContent"
			>
				<UpdateTranslate :key-word="id" :fetch="fetchTranslatedContent" />
			</ButtonDialog>

			<q-btn icon="delete" color="negative" flat round @click="confirmDelete">
				<q-tooltip>{{ $tl("remove_translate") }}</q-tooltip>
			</q-btn>
		</div>
	</div>

	<div v-if="model?.contents" v-for="value in model?.contents" class="q-pa-md">
		<q-card class="q-mb-lg">
			<q-card-section>
				<div class="text-h6 q-mb-md flex items-center">
					<q-icon name="g_translate" class="q-mr-sm" color="purple" />
					{{ $tl("translated_content_details") }}
				</div>

				<q-list>
					<q-item>
						<q-item-section>
							<q-item-label overline>{{ $tl("language") }}</q-item-label>
							<q-item-label class="text-h6">{{ value?.languageId }}</q-item-label>
						</q-item-section>
					</q-item>

					<q-item>
						<q-item-section>
							<q-item-label overline>{{ $tl("key") }}</q-item-label>
							<q-item-label class="text-h6">{{ model.key }}</q-item-label>
						</q-item-section>
					</q-item>

					<q-item>
						<q-item-section>
							<q-item-label overline>{{ $tl("value") }}</q-item-label>
							<q-item-label>{{ value?.value || "-" }}</q-item-label>
						</q-item-section>
					</q-item>
				</q-list>
			</q-card-section>
		</q-card>
	</div>

	<div v-else class="flex justify-center q-pa-xl">
		<q-spinner color="primary" size="3em" />
	</div>
</template>
