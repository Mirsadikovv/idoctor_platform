
import { ref, onMounted, onUnmounted, computed } from "vue";
import { TelegramWebApp } from "@/common/telegram";

export function useTelegramViewport() {
	const topInset = ref(0);
	const viewportHeight = ref(window.innerHeight);
	const isExpanded = ref(false);

	/**
	 * Main function to update state from Telegram WebApp
	 */
	const updateViewport = () => {
		if (TelegramWebApp.isAvailable()) {
			const webApp = TelegramWebApp.webApp;
			if (webApp) {
				// 1. Get safe area inset top
				// Fallback to CSS variable or 0 if not available
				topInset.value = webApp.contentSafeAreaInset?.top || 0;

				// 2. Get viewport height
				// Using webApp.viewportStableHeight often gives a better "keyboard-open" experience,
				// but viewportHeight is standard.
				viewportHeight.value = webApp.viewportHeight || window.innerHeight;

				// 3. Expansion state
				isExpanded.value = webApp.isExpanded || false;

				// 4. Update CSS variable for global usage
				document.documentElement.style.setProperty(
					"--tg-viewport-height",
					`${viewportHeight.value}px`
				);
				document.documentElement.style.setProperty(
					"--tg-content-safe-area-inset-top",
					`${topInset.value}px`
				);
			}
		} else {
			// Fallback for regular browser
			// We can try to read env(safe-area-inset-top) via a hidden element if strictly needed,
			// but usually 0 is fine for desktop/dev.
			topInset.value = 0;
			viewportHeight.value = window.innerHeight;
		}
	};

	/**
	 * Computed full top padding:
	 * topInset (notch) + base padding (e.g. 16px from design)
	 */
	const pageTopPaddingPx = computed(() => {
		// Base padding from design (previously pt-24 was ~96px? No, pt-24 in Tailwind is 6rem = 96px.
		// Wait, user said "pt-24 — статический отступ".
		// If pt-24 is from Quasar/standard, 24 * 4 = 96px? Or maybe it's just 24px?
		// Standard Quasar spacing: q-pt-md is 16px.
		// "pt-24" looks like Tailwind class. 24 * 0.25rem = 6rem = 96px.
		// Let's assume we want a base padding of roughly 16px-24px depending on content,
		// plus the safe area.
		// The user complained about "static indent from top".
		// Let's use a safe base like 16px.
		const basePadding = 16;
		return topInset.value + basePadding;
	});

	// Style object to apply to the page wrapper
	const containerStyle = computed(() => ({
		paddingTop: `${pageTopPaddingPx.value}px`,
		height: `calc(var(--tg-viewport-height, 100vh) - 150px)`, // Preserving logic from Page.vue
		// Note: Page.vue had 'calc(var(--app-height, 100vh) - 150px)'
		// We should ensure --app-height or --tg-viewport-height is consistent.
	}));

	onMounted(() => {
		updateViewport();

		if (TelegramWebApp.isAvailable()) {
			TelegramWebApp.webApp?.onEvent("viewportChanged", updateViewport);
		}

		window.addEventListener("resize", updateViewport);
	});

	onUnmounted(() => {
		if (TelegramWebApp.isAvailable()) {
			TelegramWebApp.webApp?.offEvent("viewportChanged", updateViewport);
		}
		window.removeEventListener("resize", updateViewport);
	});

	return {
		topInset,
		viewportHeight,
		isExpanded,
		pageTopPaddingPx,
		containerStyle,
	};
}
