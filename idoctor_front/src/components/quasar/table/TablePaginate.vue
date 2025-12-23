<script setup lang="ts">
import PerPage from "@/components/PerPage.vue";
import Paginate from "./Paginate.vue";
import ColumsPiker from "@/components/ColumsPiker.vue";

export interface Props {
	total: number;
	maxPages?: number;
	name?: string;
	pick: Record<string, boolean>;
	pikers: Record<string, boolean>;
	hidePerPage?: boolean;
}

const { maxPages, name, total, pick, pikers } = defineProps<Props>();

const emits = defineEmits<{
	(e: "page", page: number): void;
	(e: "update:pikers", pikers: Record<string, boolean>): void;
}>();
</script>

<template>
	<tr class="mobile-foot">
		<th colspan="99" class="border-t! w-full">
			<div class="flex justify-between items-center w-full!">
				<ColumsPiker
					:name="$tl('columns')"
					:model-value="pikers"
					:picker="pick"
					@update:model-value="(value) => emits('update:pikers', value)"
				/>

				<Paginate
					v-if="!hidePerPage"
					:total="total"
					:max-pages="maxPages"
					:name="name"
					@page="(page) => emits('page', page)"
				/>
				<PerPage v-if="!hidePerPage" @per-page="(page) => emits('page', page)" />
			</div>
		</th>
	</tr>
</template>

<style lang="css">
@media screen and (max-width: 1100px) {
	.mobile-foot {
		flex: 1 1 100%;
		display: flex;
		justify-content: space-between;
		width: 100%;
	}
}
</style>
