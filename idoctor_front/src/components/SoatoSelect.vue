<script setup lang="ts">
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { ref, watch, computed, onMounted } from "vue";
import { AddressService } from "@/service";
import { getId } from "@/common";
import { useLanguageStore } from "@/store/language-store";

interface Props {
	languageId?: any;
	disabled?: boolean;
}

const props = defineProps<Props>();
const organizationSoatoId = defineModel<number | undefined>();

const langStore = useLanguageStore();

// Address state management
const selectedRegion = ref<any>(null);
const selectedDistrict = ref<any>(null);
const selectedMahalla = ref<any>(null);
const regions = ref<any[]>([]);
const districts = ref<any[]>([]);
const mahallaList = ref<any[]>([]);

// Computed language ID
const currentLangID = computed(() => {
	return getId(props.languageId) || langStore.langID;
});

// Watch for language changes
watch(
	() => props.languageId,
	async () => {
		resetSelections();
		regions.value = [];
		await fetchRegions();
	},
);

// Initialize regions on component mount
onMounted(async () => {
	await fetchRegions();
});

// Address select functions
async function fetchRegions() {
	const response = await AddressService.getAllAddress(
		`parent_id=17&perpage=500`,
		currentLangID.value,
	);
	regions.value = response.data || [];
	return regions.value;
}

// Handle address selection changes
async function onRegionChange(region: any) {
	// Reset child selections
	selectedDistrict.value = null;
	selectedMahalla.value = null;
	districts.value = [];
	mahallaList.value = [];

	if (region && region.soatoId) {
		organizationSoatoId.value = region.soatoId;
		// Fetch districts for the selected region
		const response = await AddressService.getAllAddress(
			`parent_id=${region.soatoId}&perpage=500`,
			currentLangID.value,
		);
		districts.value = response.data || [];
	} else {
		organizationSoatoId.value = undefined;
	}
}

async function onDistrictChange(district: any) {
	// Reset mahalla selection
	selectedMahalla.value = null;
	mahallaList.value = [];

	if (district && district.soatoId) {
		organizationSoatoId.value = district.soatoId;
		// Fetch mahallas for the selected district
		const response = await AddressService.getAllAddress(
			`parent_id=${district.soatoId}&perpage=500`,
			currentLangID.value,
		);
		mahallaList.value = response.data || [];
	} else if (selectedRegion.value && selectedRegion.value.soatoId) {
		organizationSoatoId.value = selectedRegion.value.soatoId;
	} else {
		organizationSoatoId.value = undefined;
	}
}

async function onMahallaChange(mahalla: any) {
	if (mahalla && mahalla.soatoId) {
		organizationSoatoId.value = mahalla.soatoId;
	} else if (selectedDistrict.value && selectedDistrict.value.soatoId) {
		organizationSoatoId.value = selectedDistrict.value.soatoId;
	} else if (selectedRegion.value && selectedRegion.value.soatoId) {
		organizationSoatoId.value = selectedRegion.value.soatoId;
	} else {
		organizationSoatoId.value = undefined;
	}
}

// Reset all selections
function resetSelections() {
	selectedRegion.value = null;
	selectedDistrict.value = null;
	selectedMahalla.value = null;
	districts.value = [];
	mahallaList.value = [];
	organizationSoatoId.value = undefined;
}

// Expose reset function for parent component
defineExpose({
	resetSelections,
});
</script>

<template>
	<div class="row my-row">
		<!-- Region Select -->
		<Autocomplete
			label="region"
			v-model="selectedRegion"
			:find="fetchRegions"
			@update="onRegionChange"
			class="col-lg-4 col-md-6 col-12"
			option-label="title"
			clearable
			:disable="disabled"
		/>

		<!-- District Select -->
		<Autocomplete
			label="district"
			v-model="selectedDistrict"
			:find="async () => districts"
			@update="onDistrictChange"
			class="col-lg-4 col-md-6 col-12"
			option-label="title"
			clearable
			:disable="disabled || !selectedRegion"
		/>

		<!-- Mahalla Select -->
		<Autocomplete
			label="mahalla"
			v-model="selectedMahalla"
			:find="async () => mahallaList"
			@update="onMahallaChange"
			class="col-lg-4 col-md-6 col-12"
			option-label="title"
			clearable
			:disable="disabled || !selectedDistrict"
		/>
	</div>
</template>
