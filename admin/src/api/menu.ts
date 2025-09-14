import { get } from "@/utils/axios";
import type { ResponseData } from "@/types/common";
import type { MenuData } from "@/types/menu";

const reqMenuList = (): Promise<ResponseData<MenuData[]>> => {
  return get("/menu/list");
};

export { reqMenuList };
