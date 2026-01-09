import { useLanguageStore } from "@/store/language-store";
import { Notify } from "quasar";

const langStore = useLanguageStore();

export function ErrorNotify(message: string) {
	Notify.create({
		color: "negative",
		textColor: "white",
		icon: "sentiment_very_dissatisfied",
		message: message ? langStore.tl(message) : langStore.tl("What went wrong"),
		position: "top",
		timeout: 5000,
		actions: [{ icon: "close", color: "white" }],
	});
}
export function SuccesNotify(message: string) {
	Notify.create({
		color: "positive",
		textColor: "white",
		icon: "tag_faces",
		message: message ? langStore.tl(message) : langStore.tl("Successfull"),
		position: "top",
		timeout: 5000,
		actions: [{ icon: "close", color: "white" }],
	});
}

// @Try({
// 	async onSuccess(result) {
// 		(await import("@/common/Notify")).SuccesNotify(result.statusText);
// 	},
// 	async onError(err) {
// 		(await import("@/common/Notify")).ErrorNotify(
// 			// @ts-ignore
// 			err?.response?.data.message || err.message,
// 		);
// 	},
// })
