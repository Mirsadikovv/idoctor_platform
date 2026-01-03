/**
 * Утилиты для работы с Telegram Web App
 */

export class TelegramWebApp {
  /**
   * Проверяет, запущено ли приложение в Telegram Web App
   */
  static isAvailable(): boolean {
    return !!(window.Telegram && window.Telegram.WebApp);
  }

  /**
   * Получает Telegram ID пользователя
   */
  static getTelegramId(): number | null {
    if (!this.isAvailable()) {
      return null;
    }

    const user = window.Telegram!.WebApp.initDataUnsafe.user;
    return user?.id || null;
  }

  /**
   * Получает полную информацию о пользователе
   */
  static getUserInfo() {
    if (!this.isAvailable()) {
      return null;
    }

    return window.Telegram!.WebApp.initDataUnsafe.user || null;
  }

  /**
   * Получает сырые данные инициализации
   */
  static getInitData(): string {
    if (!this.isAvailable()) {
      return '';
    }

    return window.Telegram!.WebApp.initData;
  }

  /**
   * Инициализирует Telegram Web App
   */
  static initialize() {
    if (this.isAvailable()) {
      window.Telegram!.WebApp.ready();
      window.Telegram!.WebApp.expand();
    }
  }

  /**
   * Закрывает Web App
   */
  static close() {
    if (this.isAvailable()) {
      window.Telegram!.WebApp.close();
    }
  }

  /**
   * Проверяет валидность данных от Telegram (базовая проверка)
   */
  static isDataValid(): boolean {
    if (!this.isAvailable()) {
      return false;
    }

    const { auth_date, hash } = window.Telegram!.WebApp.initDataUnsafe;
    return !!(auth_date && hash);
  }
}