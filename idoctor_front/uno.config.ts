import {
	defineConfig,
	presetAttributify,
	presetIcons,
	presetTypography,
	presetWind3,
	transformerDirectives,
	transformerVariantGroup,
} from "unocss";
export default defineConfig({
	presets: [presetWind3(), presetAttributify(), presetIcons(), presetTypography()],
	transformers: [transformerDirectives(), transformerVariantGroup()],
	theme: {
		colors: {
			"gradient-primary": "var(--q-info)",
			"gradient-secondary": "var(--q-accent)",
		},
		animation: {
			"gradient-shift": "gradientShift 4s ease infinite",
		},
	},
	shortcuts: {
		// Градиентные утилиты
		"bg-gradient-header": "bg-gradient-to-br from-gradient-primary to-gradient-secondary",
		"bg-gradient-animated": "bg-gradient-header bg-size-200 animate-gradient-shift",

		// Размеры фона
		"bg-size-200": "bg-size-[200%_200%]",
	},
	rules: [
		// Кастомные CSS правила для анимации градиента
		[
			/^animate-gradient-shift$/,
			() => ({
				animation: "gradientShift 4s ease infinite",
			}),
		],
		[
			/^bg-size-\[(.+)\]$/,
			([, size]) => ({
				"background-size": size.replace(/_/g, " "),
			}),
		],
	],
	preflights: [
		{
			getCSS: () => `
				@keyframes gradientShift {
					0% { background-position: 0% 50%; }
					50% { background-position: 100% 50%; }
					100% { background-position: 0% 50%; }
				}
			`,
		},
	],
});
