export type keySignType =
	| "languages"
	| "partners"
	| "detail"
	| "menu"
	| "faq"
	| "appeal"
	| "application_type"
	| "appeal_act"
	| "public_offer";

export function signOptions(key: keySignType) {
	switch (key) {
		case "languages":
			return ["logo"];
		case "partners":
			return ["logo"];
		case "detail":
			return ["logo"];
		case "menu":
			return ["logo"];
		case "faq":
			return ["logo"];
		case "application_type":
			return ["logo"];
		case "public_offer":
			return ["logo"];
		case "appeal":
			return ["documents"];
		case "appeal_act":
			return ["documents"];

		default:
			return [];
	}
}
