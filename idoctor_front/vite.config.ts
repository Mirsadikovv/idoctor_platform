import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import { quasar, transformAssetUrls } from "@quasar/vite-plugin";
import { fileURLToPath } from "node:url";
import UnoCSS from "unocss/vite";
import { resolve } from "node:path";

export default defineConfig(({ mode }) => {
	const env = loadEnv(mode, "");

	const envFn = (key: string, _default: string = "") => {
		return JSON.stringify(env[key] || _default);
	};

	return {
		server: {
			port: 5171,
		},
		preview: {
			port: 5173,
			allowedHosts: true,
			host: "0.0.0.0",
		},
		plugins: [
			UnoCSS(),
			vue({
				template: { transformAssetUrls },
			}),
			quasar({
				autoImportComponentCase: "combined",
				sassVariables: resolve("src", "style", "quasar-variables.scss"),
			}),
		],
		define: {
			APP_TITLE: envFn("VITE_APP_TITLE", "APP_TITLE"),
			TOKEN_NAME: envFn("VITE_TOKEN_NAME", "token"),
			LANG_NAME: envFn("VITE_LANG_NAME", "lang"),
			API_URL: envFn("VITE_API_URL"),
		},
		resolve: {
			alias: {
				"@": fileURLToPath(new URL("./src", import.meta.url)),
				"@module": fileURLToPath(new URL("./src/modules", import.meta.url)),
				"@layout": fileURLToPath(new URL("./src/components/layouts", import.meta.url)),
			},
			extensions: [".js", ".json", ".jsx", ".mjs", ".ts", ".tsx", ".vue"],
		},
	};
});
