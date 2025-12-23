<script setup lang="ts">
import { useLanguageStore } from "@/store/language-store";
import type { QEditorProps } from "quasar";
// import { computed, useSlots } from "vue";
const { tl } = useLanguageStore();

// Define props with stricter typing and omit unnecessary modelValue override
interface Props extends Omit<QEditorProps, "modelValue"> {
	modelValue?: string;
	label?: string;
}

// Define defaults for props to reduce redundancy and ensure consistency
const props = withDefaults(defineProps<Props>(), {
	label: "content_editor",
	color: "secondary",
	outlined: true,
	lazyRules: true,
	minHeight: "5rem",
	dense: false,
	toolbarToggleColor: "secondary",
	paragraphTag: "div",
	placeholder: "content_editor",
});

// Define model with proper typing
const modelValue = defineModel<string>({
	default: "",
});
function toolbarSettigs() {
	return [
		[
			{
				label: tl("text_align"),
				icon: "format_align_left",
				fixedLabel: true,
				list: "only-icons",
				options: ["left", "center", "right", "justify"],
			},
		],
		["bold", "italic", "strike", "underline", "subscript", "superscript"],
		["token", "hr", "link", "custom_btn"],
		["print", "fullscreen"],
		[
			{
				label: tl("formatting"),
				icon: "format_underlined",
				list: "no-icons",
				options: ["p", "h1", "h2", "h3", "h4", "h5", "h6", "code"],
			},
			{
				label: tl("fontSize"),
				icon: "text_fields",
				fixedLabel: true,
				fixedIcon: true,
				list: "no-icons",
				options: ["size-1", "size-2", "size-3", "size-4", "size-5", "size-6", "size-7"],
			},
			// {
			// 	label: "defaultFont",
			// 	icon: "font_download",
			// 	fixedIcon: true,
			// 	list: "no-icons",
			// 	options: [
			// 		"default_font",
			// 		"arial",
			// 		"arial_black",
			// 		"comic_sans",
			// 		"courier_new",
			// 		"impact",
			// 		"lucida_grande",
			// 		"times_new_roman",
			// 		"verdana",
			// 	],
			// },
			"removeFormat",
		],
		// ["quote", "unordered", "ordered", "outdent", "indent"],
		["undo", "redo"],
		["viewsource"],
	];
}
// Правильное использование useSlots
// const slotsInstance = useSlots();
// const slots = computed(() => Object.keys(slotsInstance));
</script>

<template>
	<q-editor
		v-bind="props"
		v-model="modelValue"
		:placeholder="$tl(props.placeholder)"
		:toolbar="toolbarSettigs()"
	>
		<!-- <template v-for="slot in slots" :key="slot" #[slot]="slotProps: Record<string, any>">
			<slot :name="slot" v-bind="{ ...slotProps }"></slot>
		</template> -->
	</q-editor>
</template>

<style lang="css">
.q-editor__toolbar {
	height: 30px !important;
}
</style>

<!-- // toolbar: () => [ // [ // { // label: "text_align", // icon: "format_align_left", // fixedLabel:
true, // list: "only-icons", // options: ["left", "center", "right", "justify"], // }, // ], //
["bold", "italic", "strike", "underline", "subscript", "superscript"], // ["token", "hr", "link",
"custom_btn"], // ["print", "fullscreen"], // [ // { // label: "formatting", // icon:
"format_underlined", // list: "no-icons", // options: ["p", "h1", "h2", "h3", "h4", "h5", "h6",
"code"], // }, // { // label: "fontSize", // icon: "text_fields", // fixedLabel: true, // fixedIcon:
true, // list: "no-icons", // options: ["size-1", "size-2", "size-3", "size-4", "size-5", "size-6",
"size-7"], // }, // { // label: "defaultFont", // icon: "font_download", // fixedIcon: true, //
list: "no-icons", // options: [ // "default_font", // "arial", // "arial_black", // "comic_sans", //
"courier_new", // "impact", // "lucida_grande", // "times_new_roman", // "verdana", // ], // }, //
"removeFormat", // ], // // ["quote", "unordered", "ordered", "outdent", "indent"], // ["undo",
"redo"], // ["viewsource"], // ], // fonts: () => ({ // arial: "Arial", // arial_black: "Arial
Black", // comic_sans: "Comic Sans MS", // courier_new: "Courier New", // impact: "Impact", //
lucida_grande: "Lucida Grande", // times_new_roman: "Times New Roman", // verdana: "Verdana", // }), -->
