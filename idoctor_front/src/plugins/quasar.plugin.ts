import type { App } from "vue";
import { Quasar, LocalStorage, Notify, AppFullscreen } from "quasar";
import "@quasar/extras/material-icons/material-icons.css";
import "quasar/src/css/index.sass";

export const quasarPlugin = (app: App<unknown>) => {
	return app.use(Quasar, {
		plugins: {
			LocalStorage,
			Notify,
			AppFullscreen,
		},
	});
};
