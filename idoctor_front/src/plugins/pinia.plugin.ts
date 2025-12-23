import type { App } from "vue";
import { createPinia } from "pinia";

export const piniaPlugin = (app: App<unknown>) => {
  return app.use(createPinia());
};
