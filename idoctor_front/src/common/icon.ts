import type Icons from "@quasar/extras/material-icons/index.d.ts";

type RemovePrefix<T extends string> = T extends `mat${infer Rest}` ? Rest : T;

type ToSnakeCase<T extends string> = T extends `${infer Head}${infer Tail}`
  ? `${Head extends Uppercase<Head> | " "
      ? Head extends `${number}`
        ? Head
        : `_${Lowercase<Head>}`
      : Head}${ToSnakeCase<Tail>}`
  : T;

type LowercaseFirst<T extends string> = T extends `${infer First}${infer Rest}`
  ? `${Lowercase<First>}${Rest}`
  : T;

export type ProcessedIcons = ToSnakeCase<
  LowercaseFirst<RemovePrefix<keyof typeof Icons>>
>;
