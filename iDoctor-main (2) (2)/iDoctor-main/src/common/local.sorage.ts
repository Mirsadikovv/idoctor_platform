import { LocalStorage } from "quasar";

export function Token(_default: string = ""): string {
  const token = LocalStorage.getItem<string>(TOKEN_NAME);

  if (!token) {
    return _default;
  }

  return token;
}

export function HasToken(): boolean {
  return !!Token();
}

export function TokenWithBearer(): string | undefined {
  const token = Token();

  if (!token) {
    return undefined;
  }

  return `Bearer ${token}`;
}

export function SetToken(token: string) {
  LocalStorage.set(TOKEN_NAME, token);
}

export function RemoveToken() {
  LocalStorage.remove(TOKEN_NAME);
}

export function HasLang(): boolean {
  return !!Lang();
}

export function Lang(): string | null {
  return LocalStorage.getItem<string>(LANG_NAME);
}

export function SetLang(lang: string) {
  LocalStorage.set(LANG_NAME, lang);
}
