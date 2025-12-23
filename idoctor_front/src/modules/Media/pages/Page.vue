<script setup lang="ts">
import PageLoading from "@/components/PageLoading.vue";
import Table from "@/components/quasar/table/Table.vue";
import AddMedia from "./AddMedia.vue";
import Title from "@/components/Title.vue";
import Expasion from "@/components/quasar/Expasion.vue";

import { dateFormat } from "@/common";
import { signOptions, type keySignType } from "../utils/utils";
import {
	type FileQueryType,
	type MediaPageDataType,
	type MediaPartialType,
	MediaService,
} from "@/modules/Media/service/media.service";
import { computed, ref } from "vue";
import EditMedia from "./EditMedia.vue";
import type { IdType } from "@/service";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";

const baseURl = import.meta.env.VITE_API_URL;

export interface Props {
	ownerId: number | string;
	ownerCategoty: keySignType;
	defaultOpen?: boolean;
	label?: string;
	isEditable?: boolean;
	isDeletable?: boolean;
	isCreatable?: boolean;
	uniqueFiles?: string[];
	options?: string[];
	clazz?: string;
	validation?: (file: File) => Promise<boolean>;
	fetch?: (...args: any) => Promise<MediaPageDataType>;
}

const {
	ownerId,
	ownerCategoty,
	isDeletable = true,
	isEditable = true,
	isCreatable = true,
	uniqueFiles,
	options,
	clazz = "",
	validation,
	fetch = MediaService.page,
} = defineProps<Props>();

const emit = defineEmits<{
	(e: "upload", data: IdType, file: File): void;
	(e: "delete", data: IdType): void;
}>();

let models = ref<MediaPageDataType>({
	data: [],
	currentPage: 0,
	pageSize: 0,
	totalPages: 0,
	totalRows: 0,
});

const pick = {
	id: false,
	edit: false,
	name: true,
	status: true,
};
const pikers = ref({});

let createModal = ref(false);
let viewModal = ref(false);
const editModal = ref(false);
let image = ref<number>(0);
let confirm = ref(false);
let choosed = ref<MediaPartialType>({});

let selectOptions = options ? options : signOptions(ownerCategoty);

function modalClose(fn: () => Promise<void>) {
	createModal.value = false;
	editModal.value = false;
	fn();
}

async function find(_query: string) {
	const queryParams: FileQueryType = {
		owner: +ownerId,
		category: ownerCategoty,
	};
	const data = await fetch(queryParams);

	if (!data) {
		return;
	}
	models.value = data;
}
// async function view(id: number) {
// 	image.value = id;
// 	viewModal.value = true;
// }
function choose(model: MediaPartialType) {
	confirm.value = true;
	choosed.value = model;
}
async function edit(model: MediaPartialType) {
	choosed.value = model;
	editModal.value = true;
}

async function restore(model: MediaPartialType, fetch: () => void) {
	const id = model?.id as number;
	const res = await MediaService.delete(id);
	if (res.status === 204) {
		fetch();
		emit("delete", { id });
		return true;
	}
}

function downloadFile(fileId: number, filename: string) {
	const link = document.createElement("a");
	link.href = `${baseURl}file/download/${fileId}`;
	link.download = filename;
	link.target = "_blank";
	document.body.appendChild(link);
	link.click();
	document.body.removeChild(link);
}

const matchedUniqueSigns = computed(() => {
	return models.value.data
		.filter((file) => uniqueFiles?.includes(file.sign))
		.map((file) => file.sign);
});

const filteredCategories = computed(() => {
	return selectOptions?.filter((cat) => !matchedUniqueSigns.value.includes(cat));
});
</script>

<template>
	<Expasion
		:clazz="clazz ? clazz : 'mt-0'"
		:default-opened="defaultOpen"
		padding-off
		:label="$tl(label)"
		icon="photo_library"
	>
		<PageLoading :find="find" #="{ loading, fetch }">
			<div class="flex items-center justify-between no-wrap w-full px-3 py-2">
				<Title class="text-base"> {{ $tl("add_file") }} </Title>

				<q-btn
					v-if="isCreatable"
					round
					dense
					class="bg-secondary ml-5"
					color="white"
					icon="add"
					flat
					:disable="filteredCategories.length === 0"
					@click="createModal = true"
				/>
			</div>

			<Table :models="models" :loading="loading" has-order>
				<template #filename:thead="{ key }">{{ $tl(key) }}</template>
				<template #filename="{ model }">
					<a
						target="_blank"
						:href="`${baseURl}file/download/${model?.id}`"
						class="text-left! text-sm"
					>
						{{ model?.name || model?.filename }}
					</a>
				</template>

				<template #category:thead="{ key }">{{ $tl(key) }}</template>
				<template #category="{ model }">{{ $tl(model.category) }}</template>

				<template #createAt:thead="{ key }">{{ $tl(key) }}</template>
				<template #createAt="{ model }">
					{{ dateFormat(model?.createdAt, "HH:mm") }}
					|
					{{ dateFormat(model?.createdAt, "DD.MM.YYYY") }}
				</template>

				<template #sign:thead="{ key }">{{ $tl(key) }}</template>
				<template #sign="{ model }">
					{{ model?.sign ? $tl(model.sign) : "N/A" }}
				</template>

				<template #owner:thead="{ key }">{{ $tl(key) }}</template>
				<template #owner></template>

				<template #edit:thead>
					<div class="text-center">
						{{ $tl("action") }}
					</div>
				</template>
				<template #edit="{ model }">
					<div class="text-center">
						<!-- <q-btn
							:key="model.id"
							@click="view(model.id)"
							icon="visibility"
							color="secondary"
							padding="xs"
							size="12px"
							outline
							round
						/> -->
						<q-btn
							:key="model.id + '_download'"
							@click="downloadFile(model.id, model.filename)"
							class="bg-white text-secondary"
							icon="visibility"
							color="secondary"
							padding="xs"
							size="12px"
							outline
							round
						>
							<q-tooltip>{{ $tl("download") }}</q-tooltip>
						</q-btn>
						<q-btn
							v-if="isEditable"
							:key="model.id"
							class="bg-white text-secondary ml-3"
							@click="edit(model)"
							icon="edit"
							color="secondary"
							padding="xs"
							size="12px"
							outline
							round
						/>

						<q-btn
							v-if="isDeletable"
							:key="model.id"
							class="bg-white text-red-5 ml-3"
							@click="choose(model)"
							icon="delete"
							color="secondary"
							padding="xs"
							size="12px"
							outline
							round
						/>
					</div>
				</template>
				<template #tfoot="{ totalPages }">
					<TablePaginate
						v-model:pikers="pikers"
						:total="totalPages"
						:pick="pick"
						@page="fetch"
					/>
				</template>
			</Table>

			<q-dialog v-model="createModal" persistent>
				<q-card>
					<AddMedia
						:modalClose="() => modalClose(fetch)"
						:ID="+ownerId"
						:ownerCategoty="ownerCategoty"
						:categories="filteredCategories"
						:validation="validation"
						@upload="(data: IdType, file: File) => emit('upload', data, file)"
					/>
				</q-card>
			</q-dialog>
			<q-dialog v-model="editModal" persistent>
				<q-card>
					<EditMedia
						:modalClose="() => modalClose(fetch)"
						:ID="choosed.id!"
						:ownerCategoty="ownerCategoty"
						:sign="choosed.sign!"
					/>
				</q-card>
			</q-dialog>

			<q-dialog v-model="viewModal">
				<q-card class="sm:w-400px w-200px">
					<q-card-section>
						<q-img :src="`${baseURl}file/download/${image}`" alt="image" />
					</q-card-section>
				</q-card>
			</q-dialog>
			<q-dialog v-model="confirm" persistent>
				<q-card>
					<q-card-section class="row items-center text-lg">
						<span>
							{{ $tl("are_u_want_delete") }}
						</span>
					</q-card-section>

					<q-card-actions align="center" class="flex no-wrap gap-x-2">
						<q-btn
							outline
							class="full-width"
							:label="$tl('no')"
							color="secondary"
							v-close-popup
						/>
						<q-btn
							unelevated
							class="full-width"
							:label="$tl('yes')"
							color="secondary"
							v-close-popup
							@click="restore(choosed!, fetch)"
						/>
					</q-card-actions>
				</q-card>
			</q-dialog>
		</PageLoading>
	</Expasion>
</template>
