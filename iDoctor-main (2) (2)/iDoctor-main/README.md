# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

-   **Development server**: `npm run dev` или `bun dev`
-   **Production build**: `npm run build` или `bun build`
-   **Preview production**: `npm run preview` или `bun preview`

## Project Architecture

Это Vue 3 приложение с модульной архитектурой для административной панели. Проект использует:

### Technology Stack

-   **Vue 3** с Composition API и TypeScript
-   **Quasar Framework** для UI компонентов
-   **Pinia** для управления состоянием
-   **Vue Router** для маршрутизации
-   **UnoCSS** для стилизации
-   **Vite** для сборки проекта

### Module Structure

Приложение использует модульную архитектуру в папке `src/modules/`:

-   **Auth**: Аутентификация через OneID (SSO), включает профиль пользователя
-   **User**: Управление пользователями
-   **Language**: Мультиязычность и переводы
-   **TranslatedContent**: Переводимый контент
-   **Media**: Управление медиафайлами
-   **Error**: Страницы ошибок (404, Forbidden)

Каждый модуль содержит:

-   `pages/` - Vue компоненты страниц
-   `service/` - API сервисы
-   `routes.ts` - маршруты модуля
-   `components/` - компоненты модуля (если есть)

### Key Components

-   **Store Management**:
    -   `src/store/auth-store.ts` - аутентификация и разрешения
    -   `src/store/language-store.ts` - мультиязычность
-   **Route Registration**: `src/modules/index.ts` регистрирует все модули через `registerRoutes()`

-   **Global Services**: `src/common/` содержит утилиты:
    -   `permission.ts` - директивы для разрешений
    -   `guard.ts` - роутер гарды
    -   `validator.ts`, `utils.ts`, `try.ts` и др.

### Authentication & Authorization

-   Использует OneID для SSO аутентификации
-   Система разрешений на основе ролей через v-permissions директиву
-   API интеграция с `https://dmi-api.mehnat.uz/api/v1/`

### Styling & UI

-   UnoCSS с кастомными shortcuts для градиентов
-   Quasar компоненты в `src/components/quasar/`
-   Кастомные компоненты в `src/components/`
-   Montserrat шрифт включен локально

### API Integration

-   Axios настроен в `src/plugins/axios.plugin.ts`
-   Каждый модуль имеет свои API сервисы
-   @Try декоратор для обработки ошибок

### Environment Configuration

-   Настройки в `.env.example`:
    -   VITE_API_URL - основной API
    -   VITE*ONEID*\* - настройки OneID SSO
    -   VITE*EIMZO*\* - настройки электронной подписи

## Development Notes

-   Проект использует модульную архитектуру - добавляйте новые функции как отдельные модули
-   Регистрируйте новые роуты через `registerRoutes()` в `src/modules/index.ts`
-   Используйте существующие Quasar компоненты из `src/components/quasar/`
-   Следуйте паттерну service классов для API вызовов
-   Применяйте v-permissions директиву для контроля доступа
