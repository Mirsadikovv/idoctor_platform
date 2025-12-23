import "uno.css";
import "@/style/style.scss";
import { createApp, h } from "vue";
import App from "./App.vue";
import install from "./plugins/install";
import { permissionsDirective } from "./common";

const app = createApp({
	render: () => h(App),
	beforeCreate() {
		app.directive("permissions", permissionsDirective());
	},
});

install(app).then(() => app.mount("#app"));
