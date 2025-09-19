/// <reference types="vite/client" />

declare module "*.vue" {
  import type { DefineComponent } from "vue";
  const component: DefineComponent<{}, {}, any>;
  export default component;
}

declare module "file-saver" {
  export function saveAs(
    data: Blob,
    filename: string,
    options?: SaveAsOptions
  ): void;

  interface SaveAsOptions {
    autoBom?: boolean;
  }
}