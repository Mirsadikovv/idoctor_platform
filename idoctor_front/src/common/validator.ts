import { isValidDate } from "./utils";

export function formRequired(message: string = "Это поле обязательно") {
	return (value: unknown) => {
		if (typeof value === "string") {
			return value.trim().length > 0 || message;
		}
		if (Array.isArray(value)) {
			return value.length > 0 || message;
		}
		if (typeof value === "number") {
			return !isNaN(value) || message;
		}
		return !!value || message;
	};
}

export function dateValid(message: string = "Неверная дата") {
	return (value: string) => {
		// если значение пустое — вернуть true (валидация на "обязательность" не в этом валидаторе)
		if (!value) return true;

		// проверяем только валидность, не наличие
		return isValidDate(value) || message;
	};
}

export function validatePIN(message: string = "Значение должно содержать ровно 14 цифр") {
	return (value: string | null | undefined) => {
		if (!value || typeof value !== "string") {
			return message;
		}

		// Регулярное выражение для проверки только цифр длиной 14
		const lengthRegex = /^\d{14}$/;

		if (!lengthRegex.test(value)) {
			return message;
		}

		// Если всё корректно, возвращаем true
		return true;
	};
}
export function validateINN(message: string = "Значение должно содержать ровно 9 цифр") {
	return (value: string | null | undefined) => {
		if (!value || typeof value !== "string") {
			return message;
		}

		// Регулярное выражение для проверки только цифр длиной 14
		const lengthRegex = /^\d{9}$/;

		if (!lengthRegex.test(value)) {
			return message;
		}

		// Если всё корректно, возвращаем true
		return true;
	};
}

export function formNumberLength(
	min: number,
	max: number,
	minMessage: string = "Минимальное число не может быть меньше %d",
	maxMessage: string = "Максимальное число не может быть больше %d",
) {
	return (value: number) => {
		if (typeof value !== "number") {
			return "Только число";
		}

		if (value < min) {
			return minMessage.replace("%d", `${formatNumber(min)}`);
		}

		if (value > max) {
			return maxMessage.replace("%d", `${formatNumber(max)}`);
		}
		return true;
	};
}

export function formMin(
	min: number,
	message: string = "Минимальная длина не может быть меньше %d",
) {
	return (value: string) => {
		return value?.length >= min || message.replace("%d", `${formatNumber(min)}`);
	};
}

export function formMax(
	max: number,
	message: string = "Максимальная длина не может быть больше %d",
) {
	return (value: string) => {
		return value.length <= max || message.replace("%d", `${formatNumber(max)}`);
	};
}

export function formMinArray(
	min: number,
	message: string = "Минимальная длина не может быть меньше %d",
) {
	return (value: string) => {
		return value?.length >= min || message.replace("%d", `${min}`);
	};
}

export function formLazyMin(
	min: number,
	message: string = "Минимальная длина не может быть меньше %d",
) {
	return (value: string) => {
		if (value === undefined) return true;
		return value?.length >= min || message.replace("%d", `${formatNumber(min)}`);
	};
}

export function formLazyMax(
	max: number,
	message: string = "Максимальная длина не может быть больше %d",
) {
	return (value: string) => {
		if (value === undefined) return true;
		return value?.length <= max || message.replace("%d", `${formatNumber(max)}`);
	};
}

export function formNumber(message: string = "Только число!") {
	return (value: string) => {
		const number = Number(value);
		return (Number.isInteger(number) && isFinite(number)) || message;
	};
}

export function validateDate(input: string) {
	if (!input) {
		return true;
	}

	const date = new Date(input);

	if (isNaN(date.getTime())) {
		return "Некорректный формат даты. Используйте YYYY/MM/DD.";
	}

	const today = new Date();
	today.setHours(0, 0, 0, 0);
	date.setHours(0, 0, 0, 0);

	if (date > today) {
		return true;
	}

	return "Дата должна быть больше сегодняшнего дня.";
}

export function formatNumber(number: number): string {
	return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, " ");
}
