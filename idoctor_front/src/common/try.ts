type TryOptions<T> = {
	onSuccess?: (result: T) => void;
	onError?: (err: any) => void | Promise<any>;
};

export function Try<T = any>(options?: TryOptions<T>) {
	const { onSuccess, onError } = options || {};

	type AsyncFunction = (...args: any[]) => Promise<any> | ((...args: any[]) => any);

	return (target: any, _name: string, descriptor: PropertyDescriptor) => {
		const method = descriptor.value as AsyncFunction;

		const { name } = method.constructor;

		if (name !== "AsyncFunction" && name !== "Function") {
			throw new Error(`type {${name}} not supported`);
		}

		descriptor.value = (...args: any[]) => {
			try {
				const result: any = method.apply(target, args);

				if (result instanceof Promise) {
					return result
						.then((res) => {
							onSuccess && onSuccess(res);
							return res;
						})
						.catch((error: any) => {
							return onError && onError(error);
						});
				}
				return result;
			} catch (error: any) {
				return onError && onError(error);
			}
		};
	};
}
