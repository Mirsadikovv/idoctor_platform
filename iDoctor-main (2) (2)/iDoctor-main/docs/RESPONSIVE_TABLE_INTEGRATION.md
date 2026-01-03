# ResponsiveTable Integration Guide

Инструкция по интеграции адаптивных таблиц в модули проекта

## 📋 Обзор

ResponsiveTable автоматически переключается между табличным видом (desktop) и мобильными карточками при ширине экрана < 1100px, сохраняя всю функциональность существующих таблиц.

## 🚀 Быстрая интеграция

### 1. Замена импорта

**БЫЛО:**
```vue
<script setup lang="ts">
import Table from "@/components/quasar/table/Table.vue";
</script>
```

**СТАЛО:**
```vue
<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
</script>
```

### 2. Замена компонента в template

**БЫЛО:**
```vue
<Table :models="pageData" hasOrder>
  <!-- все ваши слоты остаются без изменений -->
</Table>
```

**СТАЛО:**
```vue
<ResponsiveTable :models="pageData" hasOrder>
  <!-- все ваши слоты остаются без изменений -->
</ResponsiveTable>
```

## ✅ Автоматические возможности

После интеграции **автоматически** получите:

- **Desktop (≥1100px)**: обычная таблица как раньше
- **Mobile (<1100px)**: красивые карточки с границами и анимациями
- **Все слоты работают**: thead, body, tfoot, actions
- **Типизация сохранена**: полная поддержка TypeScript generics
- **Разрешения работают**: v-permissions директивы активны

## 🎨 Кастомизация мобильного вида

### Базовая кастомизация

Добавьте слот `#card` для создания кастомного мобильного интерфейса:

```vue
<ResponsiveTable :models="pageData" hasOrder>
  <!-- Обычные слоты для desktop таблицы -->
  <template #name="{ model }">{{ model.name }}</template>
  <template #status="{ model }">
    <q-chip :color="model.active ? 'positive' : 'negative'">
      {{ model.active ? 'Активен' : 'Неактивен' }}
    </q-chip>
  </template>

  <!-- Кастомный мобильный вид -->
  <template #card="{ model, orderNumber }">
    <q-item class="q-mb-md language-item-bordered">
      <q-item-section avatar v-if="orderNumber">
        <q-avatar color="primary" text-color="white" size="sm">
          {{ orderNumber }}
        </q-avatar>
      </q-item-section>

      <q-item-section>
        <q-item-label>{{ $tl("NAME") }}: {{ model.name }}</q-item-label>
        <q-item-label>{{ $tl("DESCRIPTION") }}: {{ model.description }}</q-item-label>
        <q-item-label>
          {{ $tl("STATUS") }}:
          <q-chip
            :color="!model.deletedAt ? 'positive' : 'negative'"
            text-color="white"
            size="sm"
          >
            {{ !model.deletedAt ? $tl("ACTIVE") : $tl("DELETED") }}
          </q-chip>
        </q-item-label>
      </q-item-section>
    </q-item>

    <!-- Кнопки действий -->
    <div class="card-actions flex justify-end gap-4">
      <ButtonDialog
        v-if="!model.deletedAt"
        icon="edit"
        color="primary"
        size="sm"
        round
        flat
        :fetch="fetch"
      >
        <EditComponent :id="model.id" :fetch="fetch" />
      </ButtonDialog>
    </div>
  </template>
</ResponsiveTable>
```

### Добавление стилей для границ

```scss
<style scoped lang="scss">
.language-item-bordered {
  border: 1px solid rgba(0, 0, 0, 0.12);
  border-radius: 8px;
  padding: 12px;
}

.card-actions {
  border-top: 1px solid rgba(0, 0, 0, 0.1);
  padding-top: 12px;

  @media (max-width: 480px) {
    flex-direction: column;
    gap: 8px !important;

    :deep(.q-btn) {
      width: 100% !important;
    }
  }
}
</style>
```

## 🔧 Доступные параметры

### Props для ResponsiveTable

```typescript
interface Props<T> {
  models: PageDataType<T>;        // Данные для отображения
  loading?: boolean;              // Состояние загрузки
  pick?: Record<string, boolean>; // Фильтр колонок
  traction?: (model: T, index: number) => any; // CSS классы строк
  customRowStyle?: (model: T, index: number) => any; // Стили строк
  whithoutMaxHeight?: boolean;    // Убрать ограничение высоты
  hasOrder?: boolean;             // Показать номера строк
}
```

### Слоты для кастомизации

| Слот | Описание | Доступные данные |
|------|----------|------------------|
| `#card` | Кастомный мобильный вид | `{ model, index, fields, orderNumber }` |
| `#{field}` | Содержимое ячейки | `{ model, index }` |
| `#{field}:thead` | Заголовок колонки | `{ key }` |
| `#before` | Контент перед таблицей | `{ models, totalPages }` |
| `#tfoot` | Подвал таблицы | `{ models, totalPages }` |

## 📁 Примеры интеграции по модулям

### User модуль
```vue
<!-- src/modules/User/pages/Page.vue -->
<script setup lang="ts">
import ResponsiveTable from "@/components/quasar/table/ResponsiveTable.vue";
// остальные импорты...
</script>

<template>
  <ResponsiveTable :models="userPage" hasOrder>
    <template #fullName="{ model }">{{ model.fullName }}</template>
    <template #phoneNumber="{ model }">{{ model.phoneNumber }}</template>
    
    <!-- Кастомная мобильная карточка для пользователей -->
    <template #card="{ model, orderNumber }">
      <q-item class="q-mb-md user-item-bordered">
        <q-item-section avatar>
          <q-avatar>
            <img :src="model.avatar || '/default-avatar.png'">
          </q-avatar>
        </q-item-section>
        <q-item-section>
          <q-item-label>{{ $tl("FULL_NAME") }}: {{ model.fullName }}</q-item-label>
          <q-item-label>{{ $tl("PHONE") }}: {{ model.phoneNumber }}</q-item-label>
          <q-item-label>{{ $tl("EMAIL") }}: {{ model.email }}</q-item-label>
        </q-item-section>
      </q-item>
    </template>
  </ResponsiveTable>
</template>
```

### Role модуль
```vue
<!-- src/modules/Role/pages/Page.vue -->
<ResponsiveTable :models="rolePage" hasOrder>
  <template #card="{ model, orderNumber }">
    <q-item class="q-mb-md role-item-bordered">
      <q-item-section avatar v-if="orderNumber">
        <q-avatar color="secondary" icon="security">{{ orderNumber }}</q-avatar>
      </q-item-section>
      <q-item-section>
        <q-item-label>{{ $tl("ROLE_NAME") }}: {{ model.name }}</q-item-label>
        <q-item-label>{{ $tl("PERMISSIONS_COUNT") }}: {{ model.permissions?.length || 0 }}</q-item-label>
      </q-item-section>
    </q-item>
  </template>
</ResponsiveTable>
```

## 🎯 Best Practices

### 1. Консистентность дизайна
- Используйте одинаковую структуру мобильных карточек во всех модулях
- Применяйте единые CSS классы (`{module}-item-bordered`)
- Сохраняйте цветовую схему проекта

### 2. Адаптивные кнопки
```scss
.card-actions {
  @media (max-width: 480px) {
    flex-direction: column;
    gap: 8px !important;

    :deep(.q-btn) {
      width: 100% !important;
    }
  }
}
```

### 3. Оптимизация контента
- На мобильных показывайте только ключевую информацию
- Используйте иконки вместо длинного текста где возможно
- Группируйте связанные действия

## 🔍 Troubleshooting

### Проблема: Слоты не работают
**Решение:** Убедитесь что используете точно такие же имена слотов как в обычной таблице

### Проблема: TypeScript ошибки
**Решение:** ResponsiveTable поддерживает те же generic типы что и Table.vue

### Проблема: Стили не применяются
**Решение:** Добавьте CSS классы для границ и отступов как в примерах выше

## 📊 Результат интеграции

После интеграции все модули получат:

- ✅ **Автоматическая адаптивность** без изменения логики
- ✅ **Современный мобильный UI** с карточками и анимациями  
- ✅ **Полная обратная совместимость** с существующим кодом
- ✅ **Единый дизайн-язык** во всём приложении
- ✅ **Улучшенный UX** на мобильных устройствах

---

💡 **Совет**: Начните интеграцию с простых модулей без сложных действий, затем переходите к более комплексным таблицам.