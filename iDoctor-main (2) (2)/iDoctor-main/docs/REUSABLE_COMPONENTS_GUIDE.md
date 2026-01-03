# 📚 Руководство по переиспользуемым компонентам

## 🎯 Обзор проделанной работы

В ходе рефакторинга модуля Role были созданы переиспользуемые компоненты, которые устранили дублирование кода и могут применяться в других модулях проекта.

## 🧩 Созданные компоненты

### 1. `useAppNavigation` - Композабл для навигации

**Путь:** `src/composables/useAppNavigation.ts`

**Назначение:** Централизует логику навигации, смены языка и выхода из системы.

**Возвращает:**

-   `toggleLeftDrawer()` - переход в профиль
-   `setLang(language)` - смена языка
-   `logout()` - выход из системы

### 2. `LoadingSkeleton` - Компонент скелетона загрузки

**Путь:** `src/components/LoadingSkeleton.vue`

**Назначение:** Унифицированный скелетон для состояния загрузки.

### 3. `AppFooter` - Универсальный футер приложения

**Путь:** `src/components/AppFooter.vue`

**Назначение:** Переиспользуемый футер с навигацией, настройками и действиями.

### 4. Общие стили

**Путь:** `src/styles/telegram-app.scss`

**Назначение:** Общие стили для Telegram Web App, включая переменные, анимации и адаптивность.

---

## 🔧 Инструкции по использованию

### 1. useAppNavigation Композабл

```vue
<script setup lang="ts">
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();
</script>
```

**Когда использовать:**

-   ✅ Страницы с полноэкранным layout
-   ✅ Страницы, требующие навигацию в профиль
-   ✅ Страницы с мультиязычностью
-   ❌ Модальные окна и диалоги
-   ❌ Компоненты без навигации

### 2. LoadingSkeleton Компонент

```vue
<template>
	<PageLoading :find="loadData" #="{ loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout v-else>
			<!-- Ваш контент -->
		</q-layout>
	</PageLoading>
</template>

<script setup lang="ts">
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
</script>
```

**Когда использовать:**

-   ✅ Полноэкранные страницы с загрузкой
-   ✅ Страницы с layout структурой
-   ❌ Маленькие компоненты
-   ❌ Inline загрузчики

### 3. AppFooter Компонент

#### Базовое использование:

```vue
<template>
	<q-layout view="hHh Lpr lff">
		<q-page-container>
			<!-- Ваш контент -->
		</q-page-container>

		<AppFooter
			:username="authStore.user?.username"
			:languages="$lang.languages"
			:current-language-id="$lang._currentLang?.id"
			@toggle-drawer="toggleLeftDrawer"
			@go-to-profile="toggleLeftDrawer"
			@set-lang="setLang"
			@logout="logout"
		/>
	</q-layout>
</template>
```

#### Расширенное использование:

```vue
<AppFooter
	:username="authStore.user?.username"
	:languages="$lang.languages"
	:current-language-id="$lang._currentLang?.id"
	<!--
	Кнопки
	действий
	--
>
  :show-add-button="true"
  :add-button-route="{ name: 'CREATE_ITEM' }"
  add-button-icon="add_circle"
  
  :show-back-button="true"
  :back-button-route="{ name: 'ITEMS_PAGE' }"
  
  <!-- Расположение -->
  :center-actions="false"
  
  <!-- События -->
  @toggle-drawer="toggleLeftDrawer"
  @go-to-profile="goToProfile"
  @set-lang="setLang"
  @logout="logout"
>
  <!-- Кастомные действия -->
  <template #actions>
    <q-btn icon="save" @click="save" />
    <q-btn icon="edit" @click="edit" />
  </template>
</AppFooter>
```

#### Параметры AppFooter:

| Параметр            | Тип               | По умолчанию   | Описание                                           |
| ------------------- | ----------------- | -------------- | -------------------------------------------------- |
| `username`          | `string?`         | `undefined`    | Имя пользователя для отображения                   |
| `languages`         | `LanguageType[]?` | `undefined`    | Список доступных языков                            |
| `currentLanguageId` | `number?`         | `undefined`    | ID текущего языка                                  |
| `showAddButton`     | `boolean`         | `false`        | Показать кнопку добавления                         |
| `addButtonRoute`    | `object?`         | `undefined`    | Маршрут для кнопки добавления                      |
| `addButtonIcon`     | `string`          | `"add_circle"` | Иконка кнопки добавления                           |
| `showBackButton`    | `boolean`         | `false`        | Показать кнопку назад                              |
| `backButtonRoute`   | `object?`         | `undefined`    | Маршрут для кнопки назад                           |
| `centerActions`     | `boolean`         | `true`         | Центрировать действия или выровнять по левому краю |

#### События AppFooter:

| Событие          | Описание                   |
| ---------------- | -------------------------- |
| `@toggle-drawer` | Переключение бокового меню |
| `@go-to-profile` | Переход в профиль          |
| `@set-lang`      | Смена языка                |
| `@logout`        | Выход из системы           |

### 4. Общие стили

```vue
<style scoped>
@import "@/styles/telegram-app.scss";
</style>
```

**Подключено глобально в `main.ts`**, поэтому дополнительный импорт нужен только для scoped стилей.

---

## 📋 Полный пример страницы

```vue
<script setup lang="ts">
import { ref } from "vue";
import { ItemService, type ItemType } from "@/service";
import PageLoading from "@/components/PageLoading.vue";
import LoadingSkeleton from "@/components/LoadingSkeleton.vue";
import AppFooter from "@/components/AppFooter.vue";
import { useAppNavigation } from "@/composables/useAppNavigation";
import { useAuthStore } from "@/store/auth-store";

const authStore = useAuthStore();
const { toggleLeftDrawer, setLang, logout } = useAppNavigation();

const items = ref<ItemType[]>([]);

async function loadItems() {
	const response = await ItemService.getAll();
	if (response) items.value = response;
}
</script>

<template>
	<PageLoading :find="loadItems" #="{ loading }">
		<LoadingSkeleton v-if="loading" />

		<q-layout view="hHh Lpr lff" v-else>
			<q-page-container>
				<q-page class="bg-white text-gray-900 overflow-auto p-4 pt-24">
					<!-- Ваш контент страницы -->
					<div v-for="item in items" :key="item.id">
						{{ item.name }}
					</div>
				</q-page>
			</q-page-container>

			<AppFooter
				:username="authStore.user?.username"
				:languages="$lang.languages"
				:current-language-id="$lang._currentLang?.id"
				:show-add-button="true"
				:add-button-route="{ name: 'CREATE_ITEM' }"
				@toggle-drawer="toggleLeftDrawer"
				@go-to-profile="toggleLeftDrawer"
				@set-lang="setLang"
				@logout="logout"
			/>
		</q-layout>
	</PageLoading>
</template>

<style scoped>
@import "@/styles/telegram-app.scss";
</style>
```

---

## ⚠️ Ограничения для максимальной переиспользуемости

### AppFooter

**✅ Подходит для:**

-   Полноэкранные страницы (с `q-layout`)
-   Страницы с единообразной навигацией
-   Страницы требующие стандартные действия (профиль, настройки, выход)

**❌ НЕ подходит для:**

-   Модальные окна и диалоги
-   Встраиваемые компоненты
-   Страницы с кардинально отличающейся навигацией
-   Страницы без необходимости в footer

**Ограничения:**

1. Требует `useAppNavigation` композабл
2. Использует `$lang` и `$tl` глобальные объекты
3. Зависит от структуры авторизации проекта

### LoadingSkeleton

**✅ Подходит для:**

-   Страницы с layout структурой
-   Долгие операции загрузки

**❌ НЕ подходит для:**

-   Inline загрузка небольших элементов
-   Страницы с кардинально отличающейся структурой

**Ограничения:**

1. Фиксированная структура скелетона
2. Не настраивается под разные layouts

### useAppNavigation

**✅ Подходит для:**

-   Любые страницы с навигацией
-   Страницы с мультиязычностью

**❌ НЕ подходит для:**

-   Страницы без навигации
-   Компоненты не требующие роутинга

**Ограничения:**

1. Зависит от структуры роутинга проекта
2. Требует определенные store (auth, language)
3. Использует специфичные route names

---

## 🎯 Рекомендации по применению

### Приоритет внедрения:

1. **Высокий приоритет:**

    - `Auth/pages/Profile.vue`
    - Другие View/Edit страницы модуля Role

2. **Средний приоритет:**

    - Полноэкранные Edit/View страницы других модулей
    - Страницы с аналогичной структурой

3. **Низкий приоритет:**
    - Страницы со специфичными требованиями
    - Legacy компоненты

### Перед внедрением проверьте:

-   [ ] Страница использует полноэкранный layout (`q-layout`)
-   [ ] Необходима стандартная навигация (профиль, настройки, выход)
-   [ ] Страница не является модальным окном
-   [ ] Доступны `authStore` и языковые настройки

---

## 📊 Результат рефакторинга

**Достигнутая экономия в модуле Role:**

-   `Page.vue`: с 634 до 75 строк (-88%)
-   `Edit.vue`: с 598 до 131 строки (-78%)
-   `Create.vue`: с 80 до 75 строк (добавлен полный функционал)

**Общие выгоды:**

-   Устранено дублирование кода
-   Унифицирован интерфейс
-   Упрощена поддержка
-   Повышена переиспользуемость

---

## 🛠 Примеры применения в других модулях

### Device/pages/Edit.vue

```vue
<AppFooter
	:username="authStore.user?.username"
	:languages="$lang.languages"
	:current-language-id="$lang._currentLang?.id"
	:show-add-button="false"
	:show-back-button="true"
	:back-button-route="{ name: 'DEVICE_PAGE' }"
	@toggle-drawer="toggleLeftDrawer"
	@set-lang="setLang"
	@logout="logout"
>
  <template #actions>
    <q-btn icon="edit" color="primary" @click="enableEdit" />
  </template>
</AppFooter>
```

### User/pages/Page.vue

```vue
<AppFooter
	:username="authStore.user?.username"
	:languages="$lang.languages"
	:current-language-id="$lang._currentLang?.id"
	add-button-icon="person_add"
	:add-button-route="{ name: 'USER_CREATE' }"
	:show-add-button="true"
	@toggle-drawer="toggleLeftDrawer"
	@set-lang="setLang"
	@logout="logout"
/>
```

### Auth/pages/Profile.vue (рекомендуется)

```vue
<AppFooter
	:username="authStore.user?.username"
	:languages="$lang.languages"
	:current-language-id="$lang._currentLang?.id"
	:show-add-button="false"
	:center-actions="false"
	@toggle-drawer="() => {}"
	@go-to-profile="() => router.push({ name: 'PAGE_PROFILE' })"
	@set-lang="setLang"
	@logout="logout"
/>
```

---

## 🚀 Следующие шаги

1. **Рефакторинг Profile.vue** - максимальная выгода (~200 строк экономии)
2. **Оценка других полноэкранных страниц** в модулях
3. **При необходимости расширение AppFooter** новыми возможностями

Компоненты готовы к применению в других модулях проекта! 🎉

---

## 📞 Поддержка

При возникновении вопросов или предложений по улучшению компонентов, создайте issue в репозитории проекта.
