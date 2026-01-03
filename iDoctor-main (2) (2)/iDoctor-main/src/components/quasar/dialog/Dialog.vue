<script setup lang="ts">
import type { ProcessedIcons } from "@/common";
import Button from "./quasar/btn/Button.vue";
import IconBtn from "./quasar/btn/IconBtn.vue";
import type { QBtnProps } from "quasar";

export interface Props extends QBtnProps {
  icon?: ProcessedIcons;
}

const props = withDefaults(defineProps<Props>(), {
  color: "primary",
  replace: undefined,
  outline: undefined,
  flat: true,
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
  round: undefined,
  darkPercentage: undefined,
});

const modelValue = defineModel<boolean>();
</script>

<template>
  <Button
    v-if="!props.icon"
    v-bind="{ ...props, ...$attrs }"
    @click="modelValue = true"
  >
    {{ props.label }}
  </Button>
  <IconBtn
    v-else
    @click="modelValue = true"
    v-bind="{ ...props, ...$attrs }"
    :icon="props.icon"
    :label="$tl(label as string)"
  ></IconBtn>

  <q-dialog v-model="modelValue" persistent>
    <slot></slot>
  </q-dialog>
</template>
