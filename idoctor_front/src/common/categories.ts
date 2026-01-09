export async function searchPosition(_params: string = "") {
	return [
		"MAIN_CONTENT",
		"HEADER_NAVIGATION",
		"FOOTER_NAVIGATION",
		"FOOTER_CONTACTS",
		"FOOTER_SOCIALS",
		"APPLICATION_TYPE",
		"PARTNERS",
		"SOCIALS",
	];
}
export async function searchPositionMenu(_params: string = "") {
	return ["HEADER_NAVIGATION", "FOOTER_NAVIGATION", "FOOTER_CONTACTS"];
}
export async function searchPositionStatic(_params: string = "") {
	return ["HEADER_NAVIGATION", "FOOTER_NAVIGATION", "FOOTER_CONTACTS"];
}
export async function searchPositionPartner(_params: string = "") {
	return ["MAIN_CONTENT"];
}
export async function searchPositionFaq(_params: string = "") {
	return ["MAIN_CONTENT"];
}
export async function searchPositionApplication(_params: string = "") {
	return ["MAIN_CONTENT"];
}
export async function searchPositionPublicOffer(_params: string = "") {
	return ["MAIN_CONTENT"];
}
export async function searchDetailsCategories(_params: string = "") {
	return ["FAQ", "MENU", "PARTNERS", "APPLICATION_TYPE", "PUBLIC_OFFER"];
}
