<script setup lang="ts">
import { OrderService } from "../service";

interface Props {
	id: number;
	isRemove: boolean;
	fetch: () => void;
}

const props = withDefaults(defineProps<Props>(), {});

async function remove() {
	const response = props.isRemove
		? await OrderService.delete(props.id)
		: await OrderService.restore(props.id);

	if (response) {
		props.fetch();
	}
}
</script>

<template>
	<div class="q-pa-md">
		<div class="text-h6">{{ $tl("are_you_sure") }}</div>
		<div class="q-pa-md text-center">
			<p>
				{{ isRemove ? $tl("order_delete_confirm") : $tl("order_restore_confirm") }}
			</p>
		</div>
		<div class="row q-col-gutter-sm">
			<div class="col">
				<q-btn
					outline
					:label="$tl('cancel')"
					color="negative"
					v-close-popup
					class="w-full!"
				/>
			</div>
			<div class="col">
				<q-btn
					:label="$tl(isRemove ? 'delete' : 'restore')"
					:color="isRemove ? 'negative' : 'positive'"
					@click="remove"
					v-close-popup
					class="w-full!"
				/>
			</div>
		</div>
	</div>
</template>
