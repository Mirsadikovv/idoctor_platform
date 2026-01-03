<script setup lang="ts">
import { ref } from "vue";
import { OrderPartService, type OrderPartType } from "../service";
import PageLoading from "@/components/PageLoading.vue";
import Title from "@/components/Title.vue";

export interface Props {
	id: string | number;
}

const { id } = defineProps<Props>();

const order = ref<OrderPartType | null>(null);

async function loadOrderPart() {
	const orderData = await OrderPartService.getByID(+id);
	if (!orderData) return;

	order.value = orderData;
}
</script>

<template>
	<PageLoading :find="loadOrderPart">
		<div class="q-pa-md">
			<Title class="mb-6"> {{ $tl("order_details") }} #{{ id }} </Title>

			<div v-if="order" class="row q-col-gutter-md">
				<div class="col-12 col-md-6">
					<q-card class="q-pa-md">
						<q-card-section>
							<div class="text-h6">{{ $tl("order_information") }}</div>
						</q-card-section>
						<q-card-section>
							<div class="q-mb-sm">
								<strong>{{ $tl("price") }}:</strong>
								{{ order.price?.toLocaleString() }} сум
							</div>
						</q-card-section>
					</q-card>
				</div>
			</div>
		</div>
	</PageLoading>
</template>
