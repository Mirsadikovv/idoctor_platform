<script setup lang="ts" generic="T">
import type { PageDataType } from "@/service";
import SkeletonCards from "./SkeletonCards.vue";
import { computed } from "vue";
import { useTableData } from "@/composables/useTableData";

export interface Props<T> {
	models: PageDataType<T>;
	loading?: boolean;
	pick?: Record<string, boolean>;
	traction?: (model: T, index: number) => any;
	hasOrder?: boolean;
}

type TheadSlot = {
	[K in keyof T as `${K & string}:thead`]: (_: { key: keyof T }) => any;
};

type TafterSlot = {
	[K in keyof T as `${K & string}:tafter`]: (_: { key: keyof T }) => any;
};

type BodySlot = {
	[key in keyof T]: (_: { model: T; index: number }) => any;
};

type DynamicBodySlot = {
	[key: string]: (_: { model: T; index: number }) => any;
};

type CardSlot = {
	card: (_: { model: T; index: number; fields: string[]; orderNumber?: number }) => any;
};

type SlotBeforeFoot = {
	before: (_: { models: T[]; totalPages: number }) => any;
	tfoot: (_: { models: T[]; totalPages: number }) => any;
};

const slots = defineSlots<
	BodySlot & TheadSlot & CardSlot & SlotBeforeFoot & DynamicBodySlot & TafterSlot
>();

const {
	models = {
		data: [],
		totalPages: 0,
		currentPage: 0,
		pageSize: 0,
	},
	loading = false,
	pick,
	traction = (_: T, _i: number) => undefined,
	hasOrder = false,
} = defineProps<Props<T>>();

const { extractFields, calculateOrderNumber } = useTableData<T>();

const visibleFields = computed(() => extractFields(slots, pick));
</script>

<template>
	<div v-if="!loading" class="mobile-list-container">
		<slot
			v-if="!!slots['before']"
			name="before"
			:models="models.data"
			:totalPages="models.totalPages"
		></slot>

		<div v-if="!models.data.length" class="no-data-container">
			<q-card flat bordered class="no-data-card">
				<q-card-section class="text-center q-py-xl">
					<q-icon name="inbox" size="4rem" color="grey-5" />
					<div class="text-grey-6 q-mt-md text-h6">
						{{ $tl("no_data") }}
					</div>
				</q-card-section>
			</q-card>
		</div>

		<div v-else class="mobile-cards-container">
			<div 
				v-for="(data, index) of models.data" 
				:key="index" 
				class="mobile-card"
				:class="traction(data, index)"
			>
				<slot
					name="card"
					:model="data"
					:index="index"
					:fields="visibleFields"
					:order-number="hasOrder ? calculateOrderNumber(models, index) : undefined"
				>
					<!-- Дефолтная карточка в стиле Language -->
					<div class="default-mobile-card">
						<q-item class="q-mb-md default-item-bordered">
							<q-item-section avatar v-if="hasOrder">
								<q-avatar color="primary" text-color="white" size="sm">
									{{ calculateOrderNumber(models, index) }}
								</q-avatar>
							</q-item-section>

							<q-item-section>
								<div v-for="field in visibleFields" :key="field" class="q-mb-xs">
									<q-item-label>
										<slot :name="`${field}:thead`" :key="field">
											{{ $tl(field) }}
										</slot>:
										<slot :name="field" :model="data" :index="index">
											{{ (data as any)?.[field] }}
										</slot>
									</q-item-label>
								</div>
							</q-item-section>
						</q-item>
					</div>
				</slot>
			</div>
		</div>

		<div v-if="!!slots['tfoot']" class="mobile-footer q-mt-md">
			<slot name="tfoot" :models="models.data" :totalPages="models.totalPages"></slot>
		</div>
	</div>
	<SkeletonCards v-else />
</template>

<style scoped lang="scss">
.mobile-list-container {
	padding: 8px 0;
}

.mobile-cards-container {
	display: flex;
	flex-direction: column;
	gap: 16px;
}

.mobile-card {
	background: white;
	border-radius: 12px;
	border: 1px solid rgba(0, 0, 0, 0.1);
	box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
	padding: 16px;
	transition: all 0.3s ease;

	&:hover {
		box-shadow: 0 4px 16px rgba($primary, 0.15);
		border-color: $primary;
		transform: translateY(-2px);
	}
}

.default-mobile-card {
	.q-item {
		padding: 0;
	}
}

.default-item-bordered {
	border: 1px solid rgba(0, 0, 0, 0.12);
	border-radius: 8px;
	padding: 12px;
}

.field-list {
	.mobile-field {
		margin-bottom: 4px;

		.field-label {
			font-weight: 500;
			font-size: 0.875rem;
		}

		.field-value {
			word-break: break-word;
		}
	}
}

.no-data-container {
	.no-data-card {
		border-radius: 12px;
		border: 2px dashed rgba(0, 0, 0, 0.1);
	}
}

// Responsive adjustments
@media (max-width: 600px) {
	.mobile-list-container {
		padding: 4px 0;
	}

	.mobile-cards-container {
		gap: 8px;
	}

	.mobile-card {
		border-radius: 8px;
	}

	.field-list .mobile-field .field-row {
		flex-direction: column;
		gap: 4px;
	}

	.field-list .mobile-field .field-label {
		min-width: unset;
		font-weight: 600;
	}
}
</style>
