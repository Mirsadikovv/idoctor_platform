<script setup lang="ts">
import { ref } from "vue";
import { OrderService, type OrderType } from "../service";
import PageLoading from "@/components/PageLoading.vue";
import Title from "@/components/Title.vue";

export interface Props {
	id: string | number;
}

const { id } = defineProps<Props>();

const order = ref<OrderType | null>(null);

async function loadOrder() {
	const orderData = await OrderService.getByID(+id);
	if (!orderData) return;

	order.value = orderData;
}
</script>

<template>
	<PageLoading :find="loadOrder">
		<div class="q-pa-md">
			<Title class="mb-6"> {{ $tl("order_details") }} #{{ id }} </Title>

			<div v-if="order" class="row q-col-gutter-md">
				<!-- Order Information -->
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
							<div class="q-mb-sm">
								<strong>{{ $tl("status") }}:</strong>
								<q-chip color="primary" outline class="q-ml-sm">
									{{ order.status }}
								</q-chip>
							</div>
							<div class="q-mb-sm">
								<strong>{{ $tl("payment_status") }}:</strong>
								<q-chip
									:color="
										order.payment_status === 'paid' ? 'positive' : 'warning'
									"
									outline
									class="q-ml-sm"
								>
									{{ order.payment_status }}
								</q-chip>
							</div>
							<div class="q-mb-sm">
								<strong>{{ $tl("payment_type") }}:</strong>
								{{ order.payment_type }}
							</div>
							<div class="q-mb-sm">
								<strong>{{ $tl("created_at") }}:</strong>
								{{ new Date(order.created_at).toLocaleString() }}
							</div>
							<div v-if="order.deleted_at" class="q-mb-sm">
								<strong>{{ $tl("deleted_at") }}:</strong>
								{{ new Date(order.deleted_at).toLocaleString() }}
							</div>
						</q-card-section>
					</q-card>
				</div>

				<!-- People Information -->
				<!-- <div class="col-12 col-md-6">
					<q-card class="q-pa-md">
						<q-card-section>
							<div class="text-h6">{{ $tl("people") }}</div>
						</q-card-section>
						<q-card-section>
							<div class="q-mb-sm">
								<strong>{{ $tl("client") }}:</strong>
								<div v-if="client" class="q-mt-xs">
									<div v-if="client.phoneNumber" class="text-caption">
										{{ client.phoneNumber }}
									</div>
								</div>
								<span v-else-if="order.client_id">ID: {{ order.client_id }}</span>
								<span v-else>{{ $tl("not_assigned") }}</span>
							</div>
							<div class="q-mb-sm">
								<strong>{{ $tl("master") }}:</strong>
								<div v-if="master" class="q-mt-xs">
									<div v-if="master.phoneNumber" class="text-caption">
										{{ master.phoneNumber }}
									</div>
								</div>
								<span v-else-if="order.master_id">ID: {{ order.master_id }}</span>
								<span v-else>{{ $tl("not_assigned") }}</span>
							</div>
						</q-card-section>
					</q-card>
				</div> -->

				<!-- Problems -->
				<!-- <div class="col-12 col-md-6">
					<q-card class="q-pa-md">
						<q-card-section>
							<div class="text-h6">{{ $tl("problems") }}</div>
						</q-card-section>
						<q-card-section>
							<div v-if="problems.length > 0">
								<div v-for="problem in problems" :key="problem.id" class="q-mb-sm">
									<q-chip color="orange" outline>
										{{ problem.name }}
									</q-chip>
									<div class="text-caption">
										{{ problem.price?.toLocaleString() }} сум
									</div>
								</div>
							</div>
							<div v-else-if="order.problem_ids?.length">
								<div
									v-for="problemId in order.problem_ids"
									:key="problemId"
									class="q-mb-sm"
								>
									ID: {{ problemId }}
								</div>
							</div>
							<div v-else>
								{{ $tl("no_problems_assigned") }}
							</div>
						</q-card-section>
					</q-card>
				</div> -->

				<!-- Parts -->
				<!-- <div class="col-12 col-md-6">
					<q-card class="q-pa-md">
						<q-card-section>
							<div class="text-h6">{{ $tl("parts") }}</div>
						</q-card-section>
						<q-card-section>
							<div v-if="parts.length > 0">
								<div v-for="part in parts" :key="part.id" class="q-mb-sm">
									<q-chip color="blue" outline>
										{{ part.name }}
									</q-chip>
								</div>
							</div>
							<div v-else-if="order.part_ids?.length">
								<div v-for="partId in order.part_ids" :key="partId" class="q-mb-sm">
									ID: {{ partId }}
								</div>
							</div>
							<div v-else>
								{{ $tl("no_parts_assigned") }}
							</div>
						</q-card-section>
					</q-card>
				</div> -->
			</div>
		</div>
	</PageLoading>
</template>
