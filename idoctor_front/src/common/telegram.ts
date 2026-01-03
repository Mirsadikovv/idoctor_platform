export class TelegramWebApp {
	static isAvailable(): boolean {
		return !!window?.Telegram?.WebApp;
	}

	static get webApp() {
		return this.isAvailable() ? window.Telegram!.WebApp : null;
	}

	static getInitData(): string {
		return this.webApp?.initData ?? "";
	}

	static getUserInfo() {
		return this.webApp?.initDataUnsafe?.user ?? null;
	}

	static getTelegramId(): number | null {
		let tg = window?.Telegram?.WebApp;
		let userId = tg?.initDataUnsafe?.user?.id;
		return userId ?? null;
	}

	static initialize({ expand = false } = {}) {
		const app = this.webApp;
		if (!app) return;

		app.ready();
		if (expand) app.expand();
	}

	static close() {
		this.webApp?.close();
	}

	/**
	 * ВАЖНО: это НЕ криптопроверка.
	 * Это только "похоже, что initData присутствует".
	 */
	static hasInitData(): boolean {
		const initData = this.getInitData();
		return initData.length > 0;
	}
}
