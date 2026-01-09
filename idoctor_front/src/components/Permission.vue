<script setup lang="ts">
import { useAuthStore } from "@/store/auth-store";
import { useRoute } from "vue-router";

export interface Props {
  permissions?: string[];
  routeName: string;
}

const { permissions, routeName } = defineProps<Props>();

const authStore = useAuthStore();

const route = useRoute();

const canPage = (() => {
  if (permissions) {
    const routeName = String(route.name);
    const result = authStore.page(routeName);

    if (!result) {
      return false;
    }

    return result.some((permission) =>
      permissions.some((p) => p === permission)
    );
  }

  return !!authStore.page(routeName);
})();
</script>

<template>
  <slot v-if="canPage"></slot>
</template>
