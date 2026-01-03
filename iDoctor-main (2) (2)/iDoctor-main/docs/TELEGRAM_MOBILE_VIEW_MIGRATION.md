# Инструкция по миграции модулей на мобильный вид в стиле Telegram Web App

## 📋 Краткое описание задачи

Переделать кастомный мобильный вид ResponsiveTable в модулях для создания аккуратного интерфейса в стиле Telegram Web App с переходом на страницу детального просмотра вместо кнопок действий в карточках.

## 🎯 Что нужно сделать

### 1. Анализ существующей структуры модуля

**ОБЯЗАТЕЛЬНО проверить:**
- ✅ Какие поля отображаются в таблице
- ✅ Какое имя роута для страницы просмотра (например: `USER_VIEW`, `ROLE_VIEW`)
- ✅ Какой параметр передается в роут (обычно `id`)
- ✅ Какие кнопки действий сейчас есть в desktop версии
- ✅ Есть ли уже страница View.vue в модуле

### 2. Модификация мобильного вида в Page.vue

#### 2.1 Заменить template #card на новую структуру:

```vue
<!-- Кастомный мобильный вид -->
<template #card="{ model, orderNumber }">
  <q-item
    class="{module-name}-item-telegram"
    clickable
    :to="{ name: '{MODULE}_VIEW', params: { id: model.id } }"
  >
    <q-item-section avatar v-if="orderNumber">
      <q-avatar color="primary" text-color="white" size="md">
        {{ orderNumber }}
      </q-avatar>
    </q-item-section>

    <q-item-section>
      <q-item-label class="text-weight-bold text-h6">
        {{ /* ОСНОВНОЕ ПОЛЕ (например: полное имя, название роли) */ }}
      </q-item-label>
      <q-item-label caption class="text-body2">
        {{ /* ВТОРОСТЕПЕННОЕ ПОЛЕ (например: @username, код) */ }}
      </q-item-label>
      <q-item-label caption class="text-body2" v-if="/* УСЛОВИЕ ДЛЯ ТРЕТЬЕГО ПОЛЯ */">
        {{ /* ТРЕТЬЕ ПОЛЕ (например: телефон, email) */ }}
      </q-item-label>
      <!-- Чипы/статусы если есть -->
      <div class="q-mt-xs" v-if="/* УСЛОВИЕ ДЛЯ ЧИПА */">
        <q-chip color="primary" outline size="sm" dense>
          {{ /* СОДЕРЖИМОЕ ЧИПА */ }}
        </q-chip>
      </div>
    </q-item-section>

    <q-item-section side>
      <q-icon name="chevron_right" color="grey-6" />
    </q-item-section>
  </q-item>
</template>
```

#### 2.2 Добавить стили Telegram Web App:

```scss
<style scoped lang="scss">
.{module-name}-item-telegram {
  max-height: 120px;
  min-height: 90px;
  background: white;
  padding: 12px;
  transition: all 0.2s ease;
  cursor: pointer;

  &:hover {
    background: rgba(0, 0, 0, 0.02);
    border-color: rgba(0, 0, 0, 0.12);
    transform: translateY(-1px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }

  &:active {
    transform: translateY(0);
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
  }

  .q-item__section--avatar {
    padding-right: 16px;
  }

  .q-item__section--side {
    padding-left: 8px;
  }
}
</style>
```

### 3. Обновление страницы View.vue

#### 3.1 Добавить импорты для кнопок действий:

```vue
<script setup lang="ts">
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import IconDialog from "@/components/quasar/dialog/IconDialog.vue";
import Edit{ModuleName} from "./Edit.vue"; // Заменить на правильный компонент
import ConfirmDialog from "../components/ConfirmDialog.vue";

// Добавить функцию для обновления данных
const fetch{ModuleName} = () => {
  // Логика перезагрузки данных
};
</script>
```

#### 3.2 Добавить кнопки в заголовок:

```vue
<div class="flex! gap-x-4 items-center mb-6">
  <q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
  <q-breadcrumbs>
    <!-- breadcrumbs -->
  </q-breadcrumbs>
  <q-space />
  
  <!-- Кнопки действий -->
  <div class="flex gap-2" v-if="model.id">
    <ButtonDialog
      label="edit_{module_name}"
      icon="edit"
      color="primary"
      :fetch="fetch{ModuleName}"
    >
      <Edit{ModuleName} :id="model.id" :fetch="fetch{ModuleName}" />
    </ButtonDialog>
    
    <IconDialog
      iconColor="negative"
      icon="delete"
      tooltipText="delete_{module_name}"
      withTooltip
    >
      <ConfirmDialog :fetch="fetch{ModuleName}" :id="model.id" :isRemove="true" />
    </IconDialog>
  </div>
</div>
```

### 4. Очистка импортов в Page.vue

Убрать неиспользуемые импорты, которые теперь перенесены в View.vue:
- `IconDialog` (если не используется в других местах)
- `Edit{ModuleName}` компонент
- `ConfirmDialog`

## 📝 Примеры по типам модулей

### Модуль пользователей (User):
```vue
<q-item-label class="text-weight-bold text-h6">
  {{ model?.lastName }} {{ model?.firstName }} {{ model?.middleName }}
</q-item-label>
<q-item-label caption class="text-body2">
  @{{ model?.username || $tl("no_username") }}
</q-item-label>
<q-item-label caption class="text-body2" v-if="model?.phoneNumber">
  {{ model?.phoneNumber }}
</q-item-label>
```

### Модуль ролей (Role):
```vue
<q-item-label class="text-weight-bold text-h6">
  {{ model?.name }}
</q-item-label>
<q-item-label caption class="text-body2">
  {{ $tl("permissions_count") }}: {{ model?.permissions?.length || 0 }}
</q-item-label>
```

### Модуль языков (Language):
```vue
<q-item-label class="text-weight-bold text-h6">
  {{ model?.name }}
</q-item-label>
<q-item-label caption class="text-body2">
  {{ model?.description }}
</q-item-label>
<div class="q-mt-xs" v-if="model?.status">
  <q-chip :color="model.active ? 'positive' : 'negative'" outline size="sm" dense>
    {{ model.active ? $tl("active") : $tl("inactive") }}
  </q-chip>
</div>
```

## ✅ Чек-лист миграции

### В Page.vue:
- [ ] Заменен template #card на новую структуру
- [ ] Добавлен правильный роут в `:to`
- [ ] Убраны кнопки действий из карточки
- [ ] Добавлены стили `.{module-name}-item-telegram`
- [ ] Убраны неиспользуемые импорты
- [ ] Поля отображаются в правильном приоритете

### В View.vue:
- [ ] Добавлены импорты для ButtonDialog и IconDialog
- [ ] Добавлены кнопки редактирования и удаления
- [ ] Добавлена функция fetch{ModuleName} для обновления
- [ ] Добавлена условная проверка `v-if="model.id"`
- [ ] Кнопки размещены в правильном месте (в заголовке)

### Тестирование:
- [ ] Мобильный вид отображается корректно
- [ ] Клик по карточке открывает страницу просмотра
- [ ] Кнопки действий работают на странице просмотра
- [ ] Desktop версия не изменилась
- [ ] Анимации работают плавно

## 🚫 Ограничения

- **НЕ трогать** desktop версию таблицы
- **НЕ изменять** глобальные компоненты
- **НЕ использовать** избыточную логику
- **НЕ менять** темы и глобальные стили
- **НЕ добавлять** новые зависимости

## 📐 Структура CSS классов

Именование классов должно следовать паттерну:
- `.user-item-telegram` для модуля User
- `.role-item-telegram` для модуля Role  
- `.language-item-telegram` для модуля Language
- `{module-name}-item-telegram` для других модулей

## 💡 Советы по реализации

1. **Сначала изучите существующий код** - понимание текущей структуры критически важно
2. **Определите ключевые поля** для отображения (максимум 3-4 поля)
3. **Приоритизируйте информацию** - самое важное должно быть в `text-h6`
4. **Используйте caption для вторичной информации**
5. **Чипы только для статусов** - не перегружайте карточку
6. **Тестируйте на разных размерах** экрана

## 🔧 Типичные проблемы и решения

### Проблема: TypeScript ошибки с model.id
**Решение:** Добавить условную проверку `v-if="model.id"` для кнопок

### Проблема: Роут не найден  
**Решение:** Проверить правильность имени роута в файле routes.ts модуля

### Проблема: Стили не применяются
**Решение:** Убедиться что класс уникальный и не конфликтует с глобальными стилями

### Проблема: Анимации тормозят
**Решение:** Проверить что transition применен только к нужным свойствам