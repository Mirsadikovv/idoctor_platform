<script setup lang="ts">
import { PartService } from "../service";

interface Props {
	id: number;
	isRemove: boolean;
	fetch: () => void;
}

const props = withDefaults(defineProps<Props>(), {});

async function remove() {
	const response = props.isRemove
		? await PartService.delete(props.id)
		: await PartService.restore(props.id);

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
				{{ isRemove ? $tl("part_delete_confirm") : $tl("part_restore_confirm") }}
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
