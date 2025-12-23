<script setup lang="ts">
import type { ProcessedIcons } from "@/common";
import Permission from "@/components/Permission.vue";
import type { QBtnProps } from "quasar";
import { type VNode } from "vue";
import IconBtn from "./IconBtn.vue";
import Button from "./Button.vue";

export interface Props extends QBtnProps {
  icon?: ProcessedIcons;
  permissions?: string[];
}

const props = withDefaults(defineProps<Props>(), {
  color: "primary",
  replace: undefined,
  outline: undefined,
  flat: undefined,
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

const slots = defineSlots<{
  default: () => VNode[];
  label: () => VNode[];
  loading: () => VNode[];
}>();

const routeName = String(props.to.name);
</script>

<template>
  <Permission :permissions="props.permissions" :route-name="routeName">
    <Button
      v-if="!props.icon"
      class="bg-primary text-white"
      v-bind="{ ...props, ...$attrs }"
    >
      <template v-for="(slotFn, name) in slots" #[name]>
        <component :is="slotFn" />
      </template>
    </Button>

    <IconBtn v-else v-bind="{ ...props, ...$attrs }" :icon="props.icon">
      <template v-for="(slotFn, name) in slots" #[name]>
        <component :is="slotFn" /> </template
    ></IconBtn>
  </Permission>
</template>
