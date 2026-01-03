import { useAuthStore } from "@/store/auth-store";
import type { Directive, DirectiveBinding } from "vue";
import { useRoute, useRouter } from "vue-router";

export function permissionsDirective(): Directive {
	type BindingType = DirectiveBinding<string[]>;

	const authStore = useAuthStore();

	const route = useRoute();
	const router = useRouter();

	const action = (el: HTMLElement, binding: BindingType) => {
		const { value, arg } = binding;

		if (arg === "page") {
			const routes = router.getRoutes();

			const hasRoute = routes.some((route) => {
				const routeName = String(route.name);
				return authStore.page(routeName);
			});

			//   if (hasRoute) return true;

			return el.remove();
			//   return (el.style.display = result ? "" : "none");
		}

		const routeName = String(route.name);

		const result = authStore.page(routeName);

		if (result === undefined) {
			//   return (el.style.display = "none");
			return el.remove();
		}

		if (!result.length) {
			return true;
		}

		const res = result.some((item: string) => value.some((val) => val === item));

		if (!res) {
			return el.remove();
		}

		return true;
	};

	const permissionsDirectiveObject: Directive = {
		// beforeMount: action,
		beforeMount: action,
		// created: action,
		// mounted: action,
		// updated: action,
		// deep: true,
	};

	return permissionsDirectiveObject;
}

export default permissionsDirective;
