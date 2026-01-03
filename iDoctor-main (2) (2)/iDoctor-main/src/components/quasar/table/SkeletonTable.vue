<script setup lang="ts">
export interface Props {
	cell?: number;
}

const { cell = 5 } = defineProps<Props>();
</script>

<template>
	<q-markup-table separator="cell" flat bordered>
		<thead>
			<tr>
				<th
					v-for="n in cell"
					class="h-11!"
					style="border: 1px solid #0000001f"
				>
					<q-skeleton :key="n" animation="blink" type="text" />
				</th>
			</tr>
		</thead>

		<tbody>
			<tr v-for="n in 5" :key="n">
				<td v-for="n in cell" :key="n" class="h-14!">
					<q-skeleton animation="blink" type="text" />
				</td>
			</tr>
			<tr>
				<th colspan="99">
					<div class="flex flex-center">
						<q-skeleton
							width="210px"
							height="30px"
							animation="blink"
							type="text"
						/>
					</div>
				</th>
			</tr>
		</tbody>
	</q-markup-table>
</template>
