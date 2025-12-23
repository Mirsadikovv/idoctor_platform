<script setup lang="ts">
import { ProblemService } from "../service";

interface Props {
	id: number;
	fetch: () => void;
}

const props = withDefaults(defineProps<Props>(), {});

async function remove() {
	const response = await ProblemService.delete(props.id);
	if (response) {
		props.fetch();
	}
}
</script>

<template>
	<div class="q-pa-md">
		<div class="text-h6">{{ $tl("are_you_sure") }}</div>
		<div class="q-pa-md text-center">
			<p>{{ $tl("problem_delete_confirm") }}</p>
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
					:label="$tl('delete')"
					color="negative"
					@click="remove"
					v-close-popup
					class="w-full!"
				/>
			</div>
		</div>
	</div>
</template>
