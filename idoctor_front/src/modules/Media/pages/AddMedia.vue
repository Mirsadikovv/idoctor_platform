<script lang="ts" setup>
import { formRequired } from "@/common/validator";
import FilePicker from "@/components/quasar/form/FilePicker.vue";
import Form from "@/components/quasar/form/Form.vue";
import Title from "@/components/Title.vue";
import { type keySignType } from "../utils/utils";
import { type MediaPartialType, type MediaType, MediaService } from "../service/media.service";
import { ref } from "vue";
import type { IdType } from "@/service";
import { useQuasar } from "quasar";
import { useLanguageStore } from "@/store/language-store";
import Button from "@/components/quasar/btn/Button.vue";

export interface Props {
	modalClose: () => void;
	ID: number;
	ownerCategoty: keySignType;
	categories: string[];
	validation?: (file: File) => Promise<boolean>;
}

const { modalClose, ID, ownerCategoty, categories, validation } = defineProps<Props>();
const emit = defineEmits<{
	(e: "upload", data: IdType, file: File): void;
}>();

let model = ref<MediaPartialType>({});
const file = ref<File>();
const $q = useQuasar();
const { tl } = useLanguageStore();

async function save(newData: MediaType) {
	// Создаем объект FormData
	if (validation) {
		const isValid = await validation(file.value!);
		if (!isValid) {
			$q.notify({
				type: "negative",
				position: "top",
				message: `${tl("file_validation_error")}`,
			});
			return false;
		}
	}
	const formData = new FormData();
	formData.append("file", file.value!);
	formData.append("sign", newData.sign!);

	let ownerCat = ownerCategoty;
	const result = await MediaService.create(ownerCat, ID, formData);
	if (!result) {
		return false;
	}
	emit("upload", result, file.value!);
	modalClose();
	return true;
}
</script>
<template>
	<Form :model-value="model" :save="save" :bordered="false">
		<template #title>
			<Title class="mb-5">{{ $tl("add_media") }} </Title>
		</template>
		<template #sign="{ model }">
			<q-select
				v-model="model.sign"
				:label="$tl('sign')"
				:options="categories"
				:option-label="(option) => $tl(option)!"
				new-value-mode="add-unique"
				class="col-12"
				outlined
				use-input
			>
			</q-select>
		</template>

		<template #file>
			<FilePicker
				class="col-12"
				v-model="file"
				label="file"
				:rules="[formRequired($tl('this_field_is_required'))]"
			/>
		</template>

		<template #actions="{ loading }">
			<div class="row my-row">
				<div class="col">
					<Button ui-type="danger" v-close-popup class="w-full!">
						{{ $tl("no") }}
					</Button>
				</div>
				<div class="col">
					<Button :loading="loading" type="submit" class="w-full!">
						{{ $tl("save") }}
					</Button>
				</div>
			</div>
		</template>
	</Form>
</template>
