import type { App } from "vue";
import { quasarPlugin } from "./quasar.plugin";
import { routerPlugin } from "./router.plugin";
import { piniaPlugin } from "./pinia.plugin";
import { useLanguageStore } from "@/store/language-store";
import { useAuthStore } from "@/store/auth-store";

export default async (app: App<unknown>) => {
  quasarPlugin(app);

  piniaPlugin(app);

  const languageStore = useLanguageStore();
  const authStore = useAuthStore();
  const globalProperties = app.config.globalProperties;

  globalProperties.$tl = languageStore.tl;
  globalProperties.$tg = languageStore.tg;
  globalProperties.$lang = languageStore;
  globalProperties.$auth = authStore;
  globalProperties.$canPage = authStore.canPage;
  globalProperties.$can = (...key: string[]) => {
    return authStore.can(globalProperties.$route?.name, ...key);
  };
  await routerPlugin(app);
};
