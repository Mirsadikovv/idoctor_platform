<script setup lang="ts">
import { dateFormat, dateValid, isValidDate, parseDate } from "@/common";
import { ref, onMounted, nextTick, type Ref } from "vue";
import { useRoute, useRouter } from "vue-router";

export interface Props {
	name?: string;
	label?: string;
	maskInput?: string;
	maskDate?: string;
	dense?: boolean;
	withoutTime?: boolean;
}

const {
	name = "search",
	label = "search",
	maskDate,
	maskInput,
	dense,
	withoutTime,
} = defineProps<Props>();

const emits = defineEmits<{
	(e: "search", search: string): void;
}>();

const router = useRouter();
const route = useRoute();

const search: Ref<string> = ref("");
const hasError = ref(false);
const datePicker = ref<any>();

onMounted(async () => {
	const today = new Date();
	let defaultDate = "";

	if (name === "from_date") {
		defaultDate = `${today.getFullYear()}-01-01`;
	} else if (name === "to_date") {
		defaultDate = today.toISOString().split("T")[0];
	} else {
		defaultDate = today.toISOString().split("T")[0];
	}

	const qValue = route.query[name] as string | undefined;

	// 🧠 Agar queryda qiymat bo‘lsa — inputga shuni yozamiz
	if (qValue) {
		search.value = dateFormat(qValue, "DD-MM-YYYY") ?? "";
		emits("search", qValue);
		return;
	}
	// 🧩 to_date bo‘sh bo‘lsa, faqat bo‘sh emit qilamiz (API request ketsin)
	if (name === "to_date") {
		search.value = "";
		emits("search", "");
	}

	// 🧩 from_date bo‘sh bo‘lsa, faqat bo‘sh emit qilamiz (API request ketsin)
	if (name === "from_date") {
		search.value = "";
		emits("search", "");
	}
});

async function searchChange(item: string | number | null) {
	if (typeof item !== "string" || !isValidDate(item)) {
		hasError.value = !!item;
		return;
	}

	hasError.value = false;
	const { query } = route;

	const parsed = parseDate(item, "DD-MM-YYYY");
	if (!parsed) return;

	const year = parsed.getFullYear();
	const month = parsed.getMonth();
	const day = parsed.getDate();

	const utcDate = new Date(Date.UTC(year, month, day, 0, 0, 0));

	const newDate = withoutTime ? utcDate.toISOString().split("T")[0] : utcDate.toISOString();

	query[name] = newDate;

	search.value = dateFormat(newDate, "DD-MM-YYYY") ?? "";

	await router.push({ force: true, query: { ...query } });
	emits("search", newDate);

	await nextTick();
	datePicker.value?.hide();
}

async function onClear() {
	const { query } = route;
	const { [name]: removed, ...filtered } = query;

	search.value = "";
	hasError.value = false;

	await router.replace({ query: filtered });
	emits("search", "");
}
</script>

<template>
	<div class="relative w-full">
		<q-input
			v-model="search"
			@update:model-value="search && searchChange"
			@clear="onClear"
			:placeholder="maskDate || 'DD-MM-YYYY'"
			:mask="maskInput || '##-##-####'"
			:label="$tl(label)"
			:rules="[dateValid($tl('date_will_be_dd_mm_yyyy'))]"
			:validate="isValidDate"
			color="accent"
			debounce="700"
			clearable
			outlined
			:dense="dense"
			:error="hasError"
		>
			<template #prepend>
				<q-icon name="event" class="cursor-pointer mr-2">
					<q-popup-proxy
						ref="datePicker"
						transition-show="scale"
						transition-hide="scale"
						cover
					>
						<q-date
							v-model="search"
							@update:model-value="searchChange"
							:mask="maskDate || 'DD-MM-YYYY'"
							color="accent"
							landscape
						>
							<div class="row items-center justify-end">
								<q-btn v-close-popup :label="$tl('close')" color="accent" flat />
							</div>
						</q-date>
					</q-popup-proxy>
				</q-icon>
			</template>
		</q-input>

		<transition name="fade">
			<div
				v-if="hasError"
				class="absolute text-negative text-caption mt-1"
				style="left: 0; bottom: -18px"
			></div>
		</transition>
	</div>
</template>

<style scoped>
.relative {
	position: relative;
}
.fade-enter-active,
.fade-leave-active {
	transition: opacity 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
	opacity: 0;
}
</style>
