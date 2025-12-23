<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { UserService, type UserPartial } from "@/service";

export interface Props {
	id: number | string;
}
const { id } = defineProps<Props>();

const router = useRouter();

let model = ref<UserPartial>({});

UserService.findByID(+id).then((data) => {
	model.value = data;
});
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
