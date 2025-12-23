<script lang="ts" setup>
import { formRequired } from "@/common/validator";
import FilePicker from "@/components/quasar/form/FilePicker.vue";
import Form from "@/components/quasar/form/Form.vue";
import Title from "@/components/Title.vue";
import { type keySignType, signOptions } from "../utils/utils";
import { type MediaPartialType, type MediaType, MediaService } from "../service/media.service";
import { ref } from "vue";
import Button from "@/components/quasar/btn/Button.vue";

export interface Props {
	modalClose: () => void;
	ID: number;
	ownerCategoty: keySignType;
	sign: string;
}

const { modalClose, ID, ownerCategoty, sign } = defineProps<Props>();

let model = ref<MediaPartialType>({});
const file = ref<File>();

model.value.sign = sign;

async function save(newData: MediaType) {
	// Создаем объект FormData
	const formData = new FormData();
	formData.append("file", file.value!);
	formData.append("sign", newData.sign!);

	const result = await MediaService.update(ID, formData);
	if (!result) {
		return false;
	}
	modalClose();
	return true;
}
</script>
<template>
	<Form :model-value="model" :save="save" :bordered="false">
		<template #title>
			<Title class="mb-5">{{ $tl("edit_media") }} </Title>
		</template>

		<template #sign="{ model }">
			<q-select
				v-model="model.sign"
				:label="$tl('sign')"
				:options="signOptions(ownerCategoty)"
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
