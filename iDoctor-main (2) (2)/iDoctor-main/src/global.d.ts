interface EIMZOClient {
  API_KEYS: string[]
  NEW_API: boolean
  NEW_API2: boolean

  changeKeyPassword: (itemObject: any, success: Function, fail: Function) => void
  checkVersion: (success: Function, fail: Function) => void
  createPkcs7: (
    id: any,
    data: any,
    timestamp: any,
    success: Function,
    fail: Function,
    detached?: boolean,
    isDataBase64Encoded?: boolean,
  ) => void
  idCardIsPLuggedIn: (success: Function, fail: Function) => void
  installApiKeys: (success: Function, fail: Function) => void
  listAllUserKeys: (
    itemIdGen: Function,
    itemUiGen: Function,
    success: Function,
    fail: Function,
  ) => void
  loadKey: (itemObject: any, success: Function, fail: Function, verifyPassword?: Function) => void
}

declare global {
  interface Window {
    EIMZOClient: EIMZOClient
  }
}

export {}
