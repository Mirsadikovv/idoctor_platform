// import ExpansionItem from "@/components/sidebar/ExpansionItem.vue";
import Link from "@/components/sidebar/Link.vue";
import type { QExpansionItemProps } from "quasar";
import { h, inject, type VNode } from "vue";
import { type RouteLocationRaw } from "vue-router";

export interface ExtendedProps extends QExpansionItemProps {}

export interface LinkProps {
  to: RouteLocationRaw;
  activeLinkGroup: string;
  title: string;
  icon: string;
  active: Function;
}

export function createLinkComponent(props: LinkProps) {
  return h(Link, props);
}

// export function createExtendedComponent(
// 	props: ExtendedProps,
// 	children: () => VNode[],
// ) {
// 	return () => h(ExpansionItem, props, children);
// }

export function activeList(): Set<string> | undefined {
  return inject("activeList");
}
