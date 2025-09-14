import { get, post } from "@/utils/axios";
import type { ResponseData } from "@/types/common";
import type { BasicData } from "@/types/basic";

const reqBasicInfo = (): Promise<ResponseData<BasicData[]>> => {
  return get("/basic/info");
};

const reqBasicSubmit = (data: BasicData[]): Promise<ResponseData> => {
  return post("/basic/submit", data);
};

export { reqBasicInfo, reqBasicSubmit };
