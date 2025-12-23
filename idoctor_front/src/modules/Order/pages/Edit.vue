<script setup lang="ts">
import Form from "@/components/quasar/form/Form.vue";
import { inject, ref, onMounted } from "vue";
import { OrderService, type OrderUpdateType, type OrderType } from "../service";
import Button from "@/components/quasar/btn/Button.vue";
import Input from "@/components/quasar/form/Input.vue";
import Select from "@/components/quasar/form/Select.vue";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import Title from "@/components/Title.vue";
import { formRequired, formNumber } from "@/common/validator";
import {
	searchClients,
	searchMasters,
	searchProblems,
	searchParts,
	orderStatusOptions,
	paymentStatusOptions,
	paymentTypeOptions,
} from "../utils";

export interface Props {
	id: string | number;
	fetch: Function;
}

const { id, fetch } = defineProps<Props>();

const orderModel = ref<Partial<OrderType>>({});
const onClose = inject<() => Promise<void>>("onClose");

onMounted(async () => {
	const orderData = await OrderService.getByID(+id);
	if (orderData) {
		orderModel.value = orderData;
	}
});

async function save(model: OrderUpdateType) {
	const response = await OrderService.update(+id, model);

	if (!response) return false;

	await onClose?.();
	fetch();
	return true;
}
</script>

<template>
	<Form v-model="orderModel" :save="save">
		<template #title>
			<Title class="mb-5">
				{{ $tl("edit_order") }}
				<q-space></q-space>
			</Title>
		</template>

		<template #client_id="{ model }">
			<Autocomplete
				v-model="model.client_id"
				label="Client"
				class="col-12"
				:find="searchClients"
				option-label="username"
				option-value="id"
			/>
		</template>

		<template #master_id="{ model }">
			<Autocomplete
				v-model="model.master_id"
				label="Master"
				class="col-12"
				:find="searchMasters"
				option-label="username"
				option-value="id"
			/>
		</template>

		<template #problem_ids="{ model }">
			<Autocomplete
				v-model="model.problem_ids"
				label="Problems"
				class="col-12"
				:find="searchProblems"
				option-label="name"
				option-value="id"
				multiple
				use-chips
			/>
		</template>

		<template #part_ids="{ model }">
			<Autocomplete
				v-model="model.part_ids"
				label="Parts"
				class="col-12"
				:find="searchParts"
				option-label="name"
				option-value="id"
				multiple
				use-chips
			/>
		</template>

		<template #price="{ model }">
			<Input
				v-model.number="model.price"
				label="Price"
				class="col-12"
				:rules="[formRequired(), formNumber()]"
			/>
		</template>

		<template #status="{ model }">
			<Select
				v-model="model.status"
				:options="orderStatusOptions"
				label="Status"
				class="col-12"
				:rules="[formRequired()]"
			/>
		</template>

		<template #payment_status="{ model }">
			<Select
				v-model="model.payment_status"
				:options="paymentStatusOptions"
				label="Payment Status"
				class="col-12"
				:rules="[formRequired()]"
			/>
		</template>

		<template #payment_type="{ model }">
			<Select
				v-model="model.payment_type"
				:options="paymentTypeOptions"
				label="Payment Type"
				class="col-12"
				:rules="[formRequired()]"
			/>
		</template>

		<template #actions="{ loading }">
			<div class="col-12">
				<div class="row my-row justify-center">
					<div class="col max-w-240px!">
						<Button :loading="loading" v-close-popup ui-type="danger" class="w-full!">
							{{ $tl("cancel") }}
						</Button>
					</div>
					<div class="col max-w-240px!">
						<Button :loading="loading" type="submit" class="w-full!">
							{{ $tl("save") }}
						</Button>
					</div>
				</div>
			</div>
		</template>
	</Form>
</template>
