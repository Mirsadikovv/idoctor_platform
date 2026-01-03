<script setup lang="ts">
import { ref, onMounted, nextTick } from "vue";
import { useRoute, useRouter } from "vue-router";

export interface Props {
	name?: string;
	label?: string;
	modelValue?: string;
	dense?: boolean;
}

const { name = "year", label = "Year", dense } = defineProps<Props>();

const emits = defineEmits<{
	(e: "search", search: string): void;
}>();

const router = useRouter();
const route = useRoute();

const currentYear = new Date().getFullYear().toString();
const search = ref<string>("");
const datePicker = ref(); // popupni boshqarish uchun ref

// ✅ Komponent yuklanganda
onMounted(async () => {
	const qYear = route.query[name] as string | undefined;
	const initialYear = qYear === "" ? "" : qYear || currentYear;
	search.value = initialYear;

	// agar query’da yo‘q bo‘lsa, default yozamiz
	if (!route.query[name]) {
		await router.replace({
			query: { ...route.query, [name]: currentYear },
		});
	}

	emits("search", initialYear);
});

/**
 * 🔹 Yil tanlanganda
 */
async function searchChange(value: string | number | null) {
	if (value == null) return;

	const selectedYear = value.toString() || currentYear;
	const { query } = route;

	search.value = selectedYear;
	query[name] = selectedYear;

	await router.push({
		force: true,
		query: { ...query },
	});

	await nextTick();
	datePicker.value?.hide();

	emits("search", selectedYear);
}
/**
 * 🔹 Tozalash (yil querysiz request yuborish)
 */
async function onClear() {
	const { query } = route;
	const { [name]: removed, ...filtered } = query;

	search.value = "";

	// 🧹 year queryni butunlay olib tashlaymiz
	await router.replace({ query: filtered });

	await nextTick();
	datePicker.value?.hide();

	// 🔥 parentga signal — bu fetch() chaqiradi
	emits("search", "");
}
</script>

<template>
	<q-input
		v-model="search"
		@update:model-value="searchChange"
		@clear="onClear"
		:label="$tl ? $tl(label) : label"
		placeholder="YYYY"
		mask="####"
		color="accent"
		clearable
		outlined
		:dense="dense"
	>
		<template #prepend>
			<q-icon name="event" class="cursor-pointer mr-2">
				<q-popup-proxy
					ref="datePicker"
					transition-show="scale"
					transition-hide="scale"
					cover
				>
					<!-- ✅ faqat yillar -->
					<q-date
						v-model="search"
						@update:model-value="searchChange"
						mask="YYYY"
						color="accent"
						:default-year="Number(currentYear)"
						:default-view="'Years'"
						emit-immediately
						years-only
					/>
				</q-popup-proxy>
			</q-icon>
		</template>
	</q-input>
</template>
