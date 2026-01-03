<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { UserService, type UserPartial } from "@/service";

import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import EditUser from "./Edit.vue";
import ConfirmDialog from "../components/ConfirmDialog.vue";

export interface Props {
	id: number | string;
}
const { id } = defineProps<Props>();

const router = useRouter();

let model = ref<UserPartial>({});

const loadUser = async () => {
	const data = await UserService.findByID(+id);
	model.value = data;
};

const fetchUser = () => {
	loadUser();
};

loadUser();
</script>

<template>
	<div class="flex! gap-x-4 items-center mb-6">
		<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
		<q-breadcrumbs>
			<q-breadcrumbs-el
				:label="$tl('user_list')"
				:to="{ name: 'PAGE_USER' }"
				icon="article"
			/>
			<q-breadcrumbs-el :label="$tl('page_for_view')" />
		</q-breadcrumbs>
		<q-space />
		
		<!-- Кнопки действий -->
		<div class="flex gap-2" v-if="model.id">
			<ButtonDialog
				label="edit_user"
				icon="edit"
				color="primary"
				:fetch="fetchUser"
			>
				<EditUser :id="model.id" :fetch="fetchUser" />
			</ButtonDialog>
			
			<IconDialog
				iconColor="negative"
				icon="delete"
				tooltipText="delete_user"
				withTooltip
			>
				<ConfirmDialog :fetch="fetchUser" :id="model.id" :isRemove="true" />
			</IconDialog>
		</div>
	</div>
	<div class="bg-secondary text-white p-4 mb-4 flex justify-between items-center rounded">
		<div class="text-xl font-bold">
			{{ $tl("fullName") }}: {{ model?.lastName }} {{ model?.firstName }}
			{{ model?.middleName }}
		</div>
		<div class="text-xl font-bold">
			{{ $tl("gender") }}:
			<q-chip color="white" outline>{{ $tl(model?.gender) }} </q-chip>
		</div>
	</div>

	<q-list bordered class="rounded-borders">
		<q-item>
			<q-item-section>
				<q-item-label class="font-bold">{{ $tl("id") }}</q-item-label>
				<q-item-label caption>{{ model?.id }}</q-item-label>
			</q-item-section>
		</q-item>

		<q-separator />

		<q-item>
			<q-item-section>
				<q-item-label class="font-bold">{{ $tl("username") }}</q-item-label>
				<q-item-label caption>{{ model?.username }}</q-item-label>
			</q-item-section>
		</q-item>

		<q-separator />

		<q-item>
			<q-item-section>
				<q-item-label class="font-bold">{{ $tl("dateOfBirth") }}</q-item-label>
				<q-item-label caption>{{ model?.dateOfBirth }}</q-item-label>
			</q-item-section>
		</q-item>

		<q-separator />

		<q-item>
			<q-item-section>
				<q-item-label class="font-bold">{{ $tl("telegramUsername") }}</q-item-label>
				<q-item-label caption>{{ model?.telegramUsername }}</q-item-label>
			</q-item-section>
		</q-item>

		<q-separator />

		<q-item>
			<q-item-section>
				<q-item-label class="font-bold">{{ $tl("phoneNumber") }}</q-item-label>
				<q-item-label caption>{{ model?.phoneNumber }}</q-item-label>
			</q-item-section>
		</q-item>
	</q-list>
</template>
