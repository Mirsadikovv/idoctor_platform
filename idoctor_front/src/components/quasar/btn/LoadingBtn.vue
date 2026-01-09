<script setup lang="ts">
import { ref } from "vue";
import { type QBtnProps } from "quasar";
import Button from "./Button.vue";
import type { ProcessedIcons } from "@/common";
import IconBtn from "./IconBtn.vue";

export interface Props extends QBtnProps {
  loader?: () => Promise<unknown>;
  icon?: ProcessedIcons;
}

const props = withDefaults(defineProps<Props>(), {
  color: "primary",
  flat: true,
  round: true,
  replace: undefined,
  outline: undefined,
  unelevated: undefined,
  rounded: undefined,
  push: undefined,
  square: undefined,
  glossy: undefined,
  fab: undefined,
  fabMini: undefined,
  noCaps: undefined,
  noWrap: undefined,
  dense: undefined,
  ripple: undefined,
  stack: undefined,
  stretch: undefined,
  loading: undefined,
  disable: undefined,
  darkPercentage: undefined,
});

const confirmDialog = ref(false);
const loadingBtn = ref(false);
const emit = defineEmits(["confirm"]);

const onConfirm = async () => {
  if (props.loader) {
    try {
      loadingBtn.value = true;
      await props.loader();
    } catch (error) {
    } finally {
      loadingBtn.value = false;
    }
  }

  confirmDialog.value = false;
  emit("confirm");
};
</script>

<template>
  <span>
    <IconBtn
      v-if="props.icon"
      v-bind="props"
      :icon="props.icon!"
      :loading="loadingBtn"
      @click="confirmDialog = true"
    >
    </IconBtn>
    <Button v-else v-bind="props" :loading="loadingBtn">
      <slot></slot>
    </Button>

    <q-dialog v-model="confirmDialog" persistent>
      <q-card>
        <q-card-section class="text-center">
          {{ $tl("are_you_sure") }}
        </q-card-section>
        <q-card-actions align="right">
          <q-btn flat color="negative" v-close-popup>
            {{ $tl("cancel") }}
          </q-btn>
          <q-btn flat color="primary" @click="onConfirm">
            {{ $tl("confirm") }}
          </q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>
  </span>
</template>
