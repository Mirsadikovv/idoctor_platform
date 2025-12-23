import ExpansionItem from "@/components/quasar/sidebar/ExpansionItem.vue";
import Link from "@/components/sidebar/Link.vue";
import { routers } from "@/router/layout";
import { date, QList } from "quasar";
import { h } from "vue";
import type { RouteRecordRaw } from "vue-router";

export function buildOneIdUrl() {
	return ``;
}

export function dateFormat(
	inDate: string | Date | number | undefined | null,
	format: string = "YYYY-MM-DD HH:mm",
) {
	if (!inDate) {
		return undefined;
	}

	return date.formatDate(inDate, format) || "";
}

export function parseDate(
	inDate: string | Date | number | undefined | null,
	format: string = "YYYY-MM-DD HH:mm",
) {
	if (!inDate) {
		return undefined;
	}

	return date.extractDate(inDate as string, format) || "";
}

export async function delay(timeout: number) {
	return new Promise((resolve) => setTimeout(resolve, timeout));
}

export function currencyFormat(value?: number | string): string {
	if (value === null || value === undefined) return "";
	const number = typeof value === "string" ? parseFloat(value.replace(/\s/g, "")) : value;
	if (isNaN(number)) return "";

	return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, " ");
}

export function removeEmptyAndToFormData<T extends Record<string, any>>(obj: T): FormData {
	const formData = new FormData();

	Object.entries(obj).forEach(([key, value]) => {
		// Проверяем, что значение не null, не undefined и не пустая строка
		if (value !== null && value !== undefined && value !== "") {
			// Если значение - объект Blob или File, добавляем как есть
			if (value instanceof Blob || value instanceof File) {
				formData.append(key, value);
			} else {
				// Для остальных значений преобразуем в строку
				formData.append(key, String(value));
			}
		}
	});

	return formData;
}

export function buildSidebar() {
	return () => {
		const flatRoutes = flattenRoutes(routers());

		// Функция проверки доступа к маршруту
		const hasRouteAccess = (route: RouteRecordRaw): boolean => {
			return !!route.name;
		};

		// Функция проверки доступа к группе (проверяет, есть ли доступные дочерние элементы)
		const hasGroupAccess = (groupRoute: RouteRecordRaw): boolean => {
			if (!groupRoute.children) return false;
			return groupRoute.children.some(
				(child) => child.meta?.sidebar && hasRouteAccess(child),
			);
		};

		// Предварительная подготовка данных
		const expandedRoutes = flatRoutes.filter((r) => r.meta?.expanded);
		expandedRoutes.forEach((expanded) => {
			expanded.children?.forEach((r) => {
				if (r.meta?.sidebar) {
					r.meta.sidebar.isExpandedGroup = true;
				}
			});
		});

		// Фильтруем обычные ссылки по разрешениям
		const sidebars = flatRoutes.filter(
			(route) =>
				route.meta?.sidebar && !route.meta.sidebar.isExpandedGroup && hasRouteAccess(route),
		);

		// Фильтруем группы по наличию доступных дочерних элементов
		const accessibleExpandedRoutes = expandedRoutes.filter(hasGroupAccess);

		// Сортировка маршрутов
		const sortedRoutes = [...sidebars, ...accessibleExpandedRoutes].sort(
			(a, b) => (a.meta?.sort || 0) - (b.meta?.sort || 0),
		);

		// Вспомогательная функция для создания элементов (без директивы permissions)
		const createLink = (route: any) =>
			h(Link, {
				to: { name: route.name },
				title: route.meta?.sidebar?.label,
				icon: route.meta?.sidebar?.icon,
				activeLinkGroup: route.meta?.activeLinkGroup,
				active: (group: string) => group === route.meta?.activeLinkGroup,
			});

		const createExpandedGroup = (route: any, subRoutes: any[]) => {
			// Фильтруем и сортируем только доступные вложенные элементы
			const accessibleSubRoutes = subRoutes
				.filter(hasRouteAccess)
				.sort((a, b) => (a.meta?.sort || 0) - (b.meta?.sort || 0));

			const children = accessibleSubRoutes.map(createLink);

			return h(
				ExpansionItem,
				{
					contentInsetLevel: 0.3,
					icon: route.meta?.expanded?.icon,
					label: route.meta?.expanded?.label,
				},
				children,
			);
		};

		// Генерация элементов
		const elements = sortedRoutes.map((route) => {
			if (route.meta?.sidebar && !route.meta.sidebar.isExpandedGroup) {
				return createLink(route);
			} else {
				const subRoutes = route.children?.filter((r) => r.meta?.sidebar) || [];
				return createExpandedGroup(route, subRoutes);
			}
		});

		// Рендеринг списка
		return h(
			QList,
			{
				bordered: true,
				separator: true,
				class: "text-primary select-text! min-w-[0px]!",
			},
			() => {
				return elements;
			},
		);
	};
}

export function flattenRoutes(routes: Readonly<RouteRecordRaw[]>): RouteRecordRaw[] {
	return routes.flatMap((route: RouteRecordRaw) => {
		if (route.children && route.children.length > 0) {
			return [route, ...flattenRoutes(route.children)];
		}
		return [route];
	});
}

export function removeEmptyQueryParams(query: string): string {
	try {
		// Удаляем начальные и конечные пробелы, а также символы ?
		const cleanedQuery = query.trim().replace(/^\?+|\?+$/g, "");

		// Если строка пустая после очистки, возвращаем пустую строку
		if (!cleanedQuery) {
			return "";
		}

		const params = new URLSearchParams(cleanedQuery);

		// Удаляем параметры с пустыми значениями, null, undefined или только пробелами
		for (const [key, value] of Array.from(params.entries())) {
			if (value === "" || value === null || value === undefined || value.trim() === "") {
				params.delete(key);
			}
		}

		// Возвращаем очищенную query-строку
		return params.toString();
	} catch (error) {
		// Логируем ошибку и возвращаем пустую строку для некорректных входных данных
		console.error("Invalid query string:", error);
		return "";
	}
}

export function queryToObject(queryString: string): Record<string, string> {
	// Удаляем начальный символ '?' если он есть
	const cleanQuery: string = queryString.startsWith("?") ? queryString.slice(1) : queryString;

	// Если строка пустая, возвращаем пустой объект
	if (!cleanQuery) return {};

	// Разбиваем строку на пары ключ-значение и преобразуем в объект
	return cleanQuery.split("&").reduce((obj: Record<string, string>, pair: string) => {
		// Разбиваем пару на ключ и значение
		const [key, value = ""]: string[] = pair.split("=").map(decodeURIComponent);
		// Добавляем в объект, обрабатывая множественные значения для одного ключа
		if (key) {
			obj[key] = value;
		}
		return obj;
	}, {});
}

export function getId<T>(item: T | number): number {
	if (item && typeof item === "object") {
		// @ts-ignore
		return item.id as number;
	}
	return item as number;
}

export function getSoatoId<T>(item: T | number): number {
	if (item && typeof item === "object") {
		// @ts-ignore
		return item.soatoId as number;
	}
	return item as number;
}

export function findItemById<T extends { id?: number }>(
	items: T[],
	id: number | string | T | undefined | null,
): T | undefined {
	if (id === undefined || id === null) {
		return undefined;
	}
	if (typeof id === "object") {
		return id;
	}
	return items.find((item) => item.id === +id);
}

export function buildQuery(params: Record<string, any>): string {
	const filteredEntries = Object.entries(params).filter(
		([_, value]) =>
			value !== undefined &&
			value !== null &&
			value !== "" &&
			!(Array.isArray(value) && value.length === 0),
	);

	const searchParams = new URLSearchParams();

	for (const [key, value] of filteredEntries) {
		if (Array.isArray(value)) {
			value.forEach((v) => searchParams.append(key, String(v)));
		} else {
			searchParams.append(key, String(value));
		}
	}

	return searchParams.toString();
}

export function isValidDate(input: string): boolean {
	if (!input) {
		return false;
	}
	const dateRegex = /^(\d{2})-(\d{2})-(\d{4})$/;
	const match = input.match(dateRegex);
	if (!match) return false;

	const [_, dd, mm, yyyy] = match;
	const day = Number(dd);
	const month = Number(mm);
	const year = Number(yyyy);

	// Проверка числовых границ
	if (day < 1 || day > 31 || month < 1 || month > 12 || year < 1000 || year > 9999) return false;

	// Проверка на существующую дату
	const date = new Date(year, month - 1, day);
	return date.getDate() === day && date.getMonth() === month - 1 && date.getFullYear() === year;
}

/**
 * Универсальная функция для извлечения ID из массива объектов или чисел
 * @param items - массив объектов с полем id или массив чисел, может быть null/undefined
 * @param idField - название поля с ID (по умолчанию 'id')
 * @returns массив чисел или null
 */
export function getIds<T extends Record<string, any>>(
	items: null | (T | number)[] | number[] | undefined,
	idField: keyof T = "id" as keyof T,
): number[] | null {
	if (items === null || items === undefined) {
		return null;
	}

	if (!Array.isArray(items)) {
		return null; // Некорректный тип данных
	}

	const ids: number[] = [];

	for (const item of items) {
		if (typeof item === "number") {
			ids.push(item);
		} else if (
			typeof item === "object" &&
			item !== null &&
			idField in item &&
			typeof item[idField] === "number"
		) {
			ids.push(item[idField] as number);
		} else {
			// Некорректные данные в массиве
			return null;
		}
	}

	return ids.length > 0 ? ids : null; // Возвращаем null, если массив оказался пустым
}
