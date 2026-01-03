<script setup lang="ts">
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";
import { ref, computed, onMounted } from "vue";
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

// State
const selectedRegion = ref<any>(null);
const selectedDistrict = ref<any>(null);
const regions = ref<any[]>([]);
const districts = ref<any[]>([]);

// Computed language ID
const currentLangID = computed(() => getId(props.languageId) || langStore.langID);

// При смене currentLangID всегда будут новые регионы
const regionsComputed = computed(async () => {
  resetSelections();
  const response = await AddressService.getAllAddress(
    `parent_id=17&perpage=500`,
    currentLangID.value,
  );
  regions.value = response.data || [];
  return regions.value;
});

// Селект регионов/районов
async function onRegionChange(region: any) {
  selectedDistrict.value = null;
  districts.value = [];

  if (region?.soatoId) {
    organizationSoatoId.value = region.soatoId;
    const response = await AddressService.getAllAddress(
      `parent_id=${region.soatoId}&perpage=500`,
      currentLangID.value,
    );
    districts.value = response.data || [];
  } else {
    organizationSoatoId.value = undefined;
  }
}

function onDistrictChange(district: any) {
  if (district?.soatoId) {
    organizationSoatoId.value = district.soatoId;
  } else if (selectedRegion.value?.soatoId) {
    organizationSoatoId.value = selectedRegion.value.soatoId;
  } else {
    organizationSoatoId.value = undefined;
  }
}

// Reset
function resetSelections() {
  selectedRegion.value = null;
  selectedDistrict.value = null;
  districts.value = [];
  organizationSoatoId.value = undefined;
}

// Первичная загрузка
onMounted(async () => {
  await regionsComputed.value;
});

defineExpose({ resetSelections });
</script>

<template>
  <Autocomplete
    label="region"
    v-model="selectedRegion"
    :find=" async () => regions"
    @update="onRegionChange"
    class="col-lg-3 col-md-6 col-sm-12 col-xs-12"
    option-label="title"
    clearable
    :disable="disabled"
  />

  <Autocomplete
    label="district"
    v-model="selectedDistrict"
    :find="async () => districts"
    @update="onDistrictChange"
    class="col-lg-3 col-md-6 col-sm-12 col-xs-12"
    option-label="title"
    clearable
    :disable="disabled || !selectedRegion"
  />
</template>